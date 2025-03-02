package edge

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/edges"
	"github.com/snple/beacon/pb/nodes"
	"github.com/snple/beacon/util/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeUpService struct {
	es *EdgeService

	NodeConn *grpc.ClientConn

	token string
	lock  sync.RWMutex

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func newNodeUpService(es *EdgeService) (*NodeUpService, error) {
	var err error

	es.Logger().Sugar().Infof("link node service: %v", es.dopts.NodeOptions.Addr)

	nodeConn, err := grpc.NewClient(es.dopts.NodeOptions.Addr, es.dopts.NodeOptions.GRPCOptions...)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(es.Context())

	s := &NodeUpService{
		es:       es,
		NodeConn: nodeConn,
		ctx:      ctx,
		cancel:   cancel,
	}

	return s, nil
}

func (s *NodeUpService) start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	s.es.Logger().Sugar().Info("node service started")

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			s.try()
		}
	}
}

func (s *NodeUpService) stop() {
	s.cancel()
	s.closeWG.Wait()
}

func (s *NodeUpService) push() error {
	// login
	operation := func() error {
		err := s.login(s.ctx)
		if err != nil {
			s.es.Logger().Sugar().Errorf("node login: %v", err)
		}

		return err
	}

	err := backoff.Retry(operation, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
	if err != nil {
		s.es.Logger().Sugar().Errorf("backoff.Retry: %v", err)
		return err
	}

	s.es.Logger().Sugar().Info("node login success")

	s.es.GetSync().setNodeUpdatedLocalToRemote(s.ctx, time.Time{})

	if err := s.sync2(s.ctx); err != nil {
		return err
	}

	if err := s.syncTagValue2(s.ctx); err != nil {
		return err
	}

	s.es.Logger().Sugar().Info("push success")

	return nil
}

func (s *NodeUpService) pull() error {
	// login
	operation := func() error {
		err := s.login(s.ctx)
		if err != nil {
			s.es.Logger().Sugar().Errorf("node login: %v", err)
		}

		return err
	}

	err := backoff.Retry(operation, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
	if err != nil {
		s.es.Logger().Sugar().Errorf("backoff.Retry: %v", err)
		return err
	}

	s.es.Logger().Sugar().Info("node login success")

	s.es.GetSync().setNodeUpdatedRemoteToLocal(s.ctx, time.Time{})

	if err := s.sync1(s.ctx); err != nil {
		return err
	}

	if err := s.syncTagValue1(s.ctx); err != nil {
		return err
	}

	s.es.Logger().Sugar().Info("pull success")

	return nil
}

func (s *NodeUpService) try() {
	// login
	operation := func() error {
		err := s.login(s.ctx)
		if err != nil {
			s.es.Logger().Sugar().Infof("node login: %v", err)
		}

		return err
	}

	err := backoff.Retry(operation, backoff.WithContext(backoff.NewExponentialBackOff(), s.ctx))
	if err != nil {
		s.es.Logger().Sugar().Errorf("backoff.Retry: %v", err)
		return
	}

	s.es.Logger().Sugar().Info("node login success")

	s.NodeLink(s.ctx)
	s.link(true)
	defer s.link(false)

	if err := s.sync(s.ctx); err != nil {
		s.es.Logger().Sugar().Errorf("sync: %v", err)
		time.Sleep(time.Second * 3)
		return
	}

	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	// start ticker
	go s.ticker(ctx)

	// start realtime sync
	if s.es.dopts.SyncOptions.Realtime {
		go s.waitRemoteNodeUpdated(ctx)
		go s.waitLocalNodeUpdated(ctx)
		go s.waitRemoteTagValueUpdated(ctx)
		go s.waitLocalTagValueUpdated(ctx)
		go s.waitRemoteTagWriteUpdated(ctx)
		go s.waitLocalTagWriteUpdated(ctx)
	}

	// keep alive
	{
		ctx = metadata.SetToken(ctx, s.GetToken())

		stream, err := s.NodeServiceClient().KeepAlive(ctx, &pb.MyEmpty{})
		if err != nil {
			s.es.Logger().Sugar().Errorf("KeepAlive: %v", err)
			return
		}

		for {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.Canceled {
						return
					}
				}

				s.es.Logger().Sugar().Errorf("KeepAlive.Recv(): %v", err)
				return
			}

			s.es.Logger().Sugar().Infof("keep alive reply: %+v", reply)
		}
	}
}

func (s *NodeUpService) IsLinked() bool {
	return s.es.GetStatus().GetNodeLink() == consts.ON
}

func (s *NodeUpService) link(value bool) {
	if value {
		s.es.GetStatus().SetNodeLink(consts.ON)
	} else {
		s.es.GetStatus().SetNodeLink(consts.OFF)
	}
}

func (s *NodeUpService) ticker(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	tokenRefreshTicker := time.NewTicker(s.es.dopts.SyncOptions.TokenRefresh)
	defer tokenRefreshTicker.Stop()

	linkStatusTicker := time.NewTicker(s.es.dopts.SyncOptions.Link)
	defer linkStatusTicker.Stop()

	syncTicker := time.NewTicker(s.es.dopts.SyncOptions.Interval)
	defer syncTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.es.Logger().Sugar().Info("ticker: ctx.Done()")
			return
		case <-tokenRefreshTicker.C:
			err := s.login(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("node login: %v", err)
			}
		case <-linkStatusTicker.C:
			err := s.NodeLink(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("link node : %v", err)
			} else {
				s.es.Logger().Sugar().Info("link node ticker success")
			}

		case <-syncTicker.C:
			if err := s.sync(ctx); err != nil {
				s.es.Logger().Sugar().Errorf("sync: %v", err)
			}
		}
	}
}

func (s *NodeUpService) SyncServiceClient() nodes.SyncServiceClient {
	return nodes.NewSyncServiceClient(s.NodeConn)
}

func (s *NodeUpService) NodeServiceClient() nodes.NodeServiceClient {
	return nodes.NewNodeServiceClient(s.NodeConn)
}

func (s *NodeUpService) SlotServiceClient() nodes.SlotServiceClient {
	return nodes.NewSlotServiceClient(s.NodeConn)
}

func (s *NodeUpService) SourceServiceClient() nodes.SourceServiceClient {
	return nodes.NewSourceServiceClient(s.NodeConn)
}

func (s *NodeUpService) TagServiceClient() nodes.TagServiceClient {
	return nodes.NewTagServiceClient(s.NodeConn)
}

func (s *NodeUpService) ConstServiceClient() nodes.ConstServiceClient {
	return nodes.NewConstServiceClient(s.NodeConn)
}

func (s *NodeUpService) GetToken() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.token
}

func (s *NodeUpService) SetToken(ctx context.Context) context.Context {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return metadata.SetToken(ctx, s.token)
}

func (s *NodeUpService) NodeLink(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	request := &nodes.NodeLinkRequest{Status: consts.ON}
	_, err := s.NodeServiceClient().Link(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *NodeUpService) login(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var err error

	request := &nodes.NodeLoginRequest{
		Id:     s.es.dopts.nodeID,
		Secret: s.es.dopts.secret,
	}

	// try login
	reply, err := s.NodeServiceClient().Login(ctx, request)
	if err != nil {
		return err
	}

	if len(reply.GetToken()) == 0 {
		return errors.New("login: reply token is empty")
	}

	// set token
	s.lock.Lock()
	s.token = reply.GetToken()
	s.lock.Unlock()

	return nil
}

func (s *NodeUpService) waitRemoteNodeUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitNodeUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitNodeUpdated: %v", err)

			// retry
			time.Sleep(s.es.dopts.SyncOptions.Retry)
			continue
		}

		for {
			_, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}

				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.Canceled {
						return
					}
				}

				s.es.Logger().Sugar().Errorf("WaitNodeUpdated.Recv(): %v", err)
				return
			}

			ctx = metadata.SetToken(ctx, s.GetToken())

			err = s.sync1(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("sync1: %v", err)
			}
		}
	}
}

func (s *NodeUpService) waitLocalNodeUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY)
	defer notify.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case <-notify.Wait():
			ctx = metadata.SetToken(ctx, s.GetToken())

			err := s.sync2(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("sync2: %v", err)
			}
		}
	}
}

func (s *NodeUpService) waitRemoteTagValueUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitTagValueUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitTagValueUpdated: %v", err)

			// retry
			time.Sleep(s.es.dopts.SyncOptions.Retry)
			continue
		}

		for {
			_, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.Canceled {
						return
					}
				}

				s.es.Logger().Sugar().Errorf("WaitTagValueUpdated.Recv(): %v", err)
				return
			}

			ctx = metadata.SetToken(ctx, s.GetToken())

			err = s.syncTagValue1(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagValue1: %v", err)
			}
		}
	}
}

func (s *NodeUpService) waitLocalTagValueUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY_TV)
	defer notify.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case <-notify.Wait():
			ctx = metadata.SetToken(ctx, s.GetToken())

			err := s.syncTagValue2(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagValue2: %v", err)
			}
		}
	}
}

func (s *NodeUpService) waitRemoteTagWriteUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	for {
		ctx := metadata.SetToken(ctx, s.GetToken())

		stream, err := s.SyncServiceClient().WaitTagWriteUpdated(ctx, &pb.MyEmpty{})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.Canceled {
					return
				}
			}

			s.es.Logger().Sugar().Errorf("WaitTagWriteUpdated: %v", err)

			// retry
			time.Sleep(s.es.dopts.SyncOptions.Retry)
			continue
		}

		for {
			_, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				if code, ok := status.FromError(err); ok {
					if code.Code() == codes.Canceled {
						return
					}
				}

				s.es.Logger().Sugar().Errorf("WaitTagWriteUpdated.Recv(): %v", err)
				return
			}

			ctx = metadata.SetToken(ctx, s.GetToken())

			err = s.syncTagWrite1(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagWrite1: %v", err)
			}
		}
	}
}

func (s *NodeUpService) waitLocalTagWriteUpdated(ctx context.Context) {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	notify := s.es.GetSync().Notify(NOTIFY_TW)
	defer notify.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case <-notify.Wait():
			ctx = metadata.SetToken(ctx, s.GetToken())

			err := s.syncTagWrite2(ctx)
			if err != nil {
				s.es.Logger().Sugar().Errorf("syncTagWrite2: %v", err)
			}
		}
	}
}

func (s *NodeUpService) sync(ctx context.Context) error {
	ctx = metadata.SetToken(ctx, s.GetToken())

	if err := s.sync1(ctx); err != nil {
		return err
	}

	if err := s.sync2(ctx); err != nil {
		return err
	}

	if err := s.syncTagValue1(ctx); err != nil {
		return err
	}

	if err := s.syncTagValue2(ctx); err != nil {
		return err
	}

	if err := s.syncTagWrite1(ctx); err != nil {
		return err
	}

	return s.syncTagWrite2(ctx)
}

func (s *NodeUpService) sync1(ctx context.Context) error {
	return s.syncRemoteToLocal(ctx)
}

func (s *NodeUpService) sync2(ctx context.Context) error {
	return s.syncLocalToRemote(ctx)
}

func (s *NodeUpService) syncTagValue1(ctx context.Context) error {
	return s.syncTagValueRemoteToLocal(ctx)
}

func (s *NodeUpService) syncTagValue2(ctx context.Context) error {
	return s.syncTagValueLocalToRemote(ctx)
}

func (s *NodeUpService) syncTagWrite1(ctx context.Context) error {
	return s.syncTagWriteRemoteToLocal(ctx)
}

func (s *NodeUpService) syncTagWrite2(ctx context.Context) error {
	return s.syncTagWriteLocalToRemote(ctx)
}

func (s *NodeUpService) syncRemoteToLocal(ctx context.Context) error {
	nodeUpdated, err := s.SyncServiceClient().GetNodeUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	nodeUpdated2, err := s.es.GetSync().getNodeUpdatedRemoteToLocal(ctx)
	if err != nil {
		return err
	}

	if nodeUpdated.GetUpdated() <= nodeUpdated2.UnixMicro() {
		return nil
	}

	// node
	{
		remote, err := s.NodeServiceClient().ViewWithDeleted(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		_, err = s.es.GetNode().Sync(ctx, remote)
		if err != nil {
			return err
		}
	}

	// slot
	{
		after := nodeUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.SlotServiceClient().Pull(ctx, &nodes.SlotPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSlot() {
				_, err := s.es.GetSlot().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetSlot()) < int(limit) {
				break
			}
		}
	}

	// source
	{
		after := nodeUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.SourceServiceClient().Pull(ctx, &nodes.SourcePullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetSource() {
				_, err = s.es.GetSource().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetSource()) < int(limit) {
				break
			}
		}
	}

	// tag
	{
		after := nodeUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.TagServiceClient().Pull(ctx, &nodes.TagPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetTag() {
				_, err := s.es.GetTag().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetTag()) < int(limit) {
				break
			}
		}
	}

	// const
	{
		after := nodeUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			remotes, err := s.ConstServiceClient().Pull(ctx, &nodes.ConstPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, remote := range remotes.GetConst() {
				_, err := s.es.GetConst().Sync(ctx, remote)
				if err != nil {
					return err
				}

				after = remote.GetUpdated()
			}

			if len(remotes.GetConst()) < int(limit) {
				break
			}
		}
	}

	return s.es.GetSync().setNodeUpdatedRemoteToLocal(ctx, time.UnixMicro(nodeUpdated.GetUpdated()))
}

func (s *NodeUpService) syncLocalToRemote(ctx context.Context) error {
	nodeUpdated, err := s.es.GetSync().getNodeUpdated(ctx)
	if err != nil {
		return err
	}

	nodeUpdated2, err := s.es.GetSync().getNodeUpdatedLocalToRemote(ctx)
	if err != nil {
		return err
	}

	if nodeUpdated.UnixMicro() <= nodeUpdated2.UnixMicro() {
		return nil
	}

	// node
	{
		local, err := s.es.GetNode().ViewWithDeleted(ctx, &pb.MyEmpty{})
		if err != nil {
			return err
		}

		_, err = s.NodeServiceClient().Sync(ctx, local)
		if err != nil {
			return err
		}
	}

	// slot
	{
		after := nodeUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetSlot().Pull(ctx, &edges.SlotPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetSlot() {
				_, err = s.SlotServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetSlot()) < int(limit) {
				break
			}
		}
	}

	// source
	{
		after := nodeUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetSource().Pull(ctx, &edges.SourcePullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetSource() {
				_, err = s.SourceServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetSource()) < int(limit) {
				break
			}
		}
	}

	// tag
	{
		after := nodeUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetTag().Pull(ctx, &edges.TagPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetTag() {
				_, err = s.TagServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetTag()) < int(limit) {
				break
			}
		}
	}

	// const
	{
		after := nodeUpdated2.UnixMicro()
		limit := uint32(10)

		for {
			locals, err := s.es.GetConst().Pull(ctx, &edges.ConstPullRequest{After: after, Limit: limit})
			if err != nil {
				return err
			}

			for _, local := range locals.GetConst() {
				_, err = s.ConstServiceClient().Sync(ctx, local)
				if err != nil {
					return err
				}

				after = local.GetUpdated()
			}

			if len(locals.GetConst()) < int(limit) {
				break
			}
		}
	}

	return s.es.GetSync().setNodeUpdatedLocalToRemote(ctx, nodeUpdated)
}

func (s *NodeUpService) syncTagValueRemoteToLocal(ctx context.Context) error {
	tagValueUpdated, err := s.SyncServiceClient().GetTagValueUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	tagValueUpdated2, err := s.es.GetSync().getTagValueUpdatedRemoteToLocal(ctx)
	if err != nil {
		return err
	}

	if tagValueUpdated.GetUpdated() <= tagValueUpdated2.UnixMicro() {
		return nil
	}

	after := tagValueUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		remotes, err := s.TagServiceClient().PullValue(ctx, &nodes.TagPullValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, remote := range remotes.GetTag() {
			if remote.GetUpdated() > tagValueUpdated.GetUpdated() {
				break PULL
			}

			_, err = s.es.GetTag().SyncValue(ctx,
				&pb.TagValue{Id: remote.GetId(), Value: remote.GetValue(), Updated: remote.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncValue: %v", err)
				return err
			}

			after = remote.GetUpdated()
		}

		if len(remotes.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setTagValueUpdatedRemoteToLocal(ctx, time.UnixMicro(tagValueUpdated.GetUpdated()))
}

func (s *NodeUpService) syncTagValueLocalToRemote(ctx context.Context) error {
	tagValueUpdated, err := s.es.GetSync().getTagValueUpdated(ctx)
	if err != nil {
		return err
	}

	tagValueUpdated2, err := s.es.GetSync().getTagValueUpdatedLocalToRemote(ctx)
	if err != nil {
		return err
	}

	if tagValueUpdated.UnixMicro() <= tagValueUpdated2.UnixMicro() {
		return nil
	}

	after := tagValueUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		locals, err := s.es.GetTag().PullValue(ctx, &edges.TagPullValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, local := range locals.GetTag() {
			if local.GetUpdated() > tagValueUpdated.UnixMicro() {
				break PULL
			}

			_, err = s.TagServiceClient().SyncValue(ctx,
				&pb.TagValue{Id: local.GetId(), Value: local.GetValue(), Updated: local.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncValue: %v", err)
				return err
			}

			after = local.GetUpdated()
		}

		if len(locals.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setTagValueUpdatedLocalToRemote(ctx, tagValueUpdated)
}

func (s *NodeUpService) syncTagWriteRemoteToLocal(ctx context.Context) error {
	tagWriteUpdated, err := s.SyncServiceClient().GetTagWriteUpdated(ctx, &pb.MyEmpty{})
	if err != nil {
		return err
	}

	tagWriteUpdated2, err := s.es.GetSync().getTagWriteUpdatedRemoteToLocal(ctx)
	if err != nil {
		return err
	}

	if tagWriteUpdated.GetUpdated() <= tagWriteUpdated2.UnixMicro() {
		return nil
	}

	after := tagWriteUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		remotes, err := s.TagServiceClient().PullWrite(ctx, &nodes.TagPullValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, remote := range remotes.GetTag() {
			if remote.GetUpdated() > tagWriteUpdated.GetUpdated() {
				break PULL
			}

			_, err = s.es.GetTag().SyncWrite(ctx,
				&pb.TagValue{Id: remote.GetId(), Value: remote.GetValue(), Updated: remote.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncWrite: %v", err)
				return err
			}

			after = remote.GetUpdated()
		}

		if len(remotes.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setTagWriteUpdatedRemoteToLocal(ctx, time.UnixMicro(tagWriteUpdated.GetUpdated()))
}

func (s *NodeUpService) syncTagWriteLocalToRemote(ctx context.Context) error {
	tagWriteUpdated, err := s.es.GetSync().getTagWriteUpdated(ctx)
	if err != nil {
		return err
	}

	tagWriteUpdated2, err := s.es.GetSync().getTagWriteUpdatedLocalToRemote(ctx)
	if err != nil {
		return err
	}

	if tagWriteUpdated.UnixMicro() <= tagWriteUpdated2.UnixMicro() {
		return nil
	}

	after := tagWriteUpdated2.UnixMicro()
	limit := uint32(100)

PULL:
	for {
		locals, err := s.es.GetTag().PullWrite(ctx, &edges.TagPullValueRequest{After: after, Limit: limit})
		if err != nil {
			return err
		}

		for _, local := range locals.GetTag() {
			if local.GetUpdated() > tagWriteUpdated.UnixMicro() {
				break PULL
			}

			_, err = s.TagServiceClient().SyncWrite(ctx,
				&pb.TagValue{Id: local.GetId(), Value: local.GetValue(), Updated: local.GetUpdated()})
			if err != nil {
				s.es.Logger().Sugar().Errorf("SyncWrite: %v", err)
				return err
			}

			after = local.GetUpdated()
		}

		if len(locals.GetTag()) < int(limit) {
			break
		}
	}

	return s.es.GetSync().setTagWriteUpdatedLocalToRemote(ctx, tagWriteUpdated)
}

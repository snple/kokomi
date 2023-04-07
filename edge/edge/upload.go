package edge

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/danclive/nson-go"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/edges"
	"github.com/snple/kokomi/pb/nodes"
	"github.com/snple/kokomi/util/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UploadService struct {
	es *EdgeService

	ctx     context.Context
	cancel  func()
	closeWG sync.WaitGroup
}

func NewUploadService(es *EdgeService) (*UploadService, error) {
	ctx, cancel := context.WithCancel(es.Context())

	s := &UploadService{
		es:     es,
		ctx:    ctx,
		cancel: cancel,
	}

	return s, nil
}

func (s *UploadService) Start() {
	s.closeWG.Add(1)
	defer s.closeWG.Done()

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case ts := <-ticker.C:
			err := func(ts time.Time) error {
				ctx := context.Background()

				device, err := s.es.GetDevice().view(ctx)
				if err != nil {
					return err
				}

				if device.Status != consts.ON {
					return nil
				}

				node := s.es.GetNode()
				ctx = metadata.SetToken(ctx, node.GetToken())

				offset := 0
				limit := 100

				for {
					request := edges.ListTagRequest{
						Page: &pb.Page{
							Offset: uint32(offset),
							Limit:  uint32(limit),
						},
					}

					offset += limit

					reply, err := s.es.GetTag().List(ctx, &request)
					if err != nil {
						return err
					}

					{
						array := nson.Array{}

						array.Push(nson.U32(ts.Unix()))
						array.Push(nson.U32(0))

						for _, tag := range reply.GetTag() {
							if tag.Status != consts.ON || tag.Upload != consts.ON {
								continue
							}

							if value := s.es.GetTag().GetTagValueValue(tag.Id); value.IsSome() {
								k, err := nson.MessageIdFromHex(tag.Id)
								if err != nil {
									continue
								}

								array.Push(k)
								array.Push(value.Unwrap().Data)
							}
						}

						if len(array) == 0 {
							continue
						}

						s.es.Logger().Sugar().Debugf("upload tags: %v", len(array)/2-1)

						buffer := new(bytes.Buffer)
						err = array.Encode(buffer)
						if err != nil {
							return err
						}

						_, err = node.DataServiceClient().Upload(ctx, &nodes.DataUploadRequest{
							Id:          "",
							ContentType: 1,
							Content:     buffer.Bytes(),
							Cache:       true,
							Save:        true,
						})

						if err != nil {
							if code, ok := status.FromError(err); ok {
								if code.Code() == codes.Unavailable ||
									code.Code() == codes.Unauthenticated ||
									code.Code() == codes.FailedPrecondition {
									s.es.Logger().Sugar().Debugf("upload tags: %v", err)
								} else {
									return err
								}
							} else {
								return err
							}
						}
					}

					if len(reply.GetTag()) < limit {
						break
					}
				}

				s.es.Logger().Sugar().Debug("upload tags: send success")

				return nil
			}(ts)

			if err != nil {
				s.es.Logger().Sugar().Errorf("upload error: %v", err)
			}
		}
	}
}

func (s *UploadService) Stop() {
	s.cancel()
	s.closeWG.Wait()
}

package node

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/danclive/nson-go"
	"github.com/snple/beacon/consts"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/cores"
	"github.com/snple/beacon/util/datatype"
)

type Conn struct {
	ns *NodeService

	net.Conn

	id string

	ctx     context.Context
	cancel  context.CancelCauseFunc
	closeWG sync.WaitGroup
}

func NewConn(ns *NodeService, conn net.Conn) (*Conn, error) {
	ctx, cancel := context.WithCancelCause(ns.Context())

	c := &Conn{
		ns:   ns,
		Conn: conn,

		ctx:    ctx,
		cancel: cancel,
	}

	if err := c.auth(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Conn) Close() error {
	c.cancel(errors.New("close"))
	c.closeWG.Wait()

	// send close message
	msg := nson.Map{
		"fn": nson.String("close"),
	}
	nson.WriteMap(c.Conn, msg)

	return c.Conn.Close()
}

func (c *Conn) ID() string {
	return c.id
}

func (c *Conn) auth() error {
	// 设置读取超时
	c.Conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	req, err := nson.ReadMap(c.Conn)
	if err != nil {
		return err
	}

	// 清除读取超时
	c.Conn.SetReadDeadline(time.Time{})

	fn, err := req.GetString("fn")
	if err != nil {
		writeError(c.Conn, err)

		return err
	}

	if fn != "auth" {
		writeError(c.Conn, errors.New("invalid request, fn != auth"))

		return errors.New("invalid request")
	}

	nodeId, err := req.GetString("id")
	if err != nil {
		writeError(c.Conn, err)

		return err
	}

	secret, err := req.GetString("secret")
	if err != nil {
		writeError(c.Conn, err)

		return err
	}

	if nodeId == "" {
		writeError(c.Conn, errors.New("invalid request, nodeId is empty"))

		return errors.New("invalid request")
	}

	request := &pb.Id{Id: nodeId}

	reply, err := c.ns.Core().GetNode().View(c.ns.Context(), request)
	if err != nil {
		writeError(c.Conn, err)

		return err
	}

	if reply.GetStatus() != consts.ON {
		writeError(c.Conn, errors.New("invalid request, node is not enable"))

		return errors.New("invalid request")
	}

	if reply.GetSecret() != secret {
		writeError(c.Conn, errors.New("invalid request, secret is not valid"))

		return errors.New("invalid request")
	}

	c.id = nodeId

	resp := nson.Map{
		"fn": nson.String("auth"),
		"ok": nson.I32(0),
	}

	return nson.WriteMap(c.Conn, resp)
}

func (c *Conn) handle() {
	for {
		req, err := nson.ReadMap(c.Conn)
		if err != nil {
			return
		}

		fn, err := req.GetString("fn")
		if err != nil {
			writeError(c.Conn, err)

			return
		}

		switch fn {
		case "close":
			return
		case "ping":
			c.handlePing()
		case "set_value":
			c.handleSetValue(req)
		}
	}
}

func (c *Conn) handlePing() {
	msg := nson.Map{
		"fn": nson.String("pong"),
	}
	nson.WriteMap(c.Conn, msg)
}

func (c *Conn) handleSetValue(req nson.Map) error {
	nameValues, err := req.GetMap("name_value")
	if err != nil {
		writeError(c.Conn, err)

		return nil
	}

	errors := nson.Map{}

	for name, value := range nameValues {
		valueStr, err := datatype.EncodeNsonValue(value)
		if err != nil {
			errors[name] = nson.String(err.Error())
			continue
		}

		_, err = c.ns.Core().GetPin().SetValueByName(c.ctx,
			&cores.PinNameValue{NodeId: c.id, Name: name, Value: valueStr})
		if err != nil {
			errors[name] = nson.String(err.Error())
		}
	}

	if len(errors) > 0 {
		msg := nson.Map{
			"fn":     nson.String("set_value"),
			"ok":     nson.I32(-1),
			"errors": errors,
		}
		return nson.WriteMap(c.Conn, msg)
	}

	msg := nson.Map{
		"fn": nson.String("set_value"),
		"ok": nson.I32(0),
	}
	return nson.WriteMap(c.Conn, msg)
}

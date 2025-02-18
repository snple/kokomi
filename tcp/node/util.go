package node

import (
	"net"

	"github.com/danclive/nson-go"
)

func writeError(conn net.Conn, err error) {
	errmap := nson.Map{
		"fn":  nson.String("error"),
		"msg": nson.String(err.Error()),
	}
	nson.WriteMap(conn, errmap)
}

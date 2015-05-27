package coordinator

import (
	"net"

	"gopkg.in/nelsam/influxdb.v0/protocol"
)

type Handler interface {
	HandleRequest(*protocol.Request, net.Conn) error
}

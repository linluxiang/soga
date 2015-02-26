package network

import (
	"net"
)

type Handler interface {
	OnConnection(conn *net.Conn) error
}

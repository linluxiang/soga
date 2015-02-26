package network

import (
	//"io"
	"net"
)

type IOStream struct {
	connection net.Conn
}

func NewIOStream(connection *net.Conn) *IOStream {
	return &IOStream{*connection}
}

func (this *IOStream) Read(rbyte int) (data []byte, err error) {
	data = make([]byte, rbyte)
	_, err = this.connection.Read(data)
	return
}

func (this *IOStream) Write(buffer []byte) (wbyte int, err error) {
	wbyte, err = this.connection.Write(buffer)
	return
}

func (this *IOStream) Close() error {
	return nil
}

func (this *IOStream) Flush() error {
	return nil
}

func (this *IOStream) IsOpen() bool {
	return true
}

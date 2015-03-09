package network

import (
	"bufio"
	//"errors"
	//"fmt"
	"io"
	"io/ioutil"
	"net"

	//"encoding/binary"
)

type IOStream struct {
	connection net.Conn
	readwriter *bufio.ReadWriter
}

func NewIOStream(connection net.Conn) *IOStream {
	stream := IOStream{
		connection: connection,
		readwriter: bufio.NewReadWriter(bufio.NewReader(connection), bufio.NewWriter(connection)),
	}
	return &stream
}

func (this *IOStream) Write(buffer []byte) (nbyte int, err error) {
	nbyte = 0
	ntbyte := 0
	for nbyte < len(buffer) {
		ntbyte, err = this.connection.Write(buffer[nbyte:])
		nbyte += ntbyte
	}
	return nbyte, err
}

func (this *IOStream) Read(rbyte int) (data []byte, err error) {
	data = make([]byte, rbyte)
	_, err = io.ReadFull(this.readwriter, data)
	return
}

func (this *IOStream) ReadUntil(delimeter string) (data []byte, err error) {
	data, err = this.readwriter.ReadBytes(delimeter[0])
	return
}

func (this *IOStream) ReadUntilClose() (data []byte, err error) {
	data, err = ioutil.ReadAll(this.readwriter)
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

package network

import (
	"bufio"
	//"errors"
	//"fmt"
	"io"
	"io/ioutil"
	"net"

	"encoding/binary"
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

/*
Here I have two kinds of implementation

1. Using an infinite loop to read the byte into a buffer and when you wanna use it, take from this buffer.
	problems:

2. let the codec do the real read work, like read how many bytes or read until the end
	this coded must be determined as a parameter of server

3. when writing: make a linked list if still some bytes need to be write into buffer

4. will it stop if there is not enough read byte read? I prefer block
*/
func (this *IOStream) readLoop() {
	var msgSize int32
	for {
		binary.Read(this.readwriter, binary.BigEndian, &msgSize)
	}

}

func (this *IOStream) writeLoop() {

}

func (this *IOStream) Write(buffer []byte) (nbyte int, err error) {
	nbyte, err = this.connection.Write(buffer)
	if nbyte < len(buffer) {

	}
	return
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

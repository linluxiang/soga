package main

import (
	"fmt"
	"io"
	//"net"

	"github.com/linluxiang/soga/network"
)

type Codec struct {
	stream *network.IOStream
}

func NewCodec(stream *network.IOStream) *Codec {
	return &Codec{stream: stream}
}

func (this *Codec) Encode(v interface{}) (err error) {
	this.stream.Write(v.([]byte))
	return
}

func (this *Codec) Decode() (msg interface{}, err error) {
	data, err := this.stream.ReadUntil("\n")
	fmt.Println("received: ", data)
	return data, err
}

func (this *Codec) Close() error {
	return this.stream.Close()
}

type EchoServer struct {
	server *network.TCPServer
}

func NewEchoServer(name string) *EchoServer {
	tcpServer := network.NewTCPServer(name)
	result := &EchoServer{server: tcpServer}
	tcpServer.SetDelegate(result)
	return result
}

func (this *EchoServer) Listen(port string) error {
	err := this.server.Listen(port)
	return err
}

func (this *EchoServer) HandleStream(stream *network.IOStream) {
	codec := NewCodec(stream)
	for {
		msg, err := codec.Decode()
		if err == io.EOF {
			codec.Close()
			break
		}
		codec.Encode(msg)
	}
}

func main() {
	echoServer := NewEchoServer("test")
	err := echoServer.Listen(":5555")
	if err != nil {
		panic(err)
	}
}

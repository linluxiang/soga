package main

import (
	"fmt"
	//"io"
	//"net"

	"github.com/linluxiang/soga/network"
)

type EchoHanlder struct {
}

func (this *EchoHanlder) HandleStream(stream *network.IOStream) {
	data, _ := stream.ReadUntilClose()
	fmt.Println("received: ", data)
	stream.Write(data)
}

func main() {
	handler := &EchoHanlder{}
	server := network.NewTCPServer("test")
	server.SetDelegate(handler)
	err := server.Listen(":5555")
	if err != nil {
		panic(err)
	}
}

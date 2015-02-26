package main

import (
	"fmt"
	//"io"
	//"net"

	"soga/network"
)

type SimpleHanlder struct {
}

func (this *SimpleHanlder) HandleStream(stream *network.IOStream) {
	for {
		data, _ := stream.Read(10)
		fmt.Println("received: ", data)
		stream.Write(data)
	}
}

func main() {
	handler := &SimpleHanlder{}
	server := network.NewTCPServer("test")
	server.SetDelegate(handler)
	err := server.Listen(":5555")
	if err != nil {
		panic(err)
	}
}

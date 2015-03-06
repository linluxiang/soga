package network

import (
	"net"
)

type ConnectionHandler func(conn *net.Conn) error

type ServerDelegate interface {
	HandleStream(stream *IOStream)
}

type TCPServer struct {
	name     string
	delegate ServerDelegate
	exitchan chan bool
	listener net.Listener
	handler  Handler
}

func NewTCPServer(name string) *TCPServer {
	server := TCPServer{}
	server.name = name
	server.exitchan = make(chan bool)
	return &server
}

func NewTCPServerWithDelegate(name string, delegate ServerDelegate) *TCPServer {
	server := NewTCPServer(name)
	server.SetDelegate(delegate)
	return server
}

func (this *TCPServer) Name() string {
	return this.name
}

func (this *TCPServer) loop() {
	for {
		conn, err := this.listener.Accept()
		if err != nil {
			break
		}
		if this.delegate != nil {
			stream := NewIOStream(conn)
			go this.delegate.HandleStream(stream)
		}
	}
}

func (this *TCPServer) SetDelegate(delegate ServerDelegate) {
	this.delegate = delegate
}

func (this *TCPServer) Listen(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	this.listener = listener
	go this.loop()
	<-this.exitchan
	return nil
}

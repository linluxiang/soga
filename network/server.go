package network

import (
	"net"
)

type ConnectionHandler func(conn *net.Conn) error

type ServerDelegate interface {
	/*
		OnConnection() err
		OnMessage(conn, iostream) err
	*/
	HandleStream(stream *IOStream)
}

type TCPServer struct {
	name     string
	delegate ServerDelegate
	loopchan chan bool
	listener net.Listener
	handler  Handler
}

func NewTCPServer(name string) *TCPServer {
	server := TCPServer{}
	server.name = name
	server.loopchan = make(chan bool)
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
		conn, _ := this.listener.Accept()
		go this.handler.OnConnection(&conn)
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
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		if this.delegate != nil {
			stream := NewIOStream(conn)
			this.delegate.HandleStream(stream)
		}
	}
	return nil
}

func (this *TCPServer) Start() (err error) {
	/*
		listener, err := net.Listen(this.protocol, this.port)
		if err != nil {
			// handle error
			return err
		}
		this.listener = listener
		go this.loop()
		<-this.loopchan
	*/
	return
}

func (this *TCPServer) Stop() (err error) {
	this.loopchan <- true
	return
}

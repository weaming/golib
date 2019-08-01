package client

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	Conn    *websocket.Conn
	Url     url.URL
	Box     chan string
	Done    chan int
	Err     chan error
	Running chan int
}

func NewWSClient(host, path string) *WSClient {
	return &WSClient{
		Url:     url.URL{Scheme: "ws", Host: host, Path: path},
		Box:     make(chan string, 1),
		Done:    make(chan int, 1),
		Err:     make(chan error, 1),
		Running: make(chan int, 1),
	}
}

func (p *WSClient) connect() (err error) {
	log.Printf("connecting to %s", p.Url.String())
	p.Conn, _, err = websocket.DefaultDialer.Dial(p.Url.String(), nil)
	return
}

func (p *WSClient) run() error {
	// connect
	err := p.connect()
	if err != nil {
		return err
	}

	// read
	go func() {
		for {
			typ, message, err := p.Conn.ReadMessage()
			if err != nil {
				p.Err <- err
				return
			}
			if typ == websocket.TextMessage {
				log.Printf("ws recv: %s", message)
			} else {
				log.Printf("ws recv type %v: %v", typ, message)

			}
		}
	}()

	// write
	for {
		select {
		case <-p.Done:
			return nil
		case e := <-p.Err:
			return e
		case x := <-p.Box:
			err := p.Conn.WriteMessage(websocket.TextMessage, []byte(x))
			if err != nil || len(p.Err) > 0 {
				// put back
				p.Box <- x
				return err
			}
		}
	}
}

func (p *WSClient) RunUntilDone() {
	p.Running <- 1
	defer func() { <-p.Running }()
	for {
		err := p.run()
		// done
		if err == nil {
			return
		}
		// else retry
		log.Println(err)
	}
}

func (p *WSClient) Send(msg string) {
	select {
	case <-p.Running:
	default:
		go p.RunUntilDone()
		<-p.Running
	}
	defer func() { <-p.Running }()
	p.Box <- msg
}

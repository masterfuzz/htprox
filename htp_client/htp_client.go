///////
// htp_client listens on a local port and relays data to/from the gateway

package htp_client

import (
	"net"
	//"net/http"
	"io"
	"log"
)

type Client struct {
	endpoint    string
	listenAddr  string
	gatewayAddr string
	listenConn  net.Conn
}

func NewClient(listen string, gateway string, endpoint string) Client {
	c := Client{
		endpoint:    endpoint,
		listenAddr:  listen,
		gatewayAddr: gateway,
		listenConn:  nil}
	return c
}

func (c *Client) Run() {
	// make sure endpoint exists on gateway
	// ...
	log.Printf("Endpoint '%v' on gateway '%v' exists", c.endpoint, c.gatewayAddr)

	// Begin listening
	l, err := net.Listen("tcp", c.listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Printf("Listening on %v", c.listenAddr)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(c net.Conn) {
	log.Print("Connection accepted")
	io.Copy(c, c)
	c.Close()
}

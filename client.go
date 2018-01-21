package main

import (
	//"io"
	"log"
	"net"
	"net/http"
	"net/url"
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

		go c.handle(conn)
	}
}

func (c *Client) handle(conn net.Conn) {
	log.Print("Connection accepted")

	// Open connection on gateway
	res, err := http.PostForm("http://"+c.gatewayAddr+"/open", url.Values{"name": {c.endpoint}})
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Printf("Connection not accepted at gateway: %v", res.Status)
		conn.Close()
		return
	}
	// get id

	log.Print("Connection accepted at gateway. ID: 0")
	req, err := http.NewRequest("PUT", "http://"+c.gatewayAddr+"/send", conn)
	http.DefaultClient.Do(req)

	conn.Close()
	http.PostForm("http://"+c.gatewayAddr+"/close", url.Values{"id": {"0"}})
}

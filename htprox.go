package main

import (
	"flag"
	"log"
)

var (
	isServer  = flag.Bool("server", false, "Start Server")
	isClient  = flag.Bool("client", false, "Start Client")
	isGateway = flag.Bool("gateway", false, "Start Gateway")
	listen    = flag.String("listen", ":80", "Incomming port")
	gateAddr  = flag.String("gate-addr", ":80", "Gateway address (for client or server)")
	endpoint  = flag.String("endpoint", "default", "Endpoint to reach/register on gateway")
)

func main() {
	flag.Parse()
	log.Print("Hello")

	if *isServer {
		startServer()
	} else if *isClient {
		startClient()
	} else {
		startGateway()
	}
}

func startClient() {
	log.Print("Starting Client session")
	client := NewClient(*listen, *gateAddr, *endpoint)
	client.Run()
}

func startServer() {
	log.Print("Starting Server")
	server := NewServer(*listen, *gateAddr, *endpoint)
	server.Run()
}

func startGateway() {
	log.Print("Starting Gateway")
	gateway := NewGateway(*listen)
	gateway.Run()
}

package main

import (
	"flag"
	"log"
)

var (
	isServer   = flag.Bool("server", false, "Start Server")
	isClient   = flag.Bool("client", false, "Start Client")
	isGateway  = flag.Bool("gateway", false, "Start Gateway")
	localAddr  = flag.String("local", ":9999", "Address to listen on (client/gateway) or to connect to (server)")
	remoteAddr = flag.String("remote", ":9999", "Gateway address (for client or server)")
	endpoint   = flag.String("endpoint", "default", "Endpoint to reach/register on gateway")
	noStatus   = flag.Bool("no-status", false, "Don't server /status page (and don't parse templates)")
)

func main() {
	flag.Parse()

	if *isServer {
		startServer()
	} else if *isClient {
		startClient()
	} else if *isGateway {
		startGateway()
	} else {
		log.Fatal("Must specify one of client, server, or gateway")
	}
}

func startClient() {
	log.Print("Starting Client")
	client := NewClient(*localAddr, *remoteAddr, *endpoint)
	client.Run()
}

func startServer() {
	log.Print("Starting Server")
	server := NewServer(*localAddr, *remoteAddr, *endpoint)
	server.Run()
}

func startGateway() {
	log.Print("Starting Gateway")
	gateway := NewGateway(*localAddr)
	gateway.noStatus = *noStatus
	gateway.Run()
}

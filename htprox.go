package main

import (
    "log"
    "flag"
    htp_client "github.com/masterfuzz/htprox/htp_client"
    htp_server "github.com/masterfuzz/htprox/htp_server"
    htp_gateway "github.com/masterfuzz/htprox/htp_gateway"
)

var (
    isServer = flag.Bool("server", false, "Start HTProx Server")
    isClient = flag.Bool("client", false, "Start HTProx Client")
    isGateway = flag.Bool("gateway", false, "Start HTProx Gateway")
    listen = flag.String("listen", ":80", "Incomming port")
    gateAddr = flag.String("gate-addr", ":80", "Gateway address (for client or server)")
    endpoint = flag.String("endpoint", "default", "Endpoint to reach/register on gateway")
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
    log.Print("Starting HTProx Client session")
    client := htp_client.NewClient(*listen, *gateAddr, *endpoint)
    client.Run()
}

func startServer() {
    log.Print("Starting HTProx Server")
    server := htp_server.NewServer(*listen, *gateAddr, *endpoint)
    server.Run()
}

func startGateway() {
    log.Print("Starting HTProx Gateway")
    gateway := htp_gateway.NewGateway(*listen)
    gateway.Run()
}

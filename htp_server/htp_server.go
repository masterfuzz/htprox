package htp_server

import (
    "net"
)

type Server struct {
    endpoint string
    localAddr string
    gatewayAddr string
    gateConn net.Conn
    localConn net.Conn
}

func NewServer(local string, gateway string, endpoint string) Server {
    c := Server{
        endpoint: endpoint,
        localAddr: local,
        gatewayAddr: gateway,
        gateConn: nil,
        localConn: nil}
    return c
}

func (c Server) Run() {
    // Register with gateway
    // Wait for connections
}


package htp_client

import (
    "net"
)

type Client struct {
    endpoint string
    listenAddr string
    gatewayAddr string
    gateConn net.Conn
    listenConn net.Conn
}

func NewClient(listen string, gateway string, endpoint string) Client {
    c := Client{
        endpoint: endpoint,
        listenAddr: listen,
        gatewayAddr: gateway,
        gateConn: nil,
        listenConn: nil}
    return c
}

func (c Client) Run() {
}


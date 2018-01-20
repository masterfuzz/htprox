package htp_gateway

import (
    "net"
)

type Gateway struct {
    listenAddr string
    listenConn net.Conn
    // endpoints
}

func NewGateway(listen string) Gateway {
    c := Gateway{
        listenAddr: listen,
        listenConn: nil}
    return c
}

func (c Gateway) Run() {
}


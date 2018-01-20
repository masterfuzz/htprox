package htp_server

import (
    "net"
    "net/http"
    "net/url"
    "time"
    "log"
)

type Server struct {
    endpoint string
    localAddr string
    gatewayAddr string
    gateConn net.Conn
    localConn net.Conn
    pollRate int
}

func NewServer(local string, gateway string, endpoint string) Server {
    c := Server{
        endpoint: endpoint,
        localAddr: local,
        gatewayAddr: gateway,
        gateConn: nil,
        localConn: nil,
        pollRate: 15}
    return c
}

func (s *Server) Run() {
    // Register with gateway
    s.register()
    // Wait for connections
    log.Printf("Waiting for connections to %v from %v@%v", s.localAddr, s.gatewayAddr, s.endpoint)
    for {
        log.Print("ping")
        time.Sleep(time.Duration(s.pollRate) * time.Second)
    }
}

func (s *Server) register() {
    res, err := http.PostForm(s.gatewayAddr + "/register", url.Values{"name": {s.endpoint}})
    if err != nil { panic(err) }
    if res.StatusCode == 201 {
        log.Printf("Registered '%v' with gateway", s.endpoint)
    } else {
        log.Printf("Unable to register with gateway. Code: %v", res.Status)
        panic("unable to register")
    }
}

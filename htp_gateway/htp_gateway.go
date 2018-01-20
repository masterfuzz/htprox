package htp_gateway

import (
    "net/http"
    "log"
    "fmt"
)

type Gateway struct {
    listenAddr string
    // endpoints
}

func NewGateway(listen string) Gateway {
    c := Gateway{
        listenAddr: listen}
    return c
}

func (g Gateway) Run() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, "Hello, %q", r.URL.Path)
    })
    http.HandleFunc("/register", g.httpRegister)

    log.Printf("Listening on %v", g.listenAddr)
    log.Fatal(http.ListenAndServe(g.listenAddr, nil))
}

func (g *Gateway) httpRegister(w http.ResponseWriter, r *http.Request) {
    log.Printf("%v: %v", r.Method, r.URL.Path)
    name := r.FormValue("name")

    log.Printf("Register: %v", name)
}

func (g *Gateway) Register(endpoint string) {
}

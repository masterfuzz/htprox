package main

import (
	"log"
	"net/http"
)

type Gateway struct {
	listenAddr string
	ends       map[string]int
}

func NewGateway(listen string) Gateway {
	c := Gateway{
		listenAddr: listen,
		ends:       make(map[string]int)}
	return c
}

func (g Gateway) Run() {
	http.HandleFunc("/register", g.httpRegister)
	http.HandleFunc("/release", g.httpRelease)
	http.HandleFunc("/open", g.httpOpen)
	http.HandleFunc("/close", g.httpClose)
	http.HandleFunc("/recv", g.httpRecv)
	http.HandleFunc("/send", g.httpSend)

	log.Printf("Listening on %v", g.listenAddr)
	log.Fatal(http.ListenAndServe(g.listenAddr, nil))
}

func (g *Gateway) httpRegister(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
	name := r.FormValue("name")

	_, prs := g.ends[name]
	if prs {
		log.Printf("Endpoint %v already registered", name)
		http.Error(w, "ALREADY REGISTERED", 409)
	} else {
		g.ends[name] = 0
		log.Printf("Registered %v", name)
		http.Error(w, "REGISTERED", 201)
	}
}

func (g *Gateway) httpRelease(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
}

func (g *Gateway) httpOpen(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
}

func (g *Gateway) httpClose(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
}
func (g *Gateway) httpRecv(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
}
func (g *Gateway) httpSend(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
}

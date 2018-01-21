package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Gateway struct {
	listenAddr string
	ends       map[string]map[int]Session
	endIDs     map[string]int
	templates  *template.Template
	noStatus   bool
}

type Session struct {
	closed bool
}

func NewGateway(listen string) Gateway {
	c := Gateway{
		listenAddr: listen,
		ends:       make(map[string]map[int]Session),
		endIDs:     make(map[string]int)}
	return c
}

func (g Gateway) Run() {
	http.HandleFunc("/register", g.httpRegister)
	http.HandleFunc("/release", g.httpRelease)
	http.HandleFunc("/open", g.httpOpen)
	http.HandleFunc("/close", g.httpClose)
	http.HandleFunc("/recv", g.httpRecv)
	http.HandleFunc("/send", g.httpSend)

	if g.noStatus {
		log.Print("Not serving /status")
	} else {
		http.HandleFunc("/status", g.httpStatus)

		templates, err := template.ParseFiles("status.tmpl")
		if err != nil {
			log.Fatal("Error parsing template. ", err)
		}
		g.templates = templates
	}

	log.Printf("Listening on %v", g.listenAddr)
	log.Fatal(http.ListenAndServe(g.listenAddr, nil))
}

func (g *Gateway) httpRegister(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
	name := r.FormValue("name")

	_, prs := g.ends[name]
	if prs {
		log.Printf("Endpoint %v already registered", name)
		w.WriteHeader(409)
		fmt.Fprintf(w, "ERROR: Endpoint %v already registered", name)
	} else {
		g.ends[name] = make(map[int]Session)
		g.endIDs[name] = 0
		log.Printf("Registered %v", name)
		w.WriteHeader(201)
		fmt.Fprintf(w, "SUCCESS: Registered %v", name)
	}
}

func (g *Gateway) httpRelease(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
	name := r.FormValue("name")

	_, prs := g.ends[name]
	if prs {
		delete(g.ends, name)
		log.Printf("Endpoint %v deleted", name)
		fmt.Fprintf(w, "SUCCESS: Endpoint %v deleted", name)
	} else {
		log.Printf("Can't delete %v because it does not exist", name)
		w.WriteHeader(404)
		fmt.Fprintf(w, "ERROR: Can't delete %v because it does not exist", name)
	}
}

func (g *Gateway) httpOpen(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
	name := r.FormValue("name")

	end, prs := g.ends[name]
	if prs {
		// create session
		// weird things will happen if this is called concurrently
		g.endIDs[name] += 1
		end[g.endIDs[name]] = NewSession()
		w.WriteHeader(201)
		fmt.Fprintf(w, "%v", g.endIDs[name])
		// w.Write([]byte{byte(g.endIDs[name])})
	} else {
		log.Printf("Can't create session on %v because it doesn't exist", name)
		w.WriteHeader(404)
		fmt.Fprintf(w, "ERROR: Can't create session on %v because it doesn't exist", name)
	}
}

func NewSession() Session {
	s := Session{
		closed: false}
	return s
}

func (g *Gateway) httpClose(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
	name := r.FormValue("name")
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Printf("Invalid ID %v", r.FormValue("id"))
		w.WriteHeader(400)
		fmt.Fprintf(w, "ERROR: Can't close invalid ID '%v'", r.FormValue("id"))
		return
	}

	end, endPrs := g.ends[name]
	if endPrs {
		ses, sesPrs := end[id]
		if sesPrs {
			if ses.closed {
				log.Printf("CLOSE: Session ID %v:%v already closed.", name, id)
				w.WriteHeader(410)
				fmt.Fprintf(w, "ERROR: Session ID %v:%v already closed.", name, id)
			} else {
				fmt.Fprintf(w, "SUCCESS: Session ID %v:%v closed.", name, id)
				ses.closed = true
			}
		} else {
			log.Printf("CLOSE: Session ID %v:%v does not exist", name, id)
			w.WriteHeader(404)
			fmt.Fprintf(w, "ERROR: Session ID %v:%v does not exist", name, id)
		}
	} else {
		log.Printf("CLOSE: Endpoint '%v' does not exist", name)
		w.WriteHeader(404)
		fmt.Fprintf(w, "ERROR: Endpoint '%v' does not exist", name)
	}
}

func (g *Gateway) httpRecv(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
}
func (g *Gateway) httpSend(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
}

func (g *Gateway) httpStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v", r.Method, r.URL.Path)
	status := g.templates.Lookup("status.tmpl")
	status.Execute(w, g.ends)

}

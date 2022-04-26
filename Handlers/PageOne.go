package handlers

import (
	"log"
	"net/http"
)

type PageOne struct {
	l *log.Logger
}

// pageone initiator
func PageOneHandler(l *log.Logger) *PageOne {
	return &PageOne{l}
}

// pageone handler
func (p *PageOne) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// server log
	p.l.Println("PAGEONE :: visited")

	// respond to client
	w.Write([]byte("PAGEONE :: visited"))
}

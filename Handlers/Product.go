package handlers

import (
	"log"
	"net/http"

	models "github.com/capslock-inc/microservices/Models"
)

type Products struct {
	l *log.Logger
}

func ProductHandler(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GET(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GET(w http.ResponseWriter, r *http.Request) {
	lp := models.GetProducts()
	err := lp.ToJson(w)

	if err != nil {
		http.Error(w, "unable to marshal", http.StatusInternalServerError)
	}
}

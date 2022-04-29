package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPost {
		p.POST(w, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("put", r.URL.Path)
		rg := regexp.MustCompile(`//([0-9]+)`)
		g := rg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		idstring := g[0][1]
		id, err := strconv.Atoi(idstring)

		if err != nil {
			http.Error(w, "unable to atoi", http.StatusInternalServerError)
			return
		}

		p.l.Println("got ID :: ", id)

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

func (p *Products) POST(w http.ResponseWriter, r *http.Request) {
	Prod := &models.Product{}
	err := Prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal", http.StatusBadRequest)
	}
	p.l.Printf("product: %#v", Prod)

	models.Addproduct(Prod)

}

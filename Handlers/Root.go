package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Root struct {
	l *log.Logger
}

// root initiator
func NewHello(l *log.Logger) *Root {
	return &Root{l}
}

// root handler
func (root *Root) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	root.l.Println("ROOT :: visited")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "DATA :: %s", body)

}

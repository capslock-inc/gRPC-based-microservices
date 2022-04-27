package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handlers "github.com/capslock-inc/microservices/Handlers"
)

func main() {

	// logger
	l := log.New(os.Stdout, " Playing-around :: ", log.LstdFlags)

	// reference to handlers
	roothandler := handlers.NewHello(l)
	pageonehandler := handlers.PageOneHandler(l)
	producthandler := handlers.ProductHandler(l)

	// initialing servermux
	sm := http.NewServeMux()

	// mapping handlers with servermux
	sm.Handle("/", roothandler)
	sm.Handle("/pageone", pageonehandler)
	sm.Handle("/product", producthandler)

	// server config
	server := &http.Server{

		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// listing
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// channel
	signalchannel := make(chan os.Signal)
	signal.Notify(signalchannel, os.Interrupt)
	signal.Notify(signalchannel, os.Kill)

	sig := <-signalchannel
	l.Println("Recieved terminate :: Shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(tc)
}

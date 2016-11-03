package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hamchapman/http2-barebones-server"
)

func SubscriptionHandler(subscribeResponder http2test.SubscribeResponder, r *http.Request) {
	start := time.Now()
	t := time.NewTicker(1 * time.Second)

	subscriber := subscribeResponder.Succeed()

	sendMessages := true

	for {
		select {
		case <-t.C:
			log.Printf("elapsed %+v\n", time.Since(start))

			if sendMessages {
				e := http2test.NewEvent()
				e.Data = strings.Repeat("a", 1)
				err := subscriber.EventWriter.Send(e)
				if err != nil {
					log.Println("Abort", err)
					return
				}
			}
		}
	}
}

func main() {
	server := http2test.NewServer()
	server.Addr = ":10443"

	server.Sub("/sub", SubscriptionHandler)

	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	log.Println(err)
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received HTTP request from: %s", r.RemoteAddr)
	fmt.Fprint(w, "Greetings from test app!")
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	srv := http.Server{
		Addr: ":8080",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("srv.ListenAndServe() returned error %s", err)
		}
	}()

	log.Println("Application started.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop
	log.Printf("Stopping app on signal: %s", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("srv.Shutdown() returned %s", err)
	}

	log.Println("All done...")
}

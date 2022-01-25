package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("Received HTTP request from: %s", r.RemoteAddr)
	fmt.Fprint(w, "Greetings from test app!")
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	srv := http.Server{
		Addr: ":8080",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info().Msgf("srv.ListenAndServe() returned error %s", err)
		}
	}()

	log.Info().Msg("Application started.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop
	log.Info().Msgf("Stopping app on signal: %s", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Info().Msgf("srv.Shutdown() returned %s", err)
	}

	log.Info().Msg("All done...")
}

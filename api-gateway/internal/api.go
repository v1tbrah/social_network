package internal

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"social_network/api-gateway/config"
)

type API struct {
	server *http.Server
}

// New returns new API.
func New(cfg config.Config) (newAPI *API) {
	newAPI = &API{
		server: &http.Server{
			Addr: cfg.HTTPServHost + ":" + cfg.HTTPServPort,
		},
	}

	newAPI.server.Handler = newAPI.newRouter()

	return newAPI
}

// Run starts the API.
func (a *API) Run() {
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, time.Second*30)

		if err := a.server.Shutdown(shutdownCtx); err != nil {
			log.Fatal().Err(err).Msg("HTTP server shutdown")
		}

		log.Info().Msg("HTTP server gracefully shutdown")
		serverStopCtx()
	}()

	log.Info().Str("Addr", a.server.Addr).Msg("Starting HTTP server")
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("HTTP server ListenAndServe")
	}

	<-serverCtx.Done()
}

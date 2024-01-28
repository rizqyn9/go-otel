package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	ttrace "learn/apps/pkg/telemetry/trace"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Router *mux.Router

	traceProviderCloseFn []ttrace.CloseFunc
}

func CreateServer() *Server {
	srv := &Server{
		Router: mux.NewRouter(),
	}

	return srv
}

func (s *Server) Run(port uint) error {
	ctx := context.Background()

	httpS := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: s.cors().Handler(s.Router),
	}

	log.Info().Msgf("server serving on port %d ", port)

	go func() {
		if err := httpS.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen:%+s\n", err)
		}
	}()

	<-ctx.Done()
	log.Info().Msg("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	err := httpS.Shutdown(ctxShutDown)
	if err != nil {
		log.Fatal().Msgf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	for _, closeFn := range s.traceProviderCloseFn {
		closeFn := closeFn
		go func() {
			err = closeFn(ctxShutDown)
			if err != nil {
				log.Error().Err(err).Msgf("Unable to close trace provider")
			}
		}()
	}

	return err
}

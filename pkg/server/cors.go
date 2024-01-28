package server

import "github.com/rs/cors"

func (s *Server) cors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		MaxAge:             60, // 1 minutes
		AllowCredentials:   true,
		OptionsPassthrough: false,
		Debug:              false,
	})
}

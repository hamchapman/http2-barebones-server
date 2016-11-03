package http2test

import (
	"net/http"

	"github.com/zimbatm/httputil2"
)

type Server struct {
	Muxer
	*httputil2.MiddlewareList
	*http.Server
}

func NewServer() *Server {
	mux := newServeMux()
	ml := &httputil2.MiddlewareList{}
	ml.Handler = mux
	ml.Use(
		httputil2.CleanPathMiddleware(),
		httputil2.GzipMiddleware(-1),
		httputil2.RecoveryMiddleware(httputil2.DefaultCallback),
	)
	server := &http.Server{}
	return &Server{mux, ml, server}
}

func (s *Server) ListenAndServeTLS(certFile string, keyFile string) error {
	s.Server.Handler = s.Chain()
	return s.Server.ListenAndServeTLS(certFile, keyFile)
}

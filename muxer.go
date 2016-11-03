package http2test

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SubHandler func(s SubscribeResponder, req *http.Request)

type Registry interface {
	SubHandler(path string, handler SubHandler)
	Sub(path string, subHandler SubHandler)
}

type Muxer interface {
	http.Handler
	Registry
	AtRoot() Registry
}

type serveMux struct {
	router *httprouter.Router
}

func newServeMux() Muxer {
	return &serveMux{
		router: httprouter.New(),
	}
}

func (s *serveMux) AtRoot() Registry {
	rootMux := new(serveMux)
	*rootMux = *s
	return rootMux
}

func (s *serveMux) SubHandler(path string, handler SubHandler) {
	s.handle("SUB", path, httprouter.Handle(func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		responder := subscribeResponder{
			w: w,
		}

		handler(responder, req)
	}))
}

func mkStreamWriter(w http.ResponseWriter) StreamWriter {
	f, ok := w.(http.Flusher)
	if !ok {
		panic("Expected http.ResponseWriter to implement http.Flusher")
	}
	return streamWriter{w, f}
}

func (s *serveMux) handle(method, path string, handle httprouter.Handle) {
	s.router.Handle(method, path, handle)
}

func (s *serveMux) Sub(path string, h SubHandler) {
	s.SubHandler(path, h)
}

func (s *serveMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

var _ Muxer = new(serveMux)

package http2test

import (
	"github.com/zimbatm/httputil2"
	"net/http"
)

type Subscriber struct {
	EventWriter EventWriter
}

func NewSubscriber(eventWriter EventWriter) *Subscriber {
	subscriber := &Subscriber{
		EventWriter: eventWriter,
	}
	return subscriber
}

type SubscribeResponder interface {
	Succeed() *Subscriber
}

type subscribeResponder struct {
	w http.ResponseWriter
}

const (
	HeaderContentType = httputil2.HeaderContentType
	SubContentType    = "text/vnd.events+json"
)

func (s subscribeResponder) Succeed() *Subscriber {
	s.w.WriteHeader(http.StatusOK)
	s.w.Header().Set(HeaderContentType, SubContentType)

	sw := mkStreamWriter(s.w)
	sw.Flush()

	ew := NewEventWriter(sw)
	return NewSubscriber(ew)
}

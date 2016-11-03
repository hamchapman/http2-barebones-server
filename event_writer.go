package http2test

import (
	"encoding/json"
	"io"
	"net/http"
)

type StreamWriter interface {
	io.Writer
	http.Flusher
}

type streamWriter struct {
	io.Writer
	http.Flusher
}

type EventWriter interface {
	Send(*Event) error
}

type eventWriter struct {
	w StreamWriter
}

func NewEventWriter(w StreamWriter) EventWriter {
	return eventWriter{w}
}

const (
	LF = byte('\n')
)

func (w eventWriter) Send(e *Event) (err error) {
	var data []byte
	if data, err = json.Marshal(e); err != nil {
		return err
	}

	if _, err = w.w.Write(append(data, LF)); err != nil {
		return err
	}

	w.w.Flush()
	return nil
}

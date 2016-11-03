package http2test

type Event struct {
	Data interface{}
}

func NewEvent() *Event {
	return &Event{nil}
}

func (e *Event) typeCode() int {
	return 0
}

func (e *Event) ensureBody() interface{} {
	if e.Data == nil {
		e.Data = make(map[string]interface{})
	}
	return e.Data
}

package src

type EventListener struct {
}

func (l *EventListener) associationClosed() {

}

func NewEventListener() *EventListener {
	return &EventListener{}

}

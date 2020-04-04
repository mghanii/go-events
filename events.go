package events

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type Emitter interface {
	Emit(eventId string, event string)
	AddListener(eventId string, handler func(string, string)) string
	RemoveListener(listenerId string) error
}

func NewEmitter() Emitter {
	return &eventEmitter{listeners: make(map[uuid.UUID]*listener)}
}

type listener struct {
	id      uuid.UUID
	eventId string
	buffer  chan string
	handler func(listenerId string, event string)
}

type eventEmitter struct {
	listeners map[uuid.UUID]*listener
}

func (emitter *eventEmitter) AddListener(eventId string, handler func(string, string)) string {
	listener := listener{
		id:      uuid.New(),
		eventId: eventId,
		handler: handler,
		buffer:  make(chan string),
	}

	go func(id uuid.UUID, ch chan string) {
		for {
			msg := <-ch
			handler(id.String(), msg)
		}
	}(listener.id, listener.buffer)

	log.Printf("Listner added {id:'%v', event:'%v'}\n", listener.id, listener.eventId)
	emitter.listeners[listener.id] = &listener

	return listener.id.String()
}

func (emitter *eventEmitter) RemoveListener(listenerId string) error {

	id, err := uuid.Parse(listenerId)

	if err != nil {
		return errors.New("invalid listener id")
	}

	if l, ok := emitter.listeners[id]; ok {
		delete(emitter.listeners, id)
		log.Printf("Listner removed {id:'%v', event:'%v'}\n", l.id, l.eventId)
		return nil
	}

	return errors.New("listener doesn't exist")
}

func (emitter *eventEmitter) Emit(eventId string, event string) {
	log.Printf("Emiting event: '%v'", eventId)
	for _, listener := range emitter.listeners {
		if listener.eventId == eventId {
			go func(buffer chan string, lislistenerId uuid.UUID) {
				buffer <- event
			}(listener.buffer, listener.id)
		}
	}

}

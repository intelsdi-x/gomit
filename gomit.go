package gomit

import (
	"fmt"
)

type EventController struct {
	Emmitters map[string]*Emitter
}

type Emitter struct {
	Name          string
	Subscriptions []Subscriber
}

type Subscriber func(*Event)

type Event struct {
	Header Header
}

type Header struct {
	Name string
}

func NewEventController() *EventController {
	e := new(EventController)
	e.Emmitters = make(map[string]*Emitter)
	return e
}

func (e *EventController) RegisterEmitter(em *Emitter) {
	e.Emmitters[em.Name] = em
}

func (e *EventController) Subscribe(name string, f Subscriber) {
	if em, ok := e.Emmitters[name]; ok {
		em.Subscriptions = append(em.Subscriptions, f)
	} else {
		panic("No emitter : " + name)
	}
}

func (e *Emitter) FireEvent(event *Event) {
	fmt.Printf(" >>>> Emitter [%s] firing event [%s]\n", e.Name, event.Header.Name)
	for _, f := range e.Subscriptions {
		f(event)
	}
}

func Foo() {

}

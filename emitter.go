package gomit

import (
	"errors"
	"fmt"
	"sync"
)

// Takes registration and unregistration of Handlers and emits
// Events to be handled by the Handlers.
type Emitter struct {
	Handlers map[string]Handler

	// Used to force single writer. Reads are not locked.
	handlerMutex *sync.Mutex
}

// Something that handles the Events emitted by the Emitter.
type Handler interface {
	HandleGomitEvent(Event)
}

// Emits an Event from the Emitter. Takes an EventBody which is used
// to build an Event. Returns number of handlers that
// received the event and error if an error was raised.
func (e *Emitter) Emit(b EventBody) (int, error) {
	// int used to count the number of Handlers fired.
	var i int
	// We build an event struct to contain the Body and generate a Header.
	event := Event{Header: generateHeader(), Body: b}

	// Fire a gorountine for each handler.
	// By design the is no waiting for any Handlers to complete
	// before firing another. Therefore there is also no guarantee
	// that any Handler will predictably fire before another one.
	//
	// Any synchronizing needs to be within the Handler.
	for _, h := range e.Handlers {
		i++
		go h.HandleGomitEvent(event)
	}

	return i, nil
}

// Registers Handler with the Emitter. Takes a string for the unique name(key)
// and the handler that conforms the to Handler interface. The name(key) is used
// to unregister or check if registered.
func (e *Emitter) RegisterHandler(n string, h Handler) error {
	e.lazyLoadHandler()

	if e.IsHandlerRegistered(n) {
		return errors.New(fmt.Sprintf("%s has already been registered", n))
	}
	e.handlerMutex.Lock()
	e.Handlers[n] = h
	e.handlerMutex.Unlock()

	return nil
}

// Unregisters Handler from the Emitter. This is idempotent where if a Handler is
// not registered no error is returned.
func (e *Emitter) UnregisterHandler(n string) error {

	e.handlerMutex.Lock()
	delete(e.Handlers, n)
	e.handlerMutex.Unlock()

	return nil
}

// Returns bool on whether the Handler is registered with this Emitter.
func (e *Emitter) IsHandlerRegistered(n string) bool {
	_, x := e.Handlers[n]
	return x
}

// Return count (int) of Handlers
func (e *Emitter) HandlerCount() int {
	return len(e.Handlers)
}

// Lazy load Handlers slice/mutex for Emitter
func (e *Emitter) lazyLoadHandler() {
	// Lazy loading of Emitter handler mutex
	if e.handlerMutex == nil {
		e.handlerMutex = new(sync.Mutex)
	}

	// Lazy loading of Emitter handler map
	if e.Handlers == nil {
		e.handlerMutex.Lock()
		e.Handlers = make(map[string]Handler)
		e.handlerMutex.Unlock()
	}

}

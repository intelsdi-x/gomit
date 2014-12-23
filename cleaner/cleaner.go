package cleaner

import (
	"fmt"
	"github.com/intelsdilabs/gomit"
	"time"
)

var (
	// Supported Events
	WidgetCleanedEvent = &gomit.Event{Header: gomit.Header{Name: "Cleaner.WidgetCleaned"}}

	// Supported Emitters
	CleaningEventEmitter = &gomit.Emitter{Name: "Cleaner.CleaningEvents"}
)

type Cleaner struct {
	EventControl *gomit.EventController
}

func NewCleaner() *Cleaner {
	c := new(Cleaner)
	c.EventControl = gomit.NewEventController()
	c.EventControl.RegisterEmitter(CleaningEventEmitter)
	return c
}

func (c *Cleaner) Start() {
	time.Sleep(time.Second * 1)
	fmt.Println("Firing cleaning event")
	CleaningEventEmitter.FireEvent(WidgetCleanedEvent)
}

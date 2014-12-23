package scraper

import (
	"fmt"
	"github.com/intelsdilabs/gomit"
	"github.com/intelsdilabs/gomit/cleaner"
	"time"
)

var (
	// Supported Events
	WidgetScrapedEvent = &gomit.Event{Header: gomit.Header{Name: "Scraper.WidgetScraped"}}

	ScrapingEventEmitter = &gomit.Emitter{Name: "Scraper.ScrapingEvents"}
)

type Scraper struct {
	EventControl *gomit.EventController
	Name         string
}

func NewScraper() *Scraper {
	s := new(Scraper)
	s.EventControl = gomit.NewEventController()
	s.EventControl.RegisterEmitter(ScrapingEventEmitter)
	return s
}

func (s *Scraper) Scrape(e *gomit.Event) {
	fmt.Printf(" <<<< Scraper received event [%s]\n", e.Header.Name)
	time.Sleep(time.Second * 1)
	ScrapingEventEmitter.FireEvent(WidgetScrapedEvent)
}

func (s *Scraper) RegisterCleaner(c *cleaner.Cleaner) {
	c.EventControl.Subscribe("Cleaner.CleaningEvents", s.Scrape)
}

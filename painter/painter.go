package painter

import (
	"fmt"
	"github.com/intelsdilabs/gomit"
	"github.com/intelsdilabs/gomit/scraper"
)

type Painter struct {
	EventControl *gomit.EventController
	Name         string
}

func NewPainter() *Painter {
	p := new(Painter)
	p.EventControl = gomit.NewEventController()
	return p
}

func (p *Painter) Paint(e *gomit.Event) {
	fmt.Printf(" <<<< Painter received event [%s]\n", e.Header.Name)
}

func (p *Painter) RegisterScraper(s *scraper.Scraper) {
	s.EventControl.Subscribe("Scraper.ScrapingEvents", p.Paint)
}

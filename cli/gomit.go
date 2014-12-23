package main

import (
	"fmt"
	"github.com/intelsdilabs/gomit"
	"github.com/intelsdilabs/gomit/cleaner"
	"github.com/intelsdilabs/gomit/painter"
	"github.com/intelsdilabs/gomit/scraper"
)

func main() {
	fmt.Println("Demonstrating gomit")
	gomit.Foo()

	// Create new cleaner
	c := cleaner.NewCleaner()

	// Create scraper
	s := scraper.NewScraper()

	// Create painter
	p := painter.NewPainter()

	// Register cleaner in scraper
	s.RegisterCleaner(c)

	// Start cleaner, which will fire 3 cleaning events 1 second apart
	// This will show only the scraper responding
	fmt.Println("\n Starting with only Scraper listening to Cleaner")
	c.Start()
	fmt.Println("")

	// Register Scraper in Painter
	p.RegisterScraper(s)

	// Start cleaner, which will fire 3 cleaning events 1 second apart
	// You will see the scraper and then the painter respond
	fmt.Println("\n Starting with Painter now also listening to Scraper")
	c.Start()
}

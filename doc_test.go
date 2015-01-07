package gomit

import (
	"fmt"
	"time"
)

type Widget struct {
	EventCount int
}

func (w *Widget) HandleGomitEvent(e Event) {
	w.EventCount++
}

type RandomEventBody struct {
}

func (r *RandomEventBody) Namespace() string {
	return "random.event"
}

func Example() {
	emitter := new(Emitter)
	/*
		type Widget struct {
			EventCount int
		}

		func (w *Widget) HandleGomitEvent(e Event) {
			w.EventCount++
		}
	*/
	widget := new(Widget)

	emitter.RegisterHandler("widget1", widget)

	emitter.Emit(new(RandomEventBody))
	emitter.Emit(new(RandomEventBody))
	emitter.Emit(new(RandomEventBody))

	time.Sleep(time.Millisecond * 100)
	fmt.Println(widget.EventCount)
	// Output: 3
}

// Empty but makes the example not print the whole file
func ExampleFoo() {
}

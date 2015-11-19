/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	event_controller := new(EventController)
	/*
		type Widget struct {
			EventCount int
		}

		func (w *Widget) HandleGomitEvent(e Event) {
			w.EventCount++
		}
	*/
	widget := new(Widget)

	event_controller.RegisterHandler("widget1", widget)

	event_controller.Emit(new(RandomEventBody))
	event_controller.Emit(new(RandomEventBody))
	event_controller.Emit(new(RandomEventBody))

	time.Sleep(time.Millisecond * 100)
	fmt.Println(widget.EventCount)
	// Output: 3
}

// Empty but makes the example not print the whole file
func ExampleFoo() {
}

DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.
<!--
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
-->

# GoMit

[![Build Status](https://travis-ci.org/intelsdi-x/gomit.svg?branch=master)](https://travis-ci.org/intelsdi-x/gomit/)

GoMit (short for "go emit") provides facilities for defining, emitting, and handling events within a Go program. It's used [in Snap](https://github.com/intelsdi-x/snap) to simplify event handling. It's core principles are: 
* Speed over abstraction  
* No order guarantees  
* No persistence  

## Using GoMit
With [Go installed](https://golang.org/dl/), you can `go get` it:

```
$ go get -d github.com/intelsdi-x/gomit
```

### Examples

From [gomit_test.go](https://github.com/intelsdi-x/gomit/blob/master/gomit_test.go):

```go
type MockEventBody struct {
}

type MockThing struct {
	LastNamespace string
}

func (m *MockEventBody) Namespace() string {
	return "Mock.Event"
}
//create a function to handle the gomit event
func (m *MockThing) HandleGomitEvent(e Event) {
	m.LastNamespace = e.Namespace()
}

//create an event controller
event_controller := new(EventController)
//add registration to handler
mt := new(MockThing)
event_controller.RegisterHandler("m1", mt)
//emit event
eb := new(MockEventBody)
i, e := event_controller.Emit(eb)
//unregister handler
event_controller.UnregisterHandler("m1")
//check if handler is registered
b := event_controller.IsHandlerRegistered("m1")
```

### Roadmap
GoMit does all we need it to do and we plan to keep it that simple. If you find a bug in your own usage, please let us know through an [Issue](https://github.com/intelsdi-x/gomit/issues/new).

## Maintainers
The maintainers for GoMit are the same as [Snap](https://github.com/intelsdi-x/snap/blob/master/docs/MAINTAINERS.md).

## License
GoMit is an Open Source software released under the Apache 2.0 [License](LICENSE).

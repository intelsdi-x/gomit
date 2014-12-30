package gomit

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

type MockEventBody struct {
}

type MockThing struct {
	LastNamespace string
}

func (m *MockEventBody) Namespace() string {
	return "Mock.Event"
}

func (m *MockThing) HandleGomitEvent(e Event) {
	m.LastNamespace = e.Namespace()
}

func TestEmitter(t *testing.T) {
	Convey("gomit.Emitter", t, func() {

		Convey(".Emit", func() {
			// Ensure that we can silently emit an event when no one is handling.
			// and that the Emit() returns no error and 0 handlers.
			Convey("Emits with no Handlers", func() {
				m := new(Emitter)
				eb := new(MockEventBody)
				i, e := m.Emit(eb)

				So(i, ShouldBeZeroValue)
				So(m.HandlerCount(), ShouldEqual, 0)
				So(e, ShouldBeNil)
			})
			Convey("Emits with one Handlers", func() {
				m := new(Emitter)
				mt := new(MockThing)

				m.RegisterHandler("m1", mt)
				eb := new(MockEventBody)
				i, e := m.Emit(eb)

				So(i, ShouldEqual, 1)
				So(m.HandlerCount(), ShouldEqual, 1)
				So(e, ShouldBeNil)
			})
		})

		Convey(".RegisterHandler", func() {
			Convey("Allows registration of a single Handler", func() {
				m := new(Emitter)
				mt := new(MockThing)
				e := m.RegisterHandler("m1", mt)

				So(m.HandlerCount(), ShouldEqual, 1)
				So(e, ShouldBeNil)
			})

			Convey("Does not allow a Handler to have more than one registration", func() {
				m := new(Emitter)
				mt1 := new(MockThing)
				mt2 := new(MockThing)

				m.RegisterHandler("m1", mt1)
				// Should return error signifying it was already registered.
				e := m.RegisterHandler("m1", mt2)

				So(m.HandlerCount(), ShouldEqual, 1)
				So(e, ShouldNotBeNil)
			})
		})

		Convey(".HandlerCount", func() {
			// Some simple count testing
			Convey("Returns correct count", func() {
				m := new(Emitter)
				mt1 := new(MockThing)
				mt2 := MockThing{}
				mt3 := new(MockThing)
				mt4 := new(MockThing)
				mt5 := new(MockThing)

				e := m.RegisterHandler("m1", mt1)

				So(e, ShouldBeNil)
				So(m.HandlerCount(), ShouldEqual, 1)

				e = m.RegisterHandler("m2", &mt2)
				So(e, ShouldBeNil)
				So(m.HandlerCount(), ShouldEqual, 2)

				e = m.RegisterHandler("m3", mt3)
				So(e, ShouldBeNil)

				e = m.RegisterHandler("m4", mt4)
				So(e, ShouldBeNil)

				e = m.RegisterHandler("m5", mt5)
				So(e, ShouldBeNil)
				So(m.HandlerCount(), ShouldEqual, 5)

			})
		})

		Convey(".IsHandlerRegistered", func() {
			Convey("Returns false for Handler never registered", func() {
				m := new(Emitter)
				b := m.IsHandlerRegistered("MyMock1")

				So(b, ShouldBeFalse)
			})
			Convey("Returns true for a registered Handler", func() {
				m := new(Emitter)
				mt1 := new(MockThing)

				m.RegisterHandler("MyMock1", mt1)
				b := m.IsHandlerRegistered("MyMock1")

				So(b, ShouldBeTrue)
			})
			Convey("Returns false for a registered Handler that was unregistered", func() {
				m := new(Emitter)
				mt1 := new(MockThing)
				m.RegisterHandler("M1", mt1)
				m.UnregisterHandler("M1")
				b := m.IsHandlerRegistered("M1")

				So(b, ShouldBeFalse)
			})
		})

		Convey(".UnregsiterHandler", func() {
			Convey("Unregisters the Handler", func() {
				m := new(Emitter)
				mt1 := new(MockThing)
				m.RegisterHandler("m1", mt1)
				m.UnregisterHandler("m1")
				b := m.IsHandlerRegistered("m1")

				So(b, ShouldBeFalse)
			})
		})

	})
}

func TestHandler(t *testing.T) {
	Convey("gomit.Handler", t, func() {
		Convey("Handler is called with correct event", func() {
			m := new(Emitter)
			mt := new(MockThing)

			m.RegisterHandler("m1", mt)
			eb := new(MockEventBody)

			i, e := m.Emit(eb)
			// We have to pause to let Handlers run.
			time.Sleep(time.Millisecond * 100)

			// One handler called
			So(i, ShouldEqual, 1)
			// One handler registered
			So(m.HandlerCount(), ShouldEqual, 1)
			// MockThing should have Event namespace (handler was called)
			So(mt.LastNamespace, ShouldEqual, eb.Namespace())
			So(e, ShouldBeNil)
		})
		Convey("Should only emit to the first registered Handler", func() {
			m := new(Emitter)
			mt1 := new(MockThing)
			mt2 := new(MockThing)
			eb := new(MockEventBody)

			m.RegisterHandler("m1", mt1)
			// Should return error signifying it was already registered.
			e := m.RegisterHandler("m1", mt2)
			So(e, ShouldNotBeNil)

			i, e := m.Emit(eb)
			// We have to pause to let Handlers run.
			time.Sleep(time.Millisecond * 100)

			// One handler called
			So(i, ShouldEqual, 1)
			So(m.HandlerCount(), ShouldEqual, 1)
			// This was the first one registered and should match
			So(mt1.LastNamespace, ShouldEqual, eb.Namespace())
			// This was the second one (attempted) and should not match
			So(mt2.LastNamespace, ShouldNotEqual, eb.Namespace())
			So(e, ShouldBeNil)
		})
	})
}

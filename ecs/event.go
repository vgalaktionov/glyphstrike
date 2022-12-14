package ecs

// EID is the type of an event
type EID uint

// Event are messages, dispatched by any system, to communicate with background (event)systems.
// Events are marked by a dummy function returning their type.
type Event interface {
	EID() EID
}


// GetEventChannel retrieves a channel by event tag.
func GetEventChannel[E Event](w *World) chan Event {
	return w.events[(*new(E)).EID()]
}

// DispatchEvent sends an Event on the appropriate channel.
func DispatchEvent(w *World, evt Event) {
	go func() {
		w.events[evt.EID()] <- evt
	}()
}
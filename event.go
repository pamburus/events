package events

import "context"

// New returns a new instance of Event.
func New() Event {
	return make(chan struct{})
}

// Event is a convenient wrapper over chan struct{} with improved readability.
type Event chan struct{}

// Set sets event to signaled state.
func (e Event) Set() {
	close(e)
}

// Wait waits until event is signaled with timeout.
func (e Event) Wait(ctx context.Context) error {
	select {
	case <-e:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// WaitForever waits until event is signaled.
func (e Event) WaitForever() {
	<-e
}

// Signaled returns original channel.
func (e Event) Signaled() chan struct{} {
	return e
}

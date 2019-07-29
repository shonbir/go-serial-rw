package event

//SerialReadEvent defines the subject we are intrested in and to be notified over the channel registered
type SerialReadEvent struct {
	listeners []chan []byte
}

// NewEvents create new struct
func NewEvents() *SerialReadEvent {
	return &SerialReadEvent{
		listeners: []chan []byte{},
	}
}

//AddListener register the SerialReadEvent listener
func (e *SerialReadEvent) AddListener(listeningOnCh chan []byte) {
	if e.listeners == nil {
		e.listeners = []chan []byte{}
	}
	e.listeners = append(e.listeners, listeningOnCh)
}

//RemoveListener de-register the SerialReadEvent listener
func (e *SerialReadEvent) RemoveListener(listeningOnCh chan []byte) {
	for i := range e.listeners {
		if e.listeners[i] == listeningOnCh {
			e.listeners = append(e.listeners[:i], e.listeners[i+1:]...)
			break
		}
	}
}

// PublishEventRecieved distrubutes the recieved byte(s) to all listeners
func (e *SerialReadEvent) PublishEventRecieved(data []byte) {
	for _, channel := range e.listeners {
		go func(channel chan []byte) {
			channel <- data
		}(channel)
	}
}

package GOEventBus

// Dispatcher holds a map keyed by event name to corresponding
// handler. The handler is called when that event
// was triggered
type Dispatcher map[string]func(map[string]any)

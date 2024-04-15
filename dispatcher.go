package GOEventBus

type Dispatcher map[string]func(map[string]any) ([]byte, error)

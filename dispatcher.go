package GOEventBus

type Dispatcher map[string]func(*map[string]any) (map[string]interface{}, error)

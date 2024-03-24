package events

import (
	. "types"
)

type Event struct {
	Projection any
	Args       EventArgs
	Status     bool
}

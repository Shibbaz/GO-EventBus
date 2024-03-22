package events

import(
	. "types"
)

type Event struct{
	Projection Projection
	Args EventArgs
}
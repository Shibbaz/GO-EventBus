package events
import(
	. "types"
)

func NewEvent(projection Projection, args EventArgs) *Event{
	return &Event{
		Projection: projection,
		Args: args,
	}
}
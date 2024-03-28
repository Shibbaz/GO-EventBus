package stream

func (data *Stream) Append(node chan Stream) {
	node <- *data
}

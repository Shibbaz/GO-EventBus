package stream

func (data *Stream) Append(node chan Stream) {
	data.Nodes = append(data.Nodes, node)
}

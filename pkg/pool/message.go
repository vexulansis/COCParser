package pool

type Message struct {
	ID     int
	Client string
	Type   string
	Data   any
}

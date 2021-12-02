package sub

type Direction int

const (
	Forward Direction = iota
	Up
	Down
)

type Command struct {
	Direction Direction
	Amount    int64
}

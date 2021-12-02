package sub

type Submarine struct {
	Aim int64
	Pos Pos
}

type Pos struct {
	H int64
	V int64
}

func New() Submarine {
	return Submarine{0, Pos{0, 0}}
}

func (sub *Submarine) Navigate(cmd Command) {
	switch cmd.Direction {
	case Forward:
		sub.forward(cmd.Amount)
	case Up:
		sub.up(cmd.Amount)
	case Down:
		sub.down(cmd.Amount)
	}
}

func (sub *Submarine) forward(amount int64) {
	sub.Pos.H += amount
	sub.Pos.V += amount * sub.Aim
}

func (sub *Submarine) up(amount int64) {
	sub.Aim -= amount
}

func (sub *Submarine) down(amount int64) {
	sub.Aim += amount
}

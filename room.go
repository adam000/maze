package maze

type Room struct {
	x int
	y int

	north *Room
	south *Room
	east  *Room
	west  *Room
}

func (r *Room) ConnectTo(other *Room, direction Direction) {
	switch direction {
	case North:
		r.north = other
		other.south = r
	case South:
		r.south = other
		other.north = r
	case East:
		r.east = other
		other.west = r
	case West:
		r.west = other
		other.east = r
	}
}

func (r *Room) Go(direction Direction) *Room {
	switch direction {
	case North:
		return r.north
	case South:
		return r.south
	case East:
		return r.east
	case West:
		return r.west
	}

	return nil
}

type Location struct {
	direction Direction
	room      *Room
}

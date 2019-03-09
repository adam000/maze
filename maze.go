package maze

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

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

type Maze struct {
	width  int
	length int

	start *Room
	end   *Room
	// Rooms are arranged in SDL screen coordinate configuration
	rooms [][]Room
}

func (m Maze) Save(fileName string) error {
	// TODO
	return nil
}

func (m Maze) Print() {
	// Print header row
	header := "+"
	for i := 0; i < m.width; i++ {
		header += "-+"
	}
	fmt.Println(header)

	for j := 0; j < m.length; j++ {
		fmt.Printf("|")
		for i := 0; i < m.width; i++ {
			if m.start.x == i && m.start.y == j {
				fmt.Printf("*")
			} else if m.end.x == i && m.end.y == j {
				fmt.Printf("x")
			} else {
				fmt.Printf(" ")
			}

			if m.rooms[i][j].east == nil {
				fmt.Printf("|")
			} else {
				fmt.Printf(" ")
			}
		}

		// Now print line between rows
		fmt.Printf("\n+")
		for i := 0; i < m.width; i++ {
			if m.rooms[i][j].south == nil {
				fmt.Printf("-")
			} else {
				fmt.Printf(" ")
			}
			fmt.Printf("+")
		}
		fmt.Println()
	}
}

func New(width, length int) (Maze, error) {
	if width < 3 {
		return Maze{}, fmt.Errorf("Can't create maze with width less than 3")
	}
	if length < 3 {
		return Maze{}, fmt.Errorf("Can't create maze with length less than 3")
	}

	// Generate the maze
	maze := Maze{width: width, length: length}

	// Create all the rooms
	maze.rooms = make([][]Room, width)
	for i := 0; i < width; i++ {
		maze.rooms[i] = make([]Room, length)
		for j := 0; j < length; j++ {
			maze.rooms[i][j] = Room{x: i, y: j}
		}
	}
	maze.start = &maze.rooms[0][0]
	maze.end = &maze.rooms[width-1][length-1]

	maze.rooms[1][1].ConnectTo(&maze.rooms[1][0], North)
	maze.rooms[1][1].ConnectTo(&maze.rooms[1][2], South)

	// Validate: can reach every square from the start
	// Validate: no loops

	return maze, nil
}

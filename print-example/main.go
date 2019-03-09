package main

import "github.com/adam000/maze"

func main() {
	myMaze, err := maze.New(3, 3)
	if err != nil {
		panic(err)
	}

	myMaze.Print()
}

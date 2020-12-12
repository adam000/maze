package maze

type Generator interface {
	Generate(width, length int, params map[string]interface{}) Maze
}

type InterestingTraversalGenerator struct {
}

func (g InterestingTraversalGenerator) Generate(width, length int, params map[string]interface{}) Maze {
	// TODO, move maze to its own package so it looks like maze.New() here
	maze, err := New(width, length)

	visited := make([][]bool, width)
	for i := 0; i < width; i++ {
		visited[i] = make([]bool, length)
	}

	if err != nil {
		// TODO
		panic(err)
	}

	return maze
}

var _ Generator = (*InterestingTraversalGenerator)(nil)

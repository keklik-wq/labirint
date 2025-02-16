package main

import (
	"maze/maze"
)

func main() {
	generatedMaze := maze.NewMaze(25, 25)

	generatedMaze.Print()
	generatedMaze.Solve()
	generatedMaze.Print()
}

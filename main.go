package main

import (
	"maze/maze"
)

func main() {
	// Создание лабиринта 10x10
	generatedMaze := maze.NewMaze(25, 25)
	// Вывод лабиринта в консоль
	generatedMaze.Print()
	generatedMaze.Solve()
	generatedMaze.Print()
}

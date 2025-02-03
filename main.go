package main

import (
	"maze/maze"
)

func main() {
	// Создание лабиринта 10x10
	generatedMaze := maze.NewMaze(9, 9)

	// Вывод лабиринта в консоль
	generatedMaze.Print()
}

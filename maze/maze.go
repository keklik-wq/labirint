package maze

import (
	"fmt"
	"math/rand"
)

type Maze struct {
	Width  int
	Height int
	Grid   [][]int
}

func NewMaze(width, height int) *Maze {
	maze := &Maze{
		Width:  width,
		Height: height,
		Grid:   make([][]int, height),
	}

	for i := range maze.Grid {
		maze.Grid[i] = make([]int, width)
	}

	maze.fillWalls()

	maze.generatePerfectMaze()

	maze.createEnterAndExit()

	return maze
}

func (m *Maze) fillWalls() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			m.Grid[i][j] = 0
		}
	}
}

func (m *Maze) generatePerfectMaze() {

	startX, startY := 1, 1
	m.Grid[startY][startX] = 1

	m.carvePassages(startX, startY)
}

func (m *Maze) carvePassages(x, y int) {
	directions := rand.Perm(4)
	for _, dir := range directions {
		nx, ny := x, y
		switch dir {
		case 0:
			ny = y - 2
		case 1:
			nx = x + 2
		case 2:
			ny = y + 2
		case 3:
			nx = x - 2
		}

		if nx > 0 && nx < m.Width-1 && ny > 0 && ny < m.Height-1 && m.Grid[ny][nx] == 0 {
			m.Grid[ny][nx] = 1
			m.Grid[y+(ny-y)/2][x+(nx-x)/2] = 1

			m.carvePassages(nx, ny)
		}
	}
}

func (m *Maze) Solve() bool {
	startX, startY := 1, m.Height-1

	if m.bfs(startX, startY) {
		return true
	}

	// Рекурсивный поиск пути
	// if m.dfs(startX, startY) {
	// 	return true
	// }

	return false
}

func (m *Maze) dfs(x, y int) bool {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false
	}

	if m.Grid[x][y] != 1 {
		return false
	}

	if x == m.Width-2 && y == 0 {
		m.Grid[x][y] = 2
		return true
	}

	m.Grid[x][y] = 2

	if m.dfs(x+1, y) || m.dfs(x-1, y) || m.dfs(x, y+1) || m.dfs(x, y-1) {
		return true
	}

	m.Grid[x][y] = 1
	return false
}

func (m *Maze) bfs(startX, startY int) bool {
	queue := [][2]int{{startX, startY}}

	prev := make([][][2]int, m.Height)
	for i := range prev {
		prev[i] = make([][2]int, m.Width)
		for j := range prev[i] {
			prev[i][j] = [2]int{-1, -1}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	m.Grid[startX][startY] = 2

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		x, y := current[0], current[1]

		if x == m.Width-2 && y == 0 {
			for x != startX || y != startY {
				m.Grid[x][y] = 2
				x, y = prev[x][y][0], prev[x][y][1]
			}
			m.Grid[startX][startY] = 2
			return true
		}

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			if nx >= 0 && nx < m.Height && ny >= 0 && ny < m.Width && m.Grid[nx][ny] == 1 {
				queue = append(queue, [2]int{nx, ny})
				prev[nx][ny] = [2]int{x, y}
				m.Grid[nx][ny] = 3
			}
		}
	}

	return false
}

func (m *Maze) createEnterAndExit() {
	m.Grid[1][m.Height-1] = 1
	m.Grid[m.Width-2][0] = 1
}

func (m *Maze) Print() {
	for _, row := range m.Grid {
		for _, cell := range row {
			if (cell == 1) || (cell == 3) {
				fmt.Print("  ")
			} else if cell == 2 {
				fmt.Print("**")
			} else {
				fmt.Print("██")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

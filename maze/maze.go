package maze

import (
	"fmt"
	"math/rand"
)

// Maze представляет лабиринт
type Maze struct {
	Width  int     // Ширина лабиринта
	Height int     // Высота лабиринта
	Grid   [][]int // Сетка лабиринта (0 - стена, 1 - путь)
}

// NewMaze создает новый лабиринт с заданными размерами
func NewMaze(width, height int) *Maze {
	maze := &Maze{
		Width:  width,
		Height: height,
		Grid:   make([][]int, height),
	}

	// Инициализация сетки
	for i := range maze.Grid {
		maze.Grid[i] = make([]int, width)
	}

	// Заполнение лабиринта стенами
	maze.fillWalls()

	// Создание идеального лабиринта
	maze.generatePerfectMaze()

	maze.createEnterAndExit()

	return maze
}

// fillWalls заполняет лабиринт стенами
func (m *Maze) fillWalls() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			m.Grid[i][j] = 0 // 0 - стена
		}
	}
}

// generatePerfectMaze создает идеальный лабиринт
func (m *Maze) generatePerfectMaze() {

	// Начальная точка (вход)
	startX, startY := 1, 1
	m.Grid[startY][startX] = 1

	// Рекурсивный алгоритм для создания лабиринта
	m.carvePassages(startX, startY)
}

// carvePassages рекурсивно создает проходы в лабиринте
func (m *Maze) carvePassages(x, y int) {
	// Направления: вверх, вправо, вниз, влево
	directions := rand.Perm(4)
	//m.Print()
	for _, dir := range directions {
		nx, ny := x, y
		switch dir {
		case 0: // Вверх
			ny = y - 2
		case 1: // Вправо
			nx = x + 2
		case 2: // Вниз
			ny = y + 2
		case 3: // Влево
			nx = x - 2
		}
		//fmt.Println()
		//fmt.Println(nx, ny)

		// Проверка, что новая клетка находится в пределах лабиринта
		if nx > 0 && nx < m.Width-1 && ny > 0 && ny < m.Height-1 && m.Grid[ny][nx] == 0 {
			// Убираем стену между текущей и новой клеткой
			m.Grid[ny][nx] = 1
			m.Grid[y+(ny-y)/2][x+(nx-x)/2] = 1

			// Рекурсивно продолжаем создавать проходы
			m.carvePassages(nx, ny)
		}
	}
}

func (m *Maze) Solve() bool {
	// Начальная точка (вход)
	startX, startY := 1, m.Height-1

	// Рекурсивный поиск пути
	if m.dfs(startX, startY) {
		return true
	}

	return false
}

// dfs рекурсивно ищет путь от текущей точки до выхода
func (m *Maze) dfs(x, y int) bool {
	// Проверка, что текущая точка находится в пределах лабиринта
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false
	}

	// Проверка, что текущая точка является путем и не была посещена
	if m.Grid[x][y] != 1 {
		return false
	}

	// Если текущая точка является выходом, возвращаем true
	if x == m.Width-2 && y == 0 {
		m.Grid[x][y] = 2 // Помечаем выход
		return true
	}

	// Помечаем текущую точку как часть пути
	m.Grid[x][y] = 2

	// Рекурсивно ищем путь в четырех направлениях
	if m.dfs(x+1, y) || m.dfs(x-1, y) || m.dfs(x, y+1) || m.dfs(x, y-1) {
		return true
	}

	// Если путь не найден, снимаем пометку
	m.Grid[x][y] = 1
	return false
}

func (m *Maze) createEnterAndExit() {
	m.Grid[1][m.Height-1] = 1
	m.Grid[m.Width-2][0] = 1
}

// Print выводит лабиринт в консоль
func (m *Maze) Print() {
	for _, row := range m.Grid {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("  ") // Путь
			} else if cell == 2 {
				fmt.Print("**")
			} else {
				fmt.Print("██") // Стена
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

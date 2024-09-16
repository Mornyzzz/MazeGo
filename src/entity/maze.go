package entity

import (
	"fmt"
	"math/rand"
	"os"
)

// Структура лабиринта
type Maze struct {
	Rows              int
	Cols              int
	ActiveRow         int
	SetCounter        int
	SetMatrix         [][]int
	RightBorderMatrix [][]int
	LowBorderMatrix   [][]int
}

// Констуктор лабиринта
func NewMaze(rows int, cols int) *Maze {
	maze := &Maze{
		Rows:              rows,
		Cols:              cols,
		ActiveRow:         0,
		SetCounter:        1,
		SetMatrix:         make([][]int, rows), // Инициализация двумерного среза
		RightBorderMatrix: make([][]int, rows),
		LowBorderMatrix:   make([][]int, rows),
	}
	// Инициализация вложенных срезов
	for i := range maze.SetMatrix {
		maze.SetMatrix[i] = make([]int, cols)
		maze.RightBorderMatrix[i] = make([]int, cols)
		maze.LowBorderMatrix[i] = make([]int, cols)
	}
	return maze
}

// Получение количества строк лабиринта
func (m *Maze) RowSize() int {
	return len(m.RightBorderMatrix)
}

// Получение количества столбцов лабиринта
func (m *Maze) ColSize() int {
	return len(m.RightBorderMatrix[0])
}

// Построчная генерация лабиринта с помощью алгоритма Эйлера
func GenerateMaze(rows int, cols int) (*Maze, error) {
	if rows < 1 || cols < 1 {
		return nil, fmt.Errorf("rows and columns must be positive numbers")
	}
	if rows > 50 || cols > 50 {
		return nil, fmt.Errorf("rows and columns must be <= 50")
	}
	maze := NewMaze(rows, cols)
	for i := 0; i < rows; i++ {
		maze.assignUniqueSet()
		maze.addingVerticalWalls()
		maze.addingHorizontalWalls()
		maze.preparatingNewLine()
	}
	maze.addingEndLine()
	maze.writeToFile()
	return maze, nil
}

// Присвоение ячейки множества
func (m *Maze) assignUniqueSet() {
	for j := 0; j < m.Cols; j++ {
		if m.SetMatrix[m.ActiveRow][j] == 0 {
			m.SetMatrix[m.ActiveRow][j] = m.SetCounter
			m.SetCounter++
		}
	}
}

// Добавление вертикальных (правых) стен
func (m *Maze) addingVerticalWalls() {
	for i := 0; i < m.Cols-1; i++ {
		choise := rand.Int() % 2
		if choise == 1 || m.SetMatrix[m.ActiveRow][i] == m.SetMatrix[m.ActiveRow][i+1] {
			m.RightBorderMatrix[m.ActiveRow][i] = 1
		} else {
			m.mergeSet(i)
		}
	}
	m.RightBorderMatrix[m.ActiveRow][m.Cols-1] = 1
}

// Объединение ячеек в одно множество
func (m *Maze) mergeSet(i int) {
	x := m.SetMatrix[m.ActiveRow][i+1]
	for j := 0; j < m.Cols; j++ {
		if m.SetMatrix[m.ActiveRow][j] == x {
			m.SetMatrix[m.ActiveRow][j] = m.SetMatrix[m.ActiveRow][i]
		}
	}
}

// Добавление горизонтальных (нижних) стен
func (m *Maze) addingHorizontalWalls() {
	for i := 0; i < m.Cols; i++ {
		choise := rand.Int() % 2
		check := m.checkedHorizontalWalls(i)
		if choise == 1 && check {
			m.LowBorderMatrix[m.ActiveRow][i] = 1
		}
	}
}

// Проверка, что множества не закрыты нижними стенами
func (m *Maze) checkedHorizontalWalls(index int) bool {
	set := m.SetMatrix[m.ActiveRow][index]
	for i := 0; i < m.Cols; i++ {
		if m.SetMatrix[m.ActiveRow][i] == set && i != index && m.LowBorderMatrix[m.ActiveRow][i] == 0 {
			return true
		}
	}
	return false
}

// Подготовка новой строки лабиринта
func (m *Maze) preparatingNewLine() {
	if m.ActiveRow == m.Rows-1 {
		return
	}
	m.ActiveRow++
	for i := 0; i < m.Cols; i++ {
		if m.LowBorderMatrix[m.ActiveRow-1][i] == 0 {
			m.SetMatrix[m.ActiveRow][i] = m.SetMatrix[m.ActiveRow-1][i]
		} else {
			m.SetMatrix[m.ActiveRow][i] = 0
		}
	}
}

// Добавление последней стены лабиринта
func (m *Maze) addingEndLine() {
	for i := 0; i < m.Cols-1; i++ {
		m.LowBorderMatrix[m.ActiveRow][i] = 1
		if m.SetMatrix[m.ActiveRow][i] != m.SetMatrix[m.ActiveRow][i+1] {
			m.RightBorderMatrix[m.ActiveRow][i] = 0
			m.mergeSet(i)
		}
	}
	m.LowBorderMatrix[m.ActiveRow][m.Cols-1] = 1
	m.RightBorderMatrix[m.ActiveRow][m.Cols-1] = 1
}

// Запись лабиринта в файл
func (m *Maze) writeToFile() {
	// Открываем или создаем файл
	file, err := os.Create("maze.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Записываем данные в файл
	_, err = fmt.Fprintf(file, "%d %d\n", m.Rows, m.Cols)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	// Записываем первую матрицу
	for _, row := range m.RightBorderMatrix {
		for _, value := range row {
			_, err = fmt.Fprintf(file, "%d ", value)
			if err != nil {
				fmt.Println("Ошибка при записи в файл:", err)
				return
			}
		}
		_, err = fmt.Fprintln(file) // Переход на новую строку
	}
	_, err = fmt.Fprintf(file, "\n")
	// Записываем вторую матрицу
	for _, row := range m.LowBorderMatrix {
		for _, value := range row {
			_, err = fmt.Fprintf(file, "%d ", value)
			if err != nil {
				fmt.Println("Ошибка при записи в файл:", err)
				return
			}
		}
		_, err = fmt.Fprintln(file) // Переход на новую строку
	}
	//fmt.Println("Данные успешно записаны в файл maze.txt")
}

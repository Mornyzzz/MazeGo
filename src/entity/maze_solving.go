package entity

import (
	"fmt"
	"time"
)

// Структура решения лабиринта
type MazeSolving struct {
	mazeInfo      *Maze
	SolvingMatrix [][]int
	direction     int
	currPosX      int
	currPosY      int
	endX          int
	endY          int
}

// Получение количества строк лабиринта
func (m *MazeSolving) RowSize() int {
	return len(m.SolvingMatrix)
}

// Получение количества столбцов лабиринта
func (m *MazeSolving) ColSize() int {
	return len(m.SolvingMatrix[0])
}

// Конструктор решения лабиринта
func NewMazeSolving(maze *Maze, startX, startY, endX, endY int) *MazeSolving {
	s := &MazeSolving{
		mazeInfo:      maze,
		SolvingMatrix: make([][]int, maze.RowSize()),
		direction:     0,
		currPosX:      startX,
		currPosY:      startY,
		endX:          endX,
		endY:          endY,
	}
	for i := range s.SolvingMatrix {
		s.SolvingMatrix[i] = make([]int, maze.ColSize())
	}
	s.writeCountAvailablePaths()
	return s
}

// Решение лабиринта
func SolvingMaze(maze *Maze, startX, startY, endX, endY int) (*MazeSolving, error) {
	if maze == nil || maze.Rows < 1 || maze.Cols < 1 {
		return nil, fmt.Errorf("incorrect maze")
	}
	if startX < 0 || startX >= maze.Cols {
		return nil, fmt.Errorf("incorrect startX")
	}
	if startY < 0 || startY >= maze.Rows {
		return nil, fmt.Errorf("incorrect startY")
	}
	if endX < 0 || endX >= maze.Cols {
		return nil, fmt.Errorf("incorrect endX")
	}
	if endY < 0 || endY >= maze.Rows {
		return nil, fmt.Errorf("incorrect endY")
	}
	if endX == startX && endY == startY {
		return nil, fmt.Errorf("start and end - one point")
	}
	s := NewMazeSolving(maze, startX, startY, endX, endY)
	done := make(chan struct{})
	go func() {
		for !s.findResult() {
			for s.hasRightBorder() && !s.hasForwardBorder() && !s.findResult() { //двигаемся по правой стенке
				s.doStep()
			}
			if !s.hasRightBorder() {
				s.direction = (s.direction + 1) % 4 //поворот направо
				s.doStep()
				if !s.hasRightBorder() {
					s.direction = (s.direction + 1) % 4
					s.doStep()
				}
			} else {
				s.direction = (s.direction + 3) % 4 //поворот налево
			}
		}
		close(done)
	}()

	select {
	case <-done:
		s.doStep()
		for i := 0; i < len(s.SolvingMatrix); i++ {
			for j := 0; j < len(s.SolvingMatrix[i]); j++ {
				if s.SolvingMatrix[i][j] != 0 {
					s.SolvingMatrix[i][j] = 1
				}
			}
		}
		s.SolvingMatrix[startY][startX] = 1
		s.SolvingMatrix[endY][endX] = 1
		return s, nil
	case <-time.After(8 * time.Second):
		return nil, fmt.Errorf("not correct maze")
	}
}

// Перемещение по лабиринту на 1 ячейку по направлению
func (s *MazeSolving) doStep() {
	if s.findResult() {
		return
	}
	switch s.direction {
	case 0:
		s.currPosX-- //left
	case 1:
		s.currPosY-- //up
	case 2:
		s.currPosX++ //right
	case 3:
		s.currPosY++ // down
	}
	s.writeCountAvailablePaths()
}

// Запись количества доступных путей из ячейки в эту же ячейку
func (s *MazeSolving) writeCountAvailablePaths() {
	if s.SolvingMatrix[s.currPosY][s.currPosX] == 0 { //записываем в клетку количество доступных путей
		if !s.hasBorder(0) && s.SolvingMatrix[s.currPosY][s.currPosX-1] == 0 {
			s.SolvingMatrix[s.currPosY][s.currPosX]++
		}
		if !s.hasBorder(1) && s.SolvingMatrix[s.currPosY-1][s.currPosX] == 0 {
			s.SolvingMatrix[s.currPosY][s.currPosX]++
		}
		if !s.hasBorder(2) && s.SolvingMatrix[s.currPosY][s.currPosX+1] == 0 {
			s.SolvingMatrix[s.currPosY][s.currPosX]++
		}
		if !s.hasBorder(3) && s.SolvingMatrix[s.currPosY+1][s.currPosX] == 0 {
			s.SolvingMatrix[s.currPosY][s.currPosX]++
		}
	} else {
		s.SolvingMatrix[s.currPosY][s.currPosX]--
	}
}

// Проверка, есть ли слева стена
func (s *MazeSolving) hasLeftBorder() bool {
	return s.hasBorder((s.direction + 3) % 4)
}

// Проверка, есть ли справа стена
func (s *MazeSolving) hasRightBorder() bool {
	return s.hasBorder((s.direction + 1) % 4)
}

// Проверка, есть ли спереди стена
func (s *MazeSolving) hasForwardBorder() bool {
	return s.hasBorder(s.direction)
}

// Проверка, есть ли стена по направлению
func (s *MazeSolving) hasBorder(direction int) bool {
	switch direction {
	case 0: //left
		if s.currPosX == 0 || s.mazeInfo.RightBorderMatrix[s.currPosY][s.currPosX-1] == 1 {
			return true
		}
	case 1: //up
		if s.currPosY == 0 || s.mazeInfo.LowBorderMatrix[s.currPosY-1][s.currPosX] == 1 {
			return true
		}
	case 2: //right
		if s.mazeInfo.RightBorderMatrix[s.currPosY][s.currPosX] == 1 {
			return true
		}
	case 3: //down
		if s.mazeInfo.LowBorderMatrix[s.currPosY][s.currPosX] == 1 {
			return true
		}
	}
	return false
}

// Проверка, завершено ли выполнение лабиринта
func (s *MazeSolving) findResult() bool {
	if s.currPosX == s.endX && s.currPosY == s.endY {
		return true
	}
	return false
}

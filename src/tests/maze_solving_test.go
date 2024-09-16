package main

import (
	"maze/entity"
	"reflect"
	"testing"
)

func MazeExample() *entity.Maze {
	maze := entity.Maze{
		Rows:       10,
		Cols:       10,
		ActiveRow:  0,
		SetCounter: 1,
		SetMatrix:  make([][]int, 0),
		RightBorderMatrix: [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 1},
			{0, 1, 1, 1, 0, 0, 0, 1, 1, 1},
			{1, 0, 1, 0, 0, 1, 1, 1, 1, 1},
			{1, 0, 0, 1, 0, 0, 1, 0, 1, 1},
			{0, 0, 1, 0, 1, 0, 1, 0, 1, 1},
			{1, 0, 0, 0, 0, 1, 1, 0, 1, 1},
			{0, 0, 0, 1, 1, 0, 0, 1, 0, 1},
			{0, 0, 0, 0, 1, 0, 1, 1, 0, 1},
			{1, 0, 0, 0, 1, 1, 1, 0, 0, 1},
			{0, 1, 0, 1, 0, 1, 0, 0, 0, 1},
		},
		LowBorderMatrix: [][]int{
			{0, 1, 0, 0, 0, 1, 1, 1, 0, 0},
			{1, 1, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 0, 1, 0, 0, 1, 0},
			{1, 0, 1, 1, 1, 0, 1, 1, 1, 0},
			{0, 1, 1, 1, 0, 1, 0, 0, 0, 0},
			{1, 1, 0, 1, 0, 0, 1, 0, 1, 1},
			{0, 1, 1, 1, 1, 0, 0, 0, 1, 0},
			{1, 0, 1, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	return &maze
}

func TestSolvingMaze_NilMaze(t *testing.T) {
	_, err := entity.SolvingMaze(nil, 0, 0, 1, 1)
	if err == nil || err.Error() != "incorrect maze" {
		t.Errorf("expected error for nil maze, got %v", err)
	}
}

func TestSolvingMaze_IncorrectStartX(t *testing.T) {
	maze := MazeExample()
	_, err := entity.SolvingMaze(maze, -1, 0, 1, 1)
	if err == nil || err.Error() != "incorrect startX" {
		t.Errorf("expected error for incorrect startX, got %v", err)
	}
}

func TestSolvingMaze_IncorrectStartY(t *testing.T) {
	maze := MazeExample()
	_, err := entity.SolvingMaze(maze, 0, -1, 1, 1)
	if err == nil || err.Error() != "incorrect startY" {
		t.Errorf("expected error for incorrect startY, got %v", err)
	}
}

func TestSolvingMaze_IncorrectEndX(t *testing.T) {
	maze := MazeExample()
	_, err := entity.SolvingMaze(maze, 0, 0, 12, 1)
	if err == nil || err.Error() != "incorrect endX" {
		t.Errorf("expected error for incorrect endX, got %v", err)
	}
}

func TestSolvingMaze_IncorrectEndY(t *testing.T) {
	maze := MazeExample()
	_, err := entity.SolvingMaze(maze, 0, 0, 1, 12)
	if err == nil || err.Error() != "incorrect endY" {
		t.Errorf("expected error for incorrect endY, got %v", err)
	}
}

func TestSolvingMaze_SameStartEnd(t *testing.T) {
	maze := MazeExample()
	_, err := entity.SolvingMaze(maze, 0, 0, 0, 0)
	if err == nil || err.Error() != "start and end - one point" {
		t.Errorf("expected error for same start and end point, got %v", err)
	}
}

func TestSolvingMaze_EmptyMaze(t *testing.T) {
	maze := &entity.Maze{Rows: 0, Cols: 0}
	_, err := entity.SolvingMaze(maze, 0, 0, 0, 0)
	if err == nil || err.Error() != "incorrect maze" {
		t.Errorf("expected error for empty maze, got %v", err)
	}
}

func TestSolvingMaze_NoPath(t *testing.T) {
	maze := MazeExample()
	maze.RightBorderMatrix[9][8] = 1 //блокируем финальную точку
	_, err := entity.SolvingMaze(maze, 0, 0, 9, 9)
	if err == nil {
		t.Error("expected an error for no path, got nil")
	}
}

func TestSolvingMaze_SuccessfulPath_1(t *testing.T) {
	maze := MazeExample()
	result, err := entity.SolvingMaze(maze, 0, 0, 4, 4)
	successPath := [][]int{
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result == nil {
		t.Error("expected a result, got nil")
	}
	if !reflect.DeepEqual(result.SolvingMatrix, successPath) {
		t.Errorf("not right success path, need %v\n, got %v", successPath, result.SolvingMatrix)
	}
}

func TestSolvingMaze_SuccessfulPath_2(t *testing.T) {
	maze := MazeExample()
	result, err := entity.SolvingMaze(maze, 0, 0, 9, 9)
	successPath := [][]int{
		{1, 1, 1, 1, 1, 0, 0, 0, 1, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result == nil {
		t.Error("expected a result, got nil")
	}
	if !reflect.DeepEqual(result.SolvingMatrix, successPath) {
		t.Errorf("not right success path, need %v\n, got %v", successPath, result.SolvingMatrix)
	}
}

func TestSolvingMaze_SuccessfulPath_3(t *testing.T) {
	maze := MazeExample()
	result, err := entity.SolvingMaze(maze, 7, 0, 0, 7)
	successPath := [][]int{
		{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result == nil {
		t.Error("expected a result, got nil")
	}
	if !reflect.DeepEqual(result.SolvingMatrix, successPath) {
		t.Errorf("not right success path, need %v\n, got %v", successPath, result.SolvingMatrix)
	}
}

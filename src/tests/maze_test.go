package main

import (
	"maze/entity"
	"reflect"
	"testing"
)

func TestGenerateMaze_ValidInput(t *testing.T) {
	maze, err := entity.GenerateMaze(10, 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if maze == nil {
		t.Fatal("Expected a maze, got nil")
	}
}

func TestGenerateMaze_ZeroRows(t *testing.T) {
	maze, err := entity.GenerateMaze(0, 10)
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}
	if maze != nil {
		t.Fatal("Expected nil, got maze")
	}
}

func TestGenerateMaze_ZeroCols(t *testing.T) {
	maze, err := entity.GenerateMaze(10, 0)
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}
	if maze != nil {
		t.Fatal("Expected nil, got maze")
	}
}

func TestGenerateMaze_TooManyRows(t *testing.T) {
	maze, err := entity.GenerateMaze(51, 10)
	if err == nil {
		t.Fatal("Expected an error for too many rows, got nil")
	}
	if maze != nil {
		t.Fatal("Expected nil maze, got a maze")
	}
}

func TestGenerateMaze_TooManyCols(t *testing.T) {
	maze, err := entity.GenerateMaze(10, 51)
	if err == nil {
		t.Fatal("Expected an error for too many columns, got nil")
	}
	if maze != nil {
		t.Fatal("Expected nil maze, got a maze")
	}
}

func TestGenerateMaze_ValidBoundaryValues(t *testing.T) {
	maze1, err1 := entity.GenerateMaze(50, 50)
	if err1 != nil {
		t.Fatalf("Expected no error for boundary values, got %v", err1)
	}
	if maze1 == nil {
		t.Fatal("Expected a maze, got nil")
	}

	maze2, err2 := entity.GenerateMaze(1, 1)
	if err2 != nil {
		t.Fatalf("Expected no error for minimum valid values, got %v", err2)
	}
	if maze2 == nil {
		t.Fatal("Expected a maze, got nil")
	}
}

func TestGenerateMaze_CorrectMaze_1(t *testing.T) {
	maze, err := entity.GenerateMaze(4, 4)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if maze == nil {
		t.Fatal("Expected a maze, got nil")
	}
	result1, err1 := entity.SolvingMaze(maze, 0, 0, 3, 3)
	result2, err2 := entity.SolvingMaze(maze, 0, 0, 3, 3)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")
	}
	result1, err1 = entity.SolvingMaze(maze, 2, 2, 0, 0)
	result2, err2 = entity.SolvingMaze(maze, 2, 2, 0, 0)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")
	}
	result1, err1 = entity.SolvingMaze(maze, 3, 1, 1, 3)
	result2, err2 = entity.SolvingMaze(maze, 3, 1, 1, 3)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")
	}
}

func TestGenerateMaze_CorrectMaze_2(t *testing.T) {
	maze, err := entity.GenerateMaze(10, 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if maze == nil {
		t.Fatal("Expected a maze, got nil")
	}
	result1, err1 := entity.SolvingMaze(maze, 0, 0, 9, 9)
	result2, err2 := entity.SolvingMaze(maze, 9, 9, 0, 0)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")
	}
	result1, err1 = entity.SolvingMaze(maze, 9, 5, 8, 8)
	result2, err2 = entity.SolvingMaze(maze, 8, 8, 9, 5)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")

	}
	result1, err1 = entity.SolvingMaze(maze, 3, 3, 0, 9)
	result2, err2 = entity.SolvingMaze(maze, 0, 9, 3, 3)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")
	}
}

func TestGenerateMaze_CorrectMaze_3(t *testing.T) {
	maze, err := entity.GenerateMaze(30, 30)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if maze == nil {
		t.Fatal("Expected a maze, got nil")
	}
	result1, err1 := entity.SolvingMaze(maze, 0, 0, 29, 29)
	result2, err2 := entity.SolvingMaze(maze, 29, 29, 0, 0)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")
	}
	result1, err1 = entity.SolvingMaze(maze, 10, 10, 20, 20)
	result2, err2 = entity.SolvingMaze(maze, 20, 20, 10, 10)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")

	}
	result1, err1 = entity.SolvingMaze(maze, 10, 20, 29, 29)
	result2, err2 = entity.SolvingMaze(maze, 29, 29, 10, 20)
	if err1 != nil || err2 != nil || !reflect.DeepEqual(result1.SolvingMatrix, result2.SolvingMatrix) {
		t.Fatal("not correct maze")
	}
}

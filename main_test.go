package main

import (
	"root/chess"
	"testing"
)

func TestRook(t *testing.T) {
	data := []struct {
		currentX, currentY, nextX, nextY int8
		isValid                          bool
	}{
		{chess.Generator('B'), 3, chess.Generator('B'), 7, true},
		{chess.Generator('F'), 1, chess.Generator('G'), 3, false},
		{chess.Generator('C'), 4, chess.Generator('C'), 7, true},
		{chess.Generator('A'), 1, chess.Generator('B'), 1, true},
		{chess.Generator('E'), 3, chess.Generator('A'), 7, false},
		{chess.Generator('D'), 7, chess.Generator('A'), 7, true},
	}

	for index, d := range data {
		if result := chess.CheckRook(d.currentX, d.currentY, d.nextX, d.nextY); result != d.isValid {
			t.Errorf("Index %d | result %t | expected %t", index, result, d.isValid)
		}
	}
}

func TestGenerator(t *testing.T) {
	letters := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}

	for index, letter := range letters {
		if gen := chess.Generator(letter); gen != int8(index) {
			t.Errorf("Index: %d | result : %d | expected %d", index, gen, index)
		}
	}
}

func TestBishop(t *testing.T) {
	data := []struct {
		currentX, currentY, nextX, nextY int8
		isValid                          bool
	}{
		{chess.Generator('B'), 3, chess.Generator('E'), 6, true},
		{chess.Generator('A'), 1, chess.Generator('G'), 7, true},
		{chess.Generator('C'), 4, chess.Generator('A'), 6, true},
		{chess.Generator('D'), 3, chess.Generator('H'), 7, true},
		{chess.Generator('D'), 3, chess.Generator('A'), 0, true},
		{chess.Generator('D'), 3, chess.Generator('F'), 1, true},
	}

	for index, d := range data {
		if result := chess.CheckBishop(d.currentX, d.currentY, d.nextX, d.nextY); result != d.isValid {
			t.Errorf("Index %d | result %t | expected %t", index, result, d.isValid)
		}
	}
}

func TestKing(t *testing.T) {
	data := []struct {
		currentX, currentY, nextX, nextY int8
		isValid                          bool
	}{
		{chess.Generator('C'), 2, chess.Generator('C'), 3, true},
		{chess.Generator('C'), 2, chess.Generator('G'), 7, false},
		{chess.Generator('B'), 3, chess.Generator('D'), 3, false},
		{chess.Generator('E'), 4, chess.Generator('E'), 5, true},
		{chess.Generator('E'), 4, chess.Generator('E'), 3, true},
		{chess.Generator('E'), 4, chess.Generator('D'), 3, true},
	}

	for index, d := range data {
		if result := chess.CheckKing(d.currentX, d.currentY, d.nextX, d.nextY); result != d.isValid {
			t.Errorf("Index %d | result %t | expected %t", index, result, d.isValid)
		}
	}
}

func TestKnight(t *testing.T) {
	data := []struct {
		currentX, currentY, nextX, nextY int8
		isValid                          bool
	}{
		{chess.Generator('C'), 1, chess.Generator('D'), 3, true},
		{chess.Generator('F'), 3, chess.Generator('E'), 5, true},
		{chess.Generator('F'), 3, chess.Generator('D'), 3, false},
		{chess.Generator('F'), 3, chess.Generator('G'), 5, true},
		{chess.Generator('F'), 3, chess.Generator('D'), 2, true},
		{chess.Generator('F'), 3, chess.Generator('G'), 1, true},
		{chess.Generator('F'), 3, chess.Generator('A'), 5, false},
	}

	for index, d := range data {
		if result := chess.CheckKnight(d.currentX, d.currentY, d.nextX, d.nextY); result != d.isValid {
			t.Errorf("Index %d | result %t | expected %t", index, result, d.isValid)
		}
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chess.CheckKnight(chess.Generator('C'), 1, chess.Generator('D'), 3)
	}
}

func TestAbs(t *testing.T) {
	if result := chess.Abs(-3); result != 3 {
		t.Errorf("expected: 3 | result: %d", result)
	}
}

package main

import "testing"

func TestRook(t *testing.T) {
	data := []struct {
		currentX, currentY, nextX, nextY int8
		isValid                          bool
	}{
		{Generator('B'), 3, Generator('B'), 7, true},
		{Generator('F'), 1, Generator('G'), 3, false},
		{Generator('C'), 4, Generator('C'), 7, true},
		{Generator('A'), 1, Generator('B'), 1, true},
		{Generator('E'), 3, Generator('A'), 7, false},
		{Generator('D'), 7, Generator('A'), 7, true},
	}

	for index, d := range data {
		if result := CheckRook(d.currentX, d.currentY, d.nextX, d.nextY); result != d.isValid {
			t.Errorf("Index %d | result %t | expected %t", index, result, d.isValid)
		}
	}
}

func TestGenerator(t *testing.T) {
	letters := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}

	for index, letter := range letters {
		if gen := Generator(letter); gen != int8(index) {
			t.Errorf("Index: %d | result : %d | expected %d", index, gen, index)
		}
	}
}

func TestBishop(t *testing.T) {
	data := []struct {
		currentX, currentY, nextX, nextY int8
		isValid                          bool
	}{
		{Generator('B'), 3, Generator('E'), 6, true},
		{Generator('A'), 1, Generator('G'), 7, true},
		{Generator('C'), 4, Generator('A'), 6, true},
		{Generator('D'), 3, Generator('H'), 7, true},
		{Generator('D'), 3, Generator('A'), 0, true},
		{Generator('D'), 3, Generator('F'), 1, true},
	}

	for index, d := range data {
		if result := CheckBishop(d.currentX, d.currentY, d.nextX, d.nextY); result != d.isValid {
			t.Errorf("Index %d | result %t | expected %t", index, result, d.isValid)
		}
	}
}

func TestKing(t *testing.T) {
	data := []struct {
		currentX, currentY, nextX, nextY int8
		isValid                          bool
	}{
		{Generator('C'), 2, Generator('C'), 3, true},
		{Generator('C'), 2, Generator('G'), 7, false},
		{Generator('B'), 3, Generator('D'), 3, false},
		{Generator('E'), 4, Generator('E'), 5, true},
		{Generator('E'), 4, Generator('E'), 3, true},
		{Generator('E'), 4, Generator('D'), 3, true},
	}

	for index, d := range data {
		if result := CheckKing(d.currentX, d.currentY, d.nextX, d.nextY); result != d.isValid {
			t.Errorf("Index %d | result %t | expected %t", index, result, d.isValid)
		}
	}
}

func TestKnight(t *testing.T) {
	data := []struct {
		currentX, currentY, nextX, nextY int8
		isValid                          bool
	}{
		{Generator('C'), 1, Generator('D'), 3, true},
		{Generator('F'), 3, Generator('E'), 5, true},
		{Generator('F'), 3, Generator('D'), 3, false},
		{Generator('F'), 3, Generator('G'), 5, true},
		{Generator('F'), 3, Generator('D'), 2, true},
		{Generator('F'), 3, Generator('G'), 1, true},
		{Generator('F'), 3, Generator('A'), 5, false},
	}

	for index, d := range data {
		if result := CheckKnight(d.currentX, d.currentY, d.nextX, d.nextY); result != d.isValid {
			t.Errorf("Index %d | result %t | expected %t", index, result, d.isValid)
		}
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckKnight(Generator('C'), 1, Generator('D'), 3)
	}
}

func TestAbs(t *testing.T) {
	if result := Abs(-3); result != 3 {
		t.Errorf("expected: 3 | result: %d", result)
	}
}

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
	}
}

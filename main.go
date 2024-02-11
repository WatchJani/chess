package main

import (
	"fmt"
)

type CheckFn func(int8, int8, int8, int8) bool

var move map[string]CheckFn = map[string]CheckFn{
	"Rook":   CheckRook,
	"Bishop": CheckBishop,
	"Queen":  CheckQueen,
	"King":   CheckKing,
	"Knight": CheckKnight,
	"Pawn":   CheckPawn,
}

type Fawn struct {
	name     string
	tag      byte
	color    bool
	currentX int8
	currentY int8
}

func NewFawn(name string, tag byte, color bool, currentX, currentY int8) *Fawn {
	return &Fawn{
		name:     name,
		tag:      tag,
		color:    color,
		currentX: currentX,
		currentY: currentY,
	}
}

type Chess struct {
	move  map[string]CheckFn
	table [][]*Fawn //for print
}

func NewChess() *Chess {
	table := [][]*Fawn{
		{NewKing(false, 'A', 7), NewKnight(false, 'B', 7), NewBishop(false, 'C', 7), NewQueen(false, 'D', 7), NewKing(false, 'E', 7), NewBishop(false, 'F', 7), NewKnight(false, 'G', 7), NewKing(false, 'H', 7)},
		{NewPawn(false, 'A', 6), NewPawn(false, 'B', 6), NewPawn(false, 'C', 6), NewPawn(false, 'D', 6), NewPawn(false, 'E', 6), NewPawn(false, 'F', 6), NewPawn(false, 'G', 6), NewPawn(false, 'H', 6)},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{NewPawn(true, 'A', 1), NewPawn(true, 'B', 1), NewPawn(true, 'C', 1), NewPawn(true, 'D', 1), NewPawn(true, 'E', 1), NewPawn(true, 'F', 1), NewPawn(true, 'G', 1), NewPawn(true, 'H', 1)},
		{NewKing(true, 'A', 0), NewKnight(true, 'B', 1), NewBishop(true, 'C', 2), NewQueen(true, 'D', 3), NewKing(true, 'E', 4), NewBishop(true, 'F', 5), NewKnight(true, 'G', 6), NewKing(true, 'H', 7)},
	}

	return &Chess{
		move:  move,
		table: table,
	}
}

func (c Chess) Print() {
	for _, column := range c.table {
		for _, row := range column {
			if row != nil {
				fmt.Print(string(row.tag) + " ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (c *Chess) Move(letter byte, nextY uint8) {
	// if c.fawns(generated(generated), nextY-1) == nil {
	// 	return
	// }
}

func NewRock(color bool, letter byte, currentY int8) *Fawn {
	return NewFawn("Rook", 'R', color, Generator(letter), currentY)
}

func CheckRook(currentX, currentY, nextX, nextY int8) bool {
	return currentX == nextX || currentY == nextY
}

func NewBishop(color bool, letter byte, currentY int8) *Fawn {
	return NewFawn("Bishop", 'B', color, Generator(letter), currentY)
}

func CheckBishop(currentX, currentY, nextX, nextY int8) bool {
	return Abs(nextX+nextY) == Abs(currentX+currentY) || Abs(currentX-currentY) == Abs(nextX-nextY)
}

func NewQueen(color bool, letter byte, currentY int8) *Fawn {
	return NewFawn("Queen", 'Q', color, Generator(letter), currentY)
}

func CheckQueen(currentX, currentY, nextX, nextY int8) bool {
	return CheckRook(currentX, currentY, nextX, nextY) || CheckBishop(currentX, currentY, nextX, nextY)
}

func NewKing(color bool, letter byte, currentY int8) *Fawn {
	return NewFawn("King", 'K', color, Generator(letter), currentY)
}

func CheckKing(currentX, currentY, nextX, nextY int8) bool {
	positionX := Abs(currentX - nextX)
	positionY := Abs(currentY - nextY)

	return positionX <= 1 && positionY <= 1
}

func NewKnight(color bool, letter byte, currentY int8) *Fawn {
	return NewFawn("Knight", 'H', color, Generator(letter), currentY)
}

func CheckKnight(currentX, currentY, nextX, nextY int8) bool {
	positionX := Abs(currentX - nextX)
	positionY := Abs(currentY - nextY)

	return positionX == 2 && positionY == 1 || positionX == 1 && positionY == 2
}

func NewPawn(color bool, letter byte, currentY int8) *Fawn {
	return NewFawn("Pawn", 'P', color, Generator(letter), currentY)
}

func CheckPawn(currentX, currentY, nextX, nextY int8) bool {
	distanceX := Abs(nextX - currentX)
	distanceY := Abs(nextY - currentY)

	return ((currentX == 6 || currentX == 1) && distanceX == 2) || distanceY == 1 || distanceY == 1 && (distanceX == 1 || distanceY == -1)
}

func Abs(num int8) int8 {
	if num < 0 {
		return -num
	}

	return num
}

func Generator(letter byte) int8 {
	return int8(letter - 'A')
}

func main() {
	chess := NewChess()

	chess.Print()
}

package chess

import (
	"errors"
	"fmt"
)

//another logic, we don't need Chess table just need to know where are our Fawn, this option is better if we have real time game, with real frontend
//this implementation is more better because we have option to present

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

func NewFawn(name string, tag byte, color bool) *Fawn {
	return &Fawn{
		name:  name,
		tag:   tag,
		color: color,
	}
}

type Chess struct {
	move  map[string]CheckFn
	table [][]*Fawn //for print
	chess []byte
}

func NewChess() *Chess {
	table := [][]*Fawn{
		{NewKing(false), NewKnight(false), NewBishop(false), NewQueen(false), NewKing(false), NewBishop(false), NewKnight(false), NewKing(false)},
		{NewPawn(false), NewPawn(false), NewPawn(false), NewPawn(false), NewPawn(false), NewPawn(false), NewPawn(false), NewPawn(false)},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{NewPawn(true), NewPawn(true), NewPawn(true), NewPawn(true), NewPawn(true), NewPawn(true), NewPawn(true), NewPawn(true)},
		{NewKing(true), NewKnight(true), NewBishop(true), NewQueen(true), NewKing(true), NewBishop(true), NewKnight(true), NewKing(true)},
	}

	return &Chess{
		move:  move,
		table: table,
		chess: make([]byte, 72), //64 + 8 => \n
	}
}

// can be better, but this is test game web socket
func (c *Chess) Print() []byte {
	for i, column := range c.table {
		for j, row := range column {
			if row != nil {
				// c.chess = append(c.chess, row.tag)
				c.chess[8*i+j] = row.tag
			} else {
				// c.chess = append(c.chess, ' ')
				c.chess[8*i+j] = ' '
			}
		}
	}

	for i := 8; i < 72; i += 9 {
		copy(c.chess[i+1:], c.chess[i:])
		c.chess[i] = '\n'
	}

	return c.chess
}

func (c *Chess) Move(CurrentX, CurrentY, nextX, nextY int8) error {
	if CurrentX > 7 || CurrentY > 7 || nextX > 7 || nextY > 7 {
		return errors.New("This field is not exist!")
	}

	currentFawn := c.table[CurrentX][CurrentY]

	//if not exist
	if currentFawn == nil {
		return errors.New("This field is empty!")
	}

	//check can i make move
	if ok := move[currentFawn.name](CurrentX, CurrentY, nextX, nextY); !ok {
		return errors.New(fmt.Sprintf("Cant move this fawn [%s] on this field [%d, %d]", currentFawn.name, nextX, nextY))
	}

	//if we move next fawn on free field
	if c.table[nextX][nextY] == nil {
		c.table[CurrentX][CurrentY], c.table[nextX][nextY] = c.table[nextX][nextY], c.table[CurrentX][CurrentY] // swap //make move
	} else {
		c.table[CurrentX][CurrentY] = c.table[nextX][nextY]
	}

	return nil
}

func NewRock(color bool) *Fawn {
	return NewFawn("Rook", 'R', color)
}

func CheckRook(currentX, currentY, nextX, nextY int8) bool {
	return currentX == nextX || currentY == nextY
}

func NewBishop(color bool) *Fawn {
	return NewFawn("Bishop", 'B', color)
}

func CheckBishop(currentX, currentY, nextX, nextY int8) bool {
	return Abs(nextX+nextY) == Abs(currentX+currentY) || Abs(currentX-currentY) == Abs(nextX-nextY)
}

func NewQueen(color bool) *Fawn {
	return NewFawn("Queen", 'Q', color)
}

func CheckQueen(currentX, currentY, nextX, nextY int8) bool {
	return CheckRook(currentX, currentY, nextX, nextY) || CheckBishop(currentX, currentY, nextX, nextY)
}

func NewKing(color bool) *Fawn {
	return NewFawn("King", 'K', color)
}

func CheckKing(currentX, currentY, nextX, nextY int8) bool {
	positionX := Abs(currentX - nextX)
	positionY := Abs(currentY - nextY)

	return positionX <= 1 && positionY <= 1
}

func NewKnight(color bool) *Fawn {
	return NewFawn("Knight", 'H', color)
}

func CheckKnight(currentX, currentY, nextX, nextY int8) bool {
	positionX := Abs(currentX - nextX)
	positionY := Abs(currentY - nextY)

	return positionX == 2 && positionY == 1 || positionX == 1 && positionY == 2
}

func NewPawn(color bool) *Fawn {
	return NewFawn("Pawn", 'P', color)
}

func CheckPawn(currentX, currentY, nextX, nextY int8) bool {
	distanceX := Abs(nextX - currentX)
	distanceY := Abs(nextY - currentY)

	return ((currentX == 6 || currentX == 1) && distanceX == 2) || distanceX == 1 || distanceY == 1 && (distanceX == 1 || distanceY == -1)
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

func Parse(coordinates []byte) (int8, int8, int8, int8) {
	return int8(coordinates[1]) - '0' - 1, Generator(coordinates[0]), int8(coordinates[4]) - '0' - 1, Generator(coordinates[3])
}

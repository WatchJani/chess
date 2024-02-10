package main

import (
	"fmt"
)

// ne treba mi tabla
// trebaju mi samo figurice pozicije i nove koordinate
type Fawn struct {
	name  string
	tag   byte
	color bool
	move  func(currentX, currentY, nextX, nextY int8)
}

func Rook(currentX, currentY, nextX, nextY int8) bool {
	return currentX == nextX || currentY == nextY
}

func Bishop(currentX, currentY, nextX, nextY int8) bool {
	return Abs(nextX+nextY) == Abs(currentX+currentY) || Abs(currentX-currentY) == Abs(nextX-nextY)
}

func Queen(currentX, currentY, nextX, nextY int8) bool {
	return Rook(currentX, currentY, nextX, nextY) || Bishop(currentX, currentY, nextX, nextY)
}

func King(currentX, currentY, nextX, nextY int8) bool {
	positionX := Abs(currentX - nextX)
	positionY := Abs(currentY - nextY)

	return positionX <= 1 && positionY <= 1
}

func Knight(currentX, currentY, nextX, nextY int8) bool {
	positionX := Abs(currentX - nextX)
	positionY := Abs(currentY - nextY)

	return positionX == 2 && positionY == 1 || positionX == 1 && positionY == 2
}

func Pawn(currentX, currentY, nextX, nextY int8) bool {
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
	fmt.Println(Generator('H'))
}

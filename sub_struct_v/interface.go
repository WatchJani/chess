package sub_struct_v

type FawnInterface interface {
	Check(int8, int8) bool
	Move(int8, int8)
}

//==================================================================================================

func CheckRook(currentX, currentY, nextX, nextY int8) bool {
	return currentX == nextX || currentY == nextY
}

func CheckBishop(currentX, currentY, nextX, nextY int8) bool {
	return Abs(nextX+nextY) == Abs(currentX+currentY) || Abs(currentX-currentY) == Abs(nextX-nextY)
}

func CheckQueen(currentX, currentY, nextX, nextY int8) bool {
	return CheckRook(currentX, currentY, nextX, nextY) || CheckBishop(currentX, currentY, nextX, nextY)
}

func CheckKing(currentX, currentY, nextX, nextY int8) bool {
	positionX := Abs(currentX - nextX)
	positionY := Abs(currentY - nextY)

	return positionX <= 1 && positionY <= 1
}

func CheckKnight(currentX, currentY, nextX, nextY int8) bool {
	positionX := Abs(currentX - nextX)
	positionY := Abs(currentY - nextY)

	return positionX == 2 && positionY == 1 || positionX == 1 && positionY == 2
}

func CheckPawn(currentX, currentY, nextX, nextY int8) bool {
	distanceX := Abs(nextX - currentX)
	distanceY := Abs(nextY - currentY)

	return ((currentX == 6 || currentX == 1) && distanceX == 2) || distanceY == 1 || distanceY == 1 && (distanceX == 1 || distanceY == -1)
}

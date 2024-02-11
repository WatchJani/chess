package sub_struct_v

type Fawn struct {
	name     string
	tag      byte
	color    bool
	currentX int8
	currentY int8
}

func NewFawn(name string, tag byte, color bool, currentX, currentY int8) Fawn {
	return Fawn{
		name:     name,
		tag:      tag,
		color:    color,
		currentX: currentX,
		currentY: currentY,
	}
}

func (f *Fawn) Move(nextX, nextY int8) {
	f.currentX = nextX
	f.currentY = nextY
}

type Rook struct {
	Fawn
}

func NewRook(name string, tag byte, color bool, currentX, currentY int8) Rook {
	return Rook{
		NewFawn(name, tag, color, currentX, currentY),
	}
}

func (r Rook) Check(nextX, nextY int8) bool {
	return CheckRook(r.currentX, r.currentY, nextX, nextY)
}

//...........

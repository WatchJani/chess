package type_v

type Fawn struct {
	name     string
	tag      byte
	color    bool
	currentX int8
	currentY int8
}

type Rook Fawn

func NewRook(name string, tag byte, color bool, currentX, currentY int8) Rook {
	return Rook{
		name:     name,
		tag:      tag,
		color:    color,
		currentX: currentX,
		currentY: currentY,
	}
}

func (r Rook) Check(nextX, nextY int8) bool {
	return CheckRook(r.currentX, r.currentY, nextX, nextY)
}

func (r *Rook) Move(nextX, nextY int8) {
	r.currentX = nextX
	r.currentY = nextY
}

type Bishop Fawn

func NewBishop(name string, tag byte, color bool, currentX, currentY int8) Bishop {
	return Bishop{
		name:     name,
		tag:      tag,
		color:    color,
		currentX: currentX,
		currentY: currentY,
	}
}

func (b Bishop) Check(nextX, nextY int8) bool {
	return CheckBishop(b.currentX, b.currentY, nextX, nextY)
}

func (b *Bishop) Move(nextX, nextY int8) {
	b.currentX = nextX
	b.currentY = nextY
}

type Queen Fawn

func NewQueen(name string, tag byte, color bool, currentX, currentY int8) Queen {
	return Queen{
		name:     name,
		tag:      tag,
		color:    color,
		currentX: currentX,
		currentY: currentY,
	}
}

func (q Queen) Check(nextX, nextY int8) bool {
	return CheckRook(q.currentX, q.currentY, nextX, nextY) || CheckBishop(q.currentX, q.currentY, nextX, nextY)
}

func (q *Queen) Move(nextX, nextY int8) {
	q.currentX = nextX
	q.currentY = nextY
}

type King Fawn

func NewKing(name string, tag byte, color bool, currentX, currentY int8) King {
	return King{
		name:     name,
		tag:      tag,
		color:    color,
		currentX: currentX,
		currentY: currentY,
	}
}

func (k King) Check(nextX, nextY int8) bool {
	return CheckKing(k.currentX, k.currentY, nextX, nextY)
}

func (k *King) Move(nextX, nextY int8) {
	k.currentX = nextX
	k.currentY = nextY
}

type Knight Fawn

func NewKnight(name string, tag byte, color bool, currentX, currentY int8) Knight {
	return Knight{
		name:     name,
		tag:      tag,
		color:    color,
		currentX: currentX,
		currentY: currentY,
	}
}

func (k Knight) Check(nextX, nextY int8) bool {
	return CheckKnight(k.currentX, k.currentY, nextX, nextY)
}

func (k *Knight) Move(nextX, nextY int8) {
	k.currentX = nextX
	k.currentY = nextY
}

type Pawn Fawn

func NewPawn(name string, tag byte, color bool, currentX, currentY int8) Pawn {
	return Pawn{
		name:     name,
		tag:      tag,
		color:    color,
		currentX: currentX,
		currentY: currentY,
	}
}

func (p Pawn) Check(nextX, nextY int8) bool {
	return CheckPawn(p.currentX, p.currentY, nextX, nextY)
}

func (p *Pawn) Move(nextX, nextY int8) {
	p.currentX = nextX
	p.currentY = nextY
}

package draw

import (
	"fmt"
	"math/rand"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func NewPointFromString(ptStr string) (*Point, error) {
	intLst, err := ParseToInt(ptStr)
	if err != nil {
		return nil, err
	}

	return &Point{intLst[0], intLst[1]}, nil
}

func (p *Point) Equal(pt *Point) bool {
	return p.X == pt.X && p.Y == pt.Y
}

func (p *Point) Offset(x, y int) *Point {
	p.X = p.X + x
	p.Y = p.Y + y

	return p
}

func (p *Point) OffsetSize(s *Size) *Point {
	p.X = p.X + rand.Intn(s.Width)
	p.Y = p.Y + rand.Intn(s.Height)

	return p
}

func (p *Point) Valid() bool {
	return p.X >= 0 && p.Y >= 0
}

func (p *Point) String() string {
	return fmt.Sprintf("Point(x:%d, y:%d)", p.X, p.Y)
}

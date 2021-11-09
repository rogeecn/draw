package draw

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func NewPointFromString(ptStr string) (*Point, error) {
	ptItems := strings.Split(strings.Trim(ptStr, " ,"), ",")
	if len(ptItems) != 2 {
		return nil, errors.New("坐标字符串参数格式错误")
	}
	xStr, yStr := ptItems[0], ptItems[1]

	x, err := strconv.Atoi(xStr)
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		return nil, err
	}

	return &Point{x, y}, nil
}

func (p *Point) Equal(pt *Point) bool {
	return p.X == pt.X && p.Y == pt.Y
}

func (p *Point) Offset(x, y int) {
	p.X = p.X + x
	p.Y = p.Y + y
}

func (p *Point) OffsetRandom(x, y int, s *Size) {
	p.X = p.X + x + rand.Intn(s.Width)
	p.Y = p.Y + y + rand.Intn(s.Height)
}

func (p *Point) Valid() bool {
	return p.X >= 0 && p.Y >= 0
}

func (p *Point) String() string {
	return fmt.Sprintf("Point(x:%d, y:%d)", p.X, p.Y)
}

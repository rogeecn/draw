package draw

import "fmt"

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

func NewRect(s *Size) *Rect {
	return &Rect{0, 0, s.Width, s.Height}
}

func NewRectFromSize(lt *Point, s *Size) *Rect {
	return &Rect{lt.X, lt.Y, s.Width, s.Height}
}

func NewRectFromPoint(lt *Point, rb *Point) *Rect {
	return &Rect{lt.X, lt.Y, rb.X - lt.X, rb.Y - lt.Y}
}

func (r *Rect) Left() int {
	return r.X
}

func (r *Rect) Right() int {
	return r.X + r.Width
}

func (r *Rect) Top() int {
	return r.Y
}

func (r *Rect) Bottom() int {
	return r.Y + r.Width
}

func (r *Rect) TopLeft() *Point {
	return NewPoint(r.X, r.Y)
}

func (r *Rect) TopRight() *Point {
	return NewPoint(r.X+r.Width, r.Y)
}

func (r *Rect) BottomLeft() *Point {
	return NewPoint(r.X, r.Y+r.Height)
}

func (r *Rect) BottomRight() *Point {
	return NewPoint(r.X+r.Width, r.Y+r.Height)
}
func (r *Rect) Center() *Point {
	return NewPoint(r.X+r.Width/2, r.Y+r.Height/2)
}

func (r *Rect) Empty() bool {
	return r.Width <= 0 || r.Height <= 0
}

func (r *Rect) Location() *Point {
	return NewPoint(r.X, r.Y)
}

func (r *Rect) SetLocation(pt *Point) {
	r.X = pt.X
	r.Y = pt.Y
}

func (r *Rect) Offset(ox, oy int) {
	r.X = r.X + ox
	r.Y = r.Y + oy
}

func (r *Rect) Contains(pt *Point) bool {
	return r.X < pt.X && r.Y < pt.Y && pt.X < r.Right() && pt.Y < r.Bottom()
}

func (r *Rect) RegionLeft() *Rect {
	return NewRectFromSize(r.TopLeft(), NewSize(r.Width/2, r.Height))
}

func (r *Rect) RegionRight() *Rect {
	rect := r.RegionLeft()
	rect.SetLocation(NewPoint(r.Left()+r.Width/2, r.Top()))
	return rect
}

func (r *Rect) RegionTop() *Rect {
	return NewRectFromSize(r.TopLeft(), NewSize(r.Width, r.Height/2))
}

func (r *Rect) RegionBottom() *Rect {
	rect := r.RegionTop()
	rect.SetLocation(NewPoint(r.Left(), r.Top()+r.Width/2))
	return rect
}

func (r *Rect) RegionTopLeft() *Rect {
	return NewRectFromSize(r.TopLeft(), NewSize(r.Width/2, r.Height/2))
}

func (r *Rect) RegionTopRight() *Rect {
	rt := r.RegionTopLeft()
	rt.SetLocation(NewPoint(r.X+r.Width/2, r.Y))
	return rt
}

func (r *Rect) RegionBottomLeft() *Rect {
	rt := r.RegionTopLeft()
	rt.SetLocation(NewPoint(r.X, r.Y+r.Height/2))
	return rt
}

func (r *Rect) RegionMiddleCenter() *Rect {
	rt := r.RegionTopLeft()
	rt.SetLocation(NewPoint(r.X+r.Width/3, r.Y+r.Height/3))
	return rt
}

func (r *Rect) RegionBottomRight() *Rect {
	rt := r.RegionTopLeft()
	rt.SetLocation(NewPoint(r.X+r.Width/2, r.Y+r.Height/2))
	return rt
}

func (r *Rect) RegionCenter() *Rect {
	rt := NewRect(NewSize(r.Width/2, r.Height))
	rt.SetLocation(NewPoint(r.X+r.Width/3, r.Y))
	return rt
}

func (r *Rect) RegionMiddle() *Rect {
	rt := NewRect(NewSize(r.Width, r.Height/2))
	rt.SetLocation(NewPoint(r.X, r.Y+r.Height/3))
	return rt
}

func (r *Rect) String() string {
	return fmt.Sprintf("Rect(x:%d, y:%d, width:%d, height:%d)", r.X, r.Y, r.Width, r.Height)
}

package draw

import (
	"strconv"
	"strings"
)

func ParseToInt(raw string) ([]int, error) {
	strLst1 := strings.Split(strings.Trim(raw, ", "), ",")
	if len(strLst1) == 0 {
		return nil, ErrorEmptyData
	}
	var strLst []string
	for _, s := range strLst1 {
		ss := strings.Split(strings.Trim(s, ", "), "|")
		for _, s2 := range ss {
			strLst = append(strLst, s2)
		}
	}

	intLst := make([]int, len(strLst))
	for i, s := range strLst {
		d, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		intLst[i] = d
	}

	return intLst, nil
}

func ParsePointSize(x, y, w, h int) (*Point, *Size) {
	return NewPoint(x, y), NewSize(w, h)
}

func ParsePoint(x, y, x1, y1 int) (*Point, *Point) {
	return NewPoint(x, y), NewPoint(x1, y1)
}

func PointToRect(x, y, x1, y1 int) *Rect {
	return NewRectFromPoint(ParsePoint(x, y, x1, y1))
}

func SizeToRect(x, y, w, h int) *Rect {
	return NewRectFromSize(ParsePointSize(x, y, w, h))
}

package geom

import (
	"fmt"
	"math"
)

// Pnt 2-dim pnt
type Pnt struct {
	X float64
	Y float64
}

// Dist pnt distance
func (p Pnt) Dist(p1 Pnt) float64 {
	return math.Sqrt((p.X-p1.X)*(p.X-p1.X) + (p.Y-p1.Y)*(p.Y-p1.Y))
}

func (p Pnt) String() string {
	return "{ " + fmt.Sprint(p.X) + ", " + fmt.Sprint(p.Y) + " }"
}

type ByXThenY []Pnt

func (a ByXThenY) Len() int      { return len(a) }
func (a ByXThenY) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByXThenY) Less(i, j int) bool {
	if a[i].X < a[j].X {
		return true
	}
	if a[i].X > a[j].X {
		return false
	}
	if a[i].Y < a[j].Y {
		return true
	}
	if a[i].Y > a[j].Y {
		return false
	}
	return true
}

type ByX []Pnt

func (a ByX) Len() int      { return len(a) }
func (a ByX) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool {
	if a[i].X < a[j].X {
		return true
	}
	return false
}

type ByY []Pnt

func (a ByY) Len() int      { return len(a) }
func (a ByY) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByY) Less(i, j int) bool {
	if a[i].Y < a[j].Y {
		return true
	}
	return false
}

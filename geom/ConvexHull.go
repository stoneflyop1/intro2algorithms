package geom

import (
	"fmt"
	"sort"
)

func GetConvexHull(pnts []Pnt) []Pnt {
	sort.Sort(ByXThenY(pnts))
	count := len(pnts)
	p1 := pnts[0]
	pn := pnts[count-1]
	// x1 := p1.X
	// y1 := p1.Y
	// x2 := pn.X
	// y2 := pn.Y
	// a := y2 - y1
	// b := x1 - x2
	// c := x1*y2 - y1*x2
	ePnts := make([]Pnt, 0)
	pnts1 := make([]Pnt, 0)
	pnts2 := make([]Pnt, 0)
	for i := 1; i < count-1; i++ {
		// x3 := pnts[i].X
		// y3 := pnts[i].Y
		// d := a*x3 + b*y3 - c
		d := areaWithSymbol(p1, pn, pnts[i])
		if d > 0 {
			pnts1 = append(pnts1, pnts[i])
		} else if d < 0 {
			pnts2 = append(pnts2, pnts[i])
		}
	}
	ePnts = append(ePnts, p1)
	ePnts1 := convexHull(p1, pn, pnts1)
	ePnts = append(ePnts, ePnts1...)
	ePnts = append(ePnts, pn)
	ePnts2 := convexHull(pn, p1, pnts2)
	ePnts = append(ePnts, ePnts2...)
	return ePnts
}

func convexHull(p1, pn Pnt, pnts []Pnt) []Pnt {
	ePnts := make([]Pnt, 0)
	count := len(pnts)
	fmt.Println("total... ", p1, pn, pnts)
	if count == 0 {
		return ePnts
	}
	if count == 1 {
		ePnts = append(ePnts, pnts[0])
		fmt.Println("1 Pnt: ", pnts[0])
		return ePnts
	}
	dist0 := 0.0
	maxIndex := -1
	for i := 0; i < count; i++ {
		d := areaWithSymbol(p1, pn, pnts[i])
		d0 := d
		if d < 0 {
			d0 = -d
		}
		if d0 > dist0 {
			dist0 = d0
			maxIndex = i
		}
	}

	pmax := pnts[maxIndex]
	fmt.Println("pmax: ", pmax, len(pnts))

	pnts1, pnts2 := splitPnts(p1, pmax, pn, maxIndex, pnts)

	ps1 := convexHull(p1, pmax, pnts1)
	ePnts = append(ePnts, ps1...)
	ePnts = append(ePnts, pmax)
	ps2 := convexHull(pmax, pn, pnts2)
	ePnts = append(ePnts, ps2...)
	return ePnts
}

func splitPnts(p1, pmax, pn Pnt, maxIndex int, pnts []Pnt) ([]Pnt, []Pnt) {
	pnts1 := make([]Pnt, 0)
	pnts2 := make([]Pnt, 0)
	count := len(pnts)
	for i := 0; i < count; i++ {
		if i == maxIndex {
			continue
		}
		d := areaWithSymbol(p1, pmax, pnts[i])
		if d > 0 {
			pnts1 = append(pnts1, pnts[i])
		} else {
			d1 := areaWithSymbol(pmax, pn, pnts[i])
			if d1 > 0 {
				pnts2 = append(pnts2, pnts[i])
			}
		}
	}
	return pnts1, pnts2
}

func areaWithSymbol(p1, p2, p3 Pnt) float64 {
	x1 := p1.X
	y1 := p1.Y
	x2 := p2.X
	y2 := p2.Y
	x3 := p3.X
	y3 := p3.Y
	return x1*y2 + x3*y1 + x2*y3 - x3*y2 - x2*y1 - x1*y3
}

// GetConvexHullExtremePnts
func GetConvexHullExtremePnts(pnts []Pnt) []Pnt {

	count := len(pnts)
	eIndices := make(map[int]bool)
	for i := 0; i < count-1; i++ {
		x1 := pnts[i].X
		y1 := pnts[i].Y
		for j := i + 1; j < count; j++ {
			x2 := pnts[j].X
			y2 := pnts[j].Y
			a := y2 - y1
			b := x1 - x2
			c := x1*y2 - y1*x2
			dCount := 0
			ok := true
			d0 := 0.0
			for k := 0; k < count; k++ {
				if k == i || k == j {
					continue
				}
				x3 := pnts[k].X
				y3 := pnts[k].Y
				d := a*x3 + b*y3 - c
				if dCount == 0 {
					d0 = d
					dCount++
				} else if d0*d <= 0 { // not at the same side
					ok = false
					break
				}
			}
			if ok {
				_, isIn := eIndices[i]
				if !isIn {
					eIndices[i] = true
				}
				_, isIn = eIndices[j]
				if !isIn {
					eIndices[j] = true
				}
			}
		}
	}
	ePnts := make([]Pnt, len(eIndices))
	index := 0
	for key, _ := range eIndices {
		ePnts[index] = pnts[key]
		index++
	}
	return ePnts
}

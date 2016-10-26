package geom

import (
	"math"
	"sort"
)

// MinDistPnts min distant pnts pair, brute-force method
func MinDistPnts(pnts []Pnt) (Pnt, Pnt) {
	count := len(pnts)
	minI := 0
	minJ := 1
	dist := pnts[0].Dist(pnts[1])
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if i == 0 && j == 1 {
				continue
			}
			dist1 := pnts[i].Dist(pnts[j])
			if dist > dist1 {
				minI = i
				minJ = j
				dist = dist1
			}
		}
	}
	return pnts[minI], pnts[minJ]
}

func MinDistPnts2(pnts []Pnt) (Pnt, Pnt) {
	count := len(pnts)
	if count <= 3 {
		return MinDistPnts(pnts)
	}
	pntsX := make([]Pnt, count)
	pntsY := make([]Pnt, count)
	for i := 0; i < count; i++ {
		pntsX[i] = pnts[i]
		pntsY[i] = pnts[i]
	}
	sort.Sort(ByX(pntsX))
	sort.Sort(ByY(pntsY))
	return efficientClosestPair(pntsX, pntsY)
}

func efficientClosestPair(pntsX, pntsY []Pnt) (Pnt, Pnt) {
	count := len(pntsX)
	if count <= 3 {
		return MinDistPnts(pntsX)
	}
	// 按照中值分离数组
	l := count / 2
	r := count - l
	pl := make([]Pnt, l)
	ql := make([]Pnt, l)
	for i := 0; i < l; i++ {
		pl[i] = pntsX[i]
		ql[i] = pntsY[i]
	}
	pr := make([]Pnt, r)
	qr := make([]Pnt, r)
	for i := 0; i < r; i++ {
		pr[i] = pntsX[l+i]
		qr[i] = pntsY[l+i]
	}
	pl1, pl2 := efficientClosestPair(pl, ql)
	pr1, pr2 := efficientClosestPair(pr, qr)
	dl := pl1.Dist(pl2)
	dr := pr1.Dist(pr2)
	var pp1, pp2 Pnt
	pp1 = pl1
	pp2 = pl2
	d := dl
	if dl > dr {
		d = dr
		pp1 = pr1
		pp2 = pr2
	}

	// 处理中值附近的可能比d更小的点集，x的差在d之内
	m := pntsX[l-1].X
	s := make([]Pnt, 0)
	for i := 0; i < count; i++ {
		if math.Abs(pntsY[i].X-m) < d {
			s = append(s, pntsY[i])
		}
	}
	num := len(s)
	dmin := d * d

	for i := 0; i <= num-2; i++ {
		k := i + 1
		for k <= num-1 && (s[k].Y-s[i].Y)*(s[k].Y-s[i].Y) < dmin {
			dd := (s[k].X-s[i].X)*s[k].X - s[i].X + (s[k].Y-s[i].Y)*(s[k].Y-s[i].Y)
			if dd < dmin {
				dmin = dd
				pp1 = s[k]
				pp2 = s[i]
			}
			k++
		}
	}
	return pp1, pp2
}

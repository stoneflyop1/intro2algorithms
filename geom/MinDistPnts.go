package geom

// MinDistPnts min distant pnts pair
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

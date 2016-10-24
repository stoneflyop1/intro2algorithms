package geom

import "fmt"

func Tests() {
	fmt.Println("**************Geom****************")
	fmt.Println("----------minDistPnts---------")
	pnts := []Pnt{Pnt{X: 1, Y: 2}, Pnt{X: 2, Y: 4}, Pnt{X: 1, Y: 3}, Pnt{X: -1, Y: -5}, Pnt{X: -10, Y: -5}, Pnt{X: -1, Y: -50}, Pnt{X: -5, Y: -1}, Pnt{X: -1, Y: -60}}
	fmt.Println(pnts)
	pnt1, pnt2 := MinDistPnts(pnts)
	fmt.Println("Brute Force: ", pnt1, pnt2)

	pnt1, pnt2 = MinDistPnts2(pnts)
	fmt.Println("Divide Conquer: ", pnt1, pnt2)

	fmt.Println("-------------convexHull---------------")
	pnts0 := []Pnt{Pnt{X: 3, Y: 0}, Pnt{X: 2, Y: 0}, Pnt{X: 1, Y: 0}, Pnt{X: 0, Y: 0},
		Pnt{X: -1, Y: 0}, Pnt{X: -2, Y: 0},
		Pnt{X: 0, Y: 3}, Pnt{X: 0, Y: 2}, Pnt{X: 0, Y: 1}, Pnt{X: 0, Y: -1}, Pnt{X: 0, Y: -2}}
	fmt.Println(pnts0)
	pnts1 := GetConvexHullExtremePnts(pnts0)
	fmt.Println(pnts1)
	pnts2 := GetConvexHull(pnts0)
	fmt.Println(pnts2)
}

// run sample data for test
package main

import (
	"zjf/intro2alg/datastructs"
	"zjf/intro2alg/geom"
	"zjf/intro2alg/sorts"
	"math/rand"
)

func main() {
	sorts.Tests()
	geom.Tests()
	datastructs.Tests()
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

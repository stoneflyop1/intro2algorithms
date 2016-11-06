package datastructs

// Comparable compares two objects, used to implement collections sorting
type Comparable interface {
	compare(v interface{}) int
}

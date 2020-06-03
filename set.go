// Package set implements a simple set
// with some functions and methods.
package set

// The set is represented as a map[<type>]empty.
// The <type> can be either int or rune
type empty struct{}
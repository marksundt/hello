package main

import (
	"github.com/marksundt/samplepkg/samplepkg"
	"github.com/marksundt/samplepkg/subpkg"
)

func main() {
	sample := samplepkg.New("Test Sample Package")
	sample.Print()

	sub := subpkg.New("Test Sub Package")
	sub.Print()
}

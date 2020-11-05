package main

import (
	"bufio"
	"strconv"
	"strings"
)

func ParseWorld(r *bufio.Reader) World {
	dimensionsLine, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	dimensions := strings.Fields(dimensionsLine)
	dy, dx := validateDimension(dimensions[0]), validateDimension(dimensions[1])
	world := NewWorld(dy, dx)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			cell, _, _ := r.ReadRune()
			if cell == '*' {
				world.Awaken(y, x)
			}
		}
		r.ReadRune()
	}

	return world
}

func validateDimension(dimStr string) int {
	dim, err := strconv.Atoi(dimStr)
	if err != nil {
		panic(err)
	}
	if dim < 0 {
		panic("invalid dimension")
	}
	return dim
}

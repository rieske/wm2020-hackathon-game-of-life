package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	worldFile := "worlds/test.world"
	if len(args) == 1 {
		worldFile = args[0]
	}
	f, err := os.Open(worldFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	world := ParseWorld(bufio.NewReader(f))
	aliveRune = rune('\u2588')
	deadRune = ' '
	for {
		clearScreen()
		fmt.Println(world)
		world.Advance()
		time.Sleep(100 * time.Millisecond)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

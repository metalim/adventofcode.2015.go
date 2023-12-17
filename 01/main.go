package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go input.txt")
		os.Exit(1)
	}

	bs, err := os.ReadFile(os.Args[1])
	catch(err)
	lines := strings.Split(string(bs), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	timeStart := time.Now()
	var floor int
	for _, c := range lines[0] {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	fmt.Println("Part 1:", floor, "\tin", time.Since(timeStart))
}

func part2(lines []string) {
	timeStart := time.Now()
	var position, floor int
	for _, c := range lines[0] {
		position++
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			break
		}
	}

	fmt.Println("Part 2:", position, "\tin", time.Since(timeStart))
}

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
	var sum int
	for _, line := range lines {
		dims := toInts(strings.Split(line, "x"))
		sum += 2*dims[0]*dims[1] + 2*dims[1]*dims[2] + 2*dims[2]*dims[0] + min(dims[0]*dims[1], dims[1]*dims[2], dims[2]*dims[0])
	}

	fmt.Println("Part 1:", sum, "\tin", time.Since(timeStart))
}

func part2(lines []string) {
	timeStart := time.Now()
	var sum int
	for _, line := range lines {
		dims := toInts(strings.Split(line, "x"))
		sum += dims[0]*dims[1]*dims[2] + 2*min(dims[0]+dims[1], dims[1]+dims[2], dims[2]+dims[0])
	}

	fmt.Println("Part 2:", sum, "\tin", time.Since(timeStart))
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/iamlucassantos/advent-of-code-go/utils"
	"strconv"
	s "strings"
)

func part1() {
	data := utils.LoadFile("input1.txt")
	scanner := bufio.NewScanner(s.NewReader(data))
	horizontalPosition := 0
	depthPosition := 0
	for scanner.Scan() {
		var items []string = s.Split(scanner.Text(), " ")
		command := items[0]
		value, _ := strconv.Atoi(items[1])
		if command == "up" {
			depthPosition -= value
		} else if command == "down" {
			depthPosition += value
		} else if command == "forward" {
			horizontalPosition += value
		} else {
			fmt.Println("Invalid command")
		}
	}
	fmt.Println("Horizontal position:", horizontalPosition, "Depth position:", depthPosition)
	fmt.Println("HxD:", horizontalPosition*depthPosition)
}

func part2() {
	data := utils.LoadFile("input2.txt")
	scanner := bufio.NewScanner(s.NewReader(data))
	horizontalPosition := 0
	depthPosition := 0
	aimPosition := 0
	for scanner.Scan() {
		items := s.Split(scanner.Text(), " ")
		command := items[0]
		value, _ := strconv.Atoi(items[1])
		if command == "up" {
			aimPosition -= value
		} else if command == "down" {
			aimPosition += value
		} else if command == "forward" {
			horizontalPosition += value
			depthPosition += aimPosition * value
		}
	}
	fmt.Println("Horizontal position:", horizontalPosition, "Depth position:", depthPosition)
	fmt.Println("HxD:", horizontalPosition*depthPosition)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Choose part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		part1()
	} else if part == 2 {
		part2()
	} else {
		fmt.Println("Invalid part")
	}
}

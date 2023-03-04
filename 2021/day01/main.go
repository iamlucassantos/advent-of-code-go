package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/iamlucassantos/advent-of-code-go/utils"
	"strconv"
	"strings"
)

func part1() {
	data := utils.LoadFile("input1.txt")

	scanner := bufio.NewScanner(strings.NewReader(data))

	var previousMeasurement int
	var counter int
	for scanner.Scan() {
		currentMeasurement, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		if previousMeasurement != 0 {
			if currentMeasurement > previousMeasurement {
				counter += 1
			}
		}

		previousMeasurement = currentMeasurement
	}

	fmt.Println("Number of values higher than the previous:", counter)
}

func part2() {
	data := utils.LoadFile("input2.txt")
	fmt.Println(data)

	scanner := bufio.NewScanner(strings.NewReader(data))

	var x []int
	var counter int

	for scanner.Scan() {

		currentMeasurement, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Error during conversion")
			return
		}

		if len(x) == 3 {
			sumBefore := x[0] + x[1] + x[2]
			x = x[1:]
			sumAfter := x[0] + x[1] + currentMeasurement

			if sumBefore < sumAfter {
				counter += 1
			}
		}

		x = append(x, currentMeasurement)
	}
	fmt.Println("Number of values higher than the previous:", counter)
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

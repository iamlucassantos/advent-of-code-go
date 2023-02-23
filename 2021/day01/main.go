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

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Choose part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		part1()
	} else {
		fmt.Println("wyee")
	}

}

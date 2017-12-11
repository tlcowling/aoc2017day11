package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("11")
	if err != nil {
		log.Fatalln(err)
	}
	directions := strings.Split(string(dat), ",")

	directionCount := map[string]int{
		"n":  0,
		"ne": 0,
		"se": 0,
		"s":  0,
		"sw": 0,
		"nw": 0,
	}

	max := 0

	for _, direction := range directions {
		directionCount[direction]++
		steps := stepsAway(directionCount)
		if steps > max {
			max = steps
		}
	}

	fmt.Println(stepsAway(directionCount))
	fmt.Println(max)
}

func stepsAway(directionsCountMap map[string]int) int {
	y := directionsCountMap["n"] + directionsCountMap["nw"] - directionsCountMap["se"] - directionsCountMap["s"]
	z := directionsCountMap["s"] + directionsCountMap["sw"] - directionsCountMap["ne"] - directionsCountMap["n"]
	x := directionsCountMap["ne"] + directionsCountMap["se"] - directionsCountMap["nw"] - directionsCountMap["sw"]
	abs := math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))
	return int(abs / 2)
}

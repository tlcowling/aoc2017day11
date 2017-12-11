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

	for _, direction := range directions {
		directionCount[direction]++
	}

	y := directionCount["n"] + directionCount["nw"] - directionCount["se"] - directionCount["s"]
	z := directionCount["s"] + directionCount["sw"] - directionCount["ne"] - directionCount["n"]
	x := directionCount["ne"] + directionCount["se"] - directionCount["nw"] - directionCount["sw"]

	fmt.Println(stepsAway(x, y, z))
}

func stepsAway(x, y, z int) int {
	abs := math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))
	return int(abs / 2)
}

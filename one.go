package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func one() {
	dat, err := ioutil.ReadFile("inputs/1")
	if err != nil {
		log.Fatalln(err)
	}
	input := string(dat)
	fmt.Println(one1(input))
	fmt.Println(one2(input))
}

func one1(input string) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		this := input[i] - '0'
		next := input[(i+1)%len(input)] - '0'

		if this == next {
			sum += int(this)
		}
	}

	return sum
}

// would like to make this more general, like a pass the appropriate
// match criteria as a function to an iterator
func one2(input string) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		this := input[i] - '0'
		next := input[(i+len(input)/2)%len(input)] - '0'

		if this == next {
			sum += int(this)
		}
	}

	return sum
}

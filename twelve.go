package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/yourbasic/graph"
)

func twelve() {
	dat, err := ioutil.ReadFile("inputs/12")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(dat), "\n")
	g := graph.New(len(lines))

	for _, line := range lines {
		tokens := strings.Split(line, " <-> ")
		id := tokens[0]

		dependencies := strings.Split(tokens[1], ",")
		for _, dep := range dependencies {
			dep = strings.TrimSpace(dep)
			g.AddBoth(getInt(id), getInt(dep))
		}
	}

	comps := graph.StrongComponents(g)
	for _, group := range comps {
		for _, el := range group {
			if el == 0 {
				fmt.Println(len(group))
			}
		}
	}
	fmt.Println(len(comps))
}

func getInt(str string) int {
	idI, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln(err)
	}
	return idI
}

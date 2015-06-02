package main

import "fmt"

type person struct {
	name, lastname string
	height         int
}

func main() {

	p1, p2 := person{"peter", "buchmann", 191}, person{"erfried", "buchmann", 196}

	persons := []*person{&p1, &p2}

	outlaws := []*person{}

	outlaws = append(outlaws, persons[1])
	outlaws[0].height = 111

	for i := 0; i < len(persons); i++ {
		p := persons[i]
		fmt.Printf("%v\n", *p)

	}

}

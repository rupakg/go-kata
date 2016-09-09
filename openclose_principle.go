package main

import (
	"fmt"
)

/* Show the principle that Go types can have befhavior that are open for extension
** e.g. Legs() func which can be overriden by embedding into another type.
** But Go types can have behavior that can be closed for modification
** e.g. PrintLegs() which is a func of Cat and cannot be modified even after
** embedding into another type.
** Shared at: https://play.golang.org/p/4lNJI8qsAO
 */

type Cat struct {
	Name string
}

func (c Cat) Legs() int { return 4 }

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs.\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (o OctoCat) Legs() int { return 5 }

func main() {
	var octo OctoCat
	fmt.Println(octo.Legs()) // 5
	octo.PrintLegs()         // I have 4 legs
}

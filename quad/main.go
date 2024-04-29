package main

import piscine "piscine/quad"

import (
	"os"
	"strconv"
)

func main() {
	// string to int
    x, err := strconv.Atoi(os.Args[1])
	y, err := strconv.Atoi(os.Args[2])

    if err != nil {
        // ... handle error
        panic(err)
    }

    piscine.QuadC(x, y)
}

	//testing quadA
	// piscine.QuadA(5,3)
	// piscine.QuadA(5,1)
	// piscine.QuadA(1,1)
	// piscine.QuadA(1,5)

	// //testing quadB
	// piscine.QuadB(5,3)
	// piscine.QuadB(5,1)
	// piscine.QuadB(1,1)
	// piscine.QuadB(1,5)

	// //testing quadC
	// piscine.QuadC(5,3)
	// piscine.QuadC(5,1)
	// piscine.QuadC(1,1)
	// piscine.QuadC(1,5)

	// //testing quadD
	// piscine.QuadD(5,3)
	// piscine.QuadD(5,1)
	// piscine.QuadD(1,1)
	// piscine.QuadD(1,5)

	//testing quadE
	//piscine.QuadE(1, 2)
	//piscine.QuadC(1, 1)
	//piscine.QuadA(3, 3)
	// piscine.QuadE(1,1)
	// piscine.QuadE(1,5)


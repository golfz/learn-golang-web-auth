package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Firstname string
	}

	p1 := Person{
		Firstname: "GolFz",
	}

	p2 := Person{
		Firstname: "Tun",
	}

	xp := []Person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}

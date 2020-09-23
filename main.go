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

	fmt.Println("struct to json:", string(bs))

	xp2 := []Person{}
	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		panic(err)
	}

	fmt.Println("json to struct:", xp2)
}

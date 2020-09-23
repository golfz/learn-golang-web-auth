package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Firstname string
}

func main() {
	//type Person struct {
	//	Firstname string
	//}
	//
	//p1 := Person{
	//	Firstname: "GolFz",
	//}
	//
	//p2 := Person{
	//	Firstname: "Tun",
	//}
	//
	//xp := []Person{p1, p2}
	//
	//bs, err := json.Marshal(xp)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("struct to json:", string(bs))
	//
	//xp2 := []Person{}
	//err = json.Unmarshal(bs, &xp2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("json to struct:", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	http.ListenAndServe(":9090", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		Firstname: "GolFz",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Bad data encode:", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	p2 := Person{}
	err := json.NewDecoder(r.Body).Decode(&p2)
	if err != nil {
		log.Println("bad decode", err)
	}

	log.Println("decode", p2)
}

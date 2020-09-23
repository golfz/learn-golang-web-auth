package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.ListenAndServe(":9090", nil)
}

type Person struct {
	Name string
}

func encode(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		Name: "Surattikorn",
	}

	p2 := Person{
		Name: "Supamas",
	}

	xp := []Person{p1,p2}

	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		log.Println("encode error:", err)
	}
}

func decode(w http.ResponseWriter, r *http.Request) {
	xp := []Person{}

	err := json.NewDecoder(r.Body).Decode(&xp)
	if err != nil {
		log.Println("decode error:", err)
	}

	log.Println("decode result:", xp)
}


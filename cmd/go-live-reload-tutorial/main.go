package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type dollers float32

func (d dollers) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollers

func (d database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for k, v := range d {
		fmt.Fprintf(w, "%s, %s\n", k, v)
	}
}

type myuuid string

func (d myuuid) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, d)
}

func main() {
	go func() {
		db := database{"a": 5, "b": 6}
		log.Fatal(http.ListenAndServe("localhost:8000", db))
	}()
	go func() {
		db := database{"a": 3, "b": 2}
		log.Fatal(http.ListenAndServe("localhost:8001", db))
	}()
	go func() {
		db := database{"a": 9, "b": 7}
		log.Fatal(http.ListenAndServe("localhost:8002", db))
	}()
	go func() {
		id := myuuid(uuid.New().String())
		log.Fatal(http.ListenAndServe("localhost:8003", id))
	}()
	go func() {
		id := myuuid(uuid.New().String())
		log.Fatal(http.ListenAndServe("localhost:8004", id))
	}()
	go func() {
		id := myuuid(uuid.New().String())
		log.Fatal(http.ListenAndServe("localhost:8005", id))
	}()
	for {
	}
}

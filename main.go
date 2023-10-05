package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Go Short Server Version 0.1")

	store := NewDbInstance()
	err := store.Init()
	if err != nil {
		log.Fatal(err)
	}

	api := NewAPIServer("localhost:3579", store)

	err = api.Start()
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("POB WebApp", version(), "started at", config.Address)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}

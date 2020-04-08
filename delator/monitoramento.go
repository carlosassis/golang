package main

import (
	"fmt"
	"net/http"
)

func main() {
	wasp()
}

func wasp() {
	for {
		wasp := "http://delator.back.b2w/application/1/status"
		resp, _ := http.Get(wasp)
		fmt.Println("Wasp", resp.Status)
	}
}

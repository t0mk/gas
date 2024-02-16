package main

import (
	"fmt"
)

func main() {
	// Send request
	response, err := GetGasPrices()
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

}

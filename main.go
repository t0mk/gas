package main

import (
	"fmt"
)

func main() {
	// Send request
	response, err := GetGasPrices()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

}

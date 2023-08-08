package main

import (
	"fmt"
	"khasa/src/kql"
)

func main() {
	input := "SET test \"HELLO WORLD\" TYPE=STRING TTL=3600"
	_, err := kql.TokenizeInput(input)
	if err != nil {
		fmt.Println(err)
	}
}

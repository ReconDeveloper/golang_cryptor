package main

import (
	"fmt"
	"math/rand"
	"time"
)

func uniqueFunctionName(l int) string {

	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	bytes := []byte(str)

	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < l; i++ {

		result = append(result, bytes[r.Intn(len(bytes))])

	}

	return string(result)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 2
	max := 20
	value := rand.Intn(max-min+1) + min
	functionName := uniqueFunctionName(value)
	fmt.Printf("Function Name:\t %v \n", functionName)

}

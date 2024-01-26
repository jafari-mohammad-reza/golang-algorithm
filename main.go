package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	fmt.Println("Golang algorithm course.")
	RunNamesBinarySearch(EnglishNames)
} 

func GenRandNumber(max *int) int {
	var maxNum int
	if max != nil {
		maxNum = *max
	}else { 
		maxNum = 100
	}
	rand.Seed(time.Now().UnixNano())
    return rand.Intn(maxNum)
}
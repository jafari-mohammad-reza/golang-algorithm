package main

import (
	"fmt"
	"sort"
	"strings"
)


var EnglishNames = []string{
	"Aaron", "Abigail", "Adam",      // A
	"Barbara", "Benjamin", "Bethany",  // B
	"Caleb", "Caroline", "Charles",    // C
	"Daniel", "David", "Diana",       // D
	"Edward", "Elizabeth", "Emily",    // E
	"Fiona", "Francis", "Frederick",   // F
	"Gabriel", "Grace", "Gregory",    // G
	"Hannah", "Harold", "Heather",    // H
	"Ian", "Isaac", "Isabella",       // I
	"Jack", "Jacob", "Jasmine",       // J
	"Katherine", "Kevin", "Kimberly", // K
	"Laura", "Leonard", "Lily",       // L
	"Madison", "Mark", "Megan",       // M
	"Nancy", "Nathan", "Nicole",      // N
	"Oliver", "Olivia", "Oscar",      // O
	"Patricia", "Paul", "Penelope",   // P
	"Quentin", "Quinn", "Quincy",     // Q
	"Rachel", "Raymond", "Rebecca",   // R
	"Samuel", "Sara", "Steven",       // S
	"Teresa", "Thomas", "Tiffany",    // T
	"Ulysses", "Uma", "Uriel",        // U
	"Valerie", "Victor", "Violet",    // V
	"Walter", "Wendy", "William",     // W
	"Xander", "Xavier", "Xena",       // X
	"Yasmin", "Yosef", "Yvonne",      // Y
	"Zachary", "Zara", "Zoe", // Z
}

 

func RunNamesBinarySearch(arr []string) {
	arrLen := len(arr)
	randNum := GenRandNumber(&arrLen)
	randItem := arr[randNum]
	fmt.Printf("Randomly selected item: %s (Index: %d)\n", randItem, randNum)

	foundIndex, foundName, stepCount := NamesBinarySearch(randItem, arr)
	if foundIndex != -1 {
		fmt.Printf("Binary Search: Found '%s' at index %d with %d tries", foundName, foundIndex , stepCount)
	} else {
		fmt.Println("Binary Search: Item not found.")
	}
}


func NamesBinarySearch(name string, arr []string) (int, string, int) {
	sort.Strings(arr)
	low, high := 0, len(arr)-1
	nameLower := strings.ToLower(name)
	stepCount := 0

	for low <= high {
		stepCount++
		mid := low + (high-low)/2

		if strings.ToLower(arr[mid]) == nameLower {
			return mid, arr[mid], stepCount // Found the name
		}
		if strings.ToLower(arr[mid]) < nameLower {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, "", stepCount // Name not found
}
func AlphabetIndex(letter string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	return strings.Index(alphabet, strings.ToLower(letter))
}
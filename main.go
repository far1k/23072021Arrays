package main

import "fmt"

func main() {
	theArray := [4]int{33, 22, 11, 0}
	for index, _ := range (theArray) {
		fmt.Println(theArray[index])
		theArray[index]++
	}
	fmt.Println(theArray)
}
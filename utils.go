package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

//readInput reads the input for a given prompt
func readInput(prompt string, mandatory bool) (text string) {
	for {
		fmt.Printf("%s :", prompt)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text = scanner.Text()

		if text == "" && mandatory {
			fmt.Printf("A value for %s is required.\n", prompt)
		} else {
			break
		}
	}
	return
}

//printMessage prints the given message with a new line at the end
func printMessage(message string, args ...interface{}) {
	fmt.Printf(message, args...)
	fmt.Println()
}

func printErrorAndExit(mess string) {
	if mess != "" {
		fmt.Println(mess)
	} else {
		fmt.Println("You are hiring a boring developer! Ask him to test his code first. :)")
	}
	os.Exit(1)
}

//ContainsStr returns true if the given element is a part of the array
func ContainsStr(arr []string, ele string) bool {
	return IndexOfStr(arr, ele) != -1
}

//IndexOfStr returns the index of ele in the str array, else -1
func IndexOfStr(arr []string, ele string) int {
	for i, str := range arr {
		if str == ele {
			return i
		}
	}
	return -1
}

//EqualsIntArr compares all elements of int arr a and b
func EqualsIntArr(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if b[i] != val {
			return false
		}
	}
	return true
}

//Roll returns an integer in between min and max, inclusive of max
//returns -1 in case of an invalid input
func Roll(min int, max int) int {
	if min < 0 || max < 0 || (min > max) {
		return -1
	}
	max += 1 // since rand.Intn dosent include the max value
	return rand.Intn(max-min) + min
}

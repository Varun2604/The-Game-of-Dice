package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
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

//ContainsInt returns true if the given element is a part of the array
func ContainsInt(arr []int, ele int) bool {
	return IndexOfInt(arr, ele) != -1
}

//IndexOfStr returns the index of ele in the str array, else -1
func IndexOfInt(arr []int, ele int) int {
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

//SortInt sorts and returns the array in decending
func SortIntDesc(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	return arr
}

//PlayerRankedList Stores the player rank list
//Does insert operations in O(log(n)) complexity
//Quite useful to use this DS when the rank of the
type RankedListNode interface {
	//Rank returns a valid rank else -1
	Rank() int
	//SetRank Lets you update the rank of the current node if the score criteria is met
	SetRank(int)
	//Score returns a valid score else -1
	Score() int
	Equal(interface{}) bool
	Player() string
}

type PlayerRankedList struct {
	arr               []RankedListNode
	size              int
	latestRank        int
	maxScore          int
	rankedPlayerCount int
}

func (l *PlayerRankedList) Insert(node RankedListNode) {
	l.size += 1
	l.arr = append(l.arr, node)
	l.orderArr(l.size)
}

func (l *PlayerRankedList) InsertOrUpdate(node RankedListNode) {
	//find index of the element, and update rank from that element
	updatedIdx := -1
	for i, n := range l.arr {
		if n.Equal(node) {
			if node.Score() >= l.maxScore {
				node.SetRank(l.latestRank)
				l.latestRank += 1
				l.rankedPlayerCount += 1
			}
			l.arr[i] = node
			updatedIdx = i
			break
		}
	}
	if updatedIdx == -1 {
		//the node is not present in the array, insert it
		l.Insert(node)
		return
	}
	l.orderArr(updatedIdx)
}

// Traverse up and fix violated property
func (l *PlayerRankedList) orderArr(current int) {
	for l.greater(l.arr[current], l.arr[l.parent(current)]) {
		l.swap(current, l.parent(current))
		current = l.parent(current)
	}
}

func (l *PlayerRankedList) swap(fpos, spos int) {
	tmp := l.arr[fpos]
	l.arr[fpos] = l.arr[spos]
	l.arr[spos] = tmp
}

func (l *PlayerRankedList) parent(pos int) int {
	return pos / 2
}

func (l *PlayerRankedList) greater(lNode, rNode RankedListNode) bool {
	//when ranks are set, compare with ranks
	if (lNode.Rank() != -1) && (rNode.Rank() != -1) {
		return (lNode.Rank() < rNode.Rank())
	}
	//if no ranks, compare with score
	return (lNode.Score() > rNode.Score())
}

//List returns the list
func (l *PlayerRankedList) List() []RankedListNode {
	//TODO - return a copy of the array to avoid mutation
	return l.arr
}

//RankedPlayersCount returns the count of ranked players
func (l *PlayerRankedList) RankedPlayersCount() int {
	//TODO - return a copy of the array to avoid mutation
	return l.rankedPlayerCount
}

func NewRankedList(maxScore int) *PlayerRankedList {
	return &PlayerRankedList{
		arr:               []RankedListNode{},
		size:              -1,
		latestRank:        1,
		maxScore:          maxScore,
		rankedPlayerCount: 0,
	}
}

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

//EqualsStrArr compares all elements of int arr a and b
func EqualsStrArr(a []string, b []string) bool {
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

//PlayerRankedList Stores the player rank list
//Does insert operations in O(Log(n)) complexity (complexity in building the heap)
//Does list operation in O(nLog(n)) complexity (does lazy sorting)
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
	l.orderHeap(l.size)
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
	l.orderHeap(updatedIdx)
}

// Traverse up and fix violated property
func (l *PlayerRankedList) orderHeap(current int) {
	for l.greater(l.arr[current], l.arr[l.parent(current)]) {
		l.swap(current, l.parent(current))
		current = l.parent(current)
	}
}

// Do heap sort
func (l *PlayerRankedList) sort() {
	for i := l.size; i > 0; i-- {
		// Move current root to end
		l.swap(0, i)

		// call max heapify on the reduced heap
		l.heapify(0, i)
	}
	//rearrange in max heap format (reverse the array)
	for i := 0; i < ((l.size + 1) / 2); i++ {
		l.swap(i, (l.size - i))
	}
}

//heapify the array
func (l *PlayerRankedList) heapify(from int, to int) {
	largest := from          // Initialize largest as root
	lIdx := ((2 * from) + 1) // left = 2*i + 1
	rIdx := ((2 * from) + 2) // right = 2*i + 2

	// If left child is larger than root
	if (lIdx < to) && l.greater(l.arr[lIdx], l.arr[largest]) {
		largest = lIdx
	}

	// If right child is larger than largest so far
	if (rIdx < to) && l.greater(l.arr[rIdx], l.arr[largest]) {
		largest = rIdx
	}

	// If largest is not root
	if largest != from {
		l.swap(from, largest)
		// Recursively heapify the affected sub-tree
		l.heapify(largest, to)
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
	if (lNode.Rank() > 0) && (rNode.Rank() > 0) {
		return (lNode.Rank() < rNode.Rank())
	}
	//if no ranks, compare with score
	return (lNode.Score() > rNode.Score())
}

//List returns the list
func (l *PlayerRankedList) List() []RankedListNode {
	//TODO - return a copy of the array to avoid mutation
	l.sort()
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

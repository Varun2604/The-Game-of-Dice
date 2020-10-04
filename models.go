package main

import "fmt"

//Inputs input struct with private fields
type Inputs struct {
	totPlayers int
	maxScore   int
	players    []string
}

func (i *Inputs) TotalPlayers() int {
	return i.totPlayers
}
func (i *Inputs) ListPlayers() []string {
	return i.players
}
func (i *Inputs) MaxScore() int {
	return i.maxScore
}

//NewGameInputs creates and returns a new store for game inputs
func NewGameInputs(totPlayers int, maxScore int) *Inputs {
	players := []string{}
	for i := 1; i <= totPlayers; i++ {
		players = append(players, fmt.Sprint("Player ", i))
	}
	return &Inputs{
		totPlayers: totPlayers,
		maxScore:   maxScore,
		players:    players,
	}
}

//Log logs the ongoings of the game
type Log struct {
	rankOrder          []string
	lastTwoPlayerRolls map[string][2]int
	playerScore        map[string]int
	inputs             *Inputs
}

//RecordScore records the player's roll score
func (l *Log) RecordScore(player string, rollScore int) {
	l.updateLastRoll(player, rollScore)
	latestScore := l.updatePlayerScore(player, rollScore)
	if latestScore >= l.inputs.maxScore {
		l.rankOrder = append(l.rankOrder, player)
	}
}

//Score returns the player's score
func (l *Log) Score(player string) int {
	return l.playerScore[player]
}

//CanPlay returns true if the player can play the next turn
func (l *Log) SkipChance(player string) bool {
	return ContainsStr(l.rankOrder, player)
}

//updatePlayerScore updates the player's current score
func (l *Log) updatePlayerScore(player string, currentRollScore int) int {
	score, ok := l.playerScore[player]
	if !ok {
		score = 0
	}
	score += currentRollScore
	l.playerScore[player] = score
	return score
}

//updateLastRoll record's the player's last roll.
func (l *Log) updateLastRoll(player string, rollScore int) {
	lastTwoRolls, ok := l.lastTwoPlayerRolls[player]
	if !ok {
		lastTwoRolls = [2]int{0, 0}
	}
	//TODO - write a better logic to switch and update the last rolls
	x := lastTwoRolls[0]
	lastTwoRolls[0] = rollScore
	lastTwoRolls[1] = x
	l.lastTwoPlayerRolls[player] = lastTwoRolls // not required...
}

//LastTwoRolls returns last two rolls of the players
func (l *Log) LastTwoRolls(player string) []int {
	rolls := []int{} //do not return the ref to the actual score
	lastTwoRolls, ok := l.lastTwoPlayerRolls[player]
	if !ok {
		lastTwoRolls = [2]int{0, 0}
	}
	rolls = append(rolls, lastTwoRolls[0])
	rolls = append(rolls, lastTwoRolls[1])
	return rolls
}

//Ranks returns the ranks of the players
func (l *Log) Ranks() []string {
	ranks := []string{} //do not return the ref to the actual score
	ranks = append(ranks, l.rankOrder...)
	return ranks
}

//NewGameLog returns a new game log instance
func NewGameLog(inputs *Inputs) *Log {
	return &Log{
		inputs:             inputs,
		lastTwoPlayerRolls: make(map[string][2]int),
		playerScore:        make(map[string]int),
	}
}

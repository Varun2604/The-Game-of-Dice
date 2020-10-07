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

type PlayersScorePair struct {
	players []string
	score   int
}

//PlayerScoreDetail the node for the rank list
type PlayerScoreDetail struct {
	player string
	score  int
	rank   int
}

func (p *PlayerScoreDetail) Equals(playerName string) bool {
	return playerName == p.player
}
func (d *PlayerScoreDetail) Equal(node interface{}) bool {
	t := fmt.Sprintf("%T", node)
	if t == "*main.PlayerScoreDetail" {
		detail := node.(*PlayerScoreDetail)
		return (detail.player == d.player)
	}
	return false
}
func (d *PlayerScoreDetail) Rank() int {
	return d.rank
}
func (d *PlayerScoreDetail) Score() int {
	return d.score
}
func (d *PlayerScoreDetail) Player() string {
	return d.player
}
func (d *PlayerScoreDetail) SetRank(r int) {
	d.rank = r
}
func NewPlayerScoreDetail(player string, score int) *PlayerScoreDetail {
	return &PlayerScoreDetail{
		rank:   -1,
		player: player,
		score:  score,
	}
}

//Log logs the ongoings of the game
type Log struct {
	lastTwoPlayerRolls map[string][2]int
	playerScore        map[string]int
	inputs             *Inputs
	rankOrderedList    *PlayerRankedList
}

//RecordScore records the player's roll score
func (l *Log) RecordScore(player string, rollScore int) {
	l.updateLastRoll(player, rollScore)
	latestScore := l.updatePlayerScore(player, rollScore)
	l.updateRankedList(player, latestScore)
}

//Score returns the player's score
func (l *Log) Score(player string) int {
	return l.playerScore[player]
}

func (l *Log) updateRankedList(player string, rollScore int) {
	d := NewPlayerScoreDetail(player, rollScore)
	l.rankOrderedList.InsertOrUpdate(d)
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

//ScoreBoard returns the current rank list
func (l *Log) ScoreBoard() []RankedListNode {
	return l.rankOrderedList.List()
}

//ScoreBoard calculates and returns the current scoreboard
func (l *Log) RankedPlayerCount() int {
	return l.rankOrderedList.RankedPlayersCount()
}

//NewGameLog returns a new game log instance
func NewGameLog(inputs *Inputs) *Log {
	return &Log{
		inputs:             inputs,
		lastTwoPlayerRolls: make(map[string][2]int),
		playerScore:        make(map[string]int),
		rankOrderedList:    NewRankedList(inputs.maxScore),
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	DICE_MIN = 1
	DICS_MAX = 6
)

var CANCEL_IF_LAST_TWO_ROLLS = []int{1, 1}

func startGame() {
	printMessage("Welcome to the Game of Dice!!!")

	gameInputs, err := initGame()

	if err != nil {
		printErrorAndExit("")
	}
	gameLog := NewGameLog(gameInputs)

	playGame(gameInputs, gameLog)

	printMessage("The game has ended.")
	printMessage("Ranks of all players are as below:")
	printMessage(strings.Join(gameLog.Ranks(), "\n"))
}

func playGame(gameInputs *Inputs, log *Log) {

	players := gameInputs.ListPlayers()
	for continueRound(log.Ranks(), gameInputs.totPlayers) {
		for _, player := range players {
			readInput(fmt.Sprintf("%s its your turn (press ‘r’ to roll the dice)", player),
				false) //TODO Does it matter if the user dosent type r ?
			playTurn(player, log)
		}
	}
}

func playTurn(player string, log *Log) {
	rollScore := 6
	for continueTurn(rollScore,
		log.Score(player), log.inputs.maxScore) {

		if !skipTurn(log.LastTwoRolls(player)) {
			printMessage("Rolling for %s", player)
			rollScore = Roll(DICE_MIN, DICS_MAX)
			printMessage("%s you have rolled %d", player, rollScore)
		} else {
			printMessage("Skipping turn for %s", player)
			rollScore = 0
		}

		log.RecordScore(player, rollScore)
	}
}

//continueTurn returns true if the turn is supposed to be continued
// condition - 1. checks the total score of the player
// condition - 2. checks if the last roll score is 6
func continueTurn(lastRollScore int,
	playerTotalScore int, maxPossibleScore int) bool {
	return ((playerTotalScore < maxPossibleScore) &&
		lastRollScore == 6)
}

func skipTurn(lastTwoRolls []int) bool {
	return EqualsIntArr(lastTwoRolls, CANCEL_IF_LAST_TWO_ROLLS)
}

//continueRound returns true if the round is supposed to be continued
func continueRound(playerRanks []string, totalPlayers int) bool {
	return ((len(playerRanks) + 1) < totalPlayers)
}

//initGame fetch the game variables
func initGame() (gameInput *Inputs, err error) {
	totalPlayers := 0
	maxScore := 0
	for totalPlayers <= 0 {
		totalPlayers, err = strconv.Atoi(readInput("No. of players?", true))
		if err != nil {
			printMessage(fmt.Sprint("Message for the developer - ", err.Error()))
			err = fmt.Errorf("Internal Error, unable to process the count of players")
			return
		}
		if totalPlayers <= 0 {
			printMessage("Number of players has to be more than 0")
		}
	}

	for maxScore <= 0 {
		maxScore, err = strconv.Atoi(readInput("Max Score for Game?", true))
		if err != nil {
			printMessage(fmt.Sprint("Message for the developer - ", err.Error()))
			err = fmt.Errorf("Internal Error, unable to process the max score")
			return
		}
		if maxScore <= 0 {
			printMessage("Max score has to be more than 0")
		}
	}
	gameInput = NewGameInputs(totalPlayers, maxScore)
	return
}

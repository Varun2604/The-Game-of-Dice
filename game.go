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
	printScoreBoard(gameLog.ScoreBoard())
}

func playGame(gameInputs *Inputs, log *Log) {

	players := gameInputs.ListPlayers()
	for continueRound(log.Ranks(), gameInputs.totPlayers) {
		for _, player := range players {
			readInput(fmt.Sprintf("%s its your turn (press ‘r’ to roll the dice)", player),
				false) //TODO Does it matter if the user dosent type r ?
			playTurn(player, log)
		}
		printScoreBoard(log.ScoreBoard())
	}
}

func printScoreBoard(board []PlayersScorePair) {
	printMessage("***********+Score Board+***********")
	for i, b := range board {
		printMessage("Rank %d - \"%s\" with score %d", (i + 1), strings.Join(b.players, ", "), b.score)
	}
	printMessage("***********************************")
}

func playTurn(player string, log *Log) {
	rollScore := 0
	won := hasWon(log.Score(player), log.inputs.maxScore)
	continueTurn := !won
	for continueTurn {
		if !skipTurn(log.LastTwoRolls(player)) {
			printMessage("Rolling for %s", player)
			rollScore = Roll(DICE_MIN, DICS_MAX)
			printMessage("%s, you have rolled %d", player, rollScore)
		} else {
			printMessage("Skipping turn for %s since the last two roll score was 1", player)
			rollScore = 0
		}
		log.RecordScore(player, rollScore)
		playAgain := playAgain(rollScore)
		continueTurn = !hasWon(log.Score(player), log.inputs.maxScore) && playAgain
		if playAgain {
			printMessage("%s gets another roll since the last score was 6", player)
		}
	}
	if won {
		printMessage("%s has completed the game, skipping round", player)
	}
}

//continueTurn returns true if the turn is supposed to be continued
// condition - 1. checks the total score of the player
func playAgain(lastRollScore int) bool {
	return (lastRollScore == 6)
}

//hasWon checks of the player has won the game
func hasWon(playerTotalScore int, maxPossibleScore int) bool {
	return !(playerTotalScore < maxPossibleScore)
}

//skipTurn checks if the player has to skip turn
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

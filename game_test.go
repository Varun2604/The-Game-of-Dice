package main

import "testing"

func TestContinueRound(t *testing.T) {
	rankedPlayersCount := 2
	totalPlayerInputs := []int{4, 3, 2}
	expOp := []bool{true, true, false}
	for i, totalPlayers := range totalPlayerInputs {
		if expOp[i] != continueRound(rankedPlayersCount, totalPlayers) {
			t.Errorf("Contiue round dosen't work when player count is %d and ranked players count - %d",
				totalPlayers, rankedPlayersCount)
		}
	}
}

func TestSkipTurn(t *testing.T) {
	lastTwoRollsInputs := [][]int{
		{1, 1},
		{2, 1},
		{0, 0},
		{6, 6},
	}
	expOp := []bool{true, false, false, false}
	for i, lastTwoRolls := range lastTwoRollsInputs {
		if expOp[i] != skipTurn(lastTwoRolls) {
			t.Errorf("skipTurn mis functions for arr %v with op %v", lastTwoRolls, expOp[i])
		}
	}
}

func TestPlayAgain(t *testing.T) {
	inputs := [][]int{
		{6},
		{1},
		{2},
		{3},
		{6},
	}
	expOp := []bool{true, false, false, false, true}
	for i, inp := range inputs {
		if expOp[i] != playAgain(inp[0]) {
			t.Errorf("playAgain mis functions for inputs %v with op %v", inp, expOp[i])
		}
	}
}

func TestHasWon(t *testing.T) {
	inputs := [][]int{
		{6, 1},
		{1, 6},
		{10, 100},
		{1, 1},
		{6, 6},
	}
	expOp := []bool{true, false, false, true, true}
	for i, inp := range inputs {
		if expOp[i] != hasWon(inp[0], inp[1]) {
			t.Errorf("hasWon mis functions for inputs %v with op %v", inp, expOp[i])
		}
	}
}

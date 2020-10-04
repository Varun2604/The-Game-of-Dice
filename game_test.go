package main

import "testing"

func TestContinueRound(t *testing.T) {
	rankedPlayers := []string{"A", "B"}
	totalPlayerInputs := []int{4, 3, 2}
	expOp := []bool{true, false, false}
	for i, totalPlayers := range totalPlayerInputs {
		if expOp[i] != continueRound(rankedPlayers, totalPlayers) {
			t.Errorf("Contiue round dosen't work when player count is %d and ranked players - %v",
				totalPlayers, rankedPlayers)
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

func TestContinueTurn(t *testing.T) {
	inputs := [][]int{
		{6, 1, 3},
		{6, 2, 2},
		{6, 3, 2},
		{1, 1, 3},
		{1, 2, 2},
		{1, 3, 2},
	}
	expOp := []bool{true, false, false, false, false, false}
	for i, inp := range inputs {
		if expOp[i] != continueTurn(inp[0], inp[1], inp[2]) {
			t.Errorf("continueTurn mis functions for inputs %v with op %v", inp, expOp[i])
		}
	}
}

package main

import "testing"

func TestInput(t *testing.T) {
	input := NewGameInputs(3, 15)
	t.Run("Test input.TotalPlayers()", func(T *testing.T) {
		if input.TotalPlayers() != 3 {
			t.Error("log.TotalPlayers malfunctioning")
		}
	})
	t.Run("Test input.MaxScore()", func(T *testing.T) {
		if input.MaxScore() != 15 {
			t.Error("log.MaxScore malfunctioning")
		}
	})
	t.Run("Test input.ListPlayers()", func(T *testing.T) {
		if len(input.ListPlayers()) != 3 {
			t.Error("log.ListPlayers malfunctioning")
		}
	})
}

func TestLog(t *testing.T) {
	input := NewGameInputs(3, 15)
	log := NewGameLog(input)
	player1 := input.ListPlayers()[0]
	player2 := input.ListPlayers()[1]
	t.Run("Test log.Score()", func(T *testing.T) {
		log.RecordScore(player1, 3)
		log.RecordScore(player1, 1)
		log.RecordScore(player1, 6)
		log.RecordScore(player1, 0)
		if log.Score(player1) != 10 {
			t.Error("log.Score malfunctioning")
		}
	})
	t.Run("Test log.LastTwoRolls()", func(T *testing.T) {
		log.RecordScore(player1, 1)
		log.RecordScore(player1, 1)
		if !EqualsIntArr(log.LastTwoRolls(player1), []int{1, 1}) {
			t.Error("log.RecordScore malfunctioning")
		}
		if !EqualsIntArr(log.LastTwoRolls(player2), []int{0, 0}) {
			t.Error("log.RecordScore malfunctioning for new player")
		}
	})

	t.Run("Test log.ScoreBoard()", func(T *testing.T) {
		scoreBoard := log.ScoreBoard()
		if len(scoreBoard) != 1 {
			t.Error("log.ScoreBoard malfunctioning - records scores of un played players")
		}
		log.RecordScore(player2, 1)
		scoreBoard = log.ScoreBoard()
		if len(scoreBoard) != 2 {
			t.Error("log.ScoreBoard malfunctioning - dosen't record scores of all players")
		}
		player1Score := scoreBoard[0]
		if !ContainsStr(player1Score.players, player1) {
			t.Error("log.ScoreBoard malfunctioning - invalid rank computation")
		}
		if player1Score.score != 12 {
			t.Error("log.ScoreBoard malfunctioning - invalid score computation")
		}
		if scoreBoard[1].score != 1 {
			t.Error("log.ScoreBoard malfunctioning - invalid score computation for last played")
		}
		//TODO test case to check scoreboard copy
	})
	t.Run("Test log.Ranks()", func(T *testing.T) {
		log.RecordScore(player1, 4)
		if len(log.Ranks()) != 1 || !ContainsStr(log.Ranks(), player1) {
			t.Error("log.Ranks malfunctioning")
		}
	})
	t.Run("Test log.ScoreBoard after max score", func(T *testing.T) {
		scoreBoard := log.ScoreBoard()
		if len(scoreBoard) != 2 {
			t.Error("log.ScoreBoard malfunctioning - dosen't record scores of all players")
		}
		player1Score := scoreBoard[0]
		if !ContainsStr(player1Score.players, player1) {
			t.Error("log.ScoreBoard malfunctioning - invalid rank computation")
		}
		if player1Score.score != 16 {
			t.Error("log.ScoreBoard malfunctioning - invalid score computation")
		}
		if scoreBoard[1].score != 1 {
			t.Error("log.ScoreBoard malfunctioning - invalid score computation for last played")
		}
	})

}

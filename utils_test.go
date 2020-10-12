package main

import "testing"

func TestEqualsIntArr(t *testing.T) {
	arr1 := []int{8, 2, 4, -1, 0, 1}
	arr2 := []int{8, 2, 4, -1, 0, 1}
	if !EqualsIntArr(arr1, arr2) {
		t.Error("EqualsIntArr malfunctioning - improper comparison of two equal arrays")
	}
	arr2 = append(arr2, 5)
	if EqualsIntArr(arr1, arr2) {
		t.Error("EqualsIntArr malfunctioning - improper comparison of two unequal length arrays")
	}
	arr2 = arr2[1:]
	if EqualsIntArr(arr1, arr2) {
		t.Error("EqualsIntArr malfunctioning - improper comparison of two unequal arrays")
	}
}
func TestSortIntDesc(t *testing.T) {
	arr := []int{8, 2, 4, -1, 0, 1}
	arr = SortIntDesc(arr)
	if !EqualsIntArr(arr, []int{8, 4, 2, 1, 0, -1}) {
		t.Error("SortIntDesc malfunctioning")
	}
}

func TestContainsStr(t *testing.T) {
	arr := []string{"a", "b", "c"}
	if !ContainsStr(arr, "a") {
		t.Error("ContainsStr malfunctioning - contains an ele but returns false")
	}
	if ContainsStr(arr, "d") {
		t.Error("ContainsStr malfunctioning - dose not contain an ele but returns true")
	}
}

func TestContainsInt(t *testing.T) {
	arr := []int{0, -1, 1, 8, 10000, -58}
	if !ContainsInt(arr, 1) {
		t.Error("ContainsInt malfunctioning - contains an ele but returns false")
	}
	if !ContainsInt(arr, -58) {
		t.Error("ContainsInt malfunctioning - contains an ele but returns false")
	}
	if !ContainsInt(arr, 0) {
		t.Error("ContainsInt malfunctioning - contains an ele but returns false")
	}
	if ContainsInt(arr, 20) {
		t.Error("ContainsStr malfunctioning - dose not contain an ele but returns true")
	}
}

func TestRoll(t *testing.T) {
	min := 1
	max := 6
	test := func(iteration int) {
		r := Roll(min, max)
		if r < 1 || r > 6 {
			t.Errorf("Roll malfunctions at %dth iteration", iteration)
		}
	}
	for i := 0; i < 100; i++ {
		test(i)
	}
	min = 6
	max = 5
	r := Roll(min, max)
	if r != -1 {
		t.Error("Roll malfunctions when min > max")
	}
	min = -1
	r = Roll(min, max)
	if r != -1 {
		t.Error("Roll malfunctions when min < 0")
	}
	min = 6
	max = -1
	r = Roll(min, max)
	if r != -1 {
		t.Error("Roll malfunctions when max < 0")
	}
}

func TestPlayerRankedList(t *testing.T) {
	test := func(l *PlayerRankedList, expectedRankList []string) {
		ranklist := make([]string, 6)
		for i, r := range l.List() {
			ranklist[i] = r.Player()
		}
		if !EqualsStrArr(expectedRankList, ranklist) {
			t.Fatalf("o/p is %v, expected is %v", ranklist, expectedRankList)
		}
	}
	l := NewRankedList(20)
	p1 := &PlayerScoreDetail{
		player: "player 1", score: 10,
	}
	p2 := &PlayerScoreDetail{
		player: "player 2", score: 6,
	}
	p3 := &PlayerScoreDetail{
		player: "player 3", score: 4,
	}
	p4 := &PlayerScoreDetail{
		player: "player 4", score: 8,
	}
	p5 := &PlayerScoreDetail{
		player: "player 5", score: 9,
	}
	p6 := &PlayerScoreDetail{
		player: "player 6", score: 7,
	}
	l.InsertOrUpdate(p1)
	l.InsertOrUpdate(p2)
	l.InsertOrUpdate(p3)
	l.InsertOrUpdate(p4)
	l.InsertOrUpdate(p5)
	l.InsertOrUpdate(p6)

	expectedRankList := []string{"player 1", "player 5", "player 4", "player 6",
		"player 2", "player 3"}
	t.Run("Test List with only inserted valid scores", func(t *testing.T) {
		test(l, expectedRankList)
	})

	p4.score = 11
	l.InsertOrUpdate(p4)
	expectedRankList = []string{"player 4", "player 1", "player 5", "player 6",
		"player 2", "player 3"}
	t.Run("Test List with updated valid scores", func(t *testing.T) {
		test(l, expectedRankList)
		if l.RankedPlayersCount() != 0 {
			t.Fatalf("invalid value for ranked player count - %d", l.RankedPlayersCount())
		}
	})

	p5.score = 21
	l.InsertOrUpdate(p5)
	expectedRankList = []string{"player 5", "player 4", "player 1", "player 6",
		"player 2", "player 3"}
	t.Run("Test List having assigning player with rank", func(t *testing.T) {
		test(l, expectedRankList)
		if l.RankedPlayersCount() != 1 {
			t.Fatalf("invalid value for ranked player count - %d", l.RankedPlayersCount())
		}
	})

	p1.score = 13
	l.InsertOrUpdate(p1)
	expectedRankList = []string{"player 5", "player 1", "player 4", "player 6",
		"player 2", "player 3"}
	t.Run("Test List having a player assigned a rank", func(t *testing.T) {
		test(l, expectedRankList)
	})

	p6.score = 15
	l.InsertOrUpdate(p6)
	expectedRankList = []string{"player 5", "player 6", "player 1", "player 4",
		"player 2", "player 3"}
	t.Run("Test List having a player assigned a rank - 2", func(t *testing.T) {
		test(l, expectedRankList)
		if l.RankedPlayersCount() != 1 {
			t.Fatalf("invalid value for ranked player count - %d", l.RankedPlayersCount())
		}
	})

	p4.score = 23
	l.InsertOrUpdate(p4)
	expectedRankList = []string{"player 5", "player 4", "player 6", "player 1",
		"player 2", "player 3"}
	t.Run("Test List assigning 2nd player a rank", func(t *testing.T) {
		test(l, expectedRankList)
		if l.RankedPlayersCount() != 2 {
			t.Fatalf("invalid value for ranked player count - %d", l.RankedPlayersCount())
		}
	})

	p3.score = 16
	l.InsertOrUpdate(p3)
	expectedRankList = []string{"player 5", "player 4", "player 3", "player 6",
		"player 1", "player 2"}
	t.Run("Test List having two players assigned a rank", func(t *testing.T) {
		test(l, expectedRankList)
	})

	p1.score = 17
	l.InsertOrUpdate(p1)
	expectedRankList = []string{"player 5", "player 4", "player 1", "player 3",
		"player 6", "player 2"}
	t.Run("Test List having two players assigned a rank", func(t *testing.T) {
		test(l, expectedRankList)
		if l.RankedPlayersCount() != 2 {
			t.Fatalf("invalid value for ranked player count - %d", l.RankedPlayersCount())
		}
	})

}

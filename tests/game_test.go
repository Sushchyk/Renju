package tests

import (
	model "../model"
	"testing"
)

type testPair struct {
	game model.Game
	result bool
}

var tests = []testPair{
	{ getEmptyBoard(), false },
	{ getBoardWithHorizontalLine(), true},
	{ getBoardWithDiagonalLine(), true},
	{ getBoardWithoutRows(), false},
}

func getEmptyBoard() model.Game{
	var game model.Game;
	game.Start("Test1", "Test2")
	return game
}

func getBoardWithHorizontalLine() model.Game{
	var game model.Game;
	game.Start("Test1", "Test2")
	for i := 0; i <= 4; i++{
		game.MakeTurn(i, 0)
		game.SwitchPlayer()
		game.MakeTurn(i*2, 1)
		game.SwitchPlayer()
	}
	return game
}

func getBoardWithoutRows() model.Game{
	var game model.Game;
	game.Start("Test1", "Test2")
	for i := 0; i <= 4; i++{
		game.MakeTurn(i*2, 0)
		game.SwitchPlayer()
		game.MakeTurn(i*2, 1)
		game.SwitchPlayer()
	}
	return game
}

func getBoardWithDiagonalLine() model.Game{
	var game model.Game;
	game.Start("Test1", "Test2")
	for i := 1; i <= 5; i++{
		game.MakeTurn(i, i)
		game.SwitchPlayer()
		game.MakeTurn(10, i*2)
		game.SwitchPlayer()
	}
	return game
}

func TestGame(t *testing.T) {
	for _, pair := range tests {
		v := pair.game.IsRow()
		if v != pair.result {
			t.Error(
                "expected", pair.result,
                "got", v,
			)
		}

	}
}

type numberpair struct {
    values []int
    result bool
}

var numberTests = []numberpair{
    { []int{-1, -2}, false },
    { []int{1252, 5215}, false },
    { []int{1,1}, true },
	{ []int{1,12421}, false },
}

func TestTurn(t *testing.T) {
	var game model.Game;
	game.Start("Test1", "Test2")
    for _, pair := range numberTests {
        v := game.MakeTurn(pair.values[0], pair.values[1])
        if v != pair.result {
            t.Error(
                "For", pair.values,
                "expected", pair.result,
                "got", v,
            )
        }
    }
}
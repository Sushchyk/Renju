package main


import (
	model "rendzyu/model"
	console "rendzyu/functions"
	view "rendzyu/view"
	//"fmt"
	//"fmt"
)

func BoardInitialization(height, width int) model.Board{
	board := model.Board{Height: height, Width: width}
	board.FillBoard()
	return board
}

func GameInitialization(players [2]model.Player, board model.Board) model.Game{
	game := model.Game{Players:players, Table:board, Status: true,PlayerXPos:1, PlayerYPos:5}
	return game
}

func main() {
	board := BoardInitialization(15, 15)


	player1 := model.Player{Name:"Erik", Symbol:"x"}
	player2 := model.Player{Name:"Andrej", Symbol:"o"}

	players := [2]model.Player{player1, player2}

	game := GameInitialization(players, board)
	i := 0
	for true{
		console.Clear()
		game.CurrentPlayer = game.Players[i]
		game.Table.PlaceOnBoard(game.PlayerYPos, game.PlayerXPos, game.CurrentPlayer.Symbol)
		view.DrawBoard(game.Table.Table)
		i++
		if i == 2{
			i = 0
		}
		game.MakeTurn()
	}





}
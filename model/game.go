package model

import (
	"fmt"
	console "rendzyu/functions"
)

type Game struct {
	PlayerXPos, PlayerYPos int
	Players [2]Player
	CurrentPlayer Player
	Table Board
	Status bool
}


func (g Game) RP() {
	fmt.Println("HUJ")
}


func (g Game) GoThrough(x, y int, direction string) (int,int) {
	if direction == "right"{
		for i := x; i < len(g.Table.Table); i++{
			if g.Table.Table[y][i] == " "{
				return i, y
			} else if g.Table.Table[y][i] == "☐" {
				return -1, -1
			}
		}
	} else if direction == "left" {
		for i := x; i >0; i-- {
			if g.Table.Table[y][i] == " " {
				return i, y
			} else if g.Table.Table[y][i] == "☐" {
				return -1, -1
			}
		}
	} else if direction == "up"{
		for i := y; i > 0; i--{
			if g.Table.Table[i][x] == " "{
				return x, i
			} else if g.Table.Table[i][x] == "☐" {
				return -1, -1
			}
		}
	} else if direction == "down"{
		for i := y; i < len(g.Table.Table); i++{
			if g.Table.Table[i][x] == " "{
				return x, i
			} else if g.Table.Table[i][x] == "☐" {
				return -1, -1
			}
		}
	}
	return -1, -1
}

func (g *Game) MakeTurn(){
	L:
		for true{
			x, y := g.PlayerXPos, g.PlayerYPos
			switch console.ReadKey() {
			case "w":
				if g.Table.Table[y-1][x] == "☐"{
					continue
				} else if g.Table.Table[y-1][x] == "x" || g.Table.Table[y-1][x] == "o"{
					newX, newY := g.GoThrough(x, y, "up")
					if newX == -1 && newY == -1{
						continue
					} else{
						g.PlayerXPos, g.PlayerYPos = newX, newY
						g.Table.CleanBoard(y, x)
						g.Table.PlaceOnBoard(g.PlayerYPos, g.PlayerXPos, g.CurrentPlayer.Symbol)
						console.Clear()
						g.Table.DrawBoard()
					}

				} else {
					g.Table.CleanBoard(y, x)
					g.PlayerYPos--
					g.Table.PlaceOnBoard(y - 1, x, g.CurrentPlayer.Symbol)
					console.Clear()
					g.Table.DrawBoard()
				}
			case "s":
				if g.Table.Table[y+1][x] == "☐" {
					continue
				} else if g.Table.Table[y+1][x] == "x" || g.Table.Table[y+1][x] == "o"{
						newX, newY := g.GoThrough(x, y, "down")
					if newX == -1 && newY == -1{
						continue
					} else{
						g.PlayerXPos, g.PlayerYPos = newX, newY
						g.Table.CleanBoard(y, x)
						g.Table.PlaceOnBoard(g.PlayerYPos, g.PlayerXPos, g.CurrentPlayer.Symbol)
						console.Clear()
						g.Table.DrawBoard()
					}
				}else {
					g.Table.CleanBoard(y, x)
					g.PlayerYPos++
					g.Table.PlaceOnBoard(y + 1, x, g.CurrentPlayer.Symbol)
					console.Clear()
					g.Table.DrawBoard()
				}
			case "d":
				if g.Table.Table[y][x + 1] == "☐" {
					continue
				} else if g.Table.Table[y][x + 1] == "x" || g.Table.Table[y][x + 1] == "o"{
						newX, newY := g.GoThrough(x, y, "right")
					if newX == -1 && newY == -1{
						continue
					} else{
						g.PlayerXPos, g.PlayerYPos = newX, newY
						g.Table.CleanBoard(y, x)
						g.Table.PlaceOnBoard(g.PlayerYPos, g.PlayerXPos, g.CurrentPlayer.Symbol)
						console.Clear()
						g.Table.DrawBoard()
					}
				} else {
					g.Table.CleanBoard(y, x)
					g.PlayerXPos++
					g.Table.PlaceOnBoard(y, x + 1, g.CurrentPlayer.Symbol)
					console.Clear()
					g.Table.DrawBoard()
				}
			case "a":
				if g.Table.Table[y][x - 1] == "☐" {
					continue
				} else if g.Table.Table[y][x - 1] == "x" || g.Table.Table[y][x - 1] == "o"{
						newX, newY := g.GoThrough(x, y, "left")
					if newX == -1 && newY == -1{
						continue
					} else{
						g.PlayerXPos, g.PlayerYPos = newX, newY
						g.Table.CleanBoard(y, x)
						g.Table.PlaceOnBoard(g.PlayerYPos, g.PlayerXPos, g.CurrentPlayer.Symbol)
						console.Clear()
						g.Table.DrawBoard()
					}
				} else {
					g.Table.CleanBoard(y, x)
					g.PlayerXPos--
					g.Table.PlaceOnBoard(y, x - 1, g.CurrentPlayer.Symbol)
					console.Clear()
					g.Table.DrawBoard()
				}
			case "f":
				g.Table.PlaceOnBoard(y, x, g.CurrentPlayer.Symbol)
				console.Clear()
				g.Table.DrawBoard()
				g.PlayerXPos, g.PlayerYPos = g.Table.FindStartPosition()
				break L
			}
		}
}

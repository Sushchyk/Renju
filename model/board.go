package model

import "fmt"
//
//import (
//	"fmt"
//)

type Board struct {
	Height, Width int
	Table [17][17]string
}

//func (b Board) GetTable() [][]string{
//	return b.table
//}

//var array = make([][]int, 4)
//for i:= range array {
//	array[i] = make([]int, 5)
//}

func (b *Board) CleanBoard(x, y int){
	b.Table[x][y] = " "
}

func (b *Board) PlaceOnBoard(x, y int, symbol string) {
	b.Table[x][y] = symbol
}

func (b Board) FindStartPosition() (int, int){
	var x, y int
	L:
		for i:= 1; i < len(b.Table); i++ {
			for j:= 1; j < len(b.Table); j++ {
				if b.Table[i][j] == " "{
					x, y = j, i
					break L
				}
			}
		}
	return x, y
}

func (b *Board) FillBoard() {
	for i:= 0; i < len(b.Table); i++ {
		for j:= 0; j < len(b.Table); j++ {
			if i == 0 || i == len(b.Table) - 1{
				b.Table[i][j] = "☐"
			} else if j == 0 || j == len(b.Table) - 1{
				b.Table[i][j] = "☐"
			} else {
				b.Table[i][j] = " "
			}
		}
	}
}

func (b Board) DrawBoard() {
	for i:= 0; i < len(b.Table); i++ {
		for j := 0; j < len(b.Table); j++ {
			fmt.Print(b.Table[i][j])
		}
		fmt.Print("\n")
	}
}

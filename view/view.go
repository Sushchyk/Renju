package view

import (
	"fmt"
)

func DrawBoard(table [17][17] string) {
	for i:= 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			fmt.Print(table[i][j])
		}
		fmt.Print("\n")
	}
}

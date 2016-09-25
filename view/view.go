package view

import (
	"fmt"
	"time"
	config "../game_config"
)

func MoveCursorForTurn(x int, y int, cursorSymbol string) {
	moveCursorLeft();
	fmt.Print(" ");
	goToXY(y, x);
	fmt.Print(cursorSymbol);
}

func MoveCursorAfterTurn(x int, y int, cursorSymbol string)  {
	goToXY(y, x);
	fmt.Print(cursorSymbol);
}

func DrawTurn(playerSymbol string) {
	moveCursorLeft();
	fmt.Print(playerSymbol);
}

func DrawBoard() {
	for i:= 0; i < config.HEIGHT + 2; i++ {
		for j := 0; j < config.WIDTH + 2; j++ {
			if (i == 0 || j == 0 || i == config.HEIGHT + 1 || j == config.WIDTH + 1) {
				fmt.Print(config.BORDER_SYMBOL);
			}	else {
				fmt.Print(config.EMPTY_SYMBOL);
			}
		}
		fmt.Print("\n")
	}
}

func StartGame() {
	clearScreen();
	DrawBoard();
	time.Sleep(10 * time.Millisecond);
}

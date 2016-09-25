package view

import (
	"fmt"
	"time"
	config "../game_config"
	"os/exec"
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

func ShowWinner(playerName string) {
	goToXY(config.HEIGHT + 4, 0);
	fmt.Printf("CONGRATULATIONS!\nWINNER:\n%s", playerName);
	fmt.Println("Press 'Enter' to back in main menu");
	ReadString();
}

func DrawBoard() {
	goToXY(0, 0);
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

func GetNameOfPlayer(playerNumber int) string{
	defer exec.Command("stty", "-F", "/dev/tty", "echo")
	clearScreen();
	goToXY(0, 0);
	fmt.Printf("Enter name of Player %d: ", playerNumber);
	var playerName = ReadString();
	return playerName;
}

func StartGame() {
	clearScreen();
	DrawBoard();
	time.Sleep(10 * time.Millisecond);
}

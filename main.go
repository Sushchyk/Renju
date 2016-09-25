package main


import (
	model "./model"
	view "./view"
	//"fmt"
	//"fmt"
	"os/exec"
	"os"
	config "./game_config"
)

func readKey() string{
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	return string(b)
}

func main() {
  	var game model.Game;
	game.Start();
	view.StartGame();
	currentX, currentY := config.DEFAULT_X, config.DEFAULT_Y;
	view.MoveCursorForTurn(currentX + 1, currentY + 1, game.GetCurrentPlayer().GetSymbol());
	gameContinue := true;
	for (gameContinue) {
		moveX, moveY := 0, 0;
		action := readKey();
		switch action {
		case "s":
			moveY = 1;
		case "w":
			moveY = -1;
		case "d":
			moveX = 1;
		case "a":
			moveX = -1;
		case "f":
			if(game.MakeTurn(currentX, currentY)) {
				view.DrawTurn(game.GetCurrentPlayer().GetSymbol());
				game.SwitchPlayer();
				if (game.IsRow()) {
					gameContinue = false;
					break;
				}
				currentX, currentY = game.GetPositonAfterTurn();
				view.MoveCursorAfterTurn(currentX + 1, currentY + 1, game.GetCurrentPlayer().GetSymbol());
			}
			continue;
		}
		if (action == "w" || action == "s" || action == "d" || action == "a") {
			newX, newY := game.GetFreePosition(currentX, currentY, moveX, moveY);
			if (newX != -1 && newY != -1) {
				currentY = newY;
				currentX = newX;
				view.MoveCursorForTurn(currentX + 1, currentY + 1, game.GetCurrentPlayer().GetSymbol());
			}
		}



	}

}
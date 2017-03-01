package main

import (
	model "./model"
	view "./view"
	config "./game_config"
)

func main() {
	for true {
		view.InMenu();
		playerOneName := view.GetNameOfPlayer(1);
		playerTwoName := view.GetNameOfPlayer(2);
		var game model.Game;
		game.Start(playerOneName, playerTwoName);
		view.StartGame();
		currentX, currentY := config.DEFAULT_X, config.DEFAULT_Y;
		view.MoveCursorForTurn(currentX + 1, currentY + 1, game.GetCurrentPlayer().GetSymbol());
		L:
		for true {
			moveX, moveY := 0, 0;
			action := view.ReadKey();
			switch action {
			case config.MOVE_DOWN_KEY:
				moveY = 1;
			case config.MOVE_UP_KEY:
				moveY = -1;
			case config.MOVE_RIGHT_KEY:
				moveX = 1;
			case config.MOVE_LEFT_KEY:
				moveX = -1;
			case config.FIRE_KEY:
				if (game.MakeTurn(currentX, currentY)) {
					view.DrawTurn(game.GetCurrentPlayer().GetSymbol());
					game.SwitchPlayer();
					if (game.IsRow()) {
						break L;
					}
					currentX, currentY = game.GetFreePositonAfterTurn();
					if (currentX == -1 || currentY == -1) {
						break L;

					}
					view.MoveCursorAfterTurn(currentX + 1, currentY + 1, game.GetCurrentPlayer().GetSymbol());
				}
				continue;
			}
			if (moveX != 0 || moveY != 0) {
				newX, newY := game.GetFreePosition(currentX, currentY, moveX, moveY);
				if (newX != -1 && newY != -1) {
					currentY = newY;
					currentX = newX;
					view.MoveCursorForTurn(currentX + 1, currentY + 1, game.GetCurrentPlayer().GetSymbol());
				}
			}
		}

		if (game.IsTie()) {
			view.ShowTie();
		} else {
			view.ShowWinner(game.GetWinner().GetName());
		}

	}
}
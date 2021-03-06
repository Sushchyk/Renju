package model

import (
	config "../game_config"
)

type Game struct {
	players            [2]Player;
	currentPlayerIndex int;
	winnerIndex        int;
	board              [config.HEIGHT][config.WIDTH] string;
}

func (g *Game) Start(playerOneName, playerTwoName string) {

	g.InitializeBoard()

	// initialize players:
	g.players[0].CreatePlayer(playerOneName, config.PLAYER_ONE_SYMBOL);
	g.players[1].CreatePlayer(playerTwoName, config.PLAYER_TWO_SYMBOL);
	g.currentPlayerIndex = 0;
	g.winnerIndex = -1;
}

func (g Game) GetWinner() Player {
	return g.players[g.winnerIndex];
}

func (g Game) IsTie() bool {
	return (g.winnerIndex == -1);
}

func (g *Game) makeTie() {
	g.winnerIndex = -1;
}

func (g Game) isRowHorizontally(i, j int, playerSymbol string) bool {
	for k := 0; (j + k < config.WIDTH) && k < config.ROW_LENGTH; k++ {
		if (g.board[i][j + k] != playerSymbol) {
			break;
		}
		if (k == config.ROW_LENGTH - 1) {
			return true;
		}
	}
	return false;
}

func (g Game) isRowVertically(i, j int, playerSymbol string) bool {
	for k := 0; (i + k < config.HEIGHT) && k < config.ROW_LENGTH; k++ {
		if (g.board[i + k][j] != playerSymbol) {
			break;
		}
		if (k == config.ROW_LENGTH - 1) {
			return true;
		}
	}
	return false;
}

func (g Game) isRowDiagonally(i, j int, playerSymbol string) bool {
	for k := 0; (i + k < config.HEIGHT) && (j + k < config.WIDTH) && k < config.ROW_LENGTH; k++ {
		if (g.board[i + k][j + k] != playerSymbol) {
			break;
		}

		if (k == config.ROW_LENGTH - 1) {
			return true;
		}
	}

	for k := 0; (i + k < config.HEIGHT) && (j - k >= 0) && k < config.ROW_LENGTH; k++ {
		if (g.board[i + k][j - k] != playerSymbol) {
			break;
		}

		if (k == config.ROW_LENGTH - 1) {
			return true;
		}
	}
	return false;
}

func (g *Game) IsRow() bool {
	plOneSymb := g.players[0].GetSymbol();
	plTwoSymb := g.players[1].GetSymbol();

	for i := 0; i < config.HEIGHT; i++ {
		for j := 0; j < config.WIDTH; j++ {
			if (g.isRowHorizontally(i, j, plOneSymb) || g.isRowVertically(i, j, plOneSymb) || g.isRowDiagonally(i, j, plOneSymb)) {
				g.winnerIndex = 0;
				return true;
			}

			if (g.isRowHorizontally(i, j, plTwoSymb) || g.isRowVertically(i, j, plTwoSymb) || g.isRowDiagonally(i, j, plTwoSymb)) {
				g.winnerIndex = 1;
				return true;
			}
		}
	}

	return false;
}

func (g *Game) SwitchPlayer() {
	if (g.currentPlayerIndex == 0) {
		g.currentPlayerIndex = 1;
	} else {
		g.currentPlayerIndex = 0;
	}
}
func (g Game) GetCurrentPlayer() Player {
	return g.players[g.currentPlayerIndex];
}

func (g *Game) InitializeBoard() {
	for i := 0; i < config.HEIGHT; i++ {
		for j := 0; j < config.WIDTH; j++ {
			g.board[i][j] = config.EMPTY_SYMBOL;
		}
	}
}

func (g Game) GetFreePosition(x, y, moveX, moveY int) (int, int) {
	nextPosX := x + moveX;
	nextPosY := y + moveY;
	for (nextPosX >= 0  && nextPosY >= 0 && nextPosX < config.WIDTH && nextPosY < config.HEIGHT) {
		if (g.board[nextPosY][nextPosX] == config.EMPTY_SYMBOL) {
			return nextPosX, nextPosY;
		}
		nextPosX += moveX;
		nextPosY += moveY;
	}

	return -1, -1; // out of range
}

func (g Game) GetFreePositonAfterTurn() (int, int) {
	for i := 0; i < config.HEIGHT; i++ {
		for j := 0; j < config.WIDTH; j++ {
			if (g.board[i][j] == config.EMPTY_SYMBOL) {
				return j, i;
			}
		}

	}
	g.makeTie();
	return -1, -1;
}

func (g *Game) MakeTurn(x, y int) bool {
	if (y < 0 || y >= config.HEIGHT || x < 0 || x >= config.WIDTH) {
		return false
	}
	if (g.board[y][x] == config.EMPTY_SYMBOL) {
		g.board[y][x] = g.GetCurrentPlayer().GetSymbol();
		return true;
	}
	return false;

}

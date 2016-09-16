package view

import (
	"fmt"
	"time"
	"strconv"
)

func DrawBorders() {

	goToXY(2, 0)
	fmt.Println("ABCDEFGHIJKLMNO");
	goToXY(2, 16)
	fmt.Println("###############");

	for i := 0; i < 15; i++ {
		goToXY(0, i + 1)
		fmt.Println(strconv.Itoa(i + 1))
	}

	for j := 0; j < 15; j++ {
		goToXY(17, j + 1)
		fmt.Println("#")
	}
	goToXY(0, 20)

}

func NewGame() {
	clearScreen()
	DrawBorders()
	time.Sleep(10 * time.Millisecond)
}
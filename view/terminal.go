package view

import (
	"strconv"
	"os/exec"
	"os"
)

func goToXY(x ,y int){
	tput("cup", strconv.Itoa(y), strconv.Itoa(x)) // an initial position
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func tput(args ...string) error {
	cmd := exec.Command("tput", args...)
	cmd.Stdout = os.Stdout
	return cmd.Run();
}

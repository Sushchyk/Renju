package view

import (
	"strconv"
	"os/exec"
	"os"
)
func ReadKey() string{
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	return string(b)
}

func goToXY (x int, y int) error {
	cmd := exec.Command("tput","cup", strconv.Itoa(x), strconv.Itoa(y))
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func clearScreen() error{
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func moveCursorLeft() {
	tput("cub1");
}

func tput(args ...string) error {
	cmd := exec.Command("tput", args...)
	cmd.Stdout = os.Stdout
	return cmd.Run();
}

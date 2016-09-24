package functions

import (
	"os"
    "os/exec"
    "strconv"
)

func GoToXY (x int, y int) error {
    cmd := exec.Command("tput","cup", strconv.Itoa(x), strconv.Itoa(y))
    cmd.Stdout = os.Stdout
    return cmd.Run()
}

func Clear() error{
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func ReadKey() string{
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    // do not display entered characters on the screen
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	return string(b)
}
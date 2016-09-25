package view

import (
	"fmt"
	"os"
	"os/exec"
)

func readKey() string{
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	return string(b)
}

func drawMenu()  {
	clearScreen();
	goToXY(0, 0);
	fmt.Print("RENJU\n\n");
	fmt.Print("\t1.New game\n\t2.Help\n\t3.About\n\t4.Exit");
	goToXY(1, 0);
}

func InMenu() {
	drawMenu();
	 for true {
		 switch readKey() {
		 case "1":
			 return;
		 case "2":
			 help();
		 case "3":
			 about();
		 case "4":
			 exit();
		 }
		 drawMenu();
	 }
}

func help () {
	clearScreen();
	goToXY(0, 0);
	fmt.Print("HELP\n\n");
	fmt.Println("\tUse 'w', 'a', 's', 'd' to move, and 'f' to make turn");
	fmt.Println("\tPress any key to back in main menu")
	readKey();
}

func about() {
	clearScreen();
	goToXY(0, 0);
	fmt.Print("ABOUT\n\n");
	fmt.Println("\tMade by A.Sushchyk, E.Gimiranov (KP-42)");
	fmt.Println("\tPress any key to back in main menu")
	readKey();
}

func exit() {
	clearScreen();
	goToXY(0, 0);
	fmt.Println("GOODBYE!");
	os.Exit(0);
}

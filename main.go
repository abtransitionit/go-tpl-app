package main

import (
	"context"
	"strconv"
	"strings"

	"github.com/abtransitionit/go-log/log"
	"github.com/abtransitionit/go-res/exectx"

	"bufio"
	"fmt"
	"os"
)

func main() {
	// define the Logger Instance
	logger, err := log.GetLogger()
	if err != nil {
		panic("Failed to initialize application logging system")
	}

	// log
	logger.Info("Starting application...")

	// define the execution context
	ctx := exectx.ExeCtx{
		Ctx:    context.Background(),
		Logger: logger,
	}

	// get the menu
	menu := GetMenuItem()

	// reader used for CLI input
	reader := bufio.NewReader(os.Stdin)

	// Menu loop
	for {
		// display menu with auto-generated numbers
		fmt.Println("\nPlease choose an option:")
		fmt.Println("0. Exit")

		for i, item := range menu {
			fmt.Printf("%d. %s\n", i+1, item.Name)
		}

		// prompt user
		fmt.Print("Enter choice: ")

		// read user input from stdin
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error")
			continue
		}

		// convert user input string to integer choice
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		// exit handling
		if choice == 0 {
			fmt.Println("Bye")
			logger.Info("Exiting application")
			break
		}

		// validate user choice
		if choice < 1 || choice > len(menu) {
			fmt.Println("Invalid input")
			continue
		}

		// get menu item, based on user choice
		selectedMenuItem := menu[choice-1]

		// log
		fmt.Printf("you selected menu %d (%s)\n", choice, selectedMenuItem.Name)

		// execute selected menu item function
		selectedMenuItem.Fn(ctx)
	}
}

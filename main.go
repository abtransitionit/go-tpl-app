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

	// get the menu:Listkeys ([]string)
	menuKeys := Menu()

	// reader used for CLI input
	reader := bufio.NewReader(os.Stdin)

	// Menu loop
	for {
		// display menu:Listkeys with auto-generated numbers
		fmt.Println("\nPlease choose an option:")
		fmt.Println("0. Exit")

		for i, keyName := range menuKeys {
			fmt.Printf("%d. %s\n", i+1, keyName)
		}

		// prompt user
		fmt.Print("Enter choice: ")

		// read user input from stdin
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error")
			continue
		}

		// convert user input choice (string) to an integer
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		// handle exit
		if choice == 0 {
			fmt.Println("Bye")
			logger.Info("Exiting application")
			break
		}

		// handle other user choice
		if choice < 1 || choice > len(menuKeys) {
			fmt.Println("Invalid input")
			continue
		}

		// get menu item, based on user choice
		selectedMenuItem := menuKeys[choice-1]

		// log
		ctx.Logger.Infof("valid user choice is : %d", choice)

		// print
		fmt.Printf("you selected choice %d (%s)\n", choice, selectedMenuItem)

		// get fn from registry
		fnToExecute, err := fnList.Fetch(selectedMenuItem)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// execute selected menu item function
		fnToExecute(ctx)
		// selectedMenuItem.Fn(ctx)
		// selectedMenuItem.Fn(ctx)
	}
}

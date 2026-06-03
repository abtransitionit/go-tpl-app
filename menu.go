package main

import (
	"fmt"

	"github.com/abtransitionit/go-res/exectx"
)

// Menu returns the list of available commandsfunction
func Menu() []MenuItem {
	return []MenuItem{
		Item("Start service", startService),
		Item("Show status", stopService),
		Item("Stop service", stopService),
	}
}

func startService(ec exectx.ExeCtx) {
	ec.Logger.Info("Service started")
	fmt.Println("Service started")
}

func showStatus(ec exectx.ExeCtx) {
	ec.Logger.Info("Service started")
	fmt.Println("Service started")
}

func stopService(ec exectx.ExeCtx) {
	ec.Logger.Info("Service stopped")
	fmt.Println("Service stopped")
}

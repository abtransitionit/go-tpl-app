package main

import (
	"fmt"

	"github.com/abtransitionit/go-res/exectx"
)

type MenuItem struct {
	Name string
	Fn   func(exectx.ExeCtx)
}

// GetMenu defines all application commands
func GetMenuItem() []MenuItem {
	return []MenuItem{
		{
			Name: "Start service",
			Fn: func(ec exectx.ExeCtx) {
				ec.Logger.Info("Service started")
				fmt.Println("Service started")
			},
		},
		{
			Name: "Show status",
			Fn: func(ec exectx.ExeCtx) {
				ec.Logger.Info("Checking status")
				fmt.Println("Status: OK")
			},
		},
	}
}

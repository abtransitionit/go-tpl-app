package main

import (
	"fmt"

	"github.com/abtransitionit/go-core/registry"
	"github.com/abtransitionit/go-res/exectx"
)

// FnUnderTest define the function signature of the functions that will be added to the registry
type FnUnderTest func(exectx.ExeCtx)

var fnList = registry.NewRegistry[FnUnderTest]("menu", map[string]FnUnderTest{
	"start Service": startService,
	"status Show":   showStatus,
	"stop Service":  stopService,
})

func startService(ec exectx.ExeCtx) {
	ec.Logger.Info("starting Service ")
	fmt.Println("Service started")
}

func showStatus(ec exectx.ExeCtx) {
	ec.Logger.Info("Checking status")
	fmt.Println("Service started")
}

func stopService(ec exectx.ExeCtx) {
	ec.Logger.Info("Stopping Service")
	fmt.Println("Service stopped")
}

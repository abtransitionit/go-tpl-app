package main

import "github.com/abtransitionit/go-res/exectx"

// Item creates a MenuItem - helper function
func Item(name string, fn func(exectx.ExeCtx)) MenuItem {
	return MenuItem{
		Name: name,
		Fn:   fn,
	}
}

type MenuItem struct {
	Name string
	Fn   func(exectx.ExeCtx)
}

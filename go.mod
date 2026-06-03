module github.com/abtransitionit/go-app-test

// GO Toolchain
go 1.26

// direct dependencies added by me
require (
	github.com/abtransitionit/go-core v0.0.0
	github.com/abtransitionit/go-log v0.0.0
	github.com/abtransitionit/go-res v0.0.0
)

require (
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
)

// dev mode dependencies added by me
// removes by CI at tag step
// simplify development on active dev when working on several projects with inetr dependencies
replace (
	github.com/abtransitionit/go-core => ../go-core
	github.com/abtransitionit/go-log => ../go-log
	github.com/abtransitionit/go-res => ../go-res
)

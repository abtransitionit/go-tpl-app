module github.com/abtransitionit/go-app-test

// GO Toolchain
go 1.26

// dev mode
// removes by CI at tag step
// simplify development on active dev when working on several projects with inetr dependencies
replace github.com/abtransitionit/go-core => ../go-core

replace github.com/abtransitionit/go-yfm => ../go-yfm

replace github.com/abtransitionit/go-logx => ../go-logx

require (
	github.com/abtransitionit/go-core v0.0.0-00010101000000-000000000000
	github.com/abtransitionit/go-logx v0.0.0
	github.com/abtransitionit/go-yfm v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go-v2 v1.41.2 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.5 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.96.2 // indirect
	github.com/aws/smithy-go v1.24.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

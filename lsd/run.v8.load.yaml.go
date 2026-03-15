package lsd

import (
	"context"
	"fmt"
	"time"

	"github.com/abtransitionit/go-yfm"
)

// load a yaml from an URI into an instance of a struct
//
// Notes:
//
//   - use the function yfm.LoadFromUri
func run_load_yaml_v8() error {
	// define context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// define/get the yaml structure
	type Cfg struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	}

	// define/get the yaml location/uri
	resUri := "file:///tmp/config.yaml"
	// uri := "https://example.com/config.yaml"
	// uri := "s3://my-bucket/config.yaml"

	// create the yaml file
	// printf "name: my-app\nport: 8080" > /tmp/config.yaml

	// load the data from the yaml into the instance
	cfg, err := yfm.Load[Cfg](ctx, resUri)
	if err != nil {
		return err
	}

	// print the var
	fmt.Printf("\n%v\n", cfg)

	// handle success
	return nil
}

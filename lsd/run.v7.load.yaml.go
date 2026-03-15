// Load Save Delete
package lsd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abtransitionit/go-core/resource"
	"github.com/abtransitionit/go-yfm"
)

// load a yaml from an URI into an instance of a struct
//
// Notes:
//
//   - use the function yfm.Load
func Run_load_yaml_v7() error {
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

	// get/define a loader for that yaml
	yamlLoader, err := resource.NewLoader(resUri)
	if err != nil {
		panic(err)
	}

	// load the yaml content into the instance [of a struct]
	cfg, err := yfm.LoadFromLoader[Cfg](ctx, yamlLoader)
	if err != nil {
		log.Fatal(err)
	}

	// print the var
	fmt.Printf("\n%v\n", cfg)

	// handle success
	return nil
}

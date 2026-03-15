package lsd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abtransitionit/go-yfm"
)

func run_save_yaml_v7() error {
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

	// // get/define a saver for that yaml
	// yamlSaver, err := resource.NewSaver(resUri)
	// if err != nil {
	// 	panic(err)
	// }

	// define the data to save into that yaml
	var cfg Cfg
	cfg.Name = "titi"
	cfg.Port = 8088

	// save the data into the yaml
	err := yfm.Save(ctx, resUri, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// print the yaml data
	fmt.Printf("\n%v\n", cfg)

	// handle success
	return nil
}

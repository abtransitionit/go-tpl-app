package logx

import (
	"fmt"

	"github.com/abtransitionit/go-logx/logx"
)

func Sfield_01() error {

	// check it is called
	fmt.Println("La vie est trop belle")

	// get a structured field
	sf01f := logx.String("user", "max")
	sf02 := logx.Int("age", 42)

	// display it
	fmt.Println(sf01f)
	fmt.Println(sf02)

	// handle success
	return nil
}

package logx

import (
	"fmt"

	"github.com/abtransitionit/go-logx/logx"
)

func Nominal() error {

	// check it is called
	fmt.Println("La vie est trop belle")

	// get an instance
	// logger, err := logx.NewWithCfg(logx.Cfg{Driver: logx.DriverStd})
	logger, err := logx.GetLogger()
	if err != nil {
		return err
	}

	// use it
	logger.Info("server started")
	// logger.Error("database failed")

	// handle success
	return nil
}

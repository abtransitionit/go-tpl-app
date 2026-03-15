package logx

import (
	"fmt"

	"github.com/abtransitionit/go-logx/logx"
	"github.com/abtransitionit/go-logx/sfield"
)

func Std() error {

	// check it is called
	fmt.Println("La vie est trop belle")

	// get an instance
	// logger, err := logx.NewWithCfg(logx.Cfg{Driver: logx.DriverStd})
	logger, err := logx.GetLogger()
	logger.Info("server started")
	logger.Infof("server started on port %d", 8080)
	logger.Infosf("user login",
		sfield.String("user", "max"),
		sfield.Int("age", 42),
	)

	logger.Error("disk full")
	logger.Warnf("memory usage at %d%%", 92)
	logger.Debugsf("debug data", sfield.String("debug", "xyz"))
	// logger, err := logx.New(logx.DriverZap)
	if err != nil {
		return err
	}

	// use it
	logger.Info("server started")
	// logger.Error("database failed")

	// handle success
	return nil
}

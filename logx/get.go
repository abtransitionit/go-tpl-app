package logx

import (
	"fmt"

	"github.com/abtransitionit/go-logx/logx"
	"github.com/abtransitionit/go-logx/sfield"
)

func Get() error {

	// check it is called
	fmt.Println("starting code")

	// // logger config
	// if err := os.Setenv("APP_LOGGER_ENV", "dev"); err != nil {
	// 	panic(err)
	// }
	// if err := os.Setenv("APP_LOGGER_DRIVER", "zap"); err != nil {
	// 	panic(err)
	// }
	// get logger
	logger, _ := logx.GetLogger()
	logger.Info("server started")
	logger.Infof("server started on port %d", 8080)
	logger.Infosf("user login",
		logx.String("user", "max"),
		logx.Int("age", 42),
	)

	logger.Error("disk full")
	logger.Warnf("memory usage at %d%%", 92)
	logger.Debugsf("debug data", sfield.String("debug", "xyz"))

	// use it
	logger.Info("server started")
	// logger.Error("database failed")

	// handle success
	return nil
}

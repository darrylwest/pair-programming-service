// logger
//
// @author darryl.west@eaby.com
// @created 2018-04-08 09:26:59

package edit

import (
	"fmt"
	"os"

	"github.com/darrylwest/cassava-logger/logger"
)

var log *logger.Logger

// CreateLogger create a new console logger; use log.SetLevel( logger.WarnLevel )
func CreateLogger() *logger.Logger {
	if log == nil {
		handler, err := logger.NewStreamHandler(os.Stdout)

		if err != nil {
			fmt.Printf("%s\n", err)
			panic("logger could not be created...")
		}

		log = logger.NewLogger(handler)
	}

	return log
}

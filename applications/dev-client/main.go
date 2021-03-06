package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
	"os"
)

var hceCard types.HCECard

func main() {

	initLog()

	doUI()
}

func initLog() {

	log.SetFormatter(&log.JSONFormatter{})

	f, err := os.OpenFile("output.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {

		fmt.Println(err.Error())
	}

	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Debug("Log is initialised.")
}

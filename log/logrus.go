package main

import (
	// "fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

// 一直都有前面的前缀，不知道怎么去掉
func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Debug("test deubg")
	log.Warning("test warning")
	log.Error("error")

	//和上边的一样结果
	log.Debugln("test deubg")
	log.Warningln("test warning")
	log.Errorln("error")

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	logger := log.New()
	logger.Debug("logger debug")
	logger.Error("logger.Error")
	logger.Errorf("loggerf error")

}

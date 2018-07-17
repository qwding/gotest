package main

import (
	log "github.com/skoo87/log4go"
)

func main() {
	log.SetLevel(3)

	var name = "skoo"
	log.Debug("log4go by %s", name)
	log.Info("log4go by %s", name)
	log.Warn("log4go by %s", name)
	log.Error("log4go by %s", name)
	log.Fatal("log4go by %s", name)
}

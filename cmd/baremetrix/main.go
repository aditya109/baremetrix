package main

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/config"
	"time"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/play"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/simulator"
	logCfg "bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/logger"
	logger "github.com/sirupsen/logrus"
)

func main() {
	var timeStamp = time.Now().Format("2006-01-02_15:04:05")
	// initializing a configuration object
	configuration, err := config.GetConfiguration()
	if err != nil {
		logger.Error(err)
	}
	// initialize levelled logger
	err = logCfg.InitializeLogging(configuration, timeStamp)
	if err != nil {
		logger.Error(err)
	}
	// get list of play
	plays, err := play.GetPlays(configuration)
	if err != nil {
		logger.Error(err)
	}
	// triggering run in simulator with play
	err = simulator.Run(configuration, plays, timeStamp)
	if err != nil {
		logger.Error(err)
	}
}

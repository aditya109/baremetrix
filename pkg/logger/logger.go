package logger

import (
	"io"
	"os"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/helper"
	logger "github.com/sirupsen/logrus"
)

// InitializeLogging returns a configured logger object
func InitializeLogging(config *models.Config, timeStamp string) error {
	if err := ConfigureLogger(config.Instance.Specs.LevelledLog, timeStamp); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// ConfigureLogger configures logger based on logging configuration present in config.*.json
func ConfigureLogger(loggingConfig models.LogSpecs, timeStamp string) error {
	// declaring writers to store all the enabled the io writers
	var writers []io.Writer
	if loggingConfig.EnableLoggingToFile { // if enabled, configuring logger to log into file by filespecs of logging configuration
		logFileName, err := helper.GetFormattedFileName(models.SummarySpecificDirectives{
			FileSpecs:           loggingConfig.FileSpecs[0],
			ShouldUseDirectives: false,
			TimeStamp:           timeStamp,
		})
		if err != nil {
			logger.Errorf("error in getting filename: %v", err)
			return err
		}
		f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			logger.Errorf("error opening file: %v", err)
			return err
		}
		writers = append(writers, f)
	}

	if loggingConfig.EnableLoggingToStdout { // if enabled, configuring logger to log into stdout, according to logging configuration
		writers = append(writers, os.Stdout)
	}

	mw := io.MultiWriter(writers...)
	logger.SetOutput(mw)

	if loggingConfig.OutputFormatter == "json" { // configuring log-syntax type format - json/text
		logger.SetFormatter(&logger.JSONFormatter{})
	} else {
		logger.SetFormatter(&logger.TextFormatter{
			DisableColors: !loggingConfig.EnableColors,
			FullTimestamp: loggingConfig.EnableFullTimeStamp,
		})
	}
	return nil
}

package config

import (
	"encoding/json"
	"io/ioutil"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	helper "bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/helper"
	logger "github.com/sirupsen/logrus"
)

// GetConfiguration retrieves configuration from config file
func GetConfiguration() (*models.Config, error) {
	// declaring a config object
	var config = models.Config{}
	// getting the absolute file path of the config file
	var configFilePath, err = helper.GetAbsolutePath("/config/config.json")
	if err != nil {
		logger.Error(err)
		return &models.Config{}, err
	}

	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		logger.Error(err)
		return &models.Config{}, err
	}

	// storing the content of config file in a config object
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		logger.Error(err)
		return &models.Config{}, err
	}
	return &config, nil
}

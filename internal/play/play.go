package play

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/helper"
	logger "github.com/sirupsen/logrus"
)

// GetPlays reads the play files and provides a list of play-type objects.
func GetPlays(config *models.Config) ([]models.Play, error) {
	// declaring a list of play
	var plays []models.Play
	var playfilePaths []string

	for _, play := range config.Instance.Specs.PlaySpecs.FileSpecs {
		if play.ShouldRun {
			// extracting values off file-spec object
			var playParentDirectory = play.ContainerDirectory
			var playFileNames = play.Name
			var playFileExtension = play.Extension

			for _, filename := range playFileNames { // extracting content from play_*.json files
				var path = fmt.Sprintf("%s/%s%s", playParentDirectory, filename, playFileExtension)
				// getting absolute file path for each play file
				var playFilePath, err = helper.GetAbsolutePath(path)
				if err != nil {
					logger.Error(err)
					return []models.Play{}, err
				}
				var play models.Play
				file, err := ioutil.ReadFile(playFilePath)
				if err != nil {
					logger.Error(err)
					return []models.Play{}, err
				}
				// attempting to store parsed content from file on the play object
				err = json.Unmarshal(file, &play)
				if err != nil {
					logger.Error(err)
					return []models.Play{}, err
				}
				// append play object to list of play
				plays = append(plays, play)
				playfilePaths = append(playfilePaths, path)
			}
		}
	}
	logger.Infof("# playfiles selected for load-testing: %d", len(plays))
	for idx, filename := range playfilePaths {
		logger.Infof("%d.     %s", idx+1, filename)
	}
	return plays, nil
}

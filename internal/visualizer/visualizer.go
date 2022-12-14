package visualizer

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/visualizer/graphs"
	logger "github.com/sirupsen/logrus"
)

// CreatePlaySummaryVisualizations is the controller for graph generated based on summary list.
func CreatePlaySummaryVisualizations(config *models.Config, summary []models.Summary, play models.Play, timeStamp string, iteration int) error {
	var visCfg = config.Instance.Specs.VisualizationSpecs
	var visFileSpecs = visCfg.FileSpecs

	graphType, _ := utils.FindItemFromListWithKey(visCfg.GraphTypes, constants.RpmVersusLatency).(models.GraphType)
	if graphType.IsEnabled {
		err := graphs.GenerateRPMVersusLatencyGraph(visFileSpecs, graphType, summary, play, timeStamp, iteration)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	graphType, _ = utils.FindItemFromListWithKey(visCfg.GraphTypes, constants.RpmVersusDelayedResponseCount).(models.GraphType)
	if graphType.IsEnabled {
		err := graphs.GenerateRPMVersusDelayedResponseCountGraph(visFileSpecs, graphType, summary, play, timeStamp, iteration)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return nil
}

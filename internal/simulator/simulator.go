package simulator

import (
	"strconv"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/simulator/sidecar"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/simulator/executor"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/summary"
	logger "github.com/sirupsen/logrus"
)

var directives models.SummarySpecificDirectives
var playSummary []models.Summary

// Run simulates the scenarios from the play.
func Run(config *models.Config, plays []models.Play, timeStamp string) error {
	for _, play := range plays { // iterating through the play
		logger.Info("===============================================================")
		logger.Infof("PLAY START: %s id: %s", play.Name, play.Id)
		for i := 0; i < play.Iterations; i++ {
			// initialize a summary file
			directives = models.SummarySpecificDirectives{
				Tenant:                  play.Tenant,
				PlayName:                play.Name,
				FileSpecs:               config.Instance.Specs.SummarySpecs.FileSpecs[0],
				ShouldUseDirectives:     true,
				TimeStamp:               timeStamp,
				Iteration:               strconv.Itoa(i + 1),
				ShouldUseGraphIndicator: false,
				GraphType:               "",
			}
			err := summary.WriteSummaryToFile(config, []string{}, directives, true)
			if err != nil {
				logger.Error(err)
				return err
			}
			logger.Info("summary directory created.")
			logger.Infof("ITERATION START: %d", play.Iterations)
			if play.Acts != nil && len(play.Acts) != 0 {
				playSummary = executor.ExecuteActs(config, play, directives, i)
			}
			if play.Flows != nil && len(play.Flows) != 0 {
				playSummary = executor.ExecuteFlows(config, play, directives, i)
			}

			err = sidecar.CompleteReportingAndSummaryVisualizations(config, playSummary, play, directives, timeStamp, i)
			if err != nil {
				logger.Error(err)
				return err
			}
		}
	}
	return nil
}

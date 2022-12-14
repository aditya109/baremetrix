package summary

import (
	"encoding/csv"
	"os"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/helper"
	logger "github.com/sirupsen/logrus"
)

var summaryFileName string

// WriteSummaryToFile stores the summaries and writes into the summary
func WriteSummaryToFile(config *models.Config, params []string, directives models.SummarySpecificDirectives, doPrep bool) error {
	var (
		summaryFile   *os.File
		summaryWriter *csv.Writer
		err           error
	)
	if doPrep {
		headers := []string{
			constants.PlayId,
			constants.PlayName,
			constants.PlayIteration,
			constants.Name,
			constants.TargetId,
			constants.RateOfRequestsPerSecond,
			constants.DurationInSeconds,
			constants.VegetaTimeoutInMilliseconds,
			constants.KeepPersistentConnections,
			constants.LoadRunDuration,
			constants.TotalNumberOfRequests,
			constants.TargetDistribution,
			constants.RateOfRequestsPerMinute,
			constants.LatencyAtMeanUsage,
			constants.LatencyAt50Usage,
			constants.LatencyAt95Usage,
			constants.LatencyAt99Usage,
			constants.LatencyAtMaxUsage,
			constants.SuccessRatio,
			constants.RequestCountOverExpectedLatency,
			constants.StatusCodes,
		}

		summaryFileName, err = helper.GetFormattedFileName(directives)
		if err != nil {
			logger.Errorf("error in getting filename: %v", err)
			return err
		}
		params = headers
	}
	summaryFile, err = os.OpenFile(summaryFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Errorf("error opening file: %v", err)
		return err
	}
	summaryWriter = csv.NewWriter(summaryFile)
	err = summaryWriter.Write(params)
	if err != nil {
		logger.Errorf("error in writing header: %v", err)
		return err
	}
	summaryWriter.Flush()
	summaryFile.Close()
	return nil
}

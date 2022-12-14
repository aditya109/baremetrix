package sidecar

import (
	"fmt"
	"time"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/analyser"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/visualizer"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/prerequisite"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/summary"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

// AssembleOptions provides options object for an act part of the playfile
// based vegeta parameters
func AssembleOptions(params models.Vegeta) models.LoadTestingParameters {
	// extracting vegeta parameters from step
	return models.LoadTestingParameters{
		RateOfRequests:        params.RateOfRequests,
		DurationInSeconds:     params.DurationInSeconds,
		DuplicacyPercentage:   params.DuplicacyPercentage,
		TimeoutInMilliseconds: params.TimeoutInMilliSeconds,
		KeepAlive:             params.KeepAlive,
		Options: models.Options{
			RateOfRequests:      vegeta.Rate{Freq: params.RateOfRequests, Per: time.Second},
			Duration:            time.Duration(params.DurationInSeconds) * time.Second,
			RangeOfRand:         int64(params.RateOfRequests * params.DurationInSeconds),
			DuplicacyPercentage: params.DuplicacyPercentage,
		},
	}
}

// LogData stores and bulk write
func LogData(config *models.Config, directives models.SummarySpecificDirectives, data []string) error {
	var err = summary.WriteSummaryToFile(config, data, directives, false)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// GetPrerequisiteData conditionally evaluates Act API type, and provides the result.
func GetPrerequisiteData(config *models.Config, act models.Act) map[string][]string {
	var urls = make(map[string][]string)
	if len(act.PreOpSequence) != 0 {
		requirement := (act.Vegeta.RateOfRequests * act.Vegeta.DurationInSeconds) / len(act.PreOpSequence)
		for _, r := range act.PreOpSequence {
			switch r {
			case constants.LaGreenvnCreateAllocation:
				if act.Api == constants.LaGreenvnGetAllocationDetails {
					urls[r] = prerequisite.PreallocateGreenVNIds(config, act, 1, r)
				} else {
					urls[r] = prerequisite.PreallocateGreenVNIds(config, act, requirement, r)
				}
			case constants.LaGreenpinCreatePinAllocation:
				urls[r] = prerequisite.PreallocateGreenVNIds(config, act, requirement, r)
			}
		}
	}
	return urls
}

// CompleteReportingAndSummaryVisualizations prepares the conclusion and graph generation work.
func CompleteReportingAndSummaryVisualizations(config *models.Config,
	playSummary []models.Summary,
	play models.Play,
	directives models.SummarySpecificDirectives,
	timeStamp string,
	iteration int) error {
	conclusionBlock, err := prepareConclusionForReporting(*config, playSummary, play)
	if err != nil {
		logger.Error(err)
		return err
	}
	for _, row := range conclusionBlock {
		err = LogData(config, directives, row)
		if err != nil {
			logger.Error(err)
		}
	}
	err = visualizer.CreatePlaySummaryVisualizations(config, playSummary, play, timeStamp, iteration+1)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// CalculateCustomSuccessRatio calculates success ratio on the following formula:
// success ratio = (total-5xx)/total
func CalculateCustomSuccessRatio(codeToCountMap map[uint16]int) float64 {
	type5xxStatus := 0
	total := 0
	for status, count := range codeToCountMap {
		if status >= 500 {
			type5xxStatus += count
		}
		total += count
	}
	return (float64(total) - float64(type5xxStatus)) / float64(total)
}

// GetTargetDistributionAsString provides the target distribution for flows.
func GetTargetDistributionAsString(config *models.Config, targetDistribution []int, rpm int, flowId string) string {
	var targetPercentageDistribution string
	var flow = utils.FindItemFromListWithKey(config.Instance.Specs.Flow.FlowTypes, constants.FLOW_SCENARIO_TYPE_1).(models.FlowType).Flow
	for idx, step := range flow {
		if len(targetDistribution) >= (idx + 1) {
			targetPercentageDistribution = targetPercentageDistribution + fmt.Sprintf("%s: %0.2f%% | ", step, float64(targetDistribution[idx]*100/rpm))
		}
	}
	return targetPercentageDistribution
}

// prepareConclusionForReporting prepares the conclusion from accumulated play summary for logging
func prepareConclusionForReporting(config models.Config, playSummary []models.Summary, play models.Play) ([][]string, error) {
	conclusion, err := analyser.AnalysePlaySummary(config, playSummary, play)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return [][]string{
		{},
		{
			constants.ConclusionSectionHeading,
		},
		{
			constants.TargetId, conclusion.TargetAPIOrFlowId,
		},
		{
			constants.TenantId, conclusion.TenantId,
		},
		{
			constants.ActDuration, conclusion.ActDuration,
		},
		{
			constants.ThresholdLatency, conclusion.ThresholdLatency,
		},
		{
			constants.BreakingRpm, conclusion.BreakingRPMs,
		},
		{
			constants.MaxRpmTestedSuccessfully, conclusion.MaxRPMTestedSuccessfully,
		},
		{
			constants.LatencyMeanAtMaxRpm, conclusion.LatencyMeanAtMaxRPM,
		},
		{
			constants.LatencyP95AtMaxRpm, conclusion.LatencyP95AtMaxRPM,
		},
		{
			constants.SuccessRatioAtMaxRpm, conclusion.SuccessRatioAtMaxRPM,
		},
		{
			constants.VegetaTimeout, conclusion.VegetaTimeout,
		},
		{
			constants.InfraDetailsHeader,
		},
		{
			constants.StagingEnvInfo, config.Instance.Env.InfraDetails,
		},
		{},
	}, nil
}

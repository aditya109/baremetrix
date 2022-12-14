package analyser

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"fmt"
	"strconv"
	"time"
)

// this value is in milliseconds
var expectedLatency time.Duration

// AnalysePlaySummary provides the summary analysis which can be used to look over the metadata values.
func AnalysePlaySummary(config models.Config, summary []models.Summary, play models.Play) (models.Conclusion, error) {
	expectedLatency = time.Duration(config.Instance.Specs.SummarySpecs.ExpectedLatencyInMilliseconds) * time.Millisecond
	timeout, err := strconv.Atoi(summary[0].VegetaTimeoutInMilliSeconds)
	if err != nil {
		return models.Conclusion{}, err
	}
	maxRPMTestedSuccessfullySummary := getMaxRPMTestedSuccessfully(summary)

	var conclusion = models.Conclusion{
		TargetAPIOrFlowId:        summary[0].Id,
		TenantId:                 play.Tenant,
		ActDuration:              getPlayDuration(summary),
		ThresholdLatency:         expectedLatency.String(),
		BreakingRPMs:             getBreakingRPMs(summary),
		MaxRPMTestedSuccessfully: fmt.Sprint(maxRPMTestedSuccessfullySummary.RPM),
		LatencyP95AtMaxRPM:       maxRPMTestedSuccessfullySummary.LatencyAt95Usage.String(),
		LatencyMeanAtMaxRPM:      maxRPMTestedSuccessfullySummary.LatencyAtMeanUsage.String(),
		SuccessRatioAtMaxRPM:     fmt.Sprintf("%.3f", maxRPMTestedSuccessfullySummary.SuccessRatio),
		VegetaTimeout:            (time.Duration(timeout) * time.Millisecond).String(),
		StagingEnvironmentInfo:   config.Instance.Env.InfraDetails,
	}
	return conclusion, nil
}

// getBreakingRPMs retrieves the set of Breaking RPMs from the act summaries in form of string
// else returns NoBreakingRPMEncountered message 
func getBreakingRPMs(summary []models.Summary) string {
	var breakingRPMIndices string
	indices := findBreakingRPMIndices(summary)
	if len(indices) == 0 {
		// no breaking rpm encountered during the play
		breakingRPMIndices = constants.NoBreakingRPMEncountered
	} else {
		breakingRPMIndices = ""
		if len(indices) == 1 {
			breakingRPMIndices = fmt.Sprint(summary[indices[0]].RPM)
		}

		for _, val := range indices {
			breakingRPMIndices += fmt.Sprintf("%d,", summary[val].RPM)
		}
	}
	return breakingRPMIndices
}

// findBreakingRPMIndices finds the indices of those summaries where threshold breakforth is encountered.
func findBreakingRPMIndices(summary []models.Summary) []int {
	var indices []int
	for idx, e := range summary {
		if isCurrentRPMBreaking(e) {
			indices = append(indices, idx)
		}
	}
	return indices
}

// isCurrentRPMBreaking check whether the current summary has encountered threshold breakforth.
func isCurrentRPMBreaking(summary models.Summary) bool {
	if summary.LatencyAt50Usage > expectedLatency || summary.LatencyAt95Usage > expectedLatency || summary.LatencyAtMeanUsage > expectedLatency {
		return true
	}
	return false
}

// getPlayDuration gets total load run duration for the entire play
// in string format.
func getPlayDuration(summary []models.Summary) string {
	var duration time.Duration
	for _, actSummary := range summary {
		duration += actSummary.LoadRunDuration
	}
	return duration.String()
}

// getMaxRPMTestedSuccessfully gets the summary which contains that RPM till 
// threshold breakforth is not encountered.
func getMaxRPMTestedSuccessfully(summary []models.Summary) models.Summary {
	indexOfMaxRPMTestedSuccessfully := getIndexOfMaxRPMTestedSuccessfully(summary)
	if indexOfMaxRPMTestedSuccessfully == -1 {
		return models.Summary{}
	}
	return summary[indexOfMaxRPMTestedSuccessfully]
}

// getIndexOfMaxRPMTestedSuccessfully get the index of summary which possess maximum
// RPM which did not encounter threshold breakforth.
func getIndexOfMaxRPMTestedSuccessfully(summary []models.Summary) int {
	indices := findBreakingRPMIndices(summary)
	switch len(indices) {
	case 0:
		// no breaking rpm encountered during the play
		return len(summary) - 1
	case 1:
		if indices[0] == 0 {
			return -1 // FIXME:
		}
	}
	return indices[0] - 1
}

package executor

import (
	"fmt"
	"strconv"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/simulator/sidecar"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/targeter"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

// ExecuteActs executes the acts within a playfile
func ExecuteActs(config *models.Config, play models.Play, directives models.SummarySpecificDirectives, iteration int) []models.Summary {
	var playSummary []models.Summary
	for _, act := range play.Acts { // iterating through the acts
		parameters := sidecar.AssembleOptions(act.Vegeta)
		urls := sidecar.GetPrerequisiteData(config, act)

		// logging parameters of act
		var fields = logger.Fields{
			"play_id":            play.Id,
			"play_name":          play.Name,
			"iteration":          strconv.Itoa(iteration + 1),
			"act_name":           act.Name,
			"act_api":            act.Api,
			"rate_per_second":    strconv.Itoa(act.Vegeta.RateOfRequests),
			"duration_in_second": strconv.Itoa(act.Vegeta.DurationInSeconds),
			"timeout_in_ms":      strconv.Itoa(act.Vegeta.TimeoutInMilliSeconds),
			"vegeta_keepalive":   strconv.FormatBool(act.Vegeta.KeepAlive),
			"rpm":                strconv.Itoa(act.Vegeta.RateOfRequests * 60),
		}
		logger.WithFields(fields).Infof("RUN START: act: %s", act.Name)
		// generating a allocateTargeter for attacker
		allocateTargeter, err := targeter.NewAllocateTargeterForAct(act, urls)
		if err != nil {
			logger.Error(err)
		} else {
			var codeToCountMap = make(map[uint16]int)
			// =>=>=>=> executing act
			metrics, updatedCodeToCountMap, requestCountOverExpectedLatency := Execute(allocateTargeter, codeToCountMap, parameters)
			actSummaryCSVRow := prepareSummaryForReportingForActs(play, &playSummary, iteration+1, act, metrics, updatedCodeToCountMap, requestCountOverExpectedLatency)

			// logging target distribution
			// logger.Infof("target distribution: %s", sidecar.GetTargetDistributionAsString(targetDistribution, act.Vegeta.RateOfRequests*act.Vegeta.DurationInSeconds))
			// logging final summary of play
			fields = logger.Fields{
				"rpm":                                 fmt.Sprintf("%d", act.Vegeta.RateOfRequests*60),
				"number_of_requests":                  strconv.FormatUint(metrics.Requests, 10),
				"latency_mean":                        metrics.Latencies.Mean.String(),
				"latency_p50":                         metrics.Latencies.P50.String(),
				"latency_p95":                         metrics.Latencies.P95.String(),
				"latency_p99":                         metrics.Latencies.P99.String(),
				"latency_max":                         metrics.Latencies.Max.String(),
				"success_ratio":                       fmt.Sprintf("%.3f", sidecar.CalculateCustomSuccessRatio(updatedCodeToCountMap)),
				"request_count_over_expected_latency": requestCountOverExpectedLatency,
			}
			for code, count := range codeToCountMap {
				fields[fmt.Sprint(code)] = count
			}
			err = sidecar.LogData(config, directives, actSummaryCSVRow)
			if err != nil {
				logger.Error(err)
			}
			logger.Info("act summary recorded.")

			logger.WithFields(fields).Info("RUN COMPLETED !")
			logger.Infof("load test ran for %d requests per second for %d with %d%% duplication of payload (a and b party nums).",
				parameters.RateOfRequests, parameters.DurationInSeconds, parameters.DuplicacyPercentage)
			urls = nil
		}
	}
	return playSummary
}

// prepareSummaryForReportingForActs accumulates the collective metrics
func prepareSummaryForReportingForActs(play models.Play, playSummary *[]models.Summary, iteration int, act models.Act, metrics vegeta.Metrics, codeToCountMap map[uint16]int, overLatencyCount int) []string {
	var actSummary = models.Summary{
		PlayId:                          play.Id,
		PlayName:                        play.Name,
		Iteration:                       strconv.Itoa(iteration),
		Name:                            act.Name,
		Id:                              act.Api,
		RateOfRequests:                  strconv.Itoa(act.Vegeta.RateOfRequests),
		DurationInSeconds:               strconv.Itoa(act.Vegeta.DurationInSeconds),
		VegetaTimeoutInMilliSeconds:     strconv.Itoa(act.Vegeta.TimeoutInMilliSeconds),
		KeepPersistentConnections:       strconv.FormatBool(act.Vegeta.KeepAlive),
		LoadRunDuration:                 metrics.Duration,
		TotalNumberOfRequests:           metrics.Requests,
		TargetDistribution:              "",
		RPM:                             int64(act.Vegeta.RateOfRequests * 60),
		LatencyAtMeanUsage:              metrics.Latencies.Mean,
		LatencyAt50Usage:                metrics.Latencies.P50,
		LatencyAt95Usage:                metrics.Latencies.P95,
		LatencyAt99Usage:                metrics.Latencies.P99,
		LatencyAtMaxUsage:               metrics.Latencies.Max,
		SuccessRatio:                    sidecar.CalculateCustomSuccessRatio(codeToCountMap),
		RequestCountOverExpectedLatency: overLatencyCount,
		StatusCodes:                     codeToCountMap,
	}
	*playSummary = append(*playSummary, actSummary)

	var statusCodes string = ""
	for code, count := range codeToCountMap {
		statusCodes += fmt.Sprintf("%d=%d ", code, count)
	}

	var summaryAsStringList = []string{
		actSummary.PlayId,
		actSummary.PlayName,
		actSummary.Iteration,
		actSummary.Name,
		actSummary.Id,
		actSummary.RateOfRequests,
		actSummary.DurationInSeconds,
		actSummary.VegetaTimeoutInMilliSeconds,
		actSummary.KeepPersistentConnections,
		actSummary.DurationInSeconds,
		strconv.FormatUint(actSummary.TotalNumberOfRequests, 10),
		actSummary.TargetDistribution,
		strconv.FormatInt(actSummary.RPM, 10),
		actSummary.LatencyAtMeanUsage.String(),
		actSummary.LatencyAt50Usage.String(),
		actSummary.LatencyAt95Usage.String(),
		actSummary.LatencyAt99Usage.String(),
		actSummary.LatencyAtMaxUsage.String(),
		fmt.Sprintf("%.3f", actSummary.SuccessRatio),
		strconv.Itoa(actSummary.RequestCountOverExpectedLatency),
		statusCodes,
	}

	return summaryAsStringList
}

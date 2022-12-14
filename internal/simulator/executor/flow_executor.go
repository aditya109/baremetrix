package executor

import (
	"fmt"
	"strconv"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/flow/flows"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/flow/ingester"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/simulator/sidecar"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/targeter"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

// ExecuteFlows executes the flows within a playfile.
func ExecuteFlows(config *models.Config, play models.Play, directives models.SummarySpecificDirectives, iteration int) []models.Summary {
	var playSummary []models.Summary
	flowMapping := flows.GetFlowsMapping()
	for _, flow := range play.Flows {
		flowSets := ingester.IngestFlow(config, flow, flowMapping)

		// logging parameters of flow
		var fields = logger.Fields{
			"play_id":            play.Id,
			"play_name":          play.Name,
			"iteration":          strconv.Itoa(iteration + 1),
			"flow_name":          flow.Name,
			"flow_id":            flow.FlowId,
			"rate_per_second":    strconv.Itoa(flow.Vegeta.RateOfRequests),
			"duration_in_second": strconv.Itoa(flow.Vegeta.DurationInSeconds),
			"timeout_in_ms":      strconv.Itoa(flow.Vegeta.TimeoutInMilliSeconds),
			"vegeta_keepalive":   strconv.FormatBool(flow.Vegeta.KeepAlive),
			"rpm":                strconv.Itoa(flow.Vegeta.RateOfRequests * 60),
		}
		logger.WithFields(fields).Infof("RUN START: flow: %s", flow.Name)
		// generating an allocateTargeter for attacker
		parameters := sidecar.AssembleOptions(flow.Vegeta)
		var targetDistribution = make([]int, len(flow.Scenes))
		allocateTargeter, err := targeter.NewAllocateTargeterForFlow(flow, flowSets, &targetDistribution)
		if err != nil {
			logger.Error(err)
		} else {
			var codeToCountMap = make(map[uint16]int)
			// =>=>=>=> executing flow
			metrics, updatedCodeToCountMap, requestCountOverExpectedLatency := Execute(allocateTargeter, codeToCountMap, parameters)
			flowSummaryCSVRow := prepareSummaryForReportingForFlows(config, play, &playSummary, iteration+1, flow, metrics, updatedCodeToCountMap, requestCountOverExpectedLatency, targetDistribution)

			// logging target distribution
			logger.Infof("target distribution: %s", sidecar.GetTargetDistributionAsString(config, targetDistribution, flow.Vegeta.RateOfRequests*flow.Vegeta.DurationInSeconds, flow.FlowId))
			// logging final summary of play
			var fields = logger.Fields{
				"rpm":                                 fmt.Sprintf("%d", flow.Vegeta.RateOfRequests*60),
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
			err = sidecar.LogData(config, directives, flowSummaryCSVRow)
			if err != nil {
				logger.Error(err)
			}
			logger.Info("flow summary recorded.")

			logger.WithFields(fields).Info("RUN COMPLETED !")
			logger.Infof("load test ran for %d requests per second for %d with %d%% duplication of payload (a and b party nums).",
				parameters.RateOfRequests, parameters.DurationInSeconds, parameters.DuplicacyPercentage)
			flowSets = nil
		}
	}
	return playSummary
}

// prepareSummaryForReportingForFlows accumulates the collective metrics
func prepareSummaryForReportingForFlows(config *models.Config, play models.Play, playSummary *[]models.Summary, iteration int, flow models.Flow, metrics vegeta.Metrics, codeToCountMap map[uint16]int, overLatencyCount int, targetDistribution []int) []string {
	var flowSummary = models.Summary{
		PlayId:                          play.Id,
		PlayName:                        play.Name,
		Iteration:                       strconv.Itoa(iteration),
		Name:                            flow.Name,
		Id:                              flow.FlowId,
		RateOfRequests:                  strconv.Itoa(flow.Vegeta.RateOfRequests),
		DurationInSeconds:               strconv.Itoa(flow.Vegeta.DurationInSeconds),
		VegetaTimeoutInMilliSeconds:     strconv.Itoa(flow.Vegeta.TimeoutInMilliSeconds),
		KeepPersistentConnections:       strconv.FormatBool(flow.Vegeta.KeepAlive),
		LoadRunDuration:                 metrics.Duration,
		TotalNumberOfRequests:           metrics.Requests,
		TargetDistribution:              sidecar.GetTargetDistributionAsString(config, targetDistribution, flow.Vegeta.RateOfRequests*flow.Vegeta.DurationInSeconds, flow.FlowId),
		RPM:                             int64(flow.Vegeta.RateOfRequests * 60),
		LatencyAtMeanUsage:              metrics.Latencies.Mean,
		LatencyAt50Usage:                metrics.Latencies.P50,
		LatencyAt95Usage:                metrics.Latencies.P95,
		LatencyAt99Usage:                metrics.Latencies.P99,
		LatencyAtMaxUsage:               metrics.Latencies.Max,
		SuccessRatio:                    sidecar.CalculateCustomSuccessRatio(codeToCountMap),
		RequestCountOverExpectedLatency: overLatencyCount,
		StatusCodes:                     codeToCountMap,
	}
	*playSummary = append(*playSummary, flowSummary)

	var statusCodes string = ""
	for code, count := range codeToCountMap {
		statusCodes += fmt.Sprintf("%d=%d ", code, count)
	}

	var summaryAsStringList = []string{
		flowSummary.PlayId,
		flowSummary.PlayName,
		flowSummary.Iteration,
		flowSummary.Name,
		flowSummary.Id,
		flowSummary.RateOfRequests,
		flowSummary.DurationInSeconds,
		flowSummary.VegetaTimeoutInMilliSeconds,
		flowSummary.KeepPersistentConnections,
		flowSummary.DurationInSeconds,
		strconv.FormatUint(flowSummary.TotalNumberOfRequests, 10),
		flowSummary.TargetDistribution,
		strconv.FormatInt(flowSummary.RPM, 10),
		flowSummary.LatencyAtMeanUsage.String(),
		flowSummary.LatencyAt50Usage.String(),
		flowSummary.LatencyAt95Usage.String(),
		flowSummary.LatencyAt99Usage.String(),
		flowSummary.LatencyAtMaxUsage.String(),
		fmt.Sprintf("%.3f", flowSummary.SuccessRatio),
		strconv.Itoa(flowSummary.RequestCountOverExpectedLatency),
		statusCodes,
	}

	return summaryAsStringList
}

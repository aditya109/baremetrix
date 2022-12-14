package executor

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
	"time"
)

// Execute executes the vegeta attack with the provided targeter
func Execute(targeter vegeta.Targeter, codeToCountMap map[uint16]int, parameters models.LoadTestingParameters) (vegeta.Metrics, map[uint16]int, int) {
	// creating attacker
	attacker := vegeta.NewAttacker(vegeta.Timeout(time.Duration(parameters.TimeoutInMilliseconds)*time.Millisecond), vegeta.KeepAlive(parameters.KeepAlive))
	var metrics vegeta.Metrics
	requestCountOverExpectedLatency := 0

	// iterating through the sequence of attacks
	for res := range attacker.Attack(targeter, parameters.Options.RateOfRequests, parameters.Options.Duration, "Big Bang!") {
		metrics.Add(res)
		if res.Code != 200 {
			logger.WithFields(logger.Fields{
				"seq":         res.Seq,
				"latency":     res.Latency,
				"status_code": res.Code,
			}).Error(res.Error)
		} else {
			logger.WithFields(logger.Fields{
				"seq":         res.Seq,
				"latency":     res.Latency,
				"status_code": res.Code,
			})
		}
		codeToCountMap[res.Code] += 1
		if res.Latency >= 1*time.Second {
			requestCountOverExpectedLatency += 1
		}
	}
	metrics.Close()
	return metrics, codeToCountMap, requestCountOverExpectedLatency
}

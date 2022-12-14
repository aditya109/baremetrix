package models

import "time"

// Summary is the model which accounts for each row published at the end of each iteration.
type Summary struct {
	PlayId                          string
	PlayName                        string
	Iteration                       string
	Name                            string
	Id                              string
	RateOfRequests                  string
	DurationInSeconds               string
	VegetaTimeoutInMilliSeconds     string
	KeepPersistentConnections       string
	LoadRunDuration                 time.Duration
	TotalNumberOfRequests           uint64
	TargetDistribution              string
	RPM                             int64
	LatencyAtMeanUsage              time.Duration
	LatencyAt50Usage                time.Duration
	LatencyAt95Usage                time.Duration
	LatencyAt99Usage                time.Duration
	LatencyAtMaxUsage               time.Duration
	SuccessRatio                    float64
	RequestCountOverExpectedLatency int
	StatusCodes                     map[uint16]int
}

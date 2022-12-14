package models

import (
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

// Options contains vegeta attack-related parameters
type Options struct {
	RateOfRequests      vegeta.Rate
	Duration            time.Duration
	RangeOfRand         int64
	DuplicacyPercentage int
}

// UniqueRand contains a map of large random unique numbers
type UniqueRand struct {
	Generated map[int64]bool
}

// LoadTestingParameters parameters required by executor for proper functioning.
type LoadTestingParameters struct {
	RateOfRequests        int
	DurationInSeconds     int
	DuplicacyPercentage   int
	TimeoutInMilliseconds int
	KeepAlive             bool
	Options               Options
}

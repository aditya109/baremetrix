package models

// Conclusion containing the play summary for each playfile, which is printed on the summary file, as well.
type Conclusion struct {
	TargetAPIOrFlowId        string
	TenantId                 string
	ActDuration              string
	ThresholdLatency         string
	BreakingRPMs             string
	MaxRPMTestedSuccessfully string
	LatencyP95AtMaxRPM       string
	LatencyMeanAtMaxRPM      string
	SuccessRatioAtMaxRPM     string
	VegetaTimeout            string
	StagingEnvironmentInfo   string
}

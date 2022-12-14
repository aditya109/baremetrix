package constants

const (
	// API Names
	LaGreenvnCreateAllocation     = "LA_GREENVN_CREATE_ALLOCATION"
	LaGreenvnDeleteAllocation     = "LA_GREENVN_DELETE_ALLOCATION"
	LaGreenvnGetAllocationDetails = "LA_GREENVN_GET_ALLOCATION_DETAILS"
	LaRandomSequence              = "LA_RANDOM_SEQUENCE"
	LaExotelWebhookCallerValidate = "LA_EXOTEL_WEBHOOK_CALLER_VALIDATE"
	LaExotelWebhookCallerConnect  = "LA_EXOTEL_WEBHOOK_CALLER_CONNECT"
	LaExotelWebhookPtPostConnect  = "LA_EXOTEL_WEBHOOK_PT_POST_CONNECT"
	LaGreenpinCreatePinAllocation = "LA_GREENPIN_CREATE_PIN_ALLOCATION"
	LaGreenpinDeletePinAllocation = "LA_GREENPIN_DELETE_PIN_ALLOCATION"
	LaGreenpinGetPinAllocation    = "LA_GREENPIN_GET_PIN_ALLOCATION"

	// Summary Headers
	PlayId                          = "PlayId"
	PlayName                        = "PlayName"
	PlayIteration                   = "PlayIteration"
	Name                            = "Name"
	RateOfRequestsPerSecond         = "RateOfRequestsPerSecond"
	DurationInSeconds               = "DurationInSeconds"
	VegetaTimeoutInMilliseconds     = "VegetaTimeoutInMilliSeconds"
	KeepPersistentConnections       = "KeepPersistentConnections"
	LoadRunDuration                 = "LoadRunDuration"
	TotalNumberOfRequests           = "TotalNumberOfRequests"
	TargetDistribution              = "TargetTypeDistribution"
	RateOfRequestsPerMinute         = "CurrentRPM"
	LatencyAtMeanUsage              = "LatencyAtMeanUsage"
	LatencyAt50Usage                = "LatencyAt50Usage"
	LatencyAt95Usage                = "LatencyAt95Usage"
	LatencyAt99Usage                = "LatencyAt99Usage"
	LatencyAtMaxUsage               = "LatencyAtMaxUsage"
	SuccessRatio                    = "SuccessRatio"
	RequestCountOverExpectedLatency = "Request#OverExpectedLatency"
	StatusCodes                     = "StatusCodes"

	// Act Summary Conclusion Field Keys
	ConclusionSectionHeading = "Notes:"
	TargetId                 = "Target API/FlowId"
	TenantId                 = "Tenant ID"
	ActDuration              = "Load testing duration"
	ThresholdLatency         = "Threshold Latency"
	BreakingRpm              = "Breaking RPM"
	MaxRpmTestedSuccessfully = "Maximum RPM Tested Successfully"
	LatencyP95AtMaxRpm       = "P95 Latency at Max RPM"
	LatencyMeanAtMaxRpm      = "Mean Latency at Max RPM"
	SuccessRatioAtMaxRpm     = "Success Ratio at Max RPM"
	VegetaTimeout            = VegetaTimeoutInMilliseconds
	InfraDetailsHeader       = "Infra Details:"
	StagingEnvInfo           = "Staging Environment Information"

	// Custom messages
	NoBreakingRPMEncountered = "no breaking rpm encountered during the play"

	// Graph Type Names
	RpmVersusLatency              = "rpm_vs_latency"
	RpmVersusDelayedResponseCount = "rps_vs_delayed_responses"

	// Legends
	MeanLatency = "MeanLatency"
	P50Latency  = "P50Latency"
	P95Latency  = "P95Latency"
	P99Latency  = "P99Latency"
	MaxLatency  = "MaxLatency"

	// Flow type names
	FLOW_SCENARIO_TYPE_1 = "FLOW_SCENARIO_TYPE_1"
)

// LaGreenvnApis contains a list of GREENVN APIS
var (
	LaGreenvnApis = [...]string{
		LaGreenvnCreateAllocation,
		LaGreenvnGetAllocationDetails,
		LaGreenvnDeleteAllocation,
		// LaGreenpinCreatePinAllocation, // TBD: To be worked upon GreenPin APIs will be tested. 
		// LaGreenpinDeletePinAllocation, // TBD: To be worked upon GreenPin APIs will be tested. 
	}
)

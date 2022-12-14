package flows

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/flow/scenarios"
)

// GetFlowsMapping returns a mapping containing the mapping of flowtype to corresponding scenario package functions.
func GetFlowsMapping() map[string]interface{} {
	var flowMapping = map[string]interface{}{
		constants.FLOW_SCENARIO_TYPE_1: scenarios.GetStackListFromFlowScenarioType1,
		// add flow_id to function mappings here.
	}
	return flowMapping
}

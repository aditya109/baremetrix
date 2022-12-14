package targets

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils/targetstack"
	vegeta "github.com/tsenart/vegeta/lib"
)

// GetLAFlowScenario1 gets a target for flows of type FLOW_SCENARIO_TYPE_1.
func GetLAFlowScenario1(flowSets []targetstack.FlowStack, targetDistribution *[]int) (vegeta.Targeter, error) {
	return func(target *vegeta.Target) error {
		if target == nil {
			return vegeta.ErrNilTarget
		}
		params := utils.GetRandomTargetFromFlowSets(&flowSets)
		target.Method = params.Method
		target.URL = params.URL
		target.Body = params.Body
		target.Header = params.Header
		//logger.Infof("%s: %s target added", target.Method, target.URL)
		//logger.Infof("Payload: %v", string(target.Body))

		switch params.ApiType {
		case constants.LaGreenvnCreateAllocation:
			(*targetDistribution)[0] += 1
		case constants.LaExotelWebhookCallerValidate:
			(*targetDistribution)[1] += 1
		case constants.LaExotelWebhookCallerConnect:
			(*targetDistribution)[2] += 1
		case constants.LaExotelWebhookPtPostConnect:
			(*targetDistribution)[3] += 1
		case constants.LaGreenvnDeleteAllocation:
			(*targetDistribution)[4] += 1
		}
		return nil
	}, nil
}

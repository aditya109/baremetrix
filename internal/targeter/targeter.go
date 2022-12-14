package targeter

import (
	"fmt"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/targeter/targets"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/targeter/targets/la_greenpin"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/targeter/targets/la_greenvn"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils/targetstack"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

// NewAllocateTargeterForAct creates a new targeter of vegeta.Targeter for acts.
func NewAllocateTargeterForAct(act models.Act, urls map[string][]string) (vegeta.Targeter, error) {
	// creates a target based on method specified in the act of the step
	return func(target *vegeta.Target) error {
		if target == nil {
			return vegeta.ErrNilTarget
		}
		var targetId string
		if act.Api != constants.LaRandomSequence {
			targetId = act.Api
		} else {
			item, err := utils.GetRandomItemFromList(constants.LaGreenvnApis[:])
			if err != nil {
				logger.Error(err)
				return err
			}
			targetId = fmt.Sprint(item)
		}

		switch targetId {
		case constants.LaGreenvnCreateAllocation:
			//logger.Infof("%s target is selected", constants.LaGreenvnCreateAllocation)
			err := la_greenvn.GetLACreateGreenVNAllocationAPITarget(act, target)
			if err != nil {
				logger.Error(err)
				return err
			}
		case constants.LaGreenvnGetAllocationDetails:
			//logger.Infof("%s target is selected", constants.LaGreenvnGetAllocationDetails)
			la_greenvn.GetLAGetGreenVNAllocationDetailsAPITarget(act, urls[constants.LaGreenvnCreateAllocation][0], target)
		case constants.LaGreenvnDeleteAllocation:
			//logger.Infof("%s target is selected", constants.LaGreenvnDeleteAllocation)
			err := la_greenvn.GetLADeleteGreenVNAllocationAPITarget(act, urls[constants.LaGreenvnCreateAllocation], target)
			if err != nil {
				logger.Error(err)
				return err
			}
		case constants.LaGreenpinCreatePinAllocation:
			//logger.Infof("%s target is selected", constants.LaGreenpinCreatePinAllocation)
			err := la_greenpin.GetLAGreenPinCreatePinAllocationTarget(act, target)
			if err != nil {
				logger.Error(err)
				return err
			}
		case constants.LaGreenpinDeletePinAllocation:
			//logger.Infof("%s target is selected", constants.LaGreenvnDeleteAllocation)
			err := la_greenpin.GetLADeleteGreenPinAllocationAPITarget(act, urls[constants.LaGreenpinCreatePinAllocation], target)
			if err != nil {
				logger.Error(err)
				return err
			}
		}
		return nil
	}, nil

}

// NewAllocateTargeterForFlow creates a new targeter of vegeta.Targeter for flows.
func NewAllocateTargeterForFlow(flow models.Flow, flowSets []targetstack.FlowStack, targetDistribution *[]int) (vegeta.Targeter, error) {
	switch flow.FlowId {
	case constants.FLOW_SCENARIO_TYPE_1:
		return targets.GetLAFlowScenario1(flowSets, targetDistribution)
	}
	return nil, nil
}

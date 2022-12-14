package scenarios

import (
	"fmt"
	"strings"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/generator"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models/api/greenvncreateallocation"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/prerequisite"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils/targetstack"
	logger "github.com/sirupsen/logrus"
)

// flowStackParameters contains all the parameters for running functions for gathering flowsets.
type flowStackParameters struct {
	flow                                         models.Flow
	tuple                                        greenvncreateallocation.Tuple
	shouldAddLAGreenVNCreateAllocationTarget     bool
	shouldAddLAExotelWebhookCallerValidateTarget bool
	shouldAddLAExotelWebhookCallerConnectTarget  bool
	shouldAddLAExotelWebhookPtPostConnectTarget  bool
	shouldAddLAGreenVNDeleteAllocationTarget     bool
}

// GetStackListFromFlowScenarioType1 provides the stack-type flowset list flowing type - FLOW_SCENARIO_TYPE_1.
func GetStackListFromFlowScenarioType1(config *models.Config, flow models.Flow) []targetstack.FlowStack {
	var apiToSceneMapping = utils.GetApiTypeToSceneMapping(flow.Scenes)

	// scenario starts off with pre-requisite call to LA_GREENVN_CREATE_ALLOCATION
	createAllocationScene := apiToSceneMapping[constants.LaGreenvnCreateAllocation]

	expectedSumOfTargetsPerBatch := 0
	for _, scene := range flow.Scenes {
		expectedSumOfTargetsPerBatch += scene.ExpectedRPS
	}
	totalRequiredTargets := flow.Vegeta.RateOfRequests * flow.Vegeta.DurationInSeconds

	var targetTypeCount = map[string]int{
		constants.LaGreenvnCreateAllocation:     totalRequiredTargets * flow.Scenes[0].ExpectedRPS / 100,
		constants.LaExotelWebhookCallerValidate: totalRequiredTargets * flow.Scenes[1].ExpectedRPS / 100,
		constants.LaExotelWebhookCallerConnect:  totalRequiredTargets * flow.Scenes[2].ExpectedRPS / 100,
		constants.LaExotelWebhookPtPostConnect:  totalRequiredTargets * flow.Scenes[3].ExpectedRPS / 100,
		constants.LaGreenvnDeleteAllocation:     totalRequiredTargets * flow.Scenes[len(flow.Scenes)-1].ExpectedRPS / 100,
	}
	// assuming that LaExotelWebhookCallerValidate, LaExotelWebhookCallerConnect, LaExotelWebhookPtPostConnect has same expected_rps for the sake for avoid computational complexity
	fullStacksRequired := targetTypeCount[constants.LaExotelWebhookCallerValidate]
	// assuming that LaGreenvnCreateAllocation, LaGreenvnDeleteAllocation has same expected_rps for the sake for avoid computational complexity
	partialStacksRequired := targetTypeCount[constants.LaGreenvnCreateAllocation] - targetTypeCount[constants.LaExotelWebhookCallerValidate]
	totalStacks := fullStacksRequired + partialStacksRequired
	logger.Infof("fullStacksRequired: %d, partialStacksRequired: %d", fullStacksRequired, partialStacksRequired)

	tuplesRequired := totalStacks
	tuplesRequired += config.Instance.Specs.PrerequisiteSpecs.ConditionalBatchBufferForFlows
	logger.Infof("%d # targets required, %d tuples required", totalRequiredTargets, tuplesRequired)
	tuples := prerequisite.GetGreenVNCreateAllocationTuplesConcurrently(config, models.Act{
		Allocateurl: []models.Allocateurl{
			{
				UseCase: constants.LaGreenvnCreateAllocation,
				URL:     createAllocationScene.Allocateurl,
			},
		},
		Api:               createAllocationScene.Api,
		Method:            createAllocationScene.Method,
		Headers:           createAllocationScene.Headers,
		Endpoint:          createAllocationScene.Endpoint,
		PreOpSequence:     nil,
		DefaultParameters: createAllocationScene.DefaultParameters,
		Vegeta:            flow.Vegeta,
	}, tuplesRequired)
	logger.Infof("%d #tuples acquired", len(tuples))

	var generatedFlowStacks []targetstack.FlowStack
	accumulatedTargetCount := 0
	for _, tuple := range tuples {
		if fullStacksRequired != 0 {
			parameters := flowStackParameters{
				flow:                                     flow,
				tuple:                                    tuple,
				shouldAddLAGreenVNCreateAllocationTarget: true,
				shouldAddLAExotelWebhookCallerValidateTarget: true,
				shouldAddLAExotelWebhookCallerConnectTarget:  true,
				shouldAddLAExotelWebhookPtPostConnectTarget:  true,
				shouldAddLAGreenVNDeleteAllocationTarget:     true,
			}
			stack, err := getFlowStack(parameters)
			if err != nil {
				logger.Error(err)
			}
			accumulatedTargetCount += len(stack)
			generatedFlowStacks = append(generatedFlowStacks, stack)
			fullStacksRequired--
		}
		if fullStacksRequired == 0 && partialStacksRequired != 0 {
			parameters := flowStackParameters{
				flow:                                     flow,
				tuple:                                    tuple,
				shouldAddLAGreenVNCreateAllocationTarget: true,
				shouldAddLAExotelWebhookCallerValidateTarget: false,
				shouldAddLAExotelWebhookCallerConnectTarget:  false,
				shouldAddLAExotelWebhookPtPostConnectTarget:  false,
				shouldAddLAGreenVNDeleteAllocationTarget:     true,
			}
			stack, err := getFlowStack(parameters)
			if err != nil {
				logger.Error(err)
			}
			accumulatedTargetCount += len(stack)
			generatedFlowStacks = append(generatedFlowStacks, stack)
			partialStacksRequired--
		}
		if partialStacksRequired == 0 && fullStacksRequired == 0 {
			break
		}
	}
	logger.Infof("%d targets required, %d # tuples acquired, %d/%d # flowStacks generated, %d # targets acquired", totalRequiredTargets, len(tuples), len(generatedFlowStacks), totalStacks, accumulatedTargetCount)
	return generatedFlowStacks
}

// getFlowStack provides a stack consisting of target parameters for FLOW_SCENARIO_TYPE_1.
func getFlowStack(parameters flowStackParameters) (targetstack.FlowStack, error) {
	var (
		flow      = parameters.flow
		tuple     = parameters.tuple
		flowStack targetstack.FlowStack
	)
	callSid := utils.GenerateUUID()
	aPartyNumber := strings.ReplaceAll(tuple.Payload.ApartyNumbers[0], "+91", "0")
	greenVN := strings.ReplaceAll(tuple.Response.Data.GreenVN, "+91", "0")
	bPartyNumber := strings.ReplaceAll(tuple.Payload.BpartyNumbers[0], "+91", "0")
	startTime := utils.GetUnescapedDateTimeString()
	currentTime := utils.GetUnescapedDateTimeString()
	header := generator.GetHeaderForHTTPRequest(flow.Scenes[0].Endpoint, flow.Scenes[0].Headers)

	// logger.Infof("%s %s %s %s %s %s %s", tuple.Response.Data.GreenVNID, callSid, aPartyNumber, bPartyNumber, greenVN, startTime, currentTime)

	if parameters.shouldAddLAGreenVNDeleteAllocationTarget && len(flow.Scenes) >= 5 { // adding LA_GREENVN_DELETE_ALLOCATION target
		url := strings.ReplaceAll(strings.ReplaceAll(flow.Scenes[4].Allocateurl, "GREENVN_ID", tuple.Response.Data.GreenVNID), "SID", flow.Endpoint.Sid)
		// logger.Infof("%s: %s target added", flow.Scenes[4].Method, url)
		flowStack = flowStack.Push(models.TargetParameters{
			ApiType: constants.LaGreenvnDeleteAllocation,
			URL:     url,
			Method:  flow.Scenes[4].Method,
			Header:  header,
		})
	}
	if parameters.shouldAddLAExotelWebhookPtPostConnectTarget { // adding LA_EXOTEL_WEBHOOK_PT_POST_CONNECT target
		url := strings.ReplaceAll(flow.Scenes[3].Allocateurl, "SID", flow.Endpoint.Sid)
		url = fmt.Sprintf("%s?CallSid=%s&Direction=incoming&StartTime=%s&CallType=completed&DialWhomNumber=%s&From=%s&To=%s&CurrentTime=%s",
			url, callSid, startTime, bPartyNumber, aPartyNumber, greenVN, currentTime)
		// logger.Infof("%s: %s target added", flow.Scenes[3].Method, url)
		flowStack = flowStack.Push(models.TargetParameters{
			ApiType: constants.LaExotelWebhookPtPostConnect,
			URL:     url,
			Method:  flow.Scenes[3].Method,
			Header:  header,
		})
	}
	if parameters.shouldAddLAExotelWebhookCallerConnectTarget { // adding LA_EXOTEL_WEBHOOK_CALLER_CONNECT target
		url := strings.ReplaceAll(flow.Scenes[2].Allocateurl, "SID", flow.Endpoint.Sid)
		url = fmt.Sprintf("%s?CallSid=%s&CallStatus=ringing&Direction=incoming&CallType=call-attempt&From=%s&To=%s",
			url, callSid, aPartyNumber, greenVN)
		// logger.Infof("%s: %s target added", flow.Scenes[2].Method, url)
		flowStack = flowStack.Push(models.TargetParameters{
			ApiType: constants.LaExotelWebhookCallerConnect,
			URL:     url,
			Method:  flow.Scenes[2].Method,
			Header:  header,
		})
	}
	if parameters.shouldAddLAExotelWebhookCallerValidateTarget { // adding LA_EXOTEL_WEBHOOK_CALLER_VALIDATE target
		url := strings.ReplaceAll(flow.Scenes[1].Allocateurl, "SID", flow.Endpoint.Sid)
		url = fmt.Sprintf("%s?CallSid=%s&From=%s&To=%s", url, callSid, aPartyNumber, greenVN)
		// logger.Infof("%s: %s target added", flow.Scenes[1].Method, url)
		flowStack = flowStack.Push(models.TargetParameters{
			ApiType: constants.LaExotelWebhookCallerValidate,
			URL:     url,
			Method:  flow.Scenes[1].Method,
			Header:  header,
		})
	}
	if parameters.shouldAddLAGreenVNCreateAllocationTarget { // adding LA_GREENVN_CREATE_ALLOCATION target
		flow.Scenes[0].Allocateurl = strings.ReplaceAll(strings.ReplaceAll(flow.Scenes[0].Allocateurl, "/GREENVN_ID", ""), "SID", flow.Scenes[0].Endpoint.Sid)
		jsonPayload, _, err := generator.GetPayloadForPostRequest(flow.Scenes[0].DefaultParameters, constants.LaGreenvnCreateAllocation)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		// logger.Infof("%s: %s target added", flow.Scenes[0].Method, flow.Scenes[0].Allocateurl)
		flowStack = flowStack.Push(models.TargetParameters{
			ApiType: constants.LaGreenvnCreateAllocation,
			Method:  flow.Scenes[0].Method,
			URL:     flow.Scenes[0].Allocateurl,
			Body:    jsonPayload,
			Header:  header,
		})
	}
	//logger.Infof("%d targets added", len(flowStack))
	// logger.Info("=======================================================")
	// for _, v := range flowStack {
	// 	logger.Info(v)
	// }

	// logger.Info("=======================================================")
	return flowStack, nil
}

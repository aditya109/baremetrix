package prerequisite

import (
	"strings"
	"time"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models/api/greenpincreatepinallocation"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/generator"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models/api/greenvncreateallocation"
	logger "github.com/sirupsen/logrus"
)

// PreallocateGreenVNIds provides allocated GreenVNIds for act with non-empty pre_op_sequence
func PreallocateGreenVNIds(config *models.Config, act models.Act, specialRequirement int, apiType string) []string {
	var urls []string
	switch apiType {
	case constants.LaGreenvnCreateAllocation:
		tuples := GetGreenVNCreateAllocationTuplesConcurrently(config, act, specialRequirement)
		for _, tuple := range tuples {
			response := tuple.Response
			url := utils.FindItemFromListWithKey(act.Allocateurl, constants.LaGreenvnGetAllocationDetails).(models.Allocateurl).URL
			urls = append(urls, strings.ReplaceAll(strings.ReplaceAll(url, "GREENVN_ID", response.Data.GreenVNID), "SID", act.Endpoint.Sid))
		}
		logger.Info(urls[0])
	case constants.LaGreenpinCreatePinAllocation:
		tuples := GetGreenPinCreateAllocationTuplesConcurrently(config, act, specialRequirement)
		for _, tuple := range tuples {
			response := tuple.Response
			urls = append(urls, strings.ReplaceAll(strings.ReplaceAll(utils.FindItemFromListWithKey(act.Allocateurl, constants.LaGreenpinDeletePinAllocation).(models.Allocateurl).URL, "GREENPIN_ID", response.Data.GreenPinId), "SID", act.Endpoint.Sid))
		}
		logger.Info(urls[0])
	}

	return urls
}

// GetGreenVNCreateAllocationTuplesConcurrently provides the tuples for GreenVN create allocation API in concurrent batches.
func GetGreenVNCreateAllocationTuplesConcurrently(config *models.Config, act models.Act, specialRequirement int) []greenvncreateallocation.Tuple {
	start := time.Now()
	emptyT := 0
	var resultChannel = make(chan *greenvncreateallocation.Tuple)

	var supportedBatchSize = config.Instance.Specs.PrerequisiteSpecs.SupportedBatchSize
	var batchBuffer = config.Instance.Specs.PrerequisiteSpecs.ConditionalBatchBuffer
	var batches, iteration, expectedCount int = getPreconditionParameters(act, specialRequirement, supportedBatchSize, batchBuffer)
	logger.Infof("PREREQUISITE INFO : batch created# = %d, iteration = %d, expected# = %d", batches, iteration, expectedCount)
	var totalCollectedResults []greenvncreateallocation.Tuple

	for { // iterating http calls in sequential batches of concurrent API calls within go routines
		emptyI := 0
		start2 := time.Now()
		for i := 0; i < iteration; i++ {
			go generator.GenerateGreenVNCreateAllocation(config, act, resultChannel)
		}
		result := make([]*greenvncreateallocation.Tuple, iteration)
		for i, _ := range result {
			result[i] = <-resultChannel
			if result[i] == nil {
				emptyT++
				emptyI++
			} else {
				totalCollectedResults = append(totalCollectedResults, *result[i])
			}
		}
		logger.Infof(`# allocations successful in this iteration: %d - %d = %d in %.2fs, %d/%d completed, %d responses empty, %.2fs elapsed`,
			supportedBatchSize, emptyI, supportedBatchSize-emptyI, time.Since(start2).Seconds(), len(totalCollectedResults), expectedCount, emptyT, time.Since(start).Seconds())
		if len(totalCollectedResults) < expectedCount {
			continue
		} else {
			break
		}
	}
	return totalCollectedResults
}

// GetGreenPinCreateAllocationTuplesConcurrently provides the tuples for GreenPIN create allocation API in concurrent batches.
func GetGreenPinCreateAllocationTuplesConcurrently(config *models.Config, act models.Act, specialRequirement int) []greenpincreatepinallocation.Tuple {
	start := time.Now()
	emptyT := 0
	var resultChannel = make(chan *greenpincreatepinallocation.Tuple)

	var supportedBatchSize = config.Instance.Specs.PrerequisiteSpecs.SupportedBatchSize
	var batchBuffer = config.Instance.Specs.PrerequisiteSpecs.ConditionalBatchBuffer
	var batches, iteration, expectedCount int = getPreconditionParameters(act, specialRequirement, supportedBatchSize, batchBuffer)
	logger.Infof("PREREQUISITE INFO : batch created# = %d, iteration = %d, expected# = %d", batches, iteration, expectedCount)
	var totalCollectedResults []greenpincreatepinallocation.Tuple

	for j := 0; j < batches; j++ { // iterating http calls in sequential batches of concurrent API calls within go routines
		emptyI := 0
		start2 := time.Now()
		for i := 0; i < iteration; i++ {
			go generator.GenerateGreenPinCreateAllocation(config, act, resultChannel)
		}
		result := make([]*greenpincreatepinallocation.Tuple, iteration)
		for i, _ := range result {
			result[i] = <-resultChannel
			if result[i] == nil {
				emptyT++
				emptyI++
			} else {
				totalCollectedResults = append(totalCollectedResults, *result[i])
			}
		}
		logger.Infof(`# allocations successful in this iteration: %d - %d = %d in %.2fs, %d/%d completed, %d responses empty, %.2fs elapsed`,
			supportedBatchSize, emptyI, supportedBatchSize-emptyI, time.Since(start2).Seconds(), len(totalCollectedResults), expectedCount, emptyT, time.Since(start).Seconds())
	}
	return totalCollectedResults
}

// getPreconditionParameters provides the parameters for batch-processing of porerequisite calls.
func getPreconditionParameters(act models.Act, specialRequirement int, supportedBatchSize int, buffer int) (int, int, int) {
	var expectedCount int
	var batches int
	var iteration int
	if specialRequirement != 0 { // calculating expectedCount
		expectedCount = specialRequirement
	} else {
		expectedCount = act.Vegeta.RateOfRequests * act.Vegeta.DurationInSeconds
	}

	if expectedCount < supportedBatchSize { // calculating iteration and batch count
		iteration = expectedCount
		batches = 1
	} else {
		iteration = supportedBatchSize
		batches = expectedCount / supportedBatchSize
		if batches == 0 || expectedCount%supportedBatchSize > 0 { // adjusting batch count with buffer conditionally
			batches += buffer
		}
	}

	return batches, iteration, expectedCount
}

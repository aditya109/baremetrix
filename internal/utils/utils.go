package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils/targetstack"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

// RemoveElementAtGivenIndexFromList removes an element at given index from list
func RemoveElementAtGivenIndexFromList(list interface{}, i int) interface{} {
	switch o := list.(type) {
	case []string:
		return append(o[:i], o[i+1:]...)
	case []targetstack.FlowStack:
		return append(o[:i], o[i+1:]...)
	}
	return nil
}

// FindItemFromListWithKey is used to find a item from a list using given key.
func FindItemFromListWithKey(list interface{}, key string) interface{} {
	switch o := list.(type) {
	case []models.GraphType:
		for _, gt := range o {
			if gt.IsEnabled && gt.Name == key {
				return gt
			}
		}
	case []models.InnerPlot:
		for _, ip := range o {
			if ip.Name == key {
				return ip
			}
		}
	case []models.Allocateurl:
		for _, url := range o {
			if url.UseCase == key {
				return url
			}
		}
	case []models.FlowType:
		for _, url := range o {
			if url.Id == key {
				return url
			}
		}
	}
	return nil
}

// GetRandomItemFromList gets a random item from list.
func GetRandomItemFromList(list interface{}) (interface{}, error) {
	switch o := list.(type) {
	case []string:
		index, err := strconv.Atoi(GenerateRandomLargeNumber(0, int64(len(o))))
		if err != nil {
			logger.Error(err)
			return "", err
		}
		index = index % len(o)
		return o[index], nil
	}
	return "", nil
}

// GenerateRandomLargeNumber generates a random large number using min and max values.
func GenerateRandomLargeNumber(min int64, max int64) string {
	rand.Seed(time.Now().UnixNano())
	var r int64
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("panic occurred: %s ; parameters; min=%d max=%d", err, min, max)
		}
	}()
	r = rand.Int63n(max-min) + min
	return strconv.FormatInt(r, 10)
}

func GetApiTypeToSceneMapping(scenes []models.Scene) map[string]models.Scene {
	var apiTypeToSceneMapping = make(map[string]models.Scene)
	for _, scene := range scenes {
		apiTypeToSceneMapping[scene.Api] = scene
	}
	return apiTypeToSceneMapping
}

func GetRandomTargetFromFlowSets(flowSets *[]targetstack.FlowStack) models.TargetParameters {
	var randomIndex int
	var selectedFlowSet targetstack.FlowStack
	// logger.Infof("%d flowsets remaining", len(*flowSets))
	for {
		randomIndex, _ = strconv.Atoi(GenerateRandomLargeNumber(0, int64(len(*flowSets))))
		if randomIndex >= len(*flowSets) {
			continue
		} else {
			selectedFlowSet = (*flowSets)[randomIndex]
			break
		}

	}
	for selectedFlowSet.IsEmpty() {
		//logger.Info("selected flow stack now has 0 remaining targets")
		*flowSets = RemoveElementAtGivenIndexFromList(*flowSets, randomIndex).([]targetstack.FlowStack)
		//logger.Infof("selected flow stack at %d index is popped, %d flowsets remaining", randomIndex, len(*flowSets))
		logger.Infof("%d flowsets still remaining", len(*flowSets))
		if len(*flowSets) == 1 {
			selectedFlowSet = (*flowSets)[0]
		} else {
			logger.Infof("len: %d", len(*flowSets))
			for {
				randomIndex, _ = strconv.Atoi(GenerateRandomLargeNumber(0, int64(len(*flowSets))))
				if randomIndex >= len(*flowSets) {
					continue
				} else {
					selectedFlowSet = (*flowSets)[randomIndex]
					break
				}
			}
		}
	}
	//logger.Infof("flow stack at index %d is selected, flow stack has %d targets", randomIndex, len(selectedFlowSet))
	selectedFlowSet, poppedTarget := selectedFlowSet.Pop()
	//logger.Infof("selected flow stack now has %d remaining targets", len(selectedFlowSet))
	if !selectedFlowSet.IsEmpty() {
		(*flowSets)[randomIndex] = selectedFlowSet
	} else {
		//logger.Info("selected flow stack now has 0 remaining targets")
		*flowSets = RemoveElementAtGivenIndexFromList(*flowSets, randomIndex).([]targetstack.FlowStack)
		//logger.Infof("selected flow stack at %d index is popped, %d flowsets remaining", randomIndex, len(*flowSets))
		// logger.Infof("%d flowsets still remaining", len(*flowSets))
	}
	//logger.Infof("%d flowsets still remaining after target-selection", len(*flowSets))
	return poppedTarget
}

// GenerateUUID generates a random UUID.
func GenerateUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// GetUnescapedDateTimeString provides a URL-encoded current datetime
func GetUnescapedDateTimeString() string {
	currentTime := time.Now()
	return url.QueryEscape(fmt.Sprintf("%d-%d-%d %d:%d:%d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second()))
}

// PopRandomItemFromList pops a random element from an array
func PopRandomItemFromList(list *[]string) (string, error) {
	var randomIndex int
	var err error
	for {
		randomIndex, err = strconv.Atoi(GenerateRandomLargeNumber(0, int64(len(*list))))
		if err != nil {
			logger.Error(err)
			return "", err
		}
		if randomIndex >= len(*list) {
			continue
		} else {
			break
		}
	}
	item := (*list)[randomIndex]
	*list = RemoveElementAtGivenIndexFromList(*list, randomIndex).([]string)
	return item, nil
}

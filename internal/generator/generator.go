package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models/api/greenpincreatepinallocation"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models/api/greenvncreateallocation"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	logger "github.com/sirupsen/logrus"
)

// generatePayloadForGreenVNCreateAllocationAPI generates a random payload for GreenVN Create Allocation API.
func generatePayloadForGreenVNCreateAllocationAPI(params models.PayloadParameters) (greenvncreateallocation.Payload, error) {
	var payload = greenvncreateallocation.Payload{}
	// storing a randomly generated large number for apin
	a := utils.GenerateRandomLargeNumber(params.Seed.Apin.Min, params.Seed.Apin.Max)
	apin, err := strconv.Atoi(a)
	if err != nil {
		logger.Warn("unable to parse int while generate apin")
		return payload, errors.New("unable to parse int while generate apin")
	}
	// storing a randomly generated large number for bpin
	b := utils.GenerateRandomLargeNumber(params.Seed.Bpin.Min, params.Seed.Bpin.Max)
	bpin, err := strconv.Atoi(b)
	if err != nil {
		logger.Warn("unable to parse int while generate bpin")
		return payload, errors.New("unable to parse int while generate bpin")
	}
	// creating a payload with generated and playfile-specified parameters
	payload = greenvncreateallocation.Payload{
		ConnectionID:  utils.GenerateRandomLargeNumber(params.Seed.ConnectionId.Min, params.Seed.ConnectionId.Max),
		ApartyNumbers: append(payload.ApartyNumbers, params.CountryCode+a),
		BpartyNumbers: append(payload.BpartyNumbers, params.CountryCode+b),
		ApartyPins:    append(payload.ApartyPins, apin),
		BpartyPins:    append(payload.BpartyPins, bpin),
		Usage:         params.Usage,
		DeallocationPolicy: greenvncreateallocation.DeallocationPolicy{
			Duration: params.DeallocationPolicy.Duration,
		},
		Strictness: params.Strictness,
		Preferences: greenvncreateallocation.Preferences{
			Greenvn: params.Preferences.Greenvn,
			Region:  params.Preferences.Region,
			Type:    params.Preferences.Type,
		},
	}
	return payload, nil
}

// generatePayloadForGreenPinCreateAllocationAPI generates a random payload for GreenVN Create Allocation API.
func generatePayloadForGreenPinCreateAllocationAPI(params models.PayloadParameters) (greenpincreatepinallocation.Payload, error) {
	var payload = greenpincreatepinallocation.Payload{}
	var err error
	payload.TransactionId = utils.GenerateRandomLargeNumber(params.Seed.TransactionId.Min, params.Seed.TransactionId.Max)
	payload.A, err = generatePartyNumberRequestInfo(params)
	if err != nil {
		logger.Error(err)
		return payload, nil
	}
	payload.B, err = generatePartyNumberRequestInfo(params)
	if err != nil {
		logger.Error(err)
		return payload, nil
	}
	payload.Usage = params.Usage
	payload.DeallocationPolicy = greenvncreateallocation.DeallocationPolicy{
		Duration: params.DeallocationPolicy.Duration,
	}
	return payload, nil
}

// generatePartyNumberRequestInfo provides party number request info for GreenPin API.
func generatePartyNumberRequestInfo(params models.PayloadParameters) (greenpincreatepinallocation.PartyNumberRequestInfo, error) {
	var result greenpincreatepinallocation.PartyNumberRequestInfo
	var pinLength, err = strconv.Atoi(utils.GenerateRandomLargeNumber(params.Seed.PinLength.Min, params.Seed.PinLength.Max))
	if err != nil {
		logger.Warn("unable to parse int while generate pinLength for a party")
		return result, errors.New("unable to parse int while generate pinLength for a party")
	}
	vns := utils.GenerateRandomLargeNumber(params.Seed.VNS.Min, params.Seed.VNS.Max)
	numbers := utils.GenerateRandomLargeNumber(params.Seed.Numbers.Min, params.Seed.Numbers.Max)
	result = greenpincreatepinallocation.PartyNumberRequestInfo{
		PinLength: pinLength,
		VNS: []string{
			params.CountryCode + vns,
		},
		Numbers: []string{
			params.CountryCode + numbers,
		},
	}
	return result, nil
}

// GetPayloadForPostRequest generates a payload for CreateAllocation POST APIs
func GetPayloadForPostRequest(defaults models.PayloadParameters, apiType string) ([]byte, interface{}, error) {
	var unmarshalledPayload interface{}
	var err error
	switch apiType {
	case constants.LaGreenvnCreateAllocation:
		unmarshalledPayload, err = generatePayloadForGreenVNCreateAllocationAPI(defaults)
		if err != nil {
			logger.Error(err)
			return nil, unmarshalledPayload, err
		}
	// TBD: To be fixed when GreenPin APIs are to tested with baremetrix.
	case constants.LaGreenpinCreatePinAllocation:
		// unmarshalledPayload, err = generatePayloadForGreenPinCreateAllocationAPI(defaults)
		// if err != nil {
		// 	logger.Error(err)
		// 	return nil, unmarshalledPayload, err
		// }
		unmarshalledPayload = greenpincreatepinallocation.Payload{
			TransactionId: "562203816180075682",
			A: greenpincreatepinallocation.PartyNumberRequestInfo{
				PinLength: 6,
				VNS: []string{
					"+913337593109",
				},
				Numbers: []string{
					"+916929586705",
				},
			},
			B: greenpincreatepinallocation.PartyNumberRequestInfo{
				PinLength: 6,
				VNS: []string{
					"+912874892347",
				},
				Numbers: []string{
					"+9123454529280",
				},
			},
			Usage: "twoway",
			DeallocationPolicy: greenvncreateallocation.DeallocationPolicy{
				Duration: "10m",
			},
		}
	}

	jsonPayload, err := json.Marshal(unmarshalledPayload)
	if err != nil {
		logger.Error(err)
		return nil, unmarshalledPayload, err
	}
	return jsonPayload, unmarshalledPayload, nil
}

// GetHeaderForHTTPRequest gets the map of headers
func GetHeaderForHTTPRequest(endpoint models.Endpoint, keyValuePairs []models.Header) http.Header {
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("%s %s", endpoint.AuthorizationType, endpoint.AuthorizationToken))
	for _, h := range keyValuePairs {
		header.Set(h.Key, h.Value)
	}
	return header
}

// GenerateGreenVNCreateAllocation gets a URL with random GreenVNId
func GenerateGreenVNCreateAllocation(config *models.Config, act models.Act, result chan *greenvncreateallocation.Tuple) {
	// prepping a http client
	method := "POST"
	client := &http.Client{
		Timeout: time.Duration(config.Instance.Specs.PrerequisiteSpecs.HttpTimeoutInSeconds) * time.Second,
	}
	jsonPayload, unmarshalledPayload, err := GetPayloadForPostRequest(act.DefaultParameters, constants.LaGreenvnCreateAllocation)
	if err != nil {
		logger.Error(err)
		result <- nil
	}
	// fmt.Println(string(jsonPayload))
	request, err := http.NewRequest(method, strings.ReplaceAll(strings.Split(utils.FindItemFromListWithKey(act.Allocateurl, constants.LaGreenvnCreateAllocation).(models.Allocateurl).URL, "/GREENVN_ID")[0], "SID", act.Endpoint.Sid), strings.NewReader(string(jsonPayload)))
	if err != nil {
		logger.Error(err)
		result <- nil
	}
	request.Header = GetHeaderForHTTPRequest(act.Endpoint, act.Headers)
	var response *http.Response
	for {
		response, err = client.Do(request)
		if err != nil {
			//logger.Errorf("%v", err)
			//logger.Errorf("error occured while making client connections, retrying, please be patient....")
			result <- nil
		} else {
			var responseBody greenvncreateallocation.Response
			if response.StatusCode != 200 {
				logger.Infof("%d found:", response.StatusCode)
			}
			if responseBody.Success || response.StatusCode == http.StatusOK {
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					logger.Errorf("%v, %d", err, response.StatusCode)
					result <- nil
				}
				err = json.Unmarshal(body, &responseBody)
				if err != nil {
					logger.Errorf("%v, %d", err, response.StatusCode)
					result <- nil
				}
				result <- &greenvncreateallocation.Tuple{
					Payload:  unmarshalledPayload.(greenvncreateallocation.Payload),
					Response: responseBody,
				}
			} else {
				result <- nil
			}
			err = response.Body.Close()
			if err != nil {
				logger.Errorf("%v", err)
				result <- nil
			}
			break
		}
		client.CloseIdleConnections()
	}

}

// GenerateGreenPinCreateAllocation gets a URL with random GreenPinId
func GenerateGreenPinCreateAllocation(config *models.Config, act models.Act, result chan *greenpincreatepinallocation.Tuple) {
	// prepping a http client
	method := "POST"
	client := &http.Client{
		Timeout: time.Duration(config.Instance.Specs.PrerequisiteSpecs.HttpTimeoutInSeconds) * time.Second,
	}
	jsonPayload, unmarshalledPayload, err := GetPayloadForPostRequest(act.DefaultParameters, constants.LaGreenpinCreatePinAllocation)
	if err != nil {
		logger.Error(err)
		result <- nil
	}
	request, err := http.NewRequest(method, strings.ReplaceAll(strings.Split(utils.FindItemFromListWithKey(act.Allocateurl, constants.LaGreenpinCreatePinAllocation).(models.Allocateurl).URL, "/GREENPIN_ID")[0], "SID", act.Endpoint.Sid), strings.NewReader(string(jsonPayload)))
	if err != nil {
		logger.Error(err)
		result <- nil
	}
	request.Header = GetHeaderForHTTPRequest(act.Endpoint, act.Headers)
	response, err := client.Do(request)
	if err != nil {
		logger.Errorf("%v, %d", err, response.StatusCode)
		result <- nil
	}
	var responseBody greenpincreatepinallocation.Response
	if responseBody.Success || response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			logger.Errorf("%v, %d", err, response.StatusCode)
			result <- nil
		}
		err = json.Unmarshal(body, &responseBody)
		if err != nil {
			logger.Errorf("%v, %d", err, response.StatusCode)
			result <- nil
		}
		result <- &greenpincreatepinallocation.Tuple{
			Payload:  unmarshalledPayload.(greenpincreatepinallocation.Payload),
			Response: responseBody,
		}
	} else {
		result <- nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Errorf("%v, %d", err, response.StatusCode)
			result <- nil
		}
	}(response.Body)
}

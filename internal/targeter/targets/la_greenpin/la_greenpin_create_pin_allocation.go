package la_greenpin

import (
	"strings"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/generator"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

// GetLAGreenPinCreatePinAllocationTarget provides a vegeta target for LA_GREENPIN_CREATE_PIN_ALLOCATION api.
func GetLAGreenPinCreatePinAllocationTarget(act models.Act, target *vegeta.Target) error {
	url := strings.ReplaceAll(strings.ReplaceAll(utils.FindItemFromListWithKey(act.Allocateurl, constants.LaGreenpinCreatePinAllocation).(models.Allocateurl).URL, "/GREENVN_ID", ""), "SID", act.Endpoint.Sid)
	header := generator.GetHeaderForHTTPRequest(act.Endpoint, act.Headers)
	jsonPayload, _, err := generator.GetPayloadForPostRequest(act.DefaultParameters, constants.LaGreenpinCreatePinAllocation)
	if err != nil {
		logger.Error(err)
		return err
	}
	method := ""
	if act.Method == "UNKNOWN" {
		method = "POST"
	} else {
		method = act.Method
	}
	target.Method = method
	target.Body = jsonPayload
	target.URL = url
	target.Header = header
	return nil
}

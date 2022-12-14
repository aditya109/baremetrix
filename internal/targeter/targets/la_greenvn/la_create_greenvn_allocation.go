package la_greenvn

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/generator"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
	"strings"
)

// GetLACreateGreenVNAllocationAPITarget provides a vegeta target for LA_GREENVN_CREATE_ALLOCATION api.
func GetLACreateGreenVNAllocationAPITarget(act models.Act, target *vegeta.Target) error {
	url := strings.ReplaceAll(strings.ReplaceAll(utils.FindItemFromListWithKey(act.Allocateurl, constants.LaGreenvnCreateAllocation).(models.Allocateurl).URL, "/GREENVN_ID", ""), "SID", act.Endpoint.Sid)
	header := generator.GetHeaderForHTTPRequest(act.Endpoint, act.Headers)
	jsonPayload, _, err := generator.GetPayloadForPostRequest(act.DefaultParameters, constants.LaGreenvnCreateAllocation)
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

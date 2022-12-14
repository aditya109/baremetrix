package la_greenvn

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/generator"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	logger "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

// GetLADeleteGreenVNAllocationAPITarget provides a vegeta target for LA_GREENVN_DELETE_ALLOCATION api.
func GetLADeleteGreenVNAllocationAPITarget(act models.Act, urls []string, target *vegeta.Target) error {
	url, err := utils.PopRandomItemFromList(&urls)
	if err != nil {
		logger.Error(err)
		return err
	}
	header := generator.GetHeaderForHTTPRequest(act.Endpoint, act.Headers)
	method := ""
	if act.Method == "UNKNOWN" {
		method = "DELETE"
	} else {
		method = act.Method
	}
	target.Method = method
	target.Body = nil
	target.URL = url
	target.Header = header
	return nil
}

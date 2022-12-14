package la_greenvn

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/generator"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	vegeta "github.com/tsenart/vegeta/lib"
)

// GetLAGetGreenVNAllocationDetailsAPITarget provides a vegeta target for LA_GREENVN_GET_ALLOCATION_DETAILS api.
func GetLAGetGreenVNAllocationDetailsAPITarget(act models.Act, targetUrl string, target *vegeta.Target) {
	url := targetUrl
	header := generator.GetHeaderForHTTPRequest(act.Endpoint, act.Headers)
	method := ""
	if act.Method == "UNKNOWN" {
		method = "GET"
	} else {
		method = act.Method
	}
	target.Method = method
	target.Body = nil
	target.URL = url
	target.Header = header
}

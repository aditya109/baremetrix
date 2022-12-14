package ingester

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils/targetstack"
)

// IngestFlow return corresponding flowsets for a particular flowtype.
func IngestFlow(config *models.Config, flow models.Flow, flowMapping map[string]interface{}) []targetstack.FlowStack {
	return flowMapping[flow.FlowId].(func(*models.Config, models.Flow) []targetstack.FlowStack)(config, flow)
}

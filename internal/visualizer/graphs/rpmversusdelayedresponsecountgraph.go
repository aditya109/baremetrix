package graphs

import (
	"fmt"
	"os"
	"strconv"

	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/helper"
	logger "github.com/sirupsen/logrus"
	"github.com/wcharczuk/go-chart/v2"
)

// GenerateRPMVersusDelayedResponseCountGraph generates rpm vs delayed response count graph.
func GenerateRPMVersusDelayedResponseCountGraph(visFileSpecs []models.FileSpecs, graphType models.GraphType, summary []models.Summary, play models.Play, timeStamp string, iteration int) error {
	graph := chart.BarChart{
		Title: graphType.Title,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars:     getBars(summary),
	}
	path, err := helper.GetFormattedFileName(models.SummarySpecificDirectives{
		Tenant:                  play.Tenant,
		PlayName:                play.Name,
		FileSpecs:               visFileSpecs[0],
		ShouldUseDirectives:     true,
		TimeStamp:               timeStamp,
		ShouldUseGraphIndicator: true,
		GraphType:               graphType.Name,
		Iteration:               strconv.Itoa(iteration),
	})
	if err != nil {
		logger.Error(err)
	}
	f, err := os.Create(path)
	if err != nil {
		logger.Error(err)
	}
	err = graph.Render(chart.SVG, f)
	if err != nil {
		logger.Error(err)
	}
	path, err = helper.GetFormattedFileName(models.SummarySpecificDirectives{
		Tenant:                  play.Tenant,
		PlayName:                play.Name,
		FileSpecs:               visFileSpecs[1],
		ShouldUseDirectives:     true,
		TimeStamp:               timeStamp,
		ShouldUseGraphIndicator: true,
		GraphType:               graphType.Name,
		Iteration:               strconv.Itoa(iteration),
	})
	if err != nil {
		logger.Error(err)
	}
	f, err = os.Create(path)
	if err != nil {
		logger.Error(err)
	}
	err = graph.Render(chart.PNG, f)
	if err != nil {
		logger.Error(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	return nil
}

func getBars(summary []models.Summary) []chart.Value {
	var bars []chart.Value
	for _, actSummary := range summary {
		bars = append(bars, chart.Value{
			Value: float64(actSummary.RequestCountOverExpectedLatency),
			Label: fmt.Sprint(actSummary.RPM),
		})
	}
	return bars
}

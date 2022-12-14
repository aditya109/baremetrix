package graphs

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/constants"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/utils"
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/pkg/helper"
	logger "github.com/sirupsen/logrus"
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"strconv"
)

// GenerateRPMVersusLatencyGraph generates rpm vs latency graph.
func GenerateRPMVersusLatencyGraph(visFileSpecs []models.FileSpecs, graphType models.GraphType, summary []models.Summary, play models.Play, timeStamp string, iteration int) error {
	meanSeries := getMeanSeries(graphType, summary)
	p50Series := getP50Series(graphType, summary)
	p95Series := getP95Series(graphType, summary)
	p99Series := getP99Series(graphType, summary)
	maxSeries := getMaxSeries(graphType, summary)

	graph := chart.Chart{
		Title: graphType.Title,
		XAxis: chart.XAxis{
			Name: graphType.XAxisLabel,
		},
		YAxis: chart.YAxis{
			Name: graphType.YAxisLabel,
		},
		Series: []chart.Series{
			meanSeries,
			p50Series,
			p95Series,
			p99Series,
			maxSeries,
		},
	}

	graph.Elements = []chart.Renderable{
		chart.LegendThin(&graph),
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
	} else {
		logger.Infof("%s graph generated.", constants.RpmVersusLatency)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	return nil
}

func getMeanSeries(graphType models.GraphType, summary []models.Summary) chart.ContinuousSeries {
	xValues := []float64{0}
	yValues := []float64{0}
	for _, v := range summary {
		xValues = append(xValues, float64(v.RPM))
		yValues = append(yValues, v.LatencyAtMeanUsage.Seconds())
	}
	logger.Info(xValues)
	logger.Info(yValues)
	return chart.ContinuousSeries{
		Name:    utils.FindItemFromListWithKey(graphType.InnerPlots, constants.MeanLatency).(models.InnerPlot).Name,
		XValues: xValues,
		YValues: yValues,
	}
}

func getP50Series(graphType models.GraphType, summary []models.Summary) chart.ContinuousSeries {
	xValues := []float64{0}
	yValues := []float64{0}

	for _, v := range summary {
		xValues = append(xValues, float64(v.RPM))
		yValues = append(yValues, v.LatencyAt50Usage.Seconds())
	}

	return chart.ContinuousSeries{
		Name:    utils.FindItemFromListWithKey(graphType.InnerPlots, constants.P50Latency).(models.InnerPlot).Name,
		XValues: xValues,
		YValues: yValues,
	}
}

func getP95Series(graphType models.GraphType, summary []models.Summary) chart.ContinuousSeries {
	xValues := []float64{0}
	yValues := []float64{0}
	for _, v := range summary {
		xValues = append(xValues, float64(v.RPM))
		yValues = append(yValues, float64(v.LatencyAt95Usage.Seconds()))
	}

	return chart.ContinuousSeries{
		Name:    utils.FindItemFromListWithKey(graphType.InnerPlots, constants.P95Latency).(models.InnerPlot).Name,
		XValues: xValues,
		YValues: yValues,
	}
}

func getP99Series(graphType models.GraphType, summary []models.Summary) chart.ContinuousSeries {
	xValues := []float64{0}
	yValues := []float64{0}
	for _, v := range summary {
		xValues = append(xValues, float64(v.RPM))
		yValues = append(yValues, float64(v.LatencyAt99Usage.Seconds()))
	}

	return chart.ContinuousSeries{
		Name:    utils.FindItemFromListWithKey(graphType.InnerPlots, constants.P99Latency).(models.InnerPlot).Name,
		XValues: xValues,
		YValues: yValues,
	}
}

func getMaxSeries(graphType models.GraphType, summary []models.Summary) chart.ContinuousSeries {
	xValues := []float64{0}
	yValues := []float64{0}
	for _, v := range summary {
		xValues = append(xValues, float64(v.RPM))
		yValues = append(yValues, float64(v.LatencyAtMaxUsage.Seconds()))
	}

	return chart.ContinuousSeries{
		Name:    utils.FindItemFromListWithKey(graphType.InnerPlots, constants.MaxLatency).(models.InnerPlot).Name,
		XValues: xValues,
		YValues: yValues,
	}
}

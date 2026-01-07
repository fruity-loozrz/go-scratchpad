package charts

import (
	"fmt"
	"os"

	"github.com/vicanso/go-charts/v2"
)

const (
	DefaultWidth  = 1200
	DefaultHeight = 600
)

type LineChartOptions struct {
	Title      string
	OutputPath string
	Width      int
	Height     int
	LegendLabels []string
}

func DrawLineChart(xLabels []string, seriesData [][]float64, opts LineChartOptions) error {
	width := opts.Width
	if width == 0 {
		width = DefaultWidth
	}

	height := opts.Height
	if height == 0 {
		height = DefaultHeight
	}

	legendLabels := opts.LegendLabels
	if len(legendLabels) == 0 {
		legendLabels = []string{"Series"}
	}

	p, err := charts.LineRender(
		seriesData,
		charts.TitleTextOptionFunc(opts.Title),
		charts.XAxisDataOptionFunc(xLabels),
		charts.LegendLabelsOptionFunc(legendLabels, charts.PositionRight),
		charts.WidthOptionFunc(width),
		charts.HeightOptionFunc(height),
	)
	if err != nil {
		return fmt.Errorf("failed to render chart: %w", err)
	}

	buf, err := p.Bytes()
	if err != nil {
		return fmt.Errorf("failed to get chart bytes: %w", err)
	}

	if err := os.WriteFile(opts.OutputPath, buf, 0644); err != nil {
		return fmt.Errorf("failed to write chart file: %w", err)
	}

	return nil
}

func FormatFloatLabels(values []float64, format string) []string {
	labels := make([]string, len(values))
	for i, v := range values {
		labels[i] = fmt.Sprintf(format, v)
	}
	return labels
}

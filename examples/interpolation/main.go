package main

import (
	"fmt"
	"strings"

	"github.com/fogleman/ease"
	"github.com/fruity-loozrz/go-scratchpad/examples/charts"
	"github.com/fruity-loozrz/go-scratchpad/internal/math/interpolation"
)

const resolution = 0.01

type Example struct {
	Number      string
	Name        string
	Description string
	EasingFns   []func(float64) float64
}

var (
	xs = []float64{0, 1, 2, 3, 4, 5}
	ys = []float64{7, 3, 5, 1, 2, 4}
)

func main() {
	examples := []Example{
		example001Linear(),
		example002QuadraticIn(),
		example003QuadraticOut(),
		example004QuadraticInOut(),
		example005CubicMixed(),
		example006ElasticOut(),
		example007BounceOut(),
		example008MixedVariety(),
	}

	fmt.Printf("Generating %d interpolation examples...\n\n", len(examples))

	for _, ex := range examples {
		fmt.Printf("Generating %s: %s\n", ex.Number, ex.Name)
		if err := generateExample(ex); err != nil {
			fmt.Printf("  ERROR: %v\n", err)
		} else {
			fmt.Printf("  Success!\n")
		}
	}

	fmt.Printf("\nAll charts generated in examples/interpolation/charts/\n")
}

func example001Linear() Example {
	return Example{
		Number:      "001",
		Name:        "Linear",
		Description: "Linear interpolation between all points",
		EasingFns: []func(float64) float64{
			ease.Linear,
			ease.Linear,
			ease.Linear,
			ease.Linear,
			ease.Linear,
		},
	}
}

func example002QuadraticIn() Example {
	return Example{
		Number:      "002",
		Name:        "Quadratic In",
		Description: "Accelerating quadratic easing between all points",
		EasingFns: []func(float64) float64{
			ease.InQuad,
			ease.InQuad,
			ease.InQuad,
			ease.InQuad,
			ease.InQuad,
		},
	}
}

func example003QuadraticOut() Example {
	return Example{
		Number:      "003",
		Name:        "Quadratic Out",
		Description: "Decelerating quadratic easing between all points",
		EasingFns: []func(float64) float64{
			ease.OutQuad,
			ease.OutQuad,
			ease.OutQuad,
			ease.OutQuad,
			ease.OutQuad,
		},
	}
}

func example004QuadraticInOut() Example {
	return Example{
		Number:      "004",
		Name:        "Quadratic InOut",
		Description: "Accelerating then decelerating quadratic easing between all points",
		EasingFns: []func(float64) float64{
			ease.InOutQuad,
			ease.InOutQuad,
			ease.InOutQuad,
			ease.InOutQuad,
			ease.InOutQuad,
		},
	}
}

func example005CubicMixed() Example {
	return Example{
		Number:      "005",
		Name:        "Cubic Mixed",
		Description: "Mixed cubic easing functions (In, Out, InOut pattern)",
		EasingFns: []func(float64) float64{
			ease.InCubic,
			ease.OutCubic,
			ease.InOutCubic,
			ease.InCubic,
			ease.OutCubic,
		},
	}
}

func example006ElasticOut() Example {
	return Example{
		Number:      "006",
		Name:        "Elastic Out",
		Description: "Elastic (spring-like) decelerating easing between all points",
		EasingFns: []func(float64) float64{
			ease.OutElastic,
			ease.OutElastic,
			ease.OutElastic,
			ease.OutElastic,
			ease.OutElastic,
		},
	}
}

func example007BounceOut() Example {
	return Example{
		Number:      "007",
		Name:        "Bounce Out",
		Description: "Bouncing decelerating easing between all points",
		EasingFns: []func(float64) float64{
			ease.OutBounce,
			ease.OutBounce,
			ease.OutBounce,
			ease.OutBounce,
			ease.OutBounce,
		},
	}
}

func example008MixedVariety() Example {
	return Example{
		Number:      "008",
		Name:        "Mixed Variety",
		Description: "Various different easing functions for each segment",
		EasingFns: []func(float64) float64{
			ease.Linear,
			ease.InOutQuad,
			ease.OutElastic,
			ease.InOutCubic,
			ease.OutBounce,
		},
	}
}

func generateExample(ex Example) error {
	interpolatedX, interpolatedY := interpolate(xs, ys, ex.EasingFns, resolution)
	filename := fmt.Sprintf("%s-%s.png", ex.Number, slugify(ex.Name))
	title := fmt.Sprintf("%s - %s", ex.Number, ex.Name)
	return drawChart(interpolatedX, interpolatedY, filename, title)
}

func interpolate(xs, ys []float64, easingFns []func(float64) float64, resolution float64) ([]float64, []float64) {
	var interpolatedX, interpolatedY []float64
	for x := 0.0; x <= xs[len(xs)-1]; x += resolution {
		y, err := interpolation.PieceWiseEaseInterpolate(xs, ys, easingFns, x)
		if err != nil {
			panic(err)
		}
		interpolatedX = append(interpolatedX, x)
		interpolatedY = append(interpolatedY, y)
	}
	return interpolatedX, interpolatedY
}

func drawChart(interpolatedX, interpolatedY []float64, filename, title string) error {
	xLabels := charts.FormatFloatLabels(interpolatedX, "%.1f")
	seriesData := [][]float64{interpolatedY}
	outputPath := fmt.Sprintf("examples/interpolation/charts/%s", filename)

	return charts.DrawLineChart(xLabels, seriesData, charts.LineChartOptions{
		Title:        title,
		OutputPath:   outputPath,
		LegendLabels: []string{"Interpolated"},
	})
}

func slugify(name string) string {
	var result strings.Builder
	for _, c := range name {
		if c >= 'A' && c <= 'Z' {
			result.WriteRune(c + 32)
		} else if c >= 'a' && c <= 'z' {
			result.WriteRune(c)
		} else if c >= '0' && c <= '9' {
			result.WriteRune(c)
		} else if c == ' ' {
			result.WriteRune('-')
		}
	}
	return result.String()
}

package main

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/charts"
)

// Generate random data for bar chart
func generateBarItems() []opts.BarData {
	return []opts.BarData{
		{Value: 40}, // Tiger
		{Value: 31}, // Whale
		{Value: 25}, // Bee
		{Value: 100}, // Eagle
		{Value: 6}, // Frog
		{Value: 6}, // Gopher
		{Value: 25}, // Shark
		{Value: 11}, // Croc
	}
}

func createBarChart() {
	// Create a new bar instance
	bar := charts.NewBar()

	// Set global options
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Animal Data",
		Subtitle: "Comparing Animal Top Speeds (mph)",
	}),
	 charts.WithXAxisOpts(opts.XAxis{
        AxisLabel: &opts.AxisLabel{
            Rotate: 30,
        },
    }),
)

	// Put data into instance
	bar.SetXAxis([]string{"Tiger", "Blue Whale", "Honey Bee", "Bald Eagle", "Tree Frog", "Gopher", "Great White Shark", "Crocodile (land)"}).
	    AddSeries("Top Speed (mph)", generateBarItems())
	f, _ := os.Create("bar.html")
	_ = bar.Render(f)
}

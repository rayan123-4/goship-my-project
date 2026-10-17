package main

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/opts"
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
	items := make([opts.BarData, 0])
	for i := 0; i < 6; i++ {
		    items = append(items, opts.BarData{Value: rand.Intn(300)})
		}
		return items
}

func createBarChart() {
	// Create a new bar instance
	bar := charts.NewBar()

	// Set global options
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Animal Data"
		Subtitle: "Comparing Different Animal Average Weights"
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Tiger", "Blue Whale", "Honey Bee", "Bald Eagle", "Tree Frog", "Gopher", "Great White Shark", "Crocodile (land)"}).
	    AddSeries("Top Speed (mph", generateBarItems()).
	    AddSeries("Catagory B", generateBarItems()).
	f, _ := os.Create("bar.html")
	_ = bar.Render(f)
}

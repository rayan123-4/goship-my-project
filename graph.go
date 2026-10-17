package main

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/opts"
)

// Generate random data for bar chart
func generateBarItems() []opts.BarData {
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
	bar.SetXAxis([]string{"Tiger", "Blue Whale", "Honey Bee", "Bald Eagle", "Red-Eyed Tree Frog", "Gopher", "Great White Shark", "Crocodile (land)"}).
	    AddSeries("Top Speed (mph", speedItems()).
	    AddSeries("Catagory B", generateBarItems()).
	f, _ := os.Create("bar.html")
	_ = bar.Render(f)
}

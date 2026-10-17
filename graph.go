package main

import (
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/opts"
)

// Generate random data for bar chart
func generateBarItems() []opts.BarData {
	items := make([opts.BarData])
	for i := 0; i < 6; i++ {
		    items = append(items, opts.BarData{Value: rand.Intn(300)})
		}
		return items
}

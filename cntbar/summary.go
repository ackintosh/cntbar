package cntbar

import (
	"sort"
	"strings"

	"github.com/gizak/termui"
)

type Summary struct {
	data map[string]int
}

func NewSummary() *Summary {
	return &Summary{
		data: map[string]int{},
	}
}

func (summary *Summary) CountUp(s string) {
	s = strings.Trim(s, "\n")
	summary.data[s] = summary.data[s] + 1
}

func (summary *Summary) getData() map[string]int {
	return summary.data
}

func (summary *Summary) getSortedKeys() []string {
	originalData := summary.getData()
	var keys []string
	for k, _ := range originalData {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func (summary *Summary) GetChart() *termui.BarChart {
	var bc *termui.BarChart
	keys := summary.getSortedKeys()
	data := summary.getData()

	if len(data) == 0 {
		return bc
	}

	var bcdata []int
	for _, key := range keys {
		bcdata = append(bcdata, data[key])
	}

	bc = termui.NewBarChart()
	bc.DataLabels = keys
	bc.Data = bcdata
	bc.Border.Label = "cntbar"
	bc.Width = len(data) * 6
	bc.Height = 10
	bc.X = 3
	bc.Y = 0
	bc.BarColor = termui.ColorGreen
	bc.NumColor = termui.ColorBlack

	return bc
}

package storage

import "testing"

func TestNewPlotDrive(t *testing.T) {

	pd := NewPlotDrive("/Users/rennbon/Downloads/Plots", []string{
		"/Users/rennbon/Downloads/Plots/201910271200_200000_320",
		"/Users/rennbon/Downloads/Plots/201910271200_100000_7616",
	}, 2015)

	t.Log(pd)
	for k, v := range pd.GetPlotFiles() {
		t.Log(k, v)
	}
}

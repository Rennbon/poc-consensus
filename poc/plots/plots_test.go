package plots

import (
	"github.com/rennbon/consensus/poc"
	"testing"
)

func TestNewPlots(t *testing.T) {
	poc.NewCoreProperties(&poc.PropertiesConfig{
		NumericAccountId: "201910271200",
		SoloServer:       "localhost:10001",
		PlotPaths:        []string{"/Users/rennbon/Downloads/Plots/"},
	})
	p := NewPlots("201910271200")
	for k, v := range p.GetChunkPartStartNonces() {
		t.Log(k, v)
	}

}

func Test_collectPlotFiles(t *testing.T) {
	m := collectPlotFiles([]string{"/Users/rennbon/Downloads/Plots/"}, "201910271200")
	for k, v := range m {
		t.Log(k, v)
	}
}

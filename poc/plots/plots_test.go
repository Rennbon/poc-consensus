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
		ChunkPartNonces:  1,
	})
	p := NewPlots("201910271200")

	for _, pd := range p.GetPlotDrives() {
		for _, pf := range pd.GetPlotFiles() {
			lpch, err := pf.GetLoadedParts(3838)
			if err != nil {
				t.Error(err)
				return
			}
			for ch := range lpch {
				t.Log(ch.ChunkPartStartNonce.String(), len(ch.Scoops))
			}
			t.Log("------- finish -------")
		}
	}
}

func Test_collectPlotFiles(t *testing.T) {
	m := collectPlotFiles([]string{"/Users/rennbon/Downloads/Plots/"}, "201910271200")
	for k, v := range m {
		t.Log(k, v)
	}
}

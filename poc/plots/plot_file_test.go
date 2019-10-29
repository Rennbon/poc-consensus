package plots

import (
	"testing"
	"time"
)

func TestAA(t *testing.T) {

	t.Log(time.Now().UnixNano() / 1e6)
}

func TestNewPlotFile(t *testing.T) {
	pf := NewPlotFile("/Users/rennbon/Downloads/Plots/201910271200_100000_7616", 1)
	lpch, err := pf.GetLoadedParts(1212)
	if err != nil {
		t.Error(err)
		return
	}
	for ch := range lpch {
		t.Log(ch.ChunkPartStartNonce.String())
	}
	t.Log("finish")
}

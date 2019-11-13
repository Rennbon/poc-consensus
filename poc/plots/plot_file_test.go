package plots

import (
	"testing"
	"time"
)

func TestAA(t *testing.T) {

	t.Log(time.Now().UnixNano() / 1e6)
}

func TestNewPlotFile(t *testing.T) {
	pf := NewPlotFile("/Users/rennbon/Downloads/Plots/201910271200_200000_320", 200)
	t.Log("size:", pf.GetSize())
	t.Log("number of parts:", pf.getNumberOfParts())
	t.Log("staggeramt:", pf.GetStaggeramt())
	t.Log("plots:", pf.GetPlots())
	t.Log("start nonce:", pf.GetStartnonce())
	t.Log("num of chunks:", pf.GetNumberOfChunks())

	for k, v := range pf.getChunkPartStartNonces() {
		t.Log("chunk part start nonce:", k, " size:", v)
	}

	lpch, err := pf.GetLoadedParts(3838)
	if err != nil {
		t.Error(err)
		return
	}
	i := 0
	for ch := range lpch {
		i++
		t.Log(i, ch.ChunkPartStartNonce.String(), len(ch.Scoops))
	}
}

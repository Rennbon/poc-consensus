package plots

import (
	"math/big"
	"testing"
	"time"
)

func TestAA(t *testing.T) {

	t.Log(time.Now().UnixNano() / 1e6)
}

func TestNewPlotFile(t *testing.T) {
	signature := make([]byte, 32)
	signature[0] = 1
	signature[1] = 2
	signature[2] = 3
	signature[3] = 4
	signature[4] = 5
	pf := NewPlotFile("/Users/rennbon/Downloads/Plots/201910271200_200000_320", 200)
	t.Log("size:", pf.GetSize())
	t.Log("number of parts:", pf.getNumberOfParts())
	t.Log("staggeramt:", pf.GetStaggeramt())
	t.Log("plots:", pf.GetPlots())
	t.Log("start nonce:", pf.GetStartnonce())
	t.Log("num of chunks:", pf.GetNumberOfChunks())
	plotcal := new(PlotCalculatorImpl)
	for k, v := range pf.getChunkPartStartNonces() {
		t.Log("chunk part start nonce:", k, " size:", v)
	}

	lpch, err := pf.GetLoadedParts(3838)
	if err != nil {
		t.Error(err)
		return
	}
	i := 0
	length := 0
	for ch := range lpch {
		i++
		t.Log(i, ch.ChunkPartStartNonce.String(), len(ch.Scoops))
		if length == 0 {

			length := len(ch.Scoops) / 64
			t.Log("length:", length)
			for i := 0; i < length; i++ {
				hit := plotcal.CalculateHit2(signature, ch.Scoops[i*64:(i+1)*64])
				str, _ := time.ParseDuration(hit.Div(hit, big.NewInt(18325193796)).String() + "us")
				t.Log("nonce:", ch.ChunkPartStartNonce.String(), " deadline:", str)
			}
		}
	}
}

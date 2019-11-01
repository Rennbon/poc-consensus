package round

import (
	"github.com/rennbon/consensus/poc/plots"
	"testing"
)

func Test_calcScoopNumber(t *testing.T) {
	r := &Round{}
	signature := make([]byte, 32)
	signature[0] = 1
	signature[1] = 2
	signature[2] = 3
	signature[3] = 4
	signature[4] = 5
	scoop := r.calcScoopNumber(1230000, signature)
	t.Log(scoop)
}

func Test_calculateResult(t *testing.T) {
	t.Log("finish")
	r := &Round{}
	signature := make([]byte, 32)
	signature[0] = 1
	signature[1] = 2
	signature[2] = 3
	signature[3] = 4
	signature[4] = 5
	pf := plots.NewPlotFile("/Users/rennbon/Downloads/Plots/201910271200_100000_7616", 1)
	lpch, err := pf.GetLoadedParts(1212)
	if err != nil {
		t.Error(err)
		return
	}
	for ch := range lpch {
		t.Log(ch.ChunkPartStartNonce.String())
		r.calculateResult(ch.Scoops, signature, ch.ChunkPartStartNonce)

	}

}

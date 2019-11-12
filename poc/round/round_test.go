package round

import (
	"github.com/magiconair/properties/assert"
	"github.com/rennbon/consensus/poc/plots"
	"github.com/rennbon/consensus/util"
	"testing"
	"time"
)

func Test_calcScoopNumber(t *testing.T) {
	r := &Round{}
	oclchecker, err := util.NewOCLChecker(0, 2)
	if err != nil {
		t.Error(err)
		return
	}
	signature := make([]byte, 32)
	signature[0] = 1
	signature[1] = 2
	signature[2] = 3
	signature[3] = 4
	signature[4] = 5
	scoop := r.calcScoopNumber(1230000, signature)
	plotcal := new(plots.PlotCalculatorImpl)
	scoop2 := plotcal.CalculateScoop(signature, 1230000)
	assert.Equal(t, scoop, scoop2)
	deadline := plotcal.CalculateDeadline(201910271200, 1, signature, scoop, 18325193796, 1)
	t.Log(deadline.Int64())

	t.Log(time.ParseDuration(deadline.String() + "us"))
	/*miner := util.NewMiningPlot(201910271200, 1000)
	scoopdata := miner.GetScoop(scoop)
	deadline2 := oclchecker.FindLowest(signature, scoopdata)
	t.Log(deadline2)*/
	//oclchecker.FindLowest(signature)
	//oclchecker.FindLowest(signature, scoop)
	//plot := util.NewMiningPlot(201910271200, scoop)
	t.Log(oclchecker)
}

func Test_calculateResult(t *testing.T) {
	t.Log("finish")
	//r := &Round{}
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

		//r.calculateResult(ch.Scoops, signature, ch.ChunkPartStartNonce)

	}

}

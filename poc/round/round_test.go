package round

import (
	"github.com/magiconair/properties/assert"
	"github.com/rennbon/consensus/poc/miner"
	"github.com/rennbon/consensus/poc/plots"
	"math/big"
	"testing"
	"time"
)

func Test_calcScoopNumber(t *testing.T) {
	r := &Round{}
	oclchecker, err := miner.NewOCLChecker(0, 2)
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
	scoopNum1 := r.calcScoopNumber(1230000, signature)
	plotcal := new(plots.PlotCalculatorImpl)
	scoopNum2 := plotcal.CalculateScoop(signature, 1230000)
	assert.Equal(t, scoopNum1, scoopNum2)
	t.Log(scoopNum1)
	deadline := plotcal.CalculateDeadline(201910271200, 300000000, signature, scoopNum1, 18325193796, 2)
	d1, _ := time.ParseDuration(deadline.String() + "us")

	miner := miner.NewMiningPlot(201910271200, 300000000)
	scoopdata := miner.GetScoop(scoopNum2)
	hit := plotcal.CalculateHit2(signature, scoopdata)
	d2, _ := time.ParseDuration(hit.Div(hit, big.NewInt(18325193796)).String() + "us")
	assert.Equal(t, d1, d2)
	//opencl
	deadline2 := oclchecker.FindLowest(signature, scoopdata)
	t.Log(deadline2)
	//oclchecker.FindLowest(signature)
	//oclchecker.FindLowest(signature, scoop)
	//plot := util.NewMiningPlot(201910271200, scoop)

	t.Log(oclchecker)
}

func Test_calculateResult(t *testing.T) {

	r := &Round{}
	signature := make([]byte, 32)
	signature[0] = 1
	signature[1] = 2
	signature[2] = 3
	signature[3] = 4
	signature[4] = 5
	pf := plots.NewPlotFile("/Users/rennbon/Downloads/Plots/201910271200_100000_7616", 1)
	lpch, err := pf.GetLoadedParts(3838)
	if err != nil {
		t.Error(err)
		return
	}
	for ch := range lpch {
		//t.Log(ch.ChunkPartStartNonce.String())

		r.calculateResult(ch.Scoops, signature, 1)

	}

}

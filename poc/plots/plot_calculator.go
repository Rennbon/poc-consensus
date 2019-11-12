package plots

import (
	"encoding/binary"
	"github.com/rennbon/consensus/util"
	"math/big"
)

type PlotCalculator interface {
	CalculateGenerationSignature(lastGenSig []byte, lastGenId int64) []byte

	CalculateScoop(genSig []byte, height int64) int64

	CalculateHit1(accountId int64, nonce int, genSig []byte, scoop, pocVersion int) *big.Int

	CalculateHit2(nonce int, genSig, scoopData []byte) *big.Int

	CalculateDeadline(accountId int64, nonce int, genSig []byte, scoop int, baseTarget int64, pocVersion int) *big.Int
}

type PlotCalculatorImpl struct {
}

func (o *PlotCalculatorImpl) CalculateHit1(accountId int64, nonce int, genSig []byte, scoop, pocVersion int) *big.Int {
	plot := util.NewMiningPlot(uint64(accountId), uint64(nonce))
	shabal256 := util.NewShabal256()
	shabal256.Write(genSig)
	plot.HashScoop(shabal256, scoop)
	hash := shabal256.Sum(nil)
	return big.NewInt(0).SetBytes([]byte{hash[7], hash[6], hash[5], hash[4], hash[3], hash[2], hash[1], hash[0]})
}
func (o *PlotCalculatorImpl) CalculateHit2(nonce int, genSig, scoopData []byte) *big.Int {
	md := util.NewShabal256()
	md.Reset()
	md.Write(genSig)
	st := nonce * util.SCOOP_SIZE
	md.Write(scoopData[st : st+util.SCOOP_SIZE])
	hash := md.Sum(nil)
	return big.NewInt(0).SetBytes([]byte{hash[7], hash[6], hash[5], hash[4], hash[3], hash[2], hash[1], hash[0]})
}
func (o *PlotCalculatorImpl) CalculateScoop(genSig []byte, height int64) int64 {
	shabal256 := util.NewShabal256()
	shabal256.Write(genSig)
	buff := make([]byte, 8)
	binary.LittleEndian.PutUint64(buff, uint64(height))
	shabal256.Write(buff)
	sum := shabal256.Sum(nil)
	hashnum := big.NewInt(0).SetBytes(sum)
	return hashnum.Mod(hashnum, big.NewInt(int64(util.SCOOPS_PER_PLOT))).Int64()
}
func (o *PlotCalculatorImpl) CalculateGenerationSignature(lastGenSig []byte, lastGenId int64) []byte {
	shabal256 := util.NewShabal256()
	shabal256.Write(lastGenSig)
	buff := make([]byte, 8)
	binary.LittleEndian.PutUint64(buff, uint64(lastGenId))
	shabal256.Write(buff)
	return shabal256.Sum(nil)
}

func (o *PlotCalculatorImpl) CalculateDeadline(accountId int64, nonce int, genSig []byte, scoop int, baseTarget int64, pocVersion int) *big.Int {
	hit := o.CalculateHit1(accountId, nonce, genSig, scoop, pocVersion)
	return hit.Div(hit, big.NewInt(baseTarget))
}

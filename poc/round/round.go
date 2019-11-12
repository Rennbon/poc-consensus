package round

import (
	"bytes"
	"encoding/binary"
	"github.com/rennbon/consensus/poc/plots"
	"github.com/rennbon/consensus/util"
	"math/big"
	"time"
)

type Round struct {
	poolMining     bool
	targetDeadline int64

	timer               time.Timer
	blockNumber         int64
	finishedBlockNumber int64
	baseTarget          int64
	roundStartDate      time.Time

	lowest                *big.Int
	bestCommittedDeadline int64

	lowestCommitted *big.Int

	runningChunkPartStartNonces []*big.Int
	plots                       plots.Plots
	generationSignature         []byte

	// generationSignature
	finishedLookup []string

	networkSuccessCount int64
	networkFailCount    int64
}

func (o *Round) initNewRound(plots plots.Plots) {

}

func (o *Round) isCurrentRound(currentBlockNumber int64, currentGenerationSignature []byte) bool {
	return o.blockNumber == currentBlockNumber &&
		bytes.Equal(o.generationSignature, currentGenerationSignature)
}

func (o *Round) calcScoopNumber(blockNumber int64, generationSignature []byte) int {
	if blockNumber > 0 && generationSignature != nil {
		buf := make([]byte, 32+8)
		copy(buf[:32], generationSignature)
		binary.LittleEndian.PutUint64(buf[32:], uint64(blockNumber))
		// generate new scoop number
		md := util.NewShabal256()
		md.Write(buf)
		hashnum := big.NewInt(0).SetBytes(md.Sum(nil))
		scoopnum := hashnum.Mod(hashnum, big.NewInt(int64(util.SCOOPS_PER_PLOT))).Int64()
		return int(scoopnum)
	}
	return 0
}
func (o *Round) calculateResult(scoops, generationSignature []byte, nonce int) *big.Int {
	md := util.NewShabal256()
	md.Reset()
	md.Write(generationSignature)
	st := nonce * util.SCOOP_SIZE
	md.Write(scoops[st : st+util.SCOOP_SIZE])
	hash := md.Sum(nil)
	return big.NewInt(0).SetBytes([]byte{hash[7], hash[6], hash[5], hash[4], hash[3], hash[2], hash[1], hash[0]})
}

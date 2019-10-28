package checker

import "math/big"

type CheckerResult interface {
	GetPlotFilePath() string
	GetBlockNumber() int64
	GetResult() *big.Int
	SetResult(result *big.Int)
	GetScoops() []byte
	GetLowestNonce() int
	GetChunkPartStartNonce() *big.Int
	GetGenerationSignature() []byte
}
type checkerResult struct {
	generationSignature []byte
	chunkPartStartNonce *big.Int
	blockNumber         int64
	result              *big.Int
	plotFilePath        string
	scoops              []byte
	lowestNonce         int
}

func NewChecker(blockNumber int64, generationSignature []byte, chunkPartStartNonce *big.Int, lowestNonce int, plotFilePath string, scoops []byte) CheckerResult {
	m := &checkerResult{
		generationSignature: generationSignature,
		chunkPartStartNonce: chunkPartStartNonce,
		blockNumber:         blockNumber,
		lowestNonce:         lowestNonce,
		plotFilePath:        plotFilePath,
		scoops:              scoops,
	}
	return m
}

func (o *checkerResult) GetPlotFilePath() string {
	return o.plotFilePath
}

func (o *checkerResult) GetBlockNumber() int64 {
	return o.blockNumber
}

func (o *checkerResult) GetResult() *big.Int {
	return o.result
}

func (o *checkerResult) SetResult(result *big.Int) {
	o.result.Set(result)
}

func (o *checkerResult) GetScoops() []byte {

	return o.scoops
}

func (o *checkerResult) GetLowestNonce() int {
	return o.lowestNonce
}

func (o *checkerResult) GetChunkPartStartNonce() *big.Int {

	return o.chunkPartStartNonce
}
func (o *checkerResult) GetGenerationSignature() []byte {

	return o.generationSignature
}

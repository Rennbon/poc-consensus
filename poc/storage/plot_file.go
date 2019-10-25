package storage

import (
	"math"
	"math/big"
	"path/filepath"
	"strconv"
	"strings"
	//"path/filepath"
)

type PlotFile struct {
	chunkPartStartNonces map[string]int64 //K: bigInt V:int64    key -> size
	filePath             string
	chunkPartNonces      int64
	numOfParts           int64
	numOfChunks          int64
	fileName             string
	address              int64
	startNonce           *big.Int
	plots                int64
	staggeramt           int64
	size                 int64
	pocVersion           PocVersion
}

func NewPlotFile(path string, chunkPartNonces int64) *PlotFile {
	pf := &PlotFile{
		filePath:        path,
		chunkPartNonces: chunkPartNonces,
	}
	fileName := filepath.Base(path)
	pf.fileName = fileName
	parts := strings.Split(fileName, "_")
	pf.address, _ = strconv.ParseInt(parts[0], 10, 64)
	pf.startNonce, _ = big.NewInt(0).SetString(parts[1], 10)
	pf.plots, _ = strconv.ParseInt(parts[2], 10, 64)
	if len(parts) > 3 {
		pf.pocVersion = POC_1
		staggeramt, _ := strconv.ParseInt(parts[3], 10, 64)
		pf.numOfParts = pf.calculateNumberOfParts(staggeramt)
		pf.numOfChunks = pf.plots / staggeramt
	}

	return pf
}
func (o *PlotFile) GetSize() int64 {
	return o.size
}

func (o *PlotFile) GetFilePath() string {
	return o.filePath
}

func (o *PlotFile) GetFilename() string {
	return o.fileName
}

func (o *PlotFile) GetAddress() int64 {
	return o.address
}

func (o *PlotFile) GetStartnonce() *big.Int {
	return o.startNonce
}
func (o *PlotFile) GetPlots() int64 {
	return o.plots
}
func (o *PlotFile) GetStaggeramt() int64 {
	return o.staggeramt
}

func (o *PlotFile) GetNumberOfChunks() int64 {
	return o.numOfChunks
}

func (o *PlotFile) getNumberOfParts() int64 {
	return o.numOfParts
}

func (o *PlotFile) SetNumberOfParts(numOfParts int64) {
	o.numOfParts = numOfParts
}

func (o *PlotFile) getChunkPartStartNonces() map[string]int64 {
	return o.chunkPartStartNonces
}
func (o *PlotFile) GetPocVersion() PocVersion {
	return o.pocVersion
}
func (o *PlotFile) calculateNumberOfParts(staggeramt int64) int64 {
	maxNumberOfParts := int64(100)
	targetNoncesPerPart := int64(960000)
	if o.chunkPartNonces != 0 {
		targetNoncesPerPart = o.chunkPartNonces
	}

	// for CPU it should be much lower, ensures less idle.
	//targetNoncesPerPart = !CoreProperties.isUseOpenCl() ? targetNoncesPerPart / 10 : targetNoncesPerPart;

	// calculate numberOfParts based on target
	suggestedNumberOfParts := staggeramt/targetNoncesPerPart + 1

	// ensure stagger is dividable by numberOfParts, if not adjust numberOfParts
	for staggeramt%suggestedNumberOfParts != 0 && suggestedNumberOfParts < maxNumberOfParts {
		suggestedNumberOfParts += 1
	}
	// fallback if number of parts could not be calculated in acceptable range
	if suggestedNumberOfParts >= maxNumberOfParts {
		suggestedNumberOfParts = int64(math.Floor(math.Sqrt(float64(staggeramt))))
		for staggeramt%suggestedNumberOfParts != 0 {
			suggestedNumberOfParts--
		}
	}
	return suggestedNumberOfParts
}

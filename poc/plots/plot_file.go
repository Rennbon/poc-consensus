package plots

import (
	"github.com/rennbon/consensus/util"
	"github.com/sirupsen/logrus"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
		filePath:             path,
		chunkPartNonces:      chunkPartNonces,
		chunkPartStartNonces: make(map[string]int64),
	}
	fileName := filepath.Base(path)
	pf.fileName = fileName
	parts := strings.Split(fileName, "_")
	pf.address, _ = strconv.ParseInt(parts[0], 10, 64)
	pf.startNonce, _ = big.NewInt(0).SetString(parts[1], 10)
	pf.plots, _ = strconv.ParseInt(parts[2], 10, 64)
	if len(parts) > 3 {
		pf.pocVersion = POC_1
		pf.staggeramt, _ = strconv.ParseInt(parts[3], 10, 64)
		pf.numOfParts = pf.calculateNumberOfParts(pf.staggeramt)
		pf.numOfChunks = pf.plots / pf.staggeramt
	} else {
		pf.pocVersion = POC_2
		pf.staggeramt = pf.plots
		pf.numOfParts = pf.calculateNumberOfParts(pf.staggeramt)
		pf.numOfChunks = 1
	}
	pf.size = pf.numOfChunks * pf.staggeramt * int64(util.PLOT_SIZE)
	chunkPartSize := pf.size / pf.numOfChunks / pf.numOfParts
	for chunkNumber := int64(0); chunkNumber < pf.numOfChunks; chunkNumber++ {
		for partNumber := int64(0); partNumber < pf.numOfParts; partNumber++ {
			// register a unique key for identification
			chunkPartStartNonce := big.NewInt(0).Add(pf.startNonce, big.NewInt(0).SetInt64(chunkNumber*pf.staggeramt+partNumber*(pf.staggeramt/pf.numOfParts)))
			pf.chunkPartStartNonces[chunkPartStartNonce.String()] = chunkPartSize
		}
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

type LoadedPart struct {
	ChunkPartStartNonce *big.Int
	Scoops              []byte
	PlotFilePath        string
}

func (o *PlotFile) GetLoadedParts(scoopNumber int64) (<-chan *LoadedPart, error) {
	file, err := os.OpenFile(o.filePath, os.O_RDONLY, os.ModeSocket)
	if err != nil {
		logrus.Warn("read file failed," + err.Error())
		return nil, nil
	}
	lpch := make(chan *LoadedPart, 1)
	scoop_size := int64(util.SCOOP_SIZE)
	currentScoopPosition := scoopNumber * o.GetStaggeramt() * scoop_size
	partSize := o.GetStaggeramt() / o.getNumberOfParts()
	partBuffer := make([]byte, partSize*scoop_size)
	// optimized plotFiles only have one chunk!
	go func() {
		for chunkNumber := int64(0); chunkNumber < o.GetNumberOfChunks(); chunkNumber++ {
			currentChunkPosition := chunkNumber * o.GetStaggeramt() * scoop_size
			for partNumber := int64(0); partNumber < o.getNumberOfParts(); partNumber++ {
				//TODO 验证区块高度和签名
				//读流填充partBuffer
				file.ReadAt(partBuffer, currentScoopPosition+currentChunkPosition)
				chunkPartStartNonce := big.NewInt(0).Add(o.GetStartnonce(), big.NewInt(chunkNumber*o.GetStaggeramt()+partNumber*partSize))
				//scoops := partBuffer
				lp := &LoadedPart{
					ChunkPartStartNonce: chunkPartStartNonce,
					Scoops:              partBuffer,
					PlotFilePath:        o.GetFilePath(),
				}
				lpch <- lp
			}
		}
		close(lpch)
	}()
	return lpch, nil
}

package plots

import (
	"github.com/rennbon/consensus/poc/miner"
	"github.com/sirupsen/logrus"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type PlotFile interface {
	GetSize() int64
	GetFilePath() string
	GetFilename() string
	GetAddress() int64
	GetStartnonce() *big.Int
	GetPlots() int64
	GetStaggeramt() int64
	GetNumberOfChunks() int64
	getNumberOfParts() int64
	SetNumberOfParts(numOfParts int64)
	getChunkPartStartNonces() map[string]int64
	GetPocVersion() PocVersion
	calculateNumberOfParts(staggeramt int64) int64
	GetLoadedParts(scoopNumber int64) (<-chan *LoadedPart, error)
}

//一组scoop 为64byte
type plotFile struct {
	//块数 及其 对应的 byte数组大小
	chunkPartStartNonces map[string]int64 //K: bigInt V:int64    key -> size
	filePath             string
	//一个chunk中多少nonce
	chunkPartNonces int64
	//分割的块数
	numOfParts int64
	//v2下为 1
	numOfChunks int64
	fileName    string
	address     int64
	//开始的nonce
	startNonce *big.Int
	//nonce数
	plots int64
	//v2 等同nonce数
	staggeramt int64
	//总plot大小
	size int64
	//版本 只支持v2
	pocVersion PocVersion
}

//chunkPartNonces: 切分间隔，内部会通过算法数据晃动
func NewPlotFile(path string, chunkPartNonces int64) PlotFile {
	pf := &plotFile{
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
	pf.size = pf.numOfChunks * pf.staggeramt * int64(miner.PLOT_SIZE)
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
func (o *plotFile) GetSize() int64 {
	return o.size
}

func (o *plotFile) GetFilePath() string {
	return o.filePath
}

func (o *plotFile) GetFilename() string {
	return o.fileName
}

func (o *plotFile) GetAddress() int64 {
	return o.address
}

func (o *plotFile) GetStartnonce() *big.Int {
	return o.startNonce
}
func (o *plotFile) GetPlots() int64 {
	return o.plots
}
func (o *plotFile) GetStaggeramt() int64 {
	return o.staggeramt
}

func (o *plotFile) GetNumberOfChunks() int64 {
	return o.numOfChunks
}

func (o *plotFile) getNumberOfParts() int64 {
	return o.numOfParts
}

func (o *plotFile) SetNumberOfParts(numOfParts int64) {
	o.numOfParts = numOfParts
}

func (o *plotFile) getChunkPartStartNonces() map[string]int64 {
	return o.chunkPartStartNonces
}
func (o *plotFile) GetPocVersion() PocVersion {
	return o.pocVersion
}

// splitting into parts is not needed, but it seams to improve speed and enables us
// to have steps of nearly same size
func (o *plotFile) calculateNumberOfParts(staggeramt int64) int64 {
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

//获取指定scoop number的等差排列数据，64byte一组
func (o *plotFile) GetLoadedParts(scoopNumber int64) (<-chan *LoadedPart, error) {
	file, err := os.OpenFile(o.filePath, os.O_RDONLY, os.ModeSocket)
	if err != nil {
		logrus.Warn("read file failed," + err.Error())
		return nil, nil
	}
	lpch := make(chan *LoadedPart, 1)
	scoop_size := int64(miner.SCOOP_SIZE)
	currentScoopPosition := scoopNumber * o.GetStaggeramt() * scoop_size

	partSize := o.GetStaggeramt() / o.getNumberOfParts()
	partBuffer := make([]byte, partSize*scoop_size)
	// optimized plotFiles only have one chunk!
	go func() {
		for chunkNumber := int64(0); chunkNumber < o.GetNumberOfChunks(); chunkNumber++ {
			currentChunkPosition := chunkNumber * o.GetStaggeramt() * int64(miner.PLOT_SIZE)
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

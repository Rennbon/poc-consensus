package storage

import (
	"log"
	"math/big"
	"path/filepath"
	"strconv"
	"strings"
)

type plots struct {
	plotDrives           []*PlotDrive
	chunkPartStartNonces map[string]int64
}

func NewPlots(numericAccountId string) {
	o := plots{
		plotDrives:           make([]*PlotDrive, 0, 256),
		chunkPartStartNonces: make(map[string]int64),
	}
	plotFilesLookup := collectPlotFiles(CoreProperties.getPlotPaths(), numericAccountId)
}

func collectPlotFiles(plotDirectories []string, numericAccountId string) map[string][]string {
	//val []path
	plotFilesLookup := make(map[string][]string)
	for _, plotDirectory := range plotDirectories {

		files, _ := filepath.Glob(plotDirectory)
		plotFilePaths := make([]string, 0, len(files))
		for _, fp := range files {
			if strings.Contains(fp, numericAccountId) {
				plotFilePaths = append(plotFilePaths, fp)
			}
		}
		plotFilesLookup[plotDirectory] = plotFilePaths
	}
	return plotFilesLookup
}

/* total number of bytes of all plotFiles */
func (o *plots) GetSize() int64 {
	size := int64(0)
	for _, plotDrive := range o.plotDrives {
		size += plotDrive.GetSize()
	}
	return size
}
func (o *plots) GetPlotDrives() []*PlotDrive {
	return o.plotDrives
}
func (o *plots) printPlotFiles() {
	for _, pd := range o.GetPlotDrives() {
		for _, pf := range pd.GetPlotFiles() {
			log.Print(pf.GetFilePath())
		}
	}
}

/* gets plot file by plot file start nonce. */
func (o *plots) GetPlotFileByPlotFileStartNonce(plotFileStartNonce int64) *PlotFile {
	for _, pd := range o.GetPlotDrives() {
		for _, pf := range pd.GetPlotFiles() {
			if strings.Contains(pf.GetFilename(), strconv.FormatInt(plotFileStartNonce, 10)) {
				return pf
			}
		}
	}
	return nil
}

/* gets chunk part start nonces. */
func (o *plots) GetChunkPartStartNonces() map[string]int64 {
	return o.chunkPartStartNonces
}

/* gets plot file by chunk part start nonce. */
func (o *plots) GetPlotFileByChunkPartStartNonce(chunkPartStartNonce *big.Int) *PlotFile {
	for _, pd := range o.GetPlotDrives() {
		for _, pf := range pd.GetPlotFiles() {
			if _, ok := pf.getChunkPartStartNonces()[chunkPartStartNonce.String()]; ok {
				return pf
			}
		}
	}
	return nil
}

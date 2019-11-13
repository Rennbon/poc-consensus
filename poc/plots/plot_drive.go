package plots

import "log"

type PlotDrive interface {
	GetPlotFiles() []PlotFile
	GetDirectory() string
	collectChunkPartStartNonces() map[string]int64
	GetSize() int64
	GetDrivePocVersion() PocVersion
}
type plotDrive struct {
	plotFiles []PlotFile
	directory string
}

func NewPlotDrive(directory string, plotFilePaths []string, chunkPartNonces int64) PlotDrive {
	o := &plotDrive{
		directory: directory,
		plotFiles: make([]PlotFile, 0, len(plotFilePaths)),
	}
	for _, v := range plotFilePaths {
		pf := NewPlotFile(directory+v, chunkPartNonces)
		o.plotFiles = append(o.plotFiles, pf)
		if pf.GetStaggeramt()%pf.getNumberOfParts() != 0 {
			log.Print("could not calculate valid numOfParts" + v)
		}
	}
	return o
}

func (o *plotDrive) GetPlotFiles() []PlotFile {
	return o.plotFiles
}
func (o *plotDrive) GetDirectory() string {
	return o.directory
}
func (o *plotDrive) collectChunkPartStartNonces() map[string]int64 {
	m := make(map[string]int64)
	for _, v := range o.plotFiles {
		cpsn := v.getChunkPartStartNonces()
		expectedSize := len(m) + len(cpsn)
		for ki, vi := range cpsn {
			m[ki] = vi
		}
		if expectedSize != len(m) {
			log.Print("possible overlapping plot-file '" + v.GetFilePath() + "', please check your plots.")
		}
	}

	return m
}
func (o *plotDrive) GetSize() int64 {
	size := int64(0)
	for _, v := range o.plotFiles {
		size += v.GetSize()
	}
	return size
}
func (o *plotDrive) GetDrivePocVersion() PocVersion {
	pv := PocVersion(0)
	for _, v := range o.plotFiles {
		if pv == 0 {
			pv = v.GetPocVersion()
		} else if pv != v.GetPocVersion() {
			return 0
		}
	}
	return pv
}

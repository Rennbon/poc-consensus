package storage

import "log"

type PlotDrive struct {
	plotFiles []*PlotFile
	directory string
}

func NewPlotDrive(directory string, plotFilePaths []string, chunkPartNonces int64) *PlotDrive {
	o := &PlotDrive{
		directory: directory,
		plotFiles: make([]*PlotFile, 0, len(plotFilePaths)),
	}
	for _, v := range plotFilePaths {
		pf := NewPlotFile(v, chunkPartNonces)
		o.plotFiles = append(o.plotFiles, pf)
		if pf.GetStaggeramt()%pf.getNumberOfParts() != 0 {
			log.Print("could not calculate valid numOfParts" + v)
		}
	}
	return o
}

func (o *PlotDrive) GetPlotFiles() []*PlotFile {
	return o.plotFiles
}
func (o *PlotDrive) GetDirectory() string {
	return o.directory
}
func (o *PlotDrive) collectChunkPartStartNonces() map[string]int64 {
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
func (o *PlotDrive) GetSize() int64 {
	size := int64(0)
	for _, v := range o.plotFiles {
		size += v.GetSize()
	}
	return size
}
func (o *PlotDrive) GetDrivePocVersion() PocVersion {
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

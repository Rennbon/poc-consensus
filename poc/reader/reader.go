package reader

import (
	"github.com/ahmetb/go-linq"
	"github.com/rennbon/consensus/poc/plots"
	"github.com/rennbon/consensus/util"
	"github.com/sirupsen/logrus"
	"time"
)

type reader struct {
	numericAccountId string

	// data
	blockNumber         int64
	generationSignature []byte

	plots plots.Plots

	realCapacityLookup    map[string]int64
	realRemainingCapacity int64
	realCapacity          int64
	capacityLookup        map[string]int64

	remainingCapacity int64
	capacity          int64

	readerStartTime int64
	readerThreads   int
}

func NewReader(numericAccountId string) *reader {
	r := &reader{
		numericAccountId: numericAccountId,
		plots:            plots.NewPlots(numericAccountId),
	}
	r.capacityLookup = r.plots.GetChunkPartStartNonces()

	return r
}

func (o *reader) read(previousBlockNumber, blockNumber int64, generationSignature []byte, scoopNumber int, lastBestComittedDeadline int64, networkQuality int) {
	o.blockNumber = blockNumber
	o.generationSignature = generationSignature
	o.realCapacityLookup = make(map[string]int64)
	o.realCapacity = 0
	for chunkPartNonces, _ := range o.capacityLookup {
		realChunkPartNoncesCapacity := int64(0)
		plotFile := o.plots.GetPlotFileByChunkPartStartNonce(chunkPartNonces)
		if o.isCompatibleWithCurrentPoc(plotFile.GetPocVersion()) {
			realChunkPartNoncesCapacity = o.capacityLookup[chunkPartNonces]
		} else {
			realChunkPartNoncesCapacity = 2 * o.capacityLookup[chunkPartNonces]
		}
		o.realCapacityLookup[chunkPartNonces] = realChunkPartNoncesCapacity
		o.realCapacity += realChunkPartNoncesCapacity
	}
	o.remainingCapacity = o.plots.GetSize()
	o.capacity = o.plots.GetSize()
	o.realRemainingCapacity = o.realCapacity
	o.readerStartTime = util.ToMillisecond(time.Now())

	orderedPlotDrives := o.plots.GetPlotDrives()
	linq.From(orderedPlotDrives).Where(func(o interface{}) bool {
		return o.(*plots.PlotDrive).GetDrivePocVersion() != 0
	}).ToSlice(orderedPlotDrives)

	linq.From(orderedPlotDrives).Sort(
		func(i interface{}, j interface{}) bool {
			return i.(*plots.PlotDrive).GetSize() < j.(*plots.PlotDrive).GetSize()
		},
	).Sort(
		func(i interface{}, j interface{}) bool {
			return o.isCompatibleWithCurrentPoc(i.(*plots.PlotDrive).GetDrivePocVersion())
		},
	).ToSlice(orderedPlotDrives)

	for _, pd := range orderedPlotDrives {
		drivePocVersion := pd.GetDrivePocVersion()
		if drivePocVersion == 0 {
			logrus.Warn("Skipped '" + pd.GetDirectory() + "', different POC versions on one drive is not supported! (Workaround: put them in different directories and add them to 'plotFilePaths')")
		} else {
			if o.isCompatibleWithCurrentPoc(drivePocVersion) {
				//scoopNumber, blockNumber, generationSignature, pd
				//ReaderLoadDriveTask readerLoadDriveTask = context.getBean(ReaderLoadDriveTask.class);
				//readerLoadDriveTask.init(scoopNumber, blockNumber, generationSignature, pd);
				//readerPool.execute(readerLoadDriveTask);
			} else {
				//ReaderConvertLoadDriveTask readerConvertLoadDriveTask = context.getBean(ReaderConvertLoadDriveTask.class);
				//readerConvertLoadDriveTask.init(scoopNumber, blockNumber, generationSignature, plotDrive);
				//readerPool.execute(readerConvertLoadDriveTask);
			}
		}
	}

}
func (o *reader) isCompatibleWithCurrentPoc(pocVersion plots.PocVersion) bool {
	return plots.POC_2 == pocVersion
}

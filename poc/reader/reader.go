package reader

import "github.com/rennbon/consensus/poc/storage"

type reader struct {
	numericAccountId string

	// data
	blockNumber         int64
	generationSignature []byte

	plots storage.Plots

	realCapacityLookup    map[string]int64
	realRemainingCapacity int64
	realCapacity          int64
	capacityLookup        map[string]int64

	remainingCapacity int64
	capacity          int64

	readerStartTime int64
	readerThreads   int
}

func NewReader() *reader {

}

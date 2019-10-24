package util

import (
	"encoding/binary"
	"hash"
)

const (
	HASH_SIZE        int = 32
	HASHES_PER_SCOOP int = 2
	SCOOP_SIZE       int = HASHES_PER_SCOOP * HASH_SIZE
	SCOOPS_PER_PLOT  int = 4096 // original 1MB/plot = 16384
	PLOT_SIZE        int = SCOOPS_PER_PLOT * SCOOP_SIZE
	HASH_CAP         int = 4096
)

type MiningPlot interface {
	GetScoop(pos int) []byte
	HashScoop(shabal256 hash.Hash, pos int)
}
type miningPlot struct {
	data []byte
}

func NewMiningPlot(addr, nonce uint64) MiningPlot {
	m := &miningPlot{
		data: make([]byte, PLOT_SIZE),
	}
	buff := make([]byte, 16)
	binary.LittleEndian.PutUint64(buff[:8], addr)
	binary.LittleEndian.PutUint64(buff[8:], nonce)
	sb256 := NewShabal256()
	baseLen := len(buff)
	gendata := make([]byte, PLOT_SIZE+baseLen)
	copy(gendata[PLOT_SIZE:PLOT_SIZE+16], buff)
	for i := PLOT_SIZE; i > 0; i -= HASH_SIZE {
		sb256.Reset()
		len := PLOT_SIZE + baseLen - i
		if len > HASH_CAP {
			len = HASH_CAP
		}
		sb256.Write(gendata[i : i+len])
		sb256.Sum(gendata[i-HASH_SIZE : i])

	}
	sb256.Reset()
	sb256.Write(gendata)
	finalhash := sb256.Sum(gendata)
	for i := 0; i < PLOT_SIZE; i++ {
		m.data[i] = (byte)(gendata[i] ^ finalhash[i%HASH_SIZE])
	}
	return m
}
func (o *miningPlot) GetScoop(pos int) []byte {
	return o.data[pos*SCOOP_SIZE : (pos+1)*SCOOP_SIZE]
}
func (o *miningPlot) HashScoop(shabal256 hash.Hash, pos int) {
	shabal256.Write(o.data[pos*SCOOP_SIZE : (pos+1)*SCOOP_SIZE])
}

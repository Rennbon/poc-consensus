package util

import (
	"github.com/rennbon/consensus/poc"
	"testing"
)

func TestNewMiningPlot(t *testing.T) {
	poc.NewCoreProperties()
	plot := NewMiningPlot(201910271200, 100500)
	t.Log(plot)
}

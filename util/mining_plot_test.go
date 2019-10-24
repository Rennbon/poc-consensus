package util

import "testing"

func TestNewMiningPlot(t *testing.T) {

	plot := NewMiningPlot(0xABCDEF, 2000)
	t.Log(plot)
}

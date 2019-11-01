package util

import (
	"testing"
)

func TestNewMiningPlot(t *testing.T) {
	//poc.NewCoreProperties()
	plot := NewMiningPlot(201910271200, 100500)
	t.Log(plot)
}

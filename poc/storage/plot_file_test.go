package storage

import (
	"path/filepath"
	"testing"
)

func TestAA(t *testing.T) {

	t.Log(filepath.Base("./asb/name.jpg"))
}

func TestNewPlotFile(t *testing.T) {
	pf := NewPlotFile("/Users/rennbon/Downloads/Plots/201910271200_200000_320", 1)
	t.Log(pf)
}

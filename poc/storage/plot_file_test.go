package storage

import (
	"path/filepath"
	"testing"
)

func TestAA(t *testing.T) {

	t.Log(filepath.Base("./asb/name.jpg"))
}

func TestNewPlotFile(t *testing.T) {
	pf := NewPlotFile("/Users/rennbon/poc/test.log", 100)
	t.Log(pf)
}

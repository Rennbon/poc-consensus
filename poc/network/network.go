package network

import (
	"github.com/rennbon/consensus/poc/storage"
	"time"
)

type Network struct {
	blockNumber         int64
	timer               time.Timer
	generationSignature []byte
	mac                 string //unique system id
	plots               plots.Plots
}

func NewNetwork() *Network {

}

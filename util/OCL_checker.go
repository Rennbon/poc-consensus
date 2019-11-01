package util

import "C"
import (
	"fmt"

	"github.com/go-gl/cl/v1.2/cl"
)

const (
	BUFFER_PER_ITEM = PLOT_SIZE + 16

	MEM_PER_ITEM = 8 + 8 + BUFFER_PER_ITEM + 4 + SCOOP_SIZE
)

type OCLChecker struct {
	context       cl.Context
	queue         cl.CommandQueue
	kernel        [2]cl.Kernel
	workGroupSize [2]int64
	gensigMem     cl.Mem
	bestNum       cl.Mem
}

func NewOCLChecker(platformId uint32, deviceId int) (*OCLChecker, error) {

	ids := make([]cl.PlatformID, 1)
	numPlatforms := uint32(0)
	err := cl.GetPlatformIDs(uint32(len(ids)), &ids[0], &numPlatforms)
	if err != cl.SUCCESS {
		return nil, fmt.Errorf("getPlatformIds errcode: %d", err)
	}

	return nil, nil
}

/*func (o *OCLChecker) getPlatforms() {
	platforms, err := cl.GetPlatforms()
	if err != nil {

	}

	for k, v := range platforms {
		devices, err := v.GetDevices(cl.DeviceTypeGPU)
		if err != nil {
			continue
		}
		for ki, vi := range devices {

		}
	}
}
*/

package util

import "C"
import (
	"errors"
	"fmt"
	"github.com/go-gl/cl/v1.2/cl"
)

const (
	BUFFER_PER_ITEM = PLOT_SIZE + 16

	MEM_PER_ITEM = 8 + 8 + BUFFER_PER_ITEM + 4 + SCOOP_SIZE
)

var (
	Err_PlatformId_Invalid = errors.New("Invalid platform id")
)

type OCLChecker struct {
	context       cl.Context
	queue         cl.CommandQueue
	kernel        [2]cl.Kernel
	workGroupSize [2]int64
	gensigMem     cl.Mem
	bestNum       cl.Mem
}

func NewOCLChecker(platformId uint32, deviceId uint32) (*OCLChecker, error) {

	ids := make([]cl.PlatformID, 1)
	numPlatforms := uint32(0)
	err := cl.GetPlatformIDs(uint32(len(ids)), &ids[0], &numPlatforms)
	if err != cl.SUCCESS {
		return nil, fmt.Errorf("GetPlatformIDs errcode: %d", err)
	}
	if platformId >= numPlatforms {
		return nil, Err_PlatformId_Invalid
	}

	actualDid := uint32(0)
	err = cl.GetDeviceIDs(ids[0], cl.DEVICE_TYPE_ALL, 0, nil, &actualDid)
	if err != cl.SUCCESS {
		return nil, fmt.Errorf("GetDeviceIDs1 errcode: %d", err)
	}
	if deviceId >= actualDid {
		return nil, Err_PlatformId_Invalid
	}
	devices := make([]cl.DeviceId, actualDid)
	err = cl.GetDeviceIDs(ids[0], cl.DEVICE_TYPE_ALL, uint32(len(devices)), &devices[deviceId], &actualDid)
	if err != cl.SUCCESS {
		return nil, fmt.Errorf("GetDeviceIDs2 errcode: %d", err)
	}
	context := cl.CreateContext(nil, 1, &devices[deviceId], nil, nil, &err)
	if err != cl.SUCCESS {
		return nil, fmt.Errorf("CreateContext errcode: %d", err)
	}
	queue := cl.CreateCommandQueue(context, devices[deviceId], 0, &err)
	if err != cl.SUCCESS {
		return nil, fmt.Errorf("CreateContext errcode: %d", err)
	}
	ocl := &OCLChecker{
		context: context,
		queue:   queue,
	}

	return ocl, nil
}

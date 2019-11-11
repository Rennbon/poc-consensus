package util

import (
	"github.com/go-gl/cl/v1.2/cl"
	"math/rand"
	"testing"
	"unsafe"
)

func TestNewOCLChecker(t *testing.T) {
	oclchecker, err := NewOCLChecker(0, 2)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(oclchecker)
}

func TestCl(t *testing.T) {
	ids := make([]cl.PlatformID, 1)
	actual := uint32(0)
	err := cl.GetPlatformIDs(uint32(len(ids)), &ids[0], &actual)
	if err != cl.SUCCESS {
		t.Error(err)
		return
	}

	devices := make([]cl.DeviceId, 1)
	actualDid := uint32(0)
	err = cl.GetDeviceIDs(ids[0], cl.DEVICE_TYPE_GPU, uint32(len(devices)), &devices[0], &actualDid)
	if err != cl.SUCCESS {
		t.Error(err)
		return
	}
	t.Log("actualdid:", actualDid)
	device := devices[0]
	defer cl.ReleaseDevice(device)
	var errptr cl.ErrorCode
	context := cl.CreateContext(nil, 1, &device, nil, nil, &errptr)

	defer cl.ReleaseDevice(device)
	defer cl.ReleaseContext(context)

	contextInfos := [...]cl.ContextInfo{cl.CONTEXT_REFERENCE_COUNT, cl.CONTEXT_DEVICES, cl.CONTEXT_PROPERTIES, cl.CONTEXT_NUM_DEVICES}

	data := make([]byte, 1024)
	size := uint64(0)

	for _, info := range contextInfos {
		err := cl.GetContextInfo(context, info, 1024, unsafe.Pointer(&data[0]), &size)
		if err != cl.SUCCESS {
			t.Fail()
		}
	}
}

func TestCl2(t *testing.T) {

	data := make([]float32, 1024)
	for x := 0; x < len(data); x++ {
		data[x] = rand.Float32()*99 + 1
	}

	//Get Device
	var device cl.DeviceId
	err := cl.GetDeviceIDs(nil, cl.DEVICE_TYPE_GPU, 1, &device, nil)
	if err != cl.SUCCESS {
		t.Fatal("Failed to create device group")
	}
	var errptr *cl.ErrorCode

	//Create Computer Context
	context := cl.CreateContext(nil, 1, &device, nil, nil, errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		t.Fatal("couldnt create context")
	}
	defer cl.ReleaseContext(context)

	//Create Command Queue
	cq := cl.CreateCommandQueue(context, device, 0, errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		t.Fatal("couldnt create command queue")
	}
	defer cl.ReleaseCommandQueue(cq)

	//Create program
	srcptr := cl.Str(KernelSource)
	program := cl.CreateProgramWithSource(context, 1, &srcptr, nil, errptr)
	if errptr != nil && cl.ErrorCode(*errptr) != cl.SUCCESS {
		t.Fatal("couldnt create program")
	}
	t.Log(program)
}

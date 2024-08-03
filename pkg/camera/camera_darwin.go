package camera

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#cgo LDFLAGS: -framework AVFoundation
#cgo LDFLAGS: -framework CoreMediaIO
#include "camera_darwin.mm"
*/
import "C"
import (
	"fmt"

	"github.com/dqle/go-media-devices-state/pkg/common"
)

// IsCameraOn returns true is any camera in the system is ON
func IsCameraOnDarwin(logging bool) (bool, error) {
	isCameraOn := C.int(0)
	C.setCameraLogging(C.bool(logging))
	errCode := C.IsCameraOn(&isCameraOn)

	if errCode != common.ErrNoErr {
		var msg string
		switch errCode {
		case common.ErrOutOfMemory:
			msg = "IsCameraOn(): failed to allocate memory"
		case common.ErrAllDevicesFailed:
			msg = "IsCameraOn(): all devices failed to provide status"
		default:
			msg = fmt.Sprintf("IsCameraOn(): failed with error code: %d", errCode)
		}
		return false, fmt.Errorf("IsCameraOn(): %s", msg)
	}

	return isCameraOn == 1, nil
}

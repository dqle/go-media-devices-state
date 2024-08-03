package camera

import (
	"golang.org/x/sys/windows/registry"

	"github.com/dqle/go-media-devices-state/pkg/common"
)

func IsCameraOnWindows() (bool, error) {
	keyPath := `Software\Microsoft\Windows\CurrentVersion\CapabilityAccessManager\ConsentStore\webcam\NonPackaged`
	currentUser := registry.CURRENT_USER

	currentUserSubKeyList, err := common.GetDeviceSubKey(keyPath, currentUser)
	if err != nil {
		return false, err
	}

	isOn, err := common.GetDeviceStatus(currentUser, currentUserSubKeyList)
	if err != nil {
		return false, err
	}
	if isOn {
		return true, nil
	}

	return false, nil
}
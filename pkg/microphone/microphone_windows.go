package microphone

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

type keyList []string

func openKey(constant registry.Key, keyPath string) (registry.Key, error) {
	// Set access right
	queryValue := registry.QUERY_VALUE | registry.ENUMERATE_SUB_KEYS

	// Get microphone subkey
	currentUserKey, err := registry.OpenKey(constant, keyPath, uint32(queryValue))
	if err != nil {
		return registry.Key(0), err
	}

	return currentUserKey, nil
}

func getMicSubKey(keyPath string, k registry.Key) (keyList, error) {
	// Get CURRENT_USER or LOCAL_MACHINE subkey
	key, err := openKey(k, keyPath)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	subKeys, err := key.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	}

	return keyList(joinPath(subKeys, keyPath)), nil
}

func joinPath(subKeys []string, keyPath string) []string {
	newSubKeyList := []string{}
	for _, key := range subKeys {
		newSubKeyList = append(newSubKeyList, keyPath+key)
	}
	return newSubKeyList
}

func getMicOnStatus(k registry.Key, klist keyList) (bool, error) {
	// Get LastUsedTimeStop value for CURRENT_USER or LOCAL_MACHINE
	for _, key := range klist {
		micKey, err := openKey(k, key)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer micKey.Close()

		status, _, err := micKey.GetIntegerValue("LastUsedTimeStop")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if status == 0 {
			return true, nil
		}
	}

	return false, nil
}

// logging bool not used for Windows
func IsMicrophoneOnWindows(logging bool) (bool, error) {
	// Define where in registry we want to look for the microphone app list
	keyPath := `SOFTWARE\Microsoft\Windows\CurrentVersion\CapabilityAccessManager\ConsentStore\microphone\NonPackaged\`
	currentUser := registry.CURRENT_USER
	localMachine := registry.LOCAL_MACHINE

	// Get the list of application that uses microphone registry subkey
	currentUserSubKeyList, err := getMicSubKey(keyPath, currentUser)
	if err != nil {
		return false, err
	}

	localMachineSubKeyList, err := getMicSubKey(keyPath, localMachine)
	if err != nil {
		return false, err
	}

	isOn, err := getMicOnStatus(currentUser, currentUserSubKeyList)
	if err != nil {
		return false, err
	}
	if isOn {
		return true, nil
	}

	isOn, err = getMicOnStatus(localMachine, localMachineSubKeyList)
	if err != nil {
		return false, err
	}
	if isOn {
		return true, nil
	}

	return false, nil
}

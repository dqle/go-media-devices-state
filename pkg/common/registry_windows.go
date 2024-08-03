package common

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

type KeyList []string

func OpenKey(constant registry.Key, keyPath string) (registry.Key, error) {
	queryValue := registry.QUERY_VALUE | registry.ENUMERATE_SUB_KEYS

	key, err := registry.OpenKey(constant, keyPath, uint32(queryValue))
	if err != nil {
		return registry.Key(0), err
	}

	return key, nil
}

func GetDeviceSubKey(keyPath string, k registry.Key) (KeyList, error) {
	key, err := OpenKey(k, keyPath)
	if err != nil {
		return nil, err
	}
	defer key.Close()

	subKeys, err := key.ReadSubKeyNames(-1)
	if err != nil {
		return nil, err
	}

	return KeyList(JoinPath(subKeys, keyPath)), nil
}

func JoinPath(subKeys []string, keyPath string) []string {
	newSubKeyList := []string{}
	for _, key := range subKeys {
		newSubKeyList = append(newSubKeyList, keyPath+"\\"+key)
	}
	return newSubKeyList
}

func GetDeviceStatus(k registry.Key, klist KeyList) (bool, error) {
	for _, key := range klist {
		micKey, err := OpenKey(k, key)
		if err != nil {
			fmt.Printf("[GetMicOnStatus][OpenKey][key=%s]: %v\n", key, err)
			continue
		}
		defer micKey.Close()

		status, _, err := micKey.GetIntegerValue("LastUsedTimeStop")
		if err != nil {
			fmt.Printf("[GetMicOnStatus][GetIntegerValue('LastUsedTimeStop')]: %v\n", err)
			continue
		}
		if status == 0 {
			return true, nil
		}
	}

	return false, nil
}

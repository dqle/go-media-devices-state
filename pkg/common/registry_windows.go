package common

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

type keyList []string

func openKey(constant registry.Key, keyPath string) (registry.Key, error) {
	queryValue := registry.QUERY_VALUE | registry.ENUMERATE_SUB_KEYS

	key, err := registry.OpenKey(constant, keyPath, uint32(queryValue))
	if err != nil {
		return registry.Key(0), err
	}

	return key, nil
}

func getMicSubKey(keyPath string, k registry.Key) (keyList, error) {
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
		newSubKeyList = append(newSubKeyList, keyPath+"\\"+key)
	}
	return newSubKeyList
}

func getMicStatus(k registry.Key, klist keyList) (bool, error) {
	for _, key := range klist {
		micKey, err := openKey(k, key)
		if err != nil {
			fmt.Printf("[getMicStatus][openKey][key=%s]: %v\n", key, err)
			continue
		}
		defer micKey.Close()

		status, _, err := micKey.GetIntegerValue("LastUsedTimeStop")
		if err != nil {
			fmt.Printf("[getMicStatus][GetIntegerValue('LastUsedTimeStop')]: %v\n", err)
			continue
		}
		if status == 0 {
			return true, nil
		}
	}

	return false, nil
}

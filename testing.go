package collector

import (
	//	"strings"
	"fmt"
	"strconv"
	"strings"

	"github.com/StackExchange/wmi"
	"github.com/prometheus/common/log"
	"golang.org/x/sys/windows/registry"
)

func main() {
	version := getWindowsVersion()
	wmitesting := wmitest()
	fmt.Println("hello world")
	fmt.Println(version)
	if version > 8 {
		fmt.Println("Dagiau")
	} else {
		fmt.Println("Maziau")
	}

}

func getWindowsVersion() float64 {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Warn("Couldn't open registry", err)
	}
	defer k.Close()

	major, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		log.Warn("Couldn't open registry to determine current Windows version:", err)
	}

	log.Info("Detected Windows %d.%d\n", major)
	major_flt, err := strconv.ParseFloat(major, 64)
	log.Info(major_flt)
	return major_flt
}

type Win32_PerfRawData_Counters_ProcessorInformation struct {
	Name                  string
	AverageIdleTime       uint64
	C1TransitionsPersec   uint64
	C2TransitionsPersec   uint64
	C3TransitionsPersec   uint64
	ClockInterruptsPersec uint32
	DPCRate               uint32
}

func wmitest() {
	var dst []Win32_PerfRawData_Counters_ProcessorInformation
	q := queryAll(&dst)
	log.Info(q)
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}
	for _, data := range dst {
		log.Info(data)
		if strings.Contains(data.Name, "_Total") {
			continue
		}
	}
}

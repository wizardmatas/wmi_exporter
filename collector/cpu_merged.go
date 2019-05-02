// +build windows

package collector

import (
	"strconv"
	"strings"

	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"golang.org/x/sys/windows/registry"
)

func init() {
	Factories["cpu_merged"] = NewCPUCollectormc
}

//A function to get windows version from registry

func getWindowsVersion() float64 {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Warn("Couldn't open registry", err)
	}
	defer k.Close()

	currentv, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		log.Warn("Couldn't open registry to determine current Windows version:", err)
	}

	log.Info("Detected Windows %d.%d\n", currentv)
	currentv_flt, err := strconv.ParseFloat(currentv, 64)
	//	log.Info(currentv_flt)
	return currentv_flt
}

// A cpu_infoCollector is a Prometheus collector for WMI Win32_PerfRawData_Counters_ProcessorInformation metrics
type CPUCollectormc struct {
	CStateSecondsTotal *prometheus.Desc
	TimeTotal          *prometheus.Desc
	InterruptsTotal    *prometheus.Desc
	DPCsTotal          *prometheus.Desc
}

// NewCPUCollectormc constructs a new CPUCollectormc
func NewCPUCollectormc() (Collector, error) {
	const subsystem = "cpu_merged"
	return &CPUCollectormc{
		CStateSecondsTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "cstate_seconds_total"),
			"Time spent in low-power idle state",
			[]string{"core", "state"},
			nil,
		),
		TimeTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "time_total"),
			"Time that processor spent in different modes (idle, user, system, ...)",
			[]string{"core", "mode"},
			nil,
		),

		InterruptsTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "interrupts_total"),
			"Total number of received and serviced hardware interrupts",
			[]string{"core"},
			nil,
		),
		DPCsTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "dpcs_total"),
			"Total number of received and serviced deferred procedure calls (DPCs)",
			[]string{"core"},
			nil,
		),
	}, nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *CPUCollectormc) Collect(ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Error("failed collecting cpu_info metrics:", desc, err)
		return err
	}
	return nil
}

// Win32_PerfRawData_Counters_ProcessorInformation docs:
// - <add link to documentation here>
type Win32_PerfRawData_Counters_ProcessorInformation struct {
	Name                        string
	AverageIdleTime             uint64
	C1TransitionsPersec         uint64
	C2TransitionsPersec         uint64
	C3TransitionsPersec         uint64
	ClockInterruptsPersec       uint32
	DPCRate                     uint32
	DPCsQueuedPersec            uint32
	IdleBreakEventsPersec       uint64
	InterruptsPersec            uint32
	ParkingStatus               uint32
	PercentC1Time               uint64
	PercentC2Time               uint64
	PercentC3Time               uint64
	PercentDPCTime              uint64
	PercentIdleTime             uint64
	PercentInterruptTime        uint64
	PercentofMaximumFrequency   uint32
	PercentPerformanceLimit     uint32
	PercentPriorityTime         uint64
	PercentPrivilegedTime       uint64
	PercentPrivilegedUtility    uint64
	PercentProcessorPerformance uint64
	PercentProcessorTime        uint64
	PercentProcessorUtility     uint64
	PercentUserTime             uint64
	PerformanceLimitFlags       uint32
	ProcessorFrequency          uint32
	ProcessorStateFlags         uint32
}

type Win32_PerfRawData_PerfOS_Processor struct {
	Name                  string
	C1TransitionsPersec   uint64
	C2TransitionsPersec   uint64
	C3TransitionsPersec   uint64
	DPCRate               uint32
	DPCsQueuedPersec      uint32
	InterruptsPersec      uint32
	PercentC1Time         uint64
	PercentC2Time         uint64
	PercentC3Time         uint64
	PercentDPCTime        uint64
	PercentIdleTime       uint64
	PercentInterruptTime  uint64
	PercentPrivilegedTime uint64
	PercentProcessorTime  uint64
	PercentUserTime       uint64
}

func (c *CPUCollectormc) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	version := getWindowsVersion()
	var dst
	// Windows version by number https://docs.microsoft.com/en-us/windows/desktop/sysinfo/operating-system-version
	if version > 6.05 {
		var dst := []Win32_PerfRawData_Counters_ProcessorInformation
		q := queryAll(&dst)
	} else {
		var dst := []Win32_PerfRawData_PerfOS_Processor
	}
	q := queryAll(&dst)
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}

	for _, data := range dst {
		if strings.Contains(data.Name, "_Total") {
			continue
		}

		core := data.Name
		// These are only available from Win32_PerfRawData_Counters_ProcessorInformation, which is only available from Win2008R2+
		/*ch <- prometheus.MustNewConstMetric(
			c.ProcessorFrequency,
			prometheus.GaugeValue,
			float64(data.ProcessorFrequency),
			socket, core,
		)
		ch <- prometheus.MustNewConstMetric(
			c.MaximumFrequency,
			prometheus.GaugeValue,
			float64(data.PercentofMaximumFrequency)/100*float64(data.ProcessorFrequency),
			socket, core,
		)*/

		ch <- prometheus.MustNewConstMetric(
			c.CStateSecondsTotal,
			prometheus.CounterValue,
			float64(data.PercentC1Time)*ticksToSecondsScaleFactor,
			core, "c1",
		)
		ch <- prometheus.MustNewConstMetric(
			c.CStateSecondsTotal,
			prometheus.CounterValue,
			float64(data.PercentC2Time)*ticksToSecondsScaleFactor,
			core, "c2",
		)
		ch <- prometheus.MustNewConstMetric(
			c.CStateSecondsTotal,
			prometheus.CounterValue,
			float64(data.PercentC3Time)*ticksToSecondsScaleFactor,
			core, "c3",
		)

		ch <- prometheus.MustNewConstMetric(
			c.TimeTotal,
			prometheus.CounterValue,
			float64(data.PercentIdleTime)*ticksToSecondsScaleFactor,
			core, "idle",
		)
		ch <- prometheus.MustNewConstMetric(
			c.TimeTotal,
			prometheus.CounterValue,
			float64(data.PercentInterruptTime)*ticksToSecondsScaleFactor,
			core, "interrupt",
		)
		ch <- prometheus.MustNewConstMetric(
			c.TimeTotal,
			prometheus.CounterValue,
			float64(data.PercentDPCTime)*ticksToSecondsScaleFactor,
			core, "dpc",
		)
		ch <- prometheus.MustNewConstMetric(
			c.TimeTotal,
			prometheus.CounterValue,
			float64(data.PercentPrivilegedTime)*ticksToSecondsScaleFactor,
			core, "privileged",
		)
		ch <- prometheus.MustNewConstMetric(
			c.TimeTotal,
			prometheus.CounterValue,
			float64(data.PercentUserTime)*ticksToSecondsScaleFactor,
			core, "user",
		)

		ch <- prometheus.MustNewConstMetric(
			c.InterruptsTotal,
			prometheus.CounterValue,
			float64(data.InterruptsPersec),
			core,
		)
		ch <- prometheus.MustNewConstMetric(
			c.DPCsTotal,
			prometheus.CounterValue,
			float64(data.DPCsQueuedPersec),
			core,
		)
	}

	return nil, nil
}

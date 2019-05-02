package collector

import (
	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

func init() {
	Factories["cpu_info"] = Newcpu_infoCollector
}

// A cpu_infoCollector is a Prometheus collector for WMI Win32_PerfRawData_Counters_ProcessorInformation metrics
type cpu_infoCollector struct {
	AverageIdleTime             *prometheus.Desc
	C1TransitionsPersec         *prometheus.Desc
	C2TransitionsPersec         *prometheus.Desc
	C3TransitionsPersec         *prometheus.Desc
	ClockInterruptsPersec       *prometheus.Desc
	DPCRate                     *prometheus.Desc
	DPCsQueuedPersec            *prometheus.Desc
	IdleBreakEventsPersec       *prometheus.Desc
	InterruptsPersec            *prometheus.Desc
	ParkingStatus               *prometheus.Desc
	PercentC1Time               *prometheus.Desc
	PercentC2Time               *prometheus.Desc
	PercentC3Time               *prometheus.Desc
	PercentDPCTime              *prometheus.Desc
	PercentIdleTime             *prometheus.Desc
	PercentInterruptTime        *prometheus.Desc
	PercentofMaximumFrequency   *prometheus.Desc
	PercentPerformanceLimit     *prometheus.Desc
	PercentPriorityTime         *prometheus.Desc
	PercentPrivilegedTime       *prometheus.Desc
	PercentPrivilegedUtility    *prometheus.Desc
	PercentProcessorPerformance *prometheus.Desc
	PercentProcessorTime        *prometheus.Desc
	PercentProcessorUtility     *prometheus.Desc
	PercentUserTime             *prometheus.Desc
	PerformanceLimitFlags       *prometheus.Desc
	ProcessorFrequency          *prometheus.Desc
	ProcessorStateFlags         *prometheus.Desc
}

// Newcpu_infoCollector ...
func Newcpu_infoCollector() (Collector, error) {
	const subsystem = "cpu_info"
	return &cpu_infoCollector{
		AverageIdleTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "average_idle_time"),
			"(AverageIdleTime)",
			nil,
			nil,
		),
		C1TransitionsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "c1_transitions_persec"),
			"(C1TransitionsPersec)",
			nil,
			nil,
		),
		C2TransitionsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "c2_transitions_persec"),
			"(C2TransitionsPersec)",
			nil,
			nil,
		),
		C3TransitionsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "c3_transitions_persec"),
			"(C3TransitionsPersec)",
			nil,
			nil,
		),
		ClockInterruptsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "clock_interrupts_persec"),
			"(ClockInterruptsPersec)",
			nil,
			nil,
		),
		DPCRate: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "dpc_rate"),
			"(DPCRate)",
			nil,
			nil,
		),
		DPCsQueuedPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "dp_cs_queued_persec"),
			"(DPCsQueuedPersec)",
			nil,
			nil,
		),
		IdleBreakEventsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "idle_break_events_persec"),
			"(IdleBreakEventsPersec)",
			nil,
			nil,
		),
		InterruptsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "interrupts_persec"),
			"(InterruptsPersec)",
			nil,
			nil,
		),
		ParkingStatus: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "parking_status"),
			"(ParkingStatus)",
			nil,
			nil,
		),
		PercentC1Time: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_c1_time"),
			"(PercentC1Time)",
			nil,
			nil,
		),
		PercentC2Time: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_c2_time"),
			"(PercentC2Time)",
			nil,
			nil,
		),
		PercentC3Time: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_c3_time"),
			"(PercentC3Time)",
			nil,
			nil,
		),
		PercentDPCTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_dpc_time"),
			"(PercentDPCTime)",
			nil,
			nil,
		),
		PercentIdleTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_idle_time"),
			"(PercentIdleTime)",
			nil,
			nil,
		),
		PercentInterruptTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_interrupt_time"),
			"(PercentInterruptTime)",
			nil,
			nil,
		),
		PercentofMaximumFrequency: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percentof_maximum_frequency"),
			"(PercentofMaximumFrequency)",
			nil,
			nil,
		),
		PercentPerformanceLimit: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_performance_limit"),
			"(PercentPerformanceLimit)",
			nil,
			nil,
		),
		PercentPriorityTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_priority_time"),
			"(PercentPriorityTime)",
			nil,
			nil,
		),
		PercentPrivilegedTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_privileged_time"),
			"(PercentPrivilegedTime)",
			nil,
			nil,
		),
		PercentPrivilegedUtility: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_privileged_utility"),
			"(PercentPrivilegedUtility)",
			nil,
			nil,
		),
		PercentProcessorPerformance: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_processor_performance"),
			"(PercentProcessorPerformance)",
			nil,
			nil,
		),
		PercentProcessorTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_processor_time"),
			"(PercentProcessorTime)",
			nil,
			nil,
		),
		PercentProcessorUtility: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_processor_utility"),
			"(PercentProcessorUtility)",
			nil,
			nil,
		),
		PercentUserTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_user_time"),
			"(PercentUserTime)",
			nil,
			nil,
		),
		PerformanceLimitFlags: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "performance_limit_flags"),
			"(PerformanceLimitFlags)",
			nil,
			nil,
		),
		ProcessorFrequency: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "processor_frequency"),
			"(ProcessorFrequency)",
			nil,
			nil,
		),
		ProcessorStateFlags: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "processor_state_flags"),
			"(ProcessorStateFlags)",
			nil,
			nil,
		),
	}, nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *cpu_infoCollector) Collect(ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Error("failed collecting cpu_info metrics:", desc, err)
		return err
	}
	return nil
}

// Win32_PerfRawData_Counters_ProcessorInformation docs:
// - <add link to documentation here>
type Win32_PerfRawData_Counters_ProcessorInformation struct {
	Name string

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

func (c *cpu_infoCollector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	var dst []Win32_PerfRawData_Counters_ProcessorInformation
	q := queryAll(&dst)
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}

	ch <- prometheus.MustNewConstMetric(
		c.AverageIdleTime,
		prometheus.GaugeValue,
		float64(dst[0].AverageIdleTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.C1TransitionsPersec,
		prometheus.GaugeValue,
		float64(dst[0].C1TransitionsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.C2TransitionsPersec,
		prometheus.GaugeValue,
		float64(dst[0].C2TransitionsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.C3TransitionsPersec,
		prometheus.GaugeValue,
		float64(dst[0].C3TransitionsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ClockInterruptsPersec,
		prometheus.GaugeValue,
		float64(dst[0].ClockInterruptsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DPCRate,
		prometheus.GaugeValue,
		float64(dst[0].DPCRate),
	)

	ch <- prometheus.MustNewConstMetric(
		c.DPCsQueuedPersec,
		prometheus.GaugeValue,
		float64(dst[0].DPCsQueuedPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.IdleBreakEventsPersec,
		prometheus.GaugeValue,
		float64(dst[0].IdleBreakEventsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.InterruptsPersec,
		prometheus.GaugeValue,
		float64(dst[0].InterruptsPersec),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ParkingStatus,
		prometheus.GaugeValue,
		float64(dst[0].ParkingStatus),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentC1Time,
		prometheus.GaugeValue,
		float64(dst[0].PercentC1Time),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentC2Time,
		prometheus.GaugeValue,
		float64(dst[0].PercentC2Time),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentC3Time,
		prometheus.GaugeValue,
		float64(dst[0].PercentC3Time),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentDPCTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentDPCTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentIdleTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentIdleTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentInterruptTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentInterruptTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentofMaximumFrequency,
		prometheus.GaugeValue,
		float64(dst[0].PercentofMaximumFrequency),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentPerformanceLimit,
		prometheus.GaugeValue,
		float64(dst[0].PercentPerformanceLimit),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentPriorityTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentPriorityTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentPrivilegedTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentPrivilegedTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentPrivilegedUtility,
		prometheus.GaugeValue,
		float64(dst[0].PercentPrivilegedUtility),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentProcessorPerformance,
		prometheus.GaugeValue,
		float64(dst[0].PercentProcessorPerformance),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentProcessorTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentProcessorTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentProcessorUtility,
		prometheus.GaugeValue,
		float64(dst[0].PercentProcessorUtility),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentUserTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentUserTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PerformanceLimitFlags,
		prometheus.GaugeValue,
		float64(dst[0].PerformanceLimitFlags),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ProcessorFrequency,
		prometheus.GaugeValue,
		float64(dst[0].ProcessorFrequency),
	)

	ch <- prometheus.MustNewConstMetric(
		c.ProcessorStateFlags,
		prometheus.GaugeValue,
		float64(dst[0].ProcessorStateFlags),
	)

	return nil, nil
}

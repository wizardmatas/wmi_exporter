package collector

import (
	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

func init() {
	Factories["cpu_1"] = Newcpu_1Collector
}

// A cpu_1Collector is a Prometheus collector for WMI Win32_PerfRawData_PerfOS_Processor metrics
type cpu_1Collector struct {
	C1TransitionsPersec   *prometheus.Desc
	C2TransitionsPersec   *prometheus.Desc
	C3TransitionsPersec   *prometheus.Desc
	DPCRate               *prometheus.Desc
	DPCsQueuedPersec      *prometheus.Desc
	InterruptsPersec      *prometheus.Desc
	PercentC1Time         *prometheus.Desc
	PercentC2Time         *prometheus.Desc
	PercentC3Time         *prometheus.Desc
	PercentDPCTime        *prometheus.Desc
	PercentIdleTime       *prometheus.Desc
	PercentInterruptTime  *prometheus.Desc
	PercentPrivilegedTime *prometheus.Desc
	PercentProcessorTime  *prometheus.Desc
	PercentUserTime       *prometheus.Desc
}

// Newcpu_1Collector ...
func Newcpu_1Collector() (Collector, error) {
	const subsystem = "cpu_1"
	return &cpu_1Collector{
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
		InterruptsPersec: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "interrupts_persec"),
			"(InterruptsPersec)",
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
		PercentPrivilegedTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_privileged_time"),
			"(PercentPrivilegedTime)",
			nil,
			nil,
		),
		PercentProcessorTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_processor_time"),
			"(PercentProcessorTime)",
			nil,
			nil,
		),
		PercentUserTime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "percent_user_time"),
			"(PercentUserTime)",
			nil,
			nil,
		),
	}, nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *cpu_1Collector) Collect(ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Error("failed collecting cpu_1 metrics:", desc, err)
		return err
	}
	return nil
}

// Win32_PerfRawData_PerfOS_Processor docs:
// - <add link to documentation here>
type Win32_PerfRawData_PerfOS_Processor struct {
	Name string

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

func (c *cpu_1Collector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	var dst []Win32_PerfRawData_PerfOS_Processor
	q := queryAll(&dst)
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}

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
		c.InterruptsPersec,
		prometheus.GaugeValue,
		float64(dst[0].InterruptsPersec),
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
		c.PercentPrivilegedTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentPrivilegedTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentProcessorTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentProcessorTime),
	)

	ch <- prometheus.MustNewConstMetric(
		c.PercentUserTime,
		prometheus.GaugeValue,
		float64(dst[0].PercentUserTime),
	)

	return nil, nil
}

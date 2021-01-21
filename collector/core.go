/**
 * Copyright (c) 2020-2021, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * Collector package contains all definitions required by Prometheus exporter
 * documentation, and additional functions used to manage the exporter.
 */

package collector

import (
	"net/http"

	"github.com/intel/ipmctl_exporter/collector/nvm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

type ipmctlCollector struct {
	// reader
	metricsReader *nvm.MetricsReader
	// internal fields
	enableThresholds bool
	// performance readings
	totalMediaReads    *prometheus.Desc
	totalMediaWrites   *prometheus.Desc
	totalReadRequests  *prometheus.Desc
	totalWriteRequests *prometheus.Desc
	// sensor readings
	health                      *prometheus.Desc
	mediaTemperature            *prometheus.Desc
	controllerTemperature       *prometheus.Desc
	percentageRemaining         *prometheus.Desc
	latchedDirtyShutdownCount   *prometheus.Desc
	powerOnTime                 *prometheus.Desc
	upTime                      *prometheus.Desc
	powerCycles                 *prometheus.Desc
	fwErrorCount                *prometheus.Desc
	unlatchedDirtyShutdownCount *prometheus.Desc
	// sensor settings (thresholds)
	mtEnabled                   *prometheus.Desc
	mtUpperCriticalThreshold    *prometheus.Desc
	mtLowerCriticalThreshold    *prometheus.Desc
	mtUpperFatalThreshold       *prometheus.Desc
	mtLowerFatalThreshold       *prometheus.Desc
	mtUpperNoncriticalThreshold *prometheus.Desc
	mtLowerNoncriticalThreshold *prometheus.Desc
	ctEnabled                   *prometheus.Desc
	ctUpperCriticalThreshold    *prometheus.Desc
	ctLowerCriticalThreshold    *prometheus.Desc
	ctUpperFatalThreshold       *prometheus.Desc
	ctLowerFatalThreshold       *prometheus.Desc
	ctUpperNoncriticalThreshold *prometheus.Desc
	ctLowerNoncriticalThreshold *prometheus.Desc
	prEnabled                   *prometheus.Desc
	prUpperCriticalThreshold    *prometheus.Desc
	prLowerCriticalThreshold    *prometheus.Desc
	prUpperFatalThreshold       *prometheus.Desc
	prLowerFatalThreshold       *prometheus.Desc
	prUpperNoncriticalThreshold *prometheus.Desc
	prLowerNoncriticalThreshold *prometheus.Desc
	// metadata readings (some sort of states / additional information)
	deviceDiscoveryInfo            *prometheus.Desc
	deviceSecurityCapabilitiesInfo *prometheus.Desc
	deviceCapabilitiesInfo         *prometheus.Desc
	ipmctlExporterInfo             *prometheus.Desc
}

// Function used to get metrics description.
// If you would like to register new metric follow rules below (in terms of
// metric name):
// - always use _total postfix with Counter type, otherwise avoid using these
//   suffixes in metrics
// - always specify the units you are working with for clarity, units should be plural
// - don't put the type of the metric in the name such as gauge, counter etc.
func newIpmctlCollector(enableThresholds bool) *ipmctlCollector {
	collector := new(ipmctlCollector)
	collector.metricsReader = nvm.NewMetricsReader()
	collector.enableThresholds = enableThresholds
	collector.totalMediaReads = prometheus.NewDesc("ipmctl_total_media_reads_total",
		"Lifetime number of 64 byte reads from media on the DCPMM", nvm.DevPerformanceLabelNames, nil)
	collector.totalMediaWrites = prometheus.NewDesc("ipmctl_total_media_writes_total",
		"Lifetime number of 64 byte writes to media on the DCPMM", nvm.DevPerformanceLabelNames, nil)
	collector.totalReadRequests = prometheus.NewDesc("ipmctl_total_read_requests_total",
		"Lifetime number of DDRT read transactions the DCPMM has serviced", nvm.DevPerformanceLabelNames, nil)
	collector.totalWriteRequests = prometheus.NewDesc("ipmctl_total_write_requests_total",
		"Lifetime number of DDRT write transactions the DCPMM has serviced", nvm.DevPerformanceLabelNames, nil)
	collector.health = prometheus.NewDesc("ipmctl_health",
		"DCPMM health as reported in the SMART log", nvm.SensorLabelNames, nil)
	collector.mediaTemperature = prometheus.NewDesc("ipmctl_media_temperature_degrees_c",
		"Device media temperature in degrees Celsius", nvm.SensorLabelNames, nil)
	collector.controllerTemperature = prometheus.NewDesc("ipmctl_controller_temperature_degrees_c",
		"Device media temperature in degrees Celsius", nvm.SensorLabelNames, nil)
	collector.percentageRemaining = prometheus.NewDesc("ipmctl_percentage_remaining",
		"Amount of percentage remaining as a percentage", nvm.SensorLabelNames, nil)
	collector.latchedDirtyShutdownCount = prometheus.NewDesc("ipmctl_latched_dirty_shutdown_count_total",
		"Device shutdowns without notification", nvm.SensorLabelNames, nil)
	collector.powerOnTime = prometheus.NewDesc("ipmctl_power_on_time_total",
		"Total power-on time over the lifetime of the device", nvm.SensorLabelNames, nil)
	collector.upTime = prometheus.NewDesc("ipmctl_up_time_seconds_total",
		"Total power-on time since the last power cycle of the device", nvm.SensorLabelNames, nil)
	collector.powerCycles = prometheus.NewDesc("ipmctl_power_cycles_total",
		"Number of power cycles over the lifetime of the device", nvm.SensorLabelNames, nil)
	collector.fwErrorCount = prometheus.NewDesc("ipmctl_fw_error_count_total",
		"The total number of firmware error log entries", nvm.SensorLabelNames, nil)
	collector.unlatchedDirtyShutdownCount = prometheus.NewDesc("ipmctl_unlatched_dirty_shutdown_count_total",
		"Number of times that the FW received an unexpected power loss", nvm.SensorLabelNames, nil)
	collector.deviceDiscoveryInfo = prometheus.NewDesc("ipmctl_device_discovery_info",
		"Describes an enterprise-level view of a device", nvm.DeviceDiscoveryLabelNames, nil)
	collector.deviceSecurityCapabilitiesInfo = prometheus.NewDesc("ipmctl_device_security_capabilities_info",
		"Describes the security capabilities of a device", nvm.DeviceSecurityCapabilitiesLabelNames, nil)
	collector.deviceCapabilitiesInfo = prometheus.NewDesc("impctl_device_capabilities_info",
		"Describes the capabilities supported by a DCPMM", nvm.DeviceCapabilitiesLabelNames, nil)
	collector.ipmctlExporterInfo = prometheus.NewDesc("ipmctl_info",
		"Describes ipmctl_exporter info", nvm.IpmctlExporterLabelNames, nil)
	if enableThresholds {
		collector.mtEnabled = prometheus.NewDesc("ipmctl_media_temperature_enabled",
			"Indictes if firmware notifications are enabled when media temperature value is critical", nvm.SettingsLabelNames, nil)
		collector.mtUpperCriticalThreshold = prometheus.NewDesc("ipmctl_media_temperature_upper_critical_threshold",
			"The upper media temperature critical threshold", nvm.SettingsLabelNames, nil)
		collector.mtLowerCriticalThreshold = prometheus.NewDesc("ipmctl_media_temperature_lower_critical_threshold",
			"The lower media temperature critical threshold", nvm.SettingsLabelNames, nil)
		collector.mtUpperFatalThreshold = prometheus.NewDesc("ipmctl_media_temperature_upper_fatal_threshold",
			"The upper media temperature fatal threshold", nvm.SettingsLabelNames, nil)
		collector.mtLowerFatalThreshold = prometheus.NewDesc("ipmctl_media_temperature_lower_fatal_threshold",
			"The lower media temperature fatal threshold", nvm.SettingsLabelNames, nil)
		collector.mtUpperNoncriticalThreshold = prometheus.NewDesc("ipmctl_media_temperature_upper_noncritical_threshold",
			"The upper media temperature noncritical threshold", nvm.SettingsLabelNames, nil)
		collector.mtLowerNoncriticalThreshold = prometheus.NewDesc("ipmctl_media_temperature_lower_noncritical_threshold",
			"The lower media temperature noncritical threshold", nvm.SettingsLabelNames, nil)
		collector.ctEnabled = prometheus.NewDesc("ipmctl_controller_temperature_enabled",
			"Indictes if firmware notifications are enabled when controller temperature value is critical", nvm.SettingsLabelNames, nil)
		collector.ctUpperCriticalThreshold = prometheus.NewDesc("ipmctl_controller_temperature_upper_critical_threshold",
			"The upper controller temperature critical threshold", nvm.SettingsLabelNames, nil)
		collector.ctLowerCriticalThreshold = prometheus.NewDesc("ipmctl_controller_temperature_lower_critical_threshold",
			"The lower controller temperature critical threshold", nvm.SettingsLabelNames, nil)
		collector.ctUpperFatalThreshold = prometheus.NewDesc("ipmctl_controller_temperature_upper_fatal_threshold",
			"The upper controller temperature fatal threshold", nvm.SettingsLabelNames, nil)
		collector.ctLowerFatalThreshold = prometheus.NewDesc("ipmctl_controller_temperature_lower_fatal_threshold",
			"The lower controller temperature fatal threshold", nvm.SettingsLabelNames, nil)
		collector.ctUpperNoncriticalThreshold = prometheus.NewDesc("ipmctl_controller_temperature_upper_noncritical_threshold",
			"The upper controller temperature noncritical threshold", nvm.SettingsLabelNames, nil)
		collector.ctLowerNoncriticalThreshold = prometheus.NewDesc("ipmctl_controller_temperature_lower_noncritical_threshold",
			"The lower controller temperature noncritical threshold", nvm.SettingsLabelNames, nil)
		collector.prEnabled = prometheus.NewDesc("ipmctl_percentage_remaining_enabled",
			"Indictes if firmware notifications are enabled when percentage remaining value is critical", nvm.SettingsLabelNames, nil)
		collector.prUpperCriticalThreshold = prometheus.NewDesc("ipmctl_percentage_remaining_upper_critical_threshold",
			"The upper percentage remaining critical threshold", nvm.SettingsLabelNames, nil)
		collector.prLowerCriticalThreshold = prometheus.NewDesc("ipmctl_percentage_remaining_lower_critical_threshold",
			"The lower percentage remaining critical threshold", nvm.SettingsLabelNames, nil)
		collector.prUpperFatalThreshold = prometheus.NewDesc("ipmctl_percentage_remaining_upper_fatal_threshold",
			"The upper percentage remaining fatal threshold", nvm.SettingsLabelNames, nil)
		collector.prLowerFatalThreshold = prometheus.NewDesc("ipmctl_percentage_remaining_lower_fatal_threshold",
			"The lower percentage remaining fatal threshold", nvm.SettingsLabelNames, nil)
		collector.prUpperNoncriticalThreshold = prometheus.NewDesc("ipmctl_percentage_remaining_upper_noncritical_threshold",
			"The upper percentage remaining noncritical threshold", nvm.SettingsLabelNames, nil)
		collector.prLowerNoncriticalThreshold = prometheus.NewDesc("ipmctl_percentage_remaining_lower_noncritical_threshold",
			"The lower percentage remaining noncritical threshold", nvm.SettingsLabelNames, nil)
	}
	return collector
}

// Function called to describe all metrics exposed by ipmctl_exporter
func (collector *ipmctlCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.totalMediaReads
	ch <- collector.totalMediaWrites
	ch <- collector.totalReadRequests
	ch <- collector.totalWriteRequests
	ch <- collector.health
	ch <- collector.mediaTemperature
	ch <- collector.controllerTemperature
	ch <- collector.percentageRemaining
	ch <- collector.latchedDirtyShutdownCount
	ch <- collector.powerOnTime
	ch <- collector.upTime
	ch <- collector.powerCycles
	ch <- collector.fwErrorCount
	ch <- collector.unlatchedDirtyShutdownCount
	ch <- collector.deviceDiscoveryInfo
	ch <- collector.deviceSecurityCapabilitiesInfo
	ch <- collector.deviceCapabilitiesInfo
	ch <- collector.ipmctlExporterInfo
	if collector.enableThresholds {
		ch <- collector.mtEnabled
		ch <- collector.mtUpperCriticalThreshold
		ch <- collector.mtLowerCriticalThreshold
		ch <- collector.mtUpperFatalThreshold
		ch <- collector.mtLowerFatalThreshold
		ch <- collector.mtUpperNoncriticalThreshold
		ch <- collector.mtLowerNoncriticalThreshold
		ch <- collector.ctEnabled
		ch <- collector.ctUpperCriticalThreshold
		ch <- collector.ctLowerCriticalThreshold
		ch <- collector.ctUpperFatalThreshold
		ch <- collector.ctLowerFatalThreshold
		ch <- collector.ctUpperNoncriticalThreshold
		ch <- collector.ctLowerNoncriticalThreshold
		ch <- collector.prEnabled
		ch <- collector.prUpperCriticalThreshold
		ch <- collector.prLowerCriticalThreshold
		ch <- collector.prUpperFatalThreshold
		ch <- collector.prLowerFatalThreshold
		ch <- collector.prUpperNoncriticalThreshold
		ch <- collector.prLowerNoncriticalThreshold
	}
}

func addMetric(ch chan<- prometheus.Metric,
	desc *prometheus.Desc,
	metricType prometheus.ValueType,
	readings []nvm.MetricReading) {
	for _, reading := range readings {
		labelValues := reading.Labels.GetLabelValues()
		ch <- prometheus.MustNewConstMetric(desc,
			metricType,
			reading.MetricValue,
			labelValues...)
	}
}

// Function called to collect all data exposed by exporter as a response for
// Prometheus server request. If you would like to add new metrics, remember to
// pick up metric type wisely, a counter is a cumulative metric that represents
// a single monotonically increasing counter whose value can only increase or
// be reset to zero on restart. That is why if metric value is counting it should
// be marked as "Counter" even if it isn't persistent through the AC cycle, like
// for instance upTime metric.
func (collector *ipmctlCollector) Collect(ch chan<- prometheus.Metric) {
	reader := collector.metricsReader
	status, err := reader.GetRequiredReadings()
	if false == status {
		log.Error("ipmctl exporter - failed to read PMEM metrics due to: ", err)
		return
	}
	healthReadings, _ := reader.GetHealth()
	addMetric(ch, collector.health, prometheus.GaugeValue, healthReadings)
	mediaTemperatureReadings, _ := reader.GetMediaTemperature()
	addMetric(ch, collector.mediaTemperature, prometheus.GaugeValue, mediaTemperatureReadings)
	controllerTemperatureReadings, _ := reader.GetControllerTemperature()
	addMetric(ch, collector.controllerTemperature, prometheus.GaugeValue, controllerTemperatureReadings)
	percentageRemainingReadings, _ := reader.GetPercentageRemaining()
	addMetric(ch, collector.percentageRemaining, prometheus.GaugeValue, percentageRemainingReadings)
	LDSCReadings, _ := reader.GetLatchedDirtyShutdownCount()
	addMetric(ch, collector.latchedDirtyShutdownCount, prometheus.CounterValue, LDSCReadings)
	powerOnTimeReadings, _ := reader.GetPowerOnTime()
	addMetric(ch, collector.powerOnTime, prometheus.CounterValue, powerOnTimeReadings)
	upTimeReadings, _ := reader.GetUpTime()
	addMetric(ch, collector.upTime, prometheus.CounterValue, upTimeReadings)
	powerCyclesReadings, _ := reader.GetPowerCycles()
	addMetric(ch, collector.powerCycles, prometheus.CounterValue, powerCyclesReadings)
	fwErrorCountReadings, _ := reader.GetFwErrorCount()
	addMetric(ch, collector.fwErrorCount, prometheus.CounterValue, fwErrorCountReadings)
	UDSCReadings, _ := reader.GetUnlatchedDirtyShutdownCount()
	addMetric(ch, collector.unlatchedDirtyShutdownCount, prometheus.CounterValue, UDSCReadings)
	totalMediaReads, _ := reader.GetTotalMediaReads()
	addMetric(ch, collector.totalMediaReads, prometheus.CounterValue, totalMediaReads)
	totalMediaWrites, _ := reader.GetTotalMediaWrites()
	addMetric(ch, collector.totalMediaWrites, prometheus.CounterValue, totalMediaWrites)
	totalReadRequests, _ := reader.GetTotalReadRequests()
	addMetric(ch, collector.totalReadRequests, prometheus.CounterValue, totalReadRequests)
	totalWriteRequests, _ := reader.GetTotalWriteRequests()
	addMetric(ch, collector.totalWriteRequests, prometheus.CounterValue, totalWriteRequests)
	deviceDiscoveryInfo, _ := reader.GetDeviceDiscoveryInfo()
	addMetric(ch, collector.deviceDiscoveryInfo, prometheus.GaugeValue, deviceDiscoveryInfo)
	deviceSecurityCapabilitiesInfo, _ := reader.GetDeviceSecurityCapabilitiesInfo()
	addMetric(ch, collector.deviceSecurityCapabilitiesInfo, prometheus.GaugeValue, deviceSecurityCapabilitiesInfo)
	deviceCapabilitiesInfo, _ := reader.GetDeviceCapabilitiesInfo()
	addMetric(ch, collector.deviceCapabilitiesInfo, prometheus.GaugeValue, deviceCapabilitiesInfo)
	ipmctlExporterInfo, _ := nvm.GetIpmctlExporterInfo()
	addMetric(ch, collector.ipmctlExporterInfo, prometheus.GaugeValue, ipmctlExporterInfo)
	if collector.enableThresholds {
		mtEnabled, _ := reader.GetMTEnabled()
		addMetric(ch, collector.mtEnabled, prometheus.GaugeValue, mtEnabled)
		mtUpperCriticalThreshold, _ := reader.GetMTUpperCriticalThreshold()
		addMetric(ch, collector.mtUpperCriticalThreshold, prometheus.GaugeValue, mtUpperCriticalThreshold)
		mtLowerCriticalThreshold, _ := reader.GetMTLowerCriticalThreshold()
		addMetric(ch, collector.mtLowerCriticalThreshold, prometheus.GaugeValue, mtLowerCriticalThreshold)
		mtUpperFatalThreshold, _ := reader.GetMTUpperFatalThreshold()
		addMetric(ch, collector.mtUpperFatalThreshold, prometheus.GaugeValue, mtUpperFatalThreshold)
		mtLowerFatalThreshold, _ := reader.GetMTLowerFatalThreshold()
		addMetric(ch, collector.mtLowerFatalThreshold, prometheus.GaugeValue, mtLowerFatalThreshold)
		mtUpperNoncriticalThreshold, _ := reader.GetMTUpperNoncriticalThreshold()
		addMetric(ch, collector.mtUpperNoncriticalThreshold, prometheus.GaugeValue, mtUpperNoncriticalThreshold)
		mtLowerNoncriticalThreshold, _ := reader.GetMTLowerNoncriticalThreshold()
		addMetric(ch, collector.mtLowerNoncriticalThreshold, prometheus.GaugeValue, mtLowerNoncriticalThreshold)
		ctEnabled, _ := reader.GetCTEnabled()
		addMetric(ch, collector.ctEnabled, prometheus.GaugeValue, ctEnabled)
		ctUpperCriticalThreshold, _ := reader.GetCTUpperCriticalThreshold()
		addMetric(ch, collector.ctUpperCriticalThreshold, prometheus.GaugeValue, ctUpperCriticalThreshold)
		ctLowerCriticalThreshold, _ := reader.GetCTLowerCriticalThreshold()
		addMetric(ch, collector.ctLowerCriticalThreshold, prometheus.GaugeValue, ctLowerCriticalThreshold)
		ctUpperFatalThreshold, _ := reader.GetCTUpperFatalThreshold()
		addMetric(ch, collector.ctUpperFatalThreshold, prometheus.GaugeValue, ctUpperFatalThreshold)
		ctLowerFatalThreshold, _ := reader.GetCTLowerFatalThreshold()
		addMetric(ch, collector.ctLowerFatalThreshold, prometheus.GaugeValue, ctLowerFatalThreshold)
		ctUpperNoncriticalThreshold, _ := reader.GetCTUpperNoncriticalThreshold()
		addMetric(ch, collector.ctUpperNoncriticalThreshold, prometheus.GaugeValue, ctUpperNoncriticalThreshold)
		ctLowerNoncriticalThreshold, _ := reader.GetCTLowerNoncriticalThreshold()
		addMetric(ch, collector.ctLowerNoncriticalThreshold, prometheus.GaugeValue, ctLowerNoncriticalThreshold)
		prEnabled, _ := reader.GetPREnabled()
		addMetric(ch, collector.prEnabled, prometheus.GaugeValue, prEnabled)
		prUpperCriticalThreshold, _ := reader.GetPRUpperCriticalThreshold()
		addMetric(ch, collector.prUpperCriticalThreshold, prometheus.GaugeValue, prUpperCriticalThreshold)
		prLowerCriticalThreshold, _ := reader.GetPRLowerCriticalThreshold()
		addMetric(ch, collector.prLowerCriticalThreshold, prometheus.GaugeValue, prLowerCriticalThreshold)
		prUpperFatalThreshold, _ := reader.GetPRUpperFatalThreshold()
		addMetric(ch, collector.prUpperFatalThreshold, prometheus.GaugeValue, prUpperFatalThreshold)
		prLowerFatalThreshold, _ := reader.GetPRLowerFatalThreshold()
		addMetric(ch, collector.prLowerFatalThreshold, prometheus.GaugeValue, prLowerFatalThreshold)
		prUpperNoncriticalThreshold, _ := reader.GetPRUpperNoncriticalThreshold()
		addMetric(ch, collector.prUpperNoncriticalThreshold, prometheus.GaugeValue, prUpperNoncriticalThreshold)
		prLowerNoncriticalThreshold, _ := reader.GetPRLowerNoncriticalThreshold()
		addMetric(ch, collector.prLowerNoncriticalThreshold, prometheus.GaugeValue, prLowerNoncriticalThreshold)
	}
}

func Stop() {
	nvm.Uninit()
}

var Version string

func Run(port string, enableThresholds bool) {
	nvm.Init()
	nvm.Version = Version
	ipmctlCollector := newIpmctlCollector(enableThresholds)
	prometheus.MustRegister(ipmctlCollector)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", promhttp.Handler())
	port = ":" + port
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Error(err)
	}

}

/**
 * Copyright (c) 2020-2021, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduces wrapper for ipmctl library written in C.
 * api_perform.go file exposes external API for exporter to collect
 * some NVM performance metrics.
 */

package nvm

import (
	"fmt"
)

var DevPerformanceLabelNames = []string{
	"uid",
}

var devPerformanceTypeEnum = &devPerformanceType{
	bytesRead:    0,
	hostReads:    1,
	bytesWritten: 2,
	hostWrites:   3,
	blockReads:   4,
	blockWrites:  5,
	unknown:      0xFF,
}

type devPerformanceReading MetricReading
type devPerformanceLabels MetricLabels
type devPerformanceTypeEnumAttr enumAttr
type devPerformanceType struct {
	bytesRead    devPerformanceTypeEnumAttr
	hostReads    devPerformanceTypeEnumAttr
	bytesWritten devPerformanceTypeEnumAttr
	hostWrites   devPerformanceTypeEnumAttr
	blockReads   devPerformanceTypeEnumAttr
	blockWrites  devPerformanceTypeEnumAttr
	unknown      devPerformanceTypeEnumAttr
}

func (pl devPerformanceLabels) GetLabelValues() []string {
	return getValuesByName(DevPerformanceLabelNames, MetricLabels(pl).labels)
}

func (pl devPerformanceLabels) GetLabelNames() []string {
	return DevPerformanceLabelNames
}

func (pl devPerformanceLabels) addLabel(name string, value string) {
	MetricLabels(pl).labels[name] = value
}

func newDevPerformanceReading(dimmUID nvmUID,
	readStatus nvmStatusCodeEnumAttr,
	metricType devPerformanceTypeEnumAttr,
	metricValue nvmUint64) *devPerformanceReading {
	devPerfReading := new(devPerformanceReading)
	devPerfReading.DIMMUID = string(dimmUID)
	devPerfReading.ReadStatus = int(readStatus)
	devPerfReading.MetricType = uint8(metricType)
	devPerfReading.MetricValue = float64(metricValue)
	devPerfReading.Labels = devPerformanceLabels(*newMetricLabels())
	return devPerfReading
}

func (reader *MetricsReader) getDevicePerformanceReadings(metricType devPerformanceTypeEnumAttr) []MetricReading {
	results := make([]MetricReading, reader.deviceCount)
	for i, dev := range reader.devices {
		perf := dev.performance
		opstat := dev.performanceOpstat
		metricValue := nvmUint64(0)
		switch metricType {
		case devPerformanceTypeEnum.bytesRead:
			metricValue = perf.bytesRead
		case devPerformanceTypeEnum.hostReads:
			metricValue = perf.hostReads
		case devPerformanceTypeEnum.bytesWritten:
			metricValue = perf.bytesWritten
		case devPerformanceTypeEnum.hostWrites:
			metricValue = perf.hostWrites
		case devPerformanceTypeEnum.blockReads:
			metricValue = perf.blockReads
		case devPerformanceTypeEnum.blockWrites:
			metricValue = perf.blockWrites
		}
		devPerfReading := *newDevPerformanceReading(dev.uid, opstat, metricType, metricValue)
		devPerfReading.Labels.addLabel("uid", string(dev.uid))
		results[i] = MetricReading(devPerfReading)
	}
	return results
}

// Number of 64 byte reads from media on the DCPMM since last AC cycle
func (reader *MetricsReader) GetMediaReads() ([]MetricReading, error) {
	// stubbed - was not exposed by NVM API
	return []MetricReading{}, fmt.Errorf("stubbed")
}

// Number of 64 byte writes to media on the DCPMM since last AC cycle
func (reader *MetricsReader) GetMediaWrites() ([]MetricReading, error) {
	// stubbed - was not exposed by NVM API
	return []MetricReading{}, fmt.Errorf("stubbed")
}

// Number of DDRT read transactions the DCPMM has serviced since last AC cycle
func (reader *MetricsReader) GetReadRequests() ([]MetricReading, error) {
	// stubbed - was not exposed by NVM API
	return []MetricReading{}, fmt.Errorf("stubbed")
}

// Number of DDRT write transactions the DCPMM has serviced since last AC cycle
func (reader *MetricsReader) GetWriteRequest() ([]MetricReading, error) {
	// stubbed - was not exposed by NVM API
	return []MetricReading{}, fmt.Errorf("stubbed")
}

// Lifetime number of 64 byte reads from media on the DCPMM
func (reader *MetricsReader) GetTotalMediaReads() []MetricReading {
	metricType := devPerformanceTypeEnum.bytesRead
	return reader.getDevicePerformanceReadings(metricType)
}

// Lifetime number of 64 byte writes to media on the DCPMM
func (reader *MetricsReader) GetTotalMediaWrites() []MetricReading {
	metricType := devPerformanceTypeEnum.bytesWritten
	return reader.getDevicePerformanceReadings(metricType)
}

// Lifetime number of DDRT read transactions the DCPMM has serviced
func (reader *MetricsReader) GetTotalReadRequests() []MetricReading {
	metricType := devPerformanceTypeEnum.hostReads
	return reader.getDevicePerformanceReadings(metricType)
}

// Lifetime number of DDRT write transactions the DCPMM has serviced
func (reader *MetricsReader) GetTotalWriteRequests() []MetricReading {
	metricType := devPerformanceTypeEnum.hostWrites
	return reader.getDevicePerformanceReadings(metricType)
}

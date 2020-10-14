/**
 * Copyright (c) 2020, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduce wrapper for ipmctl library written in C.
 * api_perform.go file expose external API for exporter to collect
 * some NVM performance metrics.
 */

package nvm

import (
    "fmt"
)

var DevPerformanceLabelNames = []string {
    "uid",
}

var devPerformanceTypeEnum = &devPerformanceType {
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

func (pl devPerformanceLabels) GetLabelValues() ([]string) {
    return getValuesByName(DevPerformanceLabelNames, MetricLabels(pl).labels)
}

func (pl devPerformanceLabels) GetLabelNames() ([]string) {
    return DevPerformanceLabelNames
}

func (pl devPerformanceLabels) addLabel(name string, value string) {
    MetricLabels(pl).labels[name] = value
}

func newDevPerformanceReading(dimmUID nvmUID,
                              readStatus nvmStatusCodeEnumAttr,
                              metricType devPerformanceTypeEnumAttr,
                              metricValue nvmUint64) (*devPerformanceReading) {
    devPerfReading := new(devPerformanceReading)
    devPerfReading.DIMMUID     = string(dimmUID)
    devPerfReading.ReadStatus  = int(readStatus)
    devPerfReading.MetricType  = uint8(metricType)
    devPerfReading.MetricValue = float64(metricValue)
    devPerfReading.Labels      = devPerformanceLabels(*newMetricLabels())
    return devPerfReading
}

func getDevicePerformanceReadings(metricType devPerformanceTypeEnumAttr) ([]MetricReading, error) {
    dimmID := nvmUID("")
    opstat, count, err := GetNumberOfDevices()
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, 1)
        devPerfReading := *newDevPerformanceReading(dimmID, opstat, 0xFF, 0)
        results[0] = MetricReading(devPerfReading)
        return results, err
    }
    devices := make([]deviceDiscovery, count)
    opstat, devices, err = GetDevices(count)
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, count)
        for i:=0; i < int(count); i++ {
            devPerfReading := *newDevPerformanceReading(dimmID, opstat, 0xFF, 0)
            results[i] = MetricReading(devPerfReading)
        }
        return results, err
    }
    results    := make([]MetricReading, count)
    for i, dev := range devices {
        opstat, value, _ := GetDevicePerformance(dev.uid)
        metricValue := nvmUint64(0)
        switch metricType {
            case devPerformanceTypeEnum.bytesRead:    metricValue = value.bytesRead
            case devPerformanceTypeEnum.hostReads:    metricValue = value.hostReads
            case devPerformanceTypeEnum.bytesWritten: metricValue = value.bytesWritten
            case devPerformanceTypeEnum.hostWrites:   metricValue = value.hostWrites
            case devPerformanceTypeEnum.blockReads:   metricValue = value.blockReads
            case devPerformanceTypeEnum.blockWrites:  metricValue = value.blockWrites
        }
        devPerfReading := *newDevPerformanceReading(dev.uid, opstat, metricType, metricValue)
        devPerfReading.Labels.addLabel("uid", string(dev.uid))
        results[i] = MetricReading(devPerfReading)
    }
    return results, nil
}

// Number of 64 byte reads from media on the DCPMM since last AC cycle
func GetMediaReads() ([]MetricReading, error) {
    // stubbed - was not exposed by NVM API
    return []MetricReading{}, fmt.Errorf("stubbed")
}

// Number of 64 byte writes to media on the DCPMM since last AC cycle
func GetMediaWrites() ([]MetricReading, error) {
    // stubbed - was not exposed by NVM API
    return []MetricReading{}, fmt.Errorf("stubbed")
}

// Number of DDRT read transactions the DCPMM has serviced since last AC cycle
func GetReadRequests() ([]MetricReading, error) {
    // stubbed - was not exposed by NVM API
    return []MetricReading{}, fmt.Errorf("stubbed")
}

// Number of DDRT write transactions the DCPMM has serviced since last AC cycle
func GetWriteRequest() ([]MetricReading, error) {
    // stubbed - was not exposed by NVM API
    return []MetricReading{}, fmt.Errorf("stubbed")
}

// Lifetime number of 64 byte reads from media on the DCPMM
func GetTotalMediaReads() ([]MetricReading, error) {
    metrictType := devPerformanceTypeEnum.bytesRead
    return getDevicePerformanceReadings(metrictType)
}

// Lifetime number of 64 byte writes to media on the DCPMM
func GetTotalMediaWrites() ([]MetricReading, error) {
    metrictType := devPerformanceTypeEnum.bytesWritten
    return getDevicePerformanceReadings(metrictType)
}

// Lifetime number of DDRT read transactions the DCPMM has serviced
func GetTotalReadRequests() ([]MetricReading, error) {
    metrictType := devPerformanceTypeEnum.hostReads
    return getDevicePerformanceReadings(metrictType)
}

// Lifetime number of DDRT write transactions the DCPMM has serviced
func GetTotalWriteRequests() ([]MetricReading, error) {
    metrictType := devPerformanceTypeEnum.hostWrites
    return getDevicePerformanceReadings(metrictType)
}

/**
 * Copyright (c) 2020-2021, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduces wrapper for ipmctl library written in C.
 * api_reader.go file exposes external API for exporter to collect
 * NVM readings.
 */

package nvm

import (
	"fmt"
)

const NumberOfAvailableSensors = 10

type device struct {
	uid               nvmUID
	discovery         deviceDiscovery
	sensors           [NumberOfAvailableSensors]sensor
	performance       devicePerformance
	performanceOpstat nvmStatusCodeEnumAttr
	sensorsOpstat     [NumberOfAvailableSensors]nvmStatusCodeEnumAttr
}

type MetricsReader struct {
	deviceCount nvmUint8
	devices     []device
}

func NewMetricsReader() *MetricsReader {
	opstat, count, _ := GetNumberOfDevices()
	if nvmStatusCodeEnum.nvmSuccess != opstat {
		return &MetricsReader{
			deviceCount: 0,
			devices:     make([]device, 0),
		}
	}
	return &MetricsReader{
		deviceCount: count,
		devices:     make([]device, count),
	}
}

func (reader *MetricsReader) GetRequiredReadings() (bool, error) {
	if 0 == reader.deviceCount {
		opstat, count, _ := GetNumberOfDevices()
		if nvmStatusCodeEnum.nvmSuccess != opstat {
			return false, fmt.Errorf("Unable to get number of NVM devices")
		}
		reader.deviceCount = count
	}

	opstat, discoveries, err := GetDevices(reader.deviceCount)
	if nvmStatusCodeEnum.nvmSuccess != opstat {
		return false, err
	}
	for i := 0; i < int(reader.deviceCount); i++ {
		reader.devices[i].uid = nvmUID("8089-a2-1938-00001bf4")
		reader.devices[i].discovery = discoveries[i]
		reader.devices[i].performanceOpstat, reader.devices[i].performance, _ = GetDevicePerformance(discoveries[i].uid)
		for j := sensorTypeEnum.sensorHealth; j < NumberOfAvailableSensors; j++ {
			reader.devices[i].sensorsOpstat[j], reader.devices[i].sensors[j], _ = GetSensor(discoveries[i].uid, j)
		}
	}
	return true, nil
}

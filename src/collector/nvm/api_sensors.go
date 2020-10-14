/**
 * Copyright (c) 2020, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduce wrapper for ipmctl library written in C.
 * api_sensor.go file expose external API for exporter to collect
 * some NVM sensor readings.
 */

package nvm

var SensorLabelNames = []string {
    "uid",
}

type sensorReading MetricReading
type sensorLabels MetricLabels

func (sl sensorLabels) GetLabelValues() ([]string) {
    return getValuesByName(SensorLabelNames, MetricLabels(sl).labels)
}

func (sl sensorLabels) GetLabelNames() ([]string) {
    return SensorLabelNames
}

func (sl sensorLabels) addLabel(name string, value string) {
    MetricLabels(sl).labels[name] = value
}

func newSensorReading(dimmUID nvmUID,
                      readStatus nvmStatusCodeEnumAttr,
                      sensorType sensorTypeEnumAttr,
                      sensorValue nvmUint64) (*sensorReading) {
    sensorReading := new(sensorReading)
    sensorReading.DIMMUID     = string(dimmUID)
    sensorReading.ReadStatus  = int(readStatus)
    sensorReading.MetricType  = uint8(sensorType)
    sensorReading.MetricValue = float64(sensorValue)
    sensorReading.Labels      = sensorLabels(*newMetricLabels())
    return sensorReading
}

func getSensorReadings(sensorType sensorTypeEnumAttr) ([]MetricReading, error) {
    dimmID      := nvmUID("")
    sensorValue := nvmUint64(0)
    opstat, count, err := GetNumberOfDevices()
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, 1)
        sensorReading := *newSensorReading(dimmID, opstat, sensorType, sensorValue)
        results[0] = MetricReading(sensorReading)
        return results, err
    }
    devices := make([]deviceDiscovery, count)
    opstat, devices, err = GetDevices(count)
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, count)
        for i:=0; i < int(count); i++ {
            sensorReading := *newSensorReading(dimmID, opstat, sensorType, sensorValue)
            results[i] = MetricReading(sensorReading)
        }
        return results, err
    }
    results := make([]MetricReading, count)
    for i, dev := range devices {
        opstat, sensor, _ := GetSensor(dev.uid, sensorType)
        sensorReading := *newSensorReading(dev.uid, opstat, sensorType, sensor.reading)
        sensorReading.Labels.addLabel("uid", string(dev.uid))
        results[i] = MetricReading(sensorReading)
    }
    return results, nil
}

// DCPMM health as reported in the SMART log
func GetHealth() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorHealth
    return getSensorReadings(sensorType)
}

// Device media temperature in degrees Celsius
func GetMediaTemperature() ([]MetricReading, error) {
    sensorType  := sensorTypeEnum.sensorMediaTemperature
    return getSensorReadings(sensorType)
}

// Device media temperature in degrees Celsius
func GetControllerTemperature() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    return getSensorReadings(sensorType)
}

// Amount of percentage remaining as a percentage
func GetPercentageRemaining() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    return getSensorReadings(sensorType)
}

// Device shutdowns without notification
func GetLatchedDirtyShutdownCount() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorLatchedDirtyShutdownCount
    return getSensorReadings(sensorType)
}

// Total power-on time over the lifetime of the device
func GetPowerOnTime() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPowerontime
    return getSensorReadings(sensorType)
}

// Total power-on time since the last power cycle of the device
func GetUpTime() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorUptime
    return getSensorReadings(sensorType)
}

// Number of power cycles over the lifetime of the device
func GetPowerCycles() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPowerCycles
    return getSensorReadings(sensorType)
}

// The total number of firmware error log entries
func GetFwErrorCount() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorFWerrorlogcount
    return getSensorReadings(sensorType)
}

// Number of times that the FW received an unexpected power loss
func GetUnlatchedDirtyShutdownCount() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorUnlachedDirtyShutdownCount
    return getSensorReadings(sensorType)
}

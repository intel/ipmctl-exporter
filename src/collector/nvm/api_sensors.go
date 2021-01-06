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

func (reader *MetricsReader) getSensorReadings(sensorType sensorTypeEnumAttr) ([]MetricReading, error) {
    results := make([]MetricReading, reader.deviceCount)
    for i, dev := range reader.devices {
        sensor := dev.sensors[sensorType]
        opstat := dev.sensorsOpstat[sensorType]
        sensorReading := *newSensorReading(dev.uid, opstat, sensorType, sensor.reading)
        sensorReading.Labels.addLabel("uid", string(dev.uid))
        results[i] = MetricReading(sensorReading)
    }
    return results, nil
}

// DCPMM health as reported in the SMART log
func (reader *MetricsReader) GetHealth() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorHealth
    return reader.getSensorReadings(sensorType)
}

// Device media temperature in degrees Celsius
func (reader *MetricsReader) GetMediaTemperature() ([]MetricReading, error) {
    sensorType  := sensorTypeEnum.sensorMediaTemperature
    return reader.getSensorReadings(sensorType)
}

// Device media temperature in degrees Celsius
func (reader *MetricsReader) GetControllerTemperature() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    return reader.getSensorReadings(sensorType)
}

// Amount of percentage remaining as a percentage
func (reader *MetricsReader) GetPercentageRemaining() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    return reader.getSensorReadings(sensorType)
}

// Device shutdowns without notification
func (reader *MetricsReader) GetLatchedDirtyShutdownCount() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorLatchedDirtyShutdownCount
    return reader.getSensorReadings(sensorType)
}

// Total power-on time over the lifetime of the device
func (reader *MetricsReader) GetPowerOnTime() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPowerontime
    return reader.getSensorReadings(sensorType)
}

// Total power-on time since the last power cycle of the device
func (reader *MetricsReader) GetUpTime() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorUptime
    return reader.getSensorReadings(sensorType)
}

// Number of power cycles over the lifetime of the device
func (reader *MetricsReader) GetPowerCycles() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPowerCycles
    return reader.getSensorReadings(sensorType)
}

// The total number of firmware error log entries
func (reader *MetricsReader) GetFwErrorCount() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorFWerrorlogcount
    return reader.getSensorReadings(sensorType)
}

// Number of times that the FW received an unexpected power loss
func (reader *MetricsReader) GetUnlatchedDirtyShutdownCount() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorUnlachedDirtyShutdownCount
    return reader.getSensorReadings(sensorType)
}

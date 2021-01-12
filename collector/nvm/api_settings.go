/**
 * Copyright (c) 2020, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduce wrapper for ipmctl library written in C.
 * api_settings.go file expose external API for exporter to collect
 * some NVM sensor settings information, like for instance sensor thresholds.
 */

package nvm

var SettingsLabelNames = []string {
    "uid",
}

var sensorSettingTypeEnum = &sensorSettingType {
    enabled:                   0,
    upperCriticalThreshold:    1,
    lowerCriticalThreshold:    2,
    upperFatalThreshold:       3,
    lowerFatalThreshold:       4,
    upperNoncriticalThreshold: 5,
    lowerNoncriticalThreshold: 6,
    unknown:                   0xFF,
}

type sensorSettingsReading MetricReading
type settingsLabels MetricLabels
type sensorSettingTypeEnumAttr enumAttr
type sensorSettingType struct {
    enabled                   sensorSettingTypeEnumAttr
    upperCriticalThreshold    sensorSettingTypeEnumAttr
    lowerCriticalThreshold    sensorSettingTypeEnumAttr
    upperFatalThreshold       sensorSettingTypeEnumAttr
    lowerFatalThreshold       sensorSettingTypeEnumAttr
    upperNoncriticalThreshold sensorSettingTypeEnumAttr
    lowerNoncriticalThreshold sensorSettingTypeEnumAttr
    unknown                   sensorSettingTypeEnumAttr
}

func (sl settingsLabels) GetLabelValues() ([]string) {
    return getValuesByName(SensorLabelNames, MetricLabels(sl).labels)
}

func (sl settingsLabels) GetLabelNames() ([]string) {
    return SensorLabelNames
}

func (sl settingsLabels) addLabel(name string, value string) {
    MetricLabels(sl).labels[name] = value
}

func newSensorSettingsReading(dimmUID nvmUID,
                              readStatus nvmStatusCodeEnumAttr,
                              ssType sensorSettingTypeEnumAttr,
                              ssValue nvmUint64) (*sensorSettingsReading) {
    senSettingsReading := new(sensorSettingsReading)
    senSettingsReading.DIMMUID     = string(dimmUID)
    senSettingsReading.ReadStatus  = int(readStatus)
    senSettingsReading.MetricType  = uint8(ssType)
    senSettingsReading.MetricValue = float64(ssValue)
    senSettingsReading.Labels = settingsLabels(*newMetricLabels())
    return senSettingsReading
}

func (reader *MetricsReader) getSensorSettingsReadings(sensorType sensorTypeEnumAttr,
                                                       sensorSettingType sensorSettingTypeEnumAttr) ([]MetricReading, error) {
    results := make([]MetricReading, reader.deviceCount)
    for i, dev := range reader.devices {
        sensor := dev.sensors[sensorType]
        opstat := dev.sensorsOpstat[sensorType]
        sensorValue := nvmUint64(0)
        switch sensorSettingType {
            case sensorSettingTypeEnum.enabled:                   sensorValue = sensor.settings.enabled.toNvmUint64()
            case sensorSettingTypeEnum.upperCriticalThreshold:    sensorValue = sensor.settings.upperCriticalThreshold
            case sensorSettingTypeEnum.lowerCriticalThreshold:    sensorValue = sensor.settings.lowerCriticalThreshold
            case sensorSettingTypeEnum.upperFatalThreshold:       sensorValue = sensor.settings.upperFatalThreshold
            case sensorSettingTypeEnum.lowerFatalThreshold:       sensorValue = sensor.settings.lowerFatalThreshold
            case sensorSettingTypeEnum.upperNoncriticalThreshold: sensorValue = sensor.settings.upperNoncriticalThreshold
            case sensorSettingTypeEnum.lowerNoncriticalThreshold: sensorValue = sensor.settings.lowerNoncriticalThreshold
        }
        senSettingsReading := *newSensorSettingsReading(dev.uid, opstat, sensorSettingType, sensorValue)
        senSettingsReading.Labels.addLabel("uid", string(dev.uid))
        results[i] = MetricReading(senSettingsReading)
    }
    return results, nil
}

// Indictes if firmware notifications are enabled when media temperature sensor
// value is critical
func (reader *MetricsReader) GetMTEnabled() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorMediaTemperature
    sensorSettingType := sensorSettingTypeEnum.enabled
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper media temperature critical threshold
func (reader *MetricsReader) GetMTUpperCriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorMediaTemperature
    sensorSettingType := sensorSettingTypeEnum.upperCriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower media temperature critical threshold
func (reader *MetricsReader) GetMTLowerCriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorMediaTemperature
    sensorSettingType := sensorSettingTypeEnum.lowerCriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper media temperature fatal threshold
func (reader *MetricsReader) GetMTUpperFatalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorMediaTemperature
    sensorSettingType := sensorSettingTypeEnum.upperFatalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower media temperature fatal threshold
func (reader *MetricsReader) GetMTLowerFatalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorMediaTemperature
    sensorSettingType := sensorSettingTypeEnum.lowerFatalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper media temperature noncritical threshold
func (reader *MetricsReader) GetMTUpperNoncriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorMediaTemperature
    sensorSettingType := sensorSettingTypeEnum.upperNoncriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower media temperature noncritical threshold
func (reader *MetricsReader) GetMTLowerNoncriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorMediaTemperature
    sensorSettingType := sensorSettingTypeEnum.lowerNoncriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// Indictes if firmware notifications are enabled when controller temperature sensor
// value is critical
func (reader *MetricsReader) GetCTEnabled() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    sensorSettingType := sensorSettingTypeEnum.enabled
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper controller temperature critical threshold
func (reader *MetricsReader) GetCTUpperCriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    sensorSettingType := sensorSettingTypeEnum.upperCriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower controller temperature critical threshold
func (reader *MetricsReader) GetCTLowerCriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    sensorSettingType := sensorSettingTypeEnum.lowerCriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper controller temperature fatal threshold
func (reader *MetricsReader) GetCTUpperFatalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    sensorSettingType := sensorSettingTypeEnum.upperFatalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower controller temperature fatal threshold
func (reader *MetricsReader) GetCTLowerFatalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    sensorSettingType := sensorSettingTypeEnum.lowerFatalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper controller temperature noncritical threshold
func (reader *MetricsReader) GetCTUpperNoncriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    sensorSettingType := sensorSettingTypeEnum.upperNoncriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower controller temperature noncritical threshold
func (reader *MetricsReader) GetCTLowerNoncriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorControllerTemperature
    sensorSettingType := sensorSettingTypeEnum.lowerNoncriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// Indictes if firmware notifications are enabled when percentage remaining
// value is critical
func (reader *MetricsReader) GetPREnabled() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    sensorSettingType := sensorSettingTypeEnum.enabled
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper percentage remaining critical threshold
func (reader *MetricsReader) GetPRUpperCriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    sensorSettingType := sensorSettingTypeEnum.upperCriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower percentage remaining critical threshold
func (reader *MetricsReader) GetPRLowerCriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    sensorSettingType := sensorSettingTypeEnum.lowerCriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper percentage remaining fatal threshold
func (reader *MetricsReader) GetPRUpperFatalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    sensorSettingType := sensorSettingTypeEnum.upperFatalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower percentage remaining fatal threshold
func (reader *MetricsReader) GetPRLowerFatalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    sensorSettingType := sensorSettingTypeEnum.lowerFatalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The upper percentage remaining noncritical threshold
func (reader *MetricsReader) GetPRUpperNoncriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    sensorSettingType := sensorSettingTypeEnum.upperNoncriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

// The lower percentage remaining noncritical threshold
func (reader *MetricsReader) GetPRLowerNoncriticalThreshold() ([]MetricReading, error) {
    sensorType := sensorTypeEnum.sensorPercentageRemaining
    sensorSettingType := sensorSettingTypeEnum.lowerNoncriticalThreshold
    return reader.getSensorSettingsReadings(sensorType, sensorSettingType)
}

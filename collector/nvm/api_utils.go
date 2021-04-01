// +build cgo

/**
 * Copyright (c) 2020-2021, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduces wrapper for ipmctl library written in C.
 * api_utils.go file contains all helper functions used mostly to create
 * some internal structures
 */

package nvm

// #cgo pkg-config: libipmctl
// #include <include/nvm_management.h>
import "C"

import ()

func getValuesByName(names []string,
	dict map[string]string) []string {
	values := make([]string, len(names))
	for i, name := range names {
		value, found := dict[name]
		if found {
			values[i] = value
		} else {
			values[i] = "unknown"
		}
	}
	return values
}

func makeNVMBool(cValue C.uchar) nvmBool {
	if 0 == cValue {
		return nvmBool(false)
	}
	return nvmBool(true)
}

func makeNVMUint8Array(cValue []C.uchar) []nvmUint8 {
	result := make([]nvmUint8, len(cValue))
	for i := 0; i < len(cValue); i++ {
		result[i] = nvmUint8(cValue[i])
	}
	return result
}

func makeNVMUint16Array(cValue []C.ushort) []nvmUint16 {
	result := make([]nvmUint16, len(cValue))
	for i := 0; i < len(cValue); i++ {
		result[i] = nvmUint16(cValue[i])
	}
	return result
}

func makeNVMNfitDeviceHandle(cValue C.NVM_NFIT_DEVICE_HANDLE) []byte {
	result := make([]byte, 32)
	for i, cval := range cValue {
		result[i] = byte(cval)
	}
	return result
}

func newSensorSettings(cValue C.struct_sensor_settings) *sensorSettings {
	sensorSettings := new(sensorSettings)
	sensorSettings.enabled = makeNVMBool(cValue.enabled)
	sensorSettings.upperCriticalThreshold = nvmUint64(cValue.upper_critical_threshold)
	sensorSettings.lowerCriticalThreshold = nvmUint64(cValue.lower_critical_threshold)
	sensorSettings.upperFatalThreshold = nvmUint64(cValue.upper_fatal_threshold)
	sensorSettings.lowerFatalThreshold = nvmUint64(cValue.lower_fatal_threshold)
	sensorSettings.upperNoncriticalThreshold = nvmUint64(cValue.upper_noncritical_threshold)
	sensorSettings.lowerNoncriticalThreshold = nvmUint64(cValue.lower_noncritical_threshold)
	copy(sensorSettings.reserved[:], makeNVMUint8Array(cValue.reserved[:]))
	return sensorSettings
}

func newDeviceCapabilities(cValue C.struct_device_capabilities) *deviceCapabilities {
	devCap := new(deviceCapabilities)
	devCap.packageSparingCapable = makeNVMBool(cValue.package_sparing_capable)
	devCap.memoryModeCapable = makeNVMBool(cValue.memory_mode_capable)
	devCap.appDirectModeCapable = makeNVMBool(cValue.app_direct_mode_capable)
	copy(devCap.reserved[:], makeNVMUint8Array(cValue.reserved[:]))
	return devCap
}

func newDeviceDiscovery(cValue C.struct_device_discovery) *deviceDiscovery {
	devDisc := new(deviceDiscovery)
	devDisc.allPropertiesPopulated = makeNVMBool(cValue.all_properties_populated)
	devDisc.deviceHandle = makeNVMNfitDeviceHandle(cValue.device_handle)
	devDisc.physicalID = nvmUint16(cValue.physical_id)
	devDisc.vendorID = nvmUint16(cValue.vendor_id)
	devDisc.deviceID = nvmUint16(cValue.device_id)
	devDisc.revisionID = nvmUint16(cValue.revision_id)
	devDisc.channelPos = nvmUint16(cValue.channel_pos)
	devDisc.channelID = nvmUint16(cValue.channel_id)
	devDisc.memoryControllerID = nvmUint16(cValue.memory_controller_id)
	devDisc.socketID = nvmUint16(cValue.socket_id)
	devDisc.nodeControllerID = nvmUint16(cValue.node_controller_id)
	devDisc.memoryType = memoryTypeEnumAttr(cValue.memory_type)
	devDisc.dimmSKU = nvmUint32(cValue.dimm_sku)
	devDisc.manufacturer = nvmManufacturer(makeNVMUint8Array(cValue.manufacturer[:]))
	devDisc.serialNumber = nvmSerialNumber(makeNVMUint8Array(cValue.serial_number[:]))
	devDisc.subsystemVendorID = nvmUint16(cValue.subsystem_vendor_id)
	devDisc.subsystemDeviceID = nvmUint16(cValue.subsystem_device_id)
	devDisc.subsystemRevisionID = nvmUint16(cValue.subsystem_revision_id)
	devDisc.manufacturingInfoValid = makeNVMBool(cValue.manufacturing_info_valid)
	devDisc.manufacturingLocation = nvmUint8(cValue.manufacturing_location)
	devDisc.manufacturingDate = nvmUint16(cValue.manufacturing_date)
	devDisc.partNumber = C.GoString(&cValue.part_number[0])
	devDisc.fwRevision = nvmVersion(C.GoString(&cValue.fw_revision[0]))
	devDisc.fwAPIVersion = nvmVersion(C.GoString(&cValue.fw_api_version[0]))
	devDisc.capacity = nvmUint64(cValue.capacity)
	copy(devDisc.interfaceFormatCodes[:], makeNVMUint16Array(cValue.interface_format_codes[:]))
	devDisc.deviceCapabilities = *newDeviceCapabilities(cValue.device_capabilities)
	devDisc.uid = nvmUID(C.GoString(&cValue.uid[0]))
	devDisc.lockState = lockStateEnumAttr(cValue.lock_state)
	devDisc.manageability = manageabilityStateEnumAttr(cValue.manageability)
	devDisc.controllerRevisionID = nvmUint16(cValue.controller_revision_id)
	devDisc.masterPassphraseEnabled = makeNVMBool(cValue.master_passphrase_enabled)
	copy(devDisc.reserved[:], makeNVMUint8Array(cValue.reserved[:]))
	return devDisc
}

func newSensor(cValue C.struct_sensor) *sensor {
	sensor := new(sensor)
	sensor.stype = sensorTypeEnumAttr(cValue._type)
	sensor.units = sensorUnitsEnumAttr(cValue.units)
	sensor.currentState = sensorStatusEnumAttr(cValue.current_state)
	sensor.reading = nvmUint64(cValue.reading)
	sensor.settings = *newSensorSettings(cValue.settings)
	sensor.lowerCriticalSettable = makeNVMBool(cValue.lower_critical_settable)
	sensor.upperCriticalSettable = makeNVMBool(cValue.upper_critical_settable)
	sensor.lowerCriticalSupport = makeNVMBool(cValue.lower_critical_support)
	sensor.upperCriticalSupport = makeNVMBool(cValue.upper_critical_support)
	sensor.lowerFatalSettable = makeNVMBool(cValue.lower_fatal_settable)
	sensor.upperFatalSettable = makeNVMBool(cValue.upper_fatal_settable)
	sensor.lowerFatalSupport = makeNVMBool(cValue.lower_fatal_support)
	sensor.upperFatalSupport = makeNVMBool(cValue.upper_fatal_support)
	sensor.lowerNoncriticalSettable = makeNVMBool(cValue.lower_noncritical_settable)
	sensor.upperNoncriticalSettable = makeNVMBool(cValue.upper_noncritical_settable)
	sensor.lowerNoncriticalSupport = makeNVMBool(cValue.lower_noncritical_support)
	sensor.upperNoncriticalSupport = makeNVMBool(cValue.upper_noncritical_support)
	copy(sensor.reserved[:], makeNVMUint8Array(cValue.reserved[:]))
	return sensor
}

func newDevicePerformance(cValue C.struct_device_performance) *devicePerformance {
	devPerf := new(devicePerformance)
	devPerf.time = timeT(nvmUint64(cValue.time))
	devPerf.bytesRead = nvmUint64(cValue.bytes_read)
	devPerf.hostReads = nvmUint64(cValue.host_reads)
	devPerf.bytesWritten = nvmUint64(cValue.bytes_written)
	devPerf.hostWrites = nvmUint64(cValue.host_writes)
	devPerf.blockReads = nvmUint64(cValue.block_reads)
	devPerf.blockWrites = nvmUint64(cValue.block_writes)
	copy(devPerf.reserved[:], makeNVMUint8Array(cValue.reserved[:]))
	return devPerf
}

func newMetricLabels() *MetricLabels {
	ml := new(MetricLabels)
	ml.labels = make(map[string]string)
	return ml
}

func getMemoryTypeName(memType memoryTypeEnumAttr) string {
	switch memType {
	case memoryTypeEnum.memoryTypeUnknown:
		return "unknown"
	case memoryTypeEnum.memoryTypeDDR4:
		return "ddr4"
	case memoryTypeEnum.memoryTypeNVMDIMM:
		return "nvm"
	}
	return "unknown"
}

func getLockstateName(lockState lockStateEnumAttr) string {
	switch lockState {
	case lockStateEnum.lockStateUnknown:
		return "unknown"
	case lockStateEnum.lockStateDisable:
		return "disable"
	case lockStateEnum.lockStateUnlocked:
		return "unlocked"
	case lockStateEnum.lockStateLocked:
		return "locked"
	case lockStateEnum.lockStateFrozen:
		return "frozen"
	case lockStateEnum.lockStatePassphraseLimit:
		return "passphrase_limit"
	case lockStateEnum.lockStateNotSupported:
		return "not_supported"
	}
	return "unknown"
}

func getManageabilityName(manageability manageabilityStateEnumAttr) string {
	switch manageability {
	case manageabilityStateEnum.managementUnknown:
		return "unknown"
	case manageabilityStateEnum.managementValidConfig:
		return "valid_config"
	case manageabilityStateEnum.managementInvalidConfig:
		return "invalid_config"
	case manageabilityStateEnum.managementNonFunctional:
		return "non_functional"
	}
	return "unknown"
}

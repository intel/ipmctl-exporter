/**
 * Copyright (c) 2020, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduce wrapper for ipmctl library written in C.
 * api_metadata.go file expose some information about DCPMM taken mostly
 * during DIMM dicovery process. Adding all the information to the labels
 * section of each metric is not a good idea, because some of them may changed
 * during the execution (like for instance fw revision number), that is why
 * these information were exposed as a seperate metrics, were value is always
 * set to "1" ("0" will be set only if for some reasones gathering the
 * data will be impossible). This solution was borrowed from node exporter,
 * and sample useage was introduce here: https://tiantiankan.me/a/5ce42184714c6a22ca1f62f5
 */

package nvm

import (
    "fmt"
)

//Variable initialised during build
var Version string

var DeviceDiscoveryLabelNames = []string {
    "uid",
    "physical_id",
    "vendor_id",
    "device_id",
    "revision_id",
    "channel_pos",
    "channel_id",
    "memory_controller_id",
    "socket_id",
    "node_controller_id",
    "memory_type",
    "sku",
    "manufacturer",
    "serial_number",
    "subsystem_vendor_id",
    "subsystem_device_id",
    "subsystem_revision_id",
    "manufacturing_info_valid",
    "manufacturing_location",
    "manufacturing_date",
    "part_number",
    "fw_revision",
    "fw_api_version",
    "capacity",
    "interface_format_codes",
    "lock_state",
    "manageability",
    "controller_revision_id",
    "master_passphrase_enabled",
}

var DeviceSecurityCapabilitiesLabelNames = []string {
    "uid",
    "passphrase_capable",
    "unlock_device_capable",
    "erase_crypto_capable",
    "master_passphrase_capable",
}

var DeviceCapabilitiesLabelNames = []string {
    "uid",
    "package_sparing_capable",
    "memory_mode_capable",
    "app_direct_mode_capable",
}

var IpmctlExporterLabelNames = []string {
    "version",
}

type deviceDiscoveryReading            MetricReading
type deviceSecurityCapabilitiesReading MetricReading
type deviceCapabilitiesReading         MetricReading
type IpmctlExporterReading             MetricReading

type deviceDiscoveryLabels             MetricLabels
type deviceSecurityCapabilitiesLabels  MetricLabels
type deviceCapabilitiesLabels          MetricLabels
type IpmctlExporterLabels              MetricLabels

func (ddl deviceDiscoveryLabels) GetLabelValues() ([]string) {
    return getValuesByName(DeviceDiscoveryLabelNames, MetricLabels(ddl).labels)
}

func (dscl deviceSecurityCapabilitiesLabels) GetLabelValues() ([]string) {
    return getValuesByName(DeviceSecurityCapabilitiesLabelNames, MetricLabels(dscl).labels)
}

func (dcl deviceCapabilitiesLabels) GetLabelValues() ([]string) {
    return getValuesByName(DeviceCapabilitiesLabelNames, MetricLabels(dcl).labels)
}

func (iel IpmctlExporterLabels) GetLabelValues() ([]string) {
    return getValuesByName(IpmctlExporterLabelNames, MetricLabels(iel).labels)
}

func (ddl deviceDiscoveryLabels) GetLabelNames() ([]string) {
    return DeviceDiscoveryLabelNames
}

func (dscl deviceSecurityCapabilitiesLabels) GetLabelNames() ([]string) {
    return DeviceSecurityCapabilitiesLabelNames
}

func (dcl deviceCapabilitiesLabels) GetLabelNames() ([]string) {
    return DeviceCapabilitiesLabelNames
}

func (iel IpmctlExporterLabels) GetLabelNames() ([]string) {
    return IpmctlExporterLabelNames
}

func (ddl deviceDiscoveryLabels) addLabel(name string, value string) {
    MetricLabels(ddl).labels[name] = value
}

func (dscl deviceSecurityCapabilitiesLabels) addLabel(name string, value string) {
    MetricLabels(dscl).labels[name] = value
}

func (dcl deviceCapabilitiesLabels) addLabel(name string, value string) {
    MetricLabels(dcl).labels[name] = value
}

func (iel IpmctlExporterLabels) addLabel(name string, value string) {
    MetricLabels(iel).labels[name] = value
}

func newDeviceDiscoveryReading(dimmUID nvmUID,
                               ddValue nvmUint64) (*deviceDiscoveryReading) {
    deviceDiscoveryReading := new(deviceDiscoveryReading)
    deviceDiscoveryReading.DIMMUID     = string(dimmUID)
    deviceDiscoveryReading.ReadStatus  = int(0)
    deviceDiscoveryReading.MetricType  = uint8(0)
    deviceDiscoveryReading.MetricValue = float64(ddValue)
    deviceDiscoveryReading.Labels      = deviceDiscoveryLabels(*newMetricLabels())
    return deviceDiscoveryReading
}

func newDeviceSecurityCapabilitiesReading(dimmUID nvmUID,
                                          dscValue nvmUint64) (*deviceSecurityCapabilitiesReading) {
    devSecCapabilitiesReading := new(deviceSecurityCapabilitiesReading)
    devSecCapabilitiesReading.DIMMUID     = string(dimmUID)
    devSecCapabilitiesReading.ReadStatus  = int(0)
    devSecCapabilitiesReading.MetricType  = uint8(0)
    devSecCapabilitiesReading.MetricValue = float64(dscValue)
    devSecCapabilitiesReading.Labels      = deviceSecurityCapabilitiesLabels(*newMetricLabels())
    return devSecCapabilitiesReading
}

func newDeviceCapabilitiesReading(dimmUID nvmUID,
                                  dcValue nvmUint64) (*deviceCapabilitiesReading) {
    devCapabilitiesReading := new(deviceCapabilitiesReading)
    devCapabilitiesReading.DIMMUID     = string(dimmUID)
    devCapabilitiesReading.ReadStatus  = int(0)
    devCapabilitiesReading.MetricType  = uint8(0)
    devCapabilitiesReading.MetricValue = float64(dcValue)
    devCapabilitiesReading.Labels      = deviceCapabilitiesLabels(*newMetricLabels())
    return devCapabilitiesReading
}

func newIpmctlExporterReading(dimmUID nvmUID,
                              ieValue nvmUint64) (*IpmctlExporterReading) {
    ipmctlExpReading := new(IpmctlExporterReading)
    ipmctlExpReading.DIMMUID     = string(dimmUID)
    ipmctlExpReading.ReadStatus  = int(0)
    ipmctlExpReading.MetricType  = uint8(0)
    ipmctlExpReading.MetricValue = float64(ieValue)
    ipmctlExpReading.Labels      = IpmctlExporterLabels(*newMetricLabels())
    return ipmctlExpReading
}

func GetDeviceDiscoveryInfo() ([]MetricReading, error) {
    dimmID      := nvmUID("")
    opstat, count, err := GetNumberOfDevices()
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, 1)
        devDiscoveryReading := *newDeviceDiscoveryReading(dimmID, 0)
        results[0] = MetricReading(devDiscoveryReading)
        return results, err
    }
    devices := make([]deviceDiscovery, count)
    opstat, devices, err = GetDevices(count)
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, count)
        for i:=0; i < int(count); i++ {
            devDiscoveryReading := *newDeviceDiscoveryReading(dimmID, 0)
            results[i] = MetricReading(devDiscoveryReading)
        }
        return results, err
    }
    results := make([]MetricReading, count)
    for i, dev := range devices {
        devDiscoveryReading := *newDeviceDiscoveryReading(dimmID, 1)
        devDiscoveryReading.Labels.addLabel("uid", string(dev.uid))
        devDiscoveryReading.Labels.addLabel("physical_id", dev.physicalID.toString(16))
        devDiscoveryReading.Labels.addLabel("vendor_id", dev.vendorID.toString(16))
        devDiscoveryReading.Labels.addLabel("device_id", dev.deviceID.toString(16))
        devDiscoveryReading.Labels.addLabel("revision_id", dev.revisionID.toString(16))
        devDiscoveryReading.Labels.addLabel("channel_pos", dev.channelPos.toString(16))
        devDiscoveryReading.Labels.addLabel("channel_id", dev.channelID.toString(16))
        devDiscoveryReading.Labels.addLabel("memory_controller_id", dev.memoryControllerID.toString(16))
        devDiscoveryReading.Labels.addLabel("socket_id", dev.socketID.toString(16))
        devDiscoveryReading.Labels.addLabel("node_controller_id", dev.nodeControllerID.toString(16))
        devDiscoveryReading.Labels.addLabel("memory_type", getMemoryTypeName(dev.memoryType))
        devDiscoveryReading.Labels.addLabel("sku", dev.dimmSKU.toString(16))
        devDiscoveryReading.Labels.addLabel("manufacturer", bytesToString([]nvmUint8(dev.manufacturer)))
        devDiscoveryReading.Labels.addLabel("serial_number", bytesToString([]nvmUint8(dev.serialNumber)))
        devDiscoveryReading.Labels.addLabel("subsystem_vendor_id", dev.subsystemVendorID.toString(16))
        devDiscoveryReading.Labels.addLabel("subsystem_device_id", dev.subsystemDeviceID.toString(16))
        devDiscoveryReading.Labels.addLabel("subsystem_revision_id", dev.subsystemRevisionID.toString(16))
        devDiscoveryReading.Labels.addLabel("manufacturing_info_valid", dev.manufacturingInfoValid.toString(10))
        devDiscoveryReading.Labels.addLabel("manufacturing_location", dev.manufacturingLocation.toString(16))
        devDiscoveryReading.Labels.addLabel("manufacturing_date", dev.manufacturingDate.toString(10))
        devDiscoveryReading.Labels.addLabel("part_number", dev.partNumber)
        devDiscoveryReading.Labels.addLabel("fw_revision", string(dev.fwRevision))
        devDiscoveryReading.Labels.addLabel("fw_api_version", string(dev.fwAPIVersion))
        devDiscoveryReading.Labels.addLabel("capacity", dev.capacity.toString(10))
        devDiscoveryReading.Labels.addLabel("interface_format_codes", uint16ToString(dev.interfaceFormatCodes[:]))
        devDiscoveryReading.Labels.addLabel("lock_state", getLockstateName(dev.lockState))
        devDiscoveryReading.Labels.addLabel("manageability", getManageabilityName(dev.manageability))
        devDiscoveryReading.Labels.addLabel("controller_revision_id", dev.controllerRevisionID.toString(16))
        devDiscoveryReading.Labels.addLabel("master_passphrase_enabled", dev.masterPassphraseEnabled.toString(10))
        results[i] = MetricReading(devDiscoveryReading)
    }
    return results, nil
}

func GetDeviceSecurityCapabilitiesInfo() ([]MetricReading, error) {
    dimmID      := nvmUID("")
    opstat, count, err := GetNumberOfDevices()
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, 1)
        devSecCapReading := *newDeviceSecurityCapabilitiesReading(dimmID, 0)
        results[0] = MetricReading(devSecCapReading)
        return results, err
    }
    devices := make([]deviceDiscovery, count)
    opstat, devices, err = GetDevices(count)
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, count)
        for i:=0; i < int(count); i++ {
            devSecCapsReading := *newDeviceSecurityCapabilitiesReading(dimmID, 0)
            results[i] = MetricReading(devSecCapsReading)
        }
        return results, err
    }
    results := make([]MetricReading, count)
    for i, dev := range devices {
        devSecCapsReading := *newDeviceSecurityCapabilitiesReading(dimmID, 1)
        devSecCapsReading.Labels.addLabel("uid", string(dev.uid))
        devSecCapsReading.Labels.addLabel("passphrase_capable", dev.securityCapabilities.passphraseCapable.toString(10))
        devSecCapsReading.Labels.addLabel("unlock_device_capable", dev.securityCapabilities.unlockDeviceCapable.toString(10))
        devSecCapsReading.Labels.addLabel("erase_crypto_capable", dev.securityCapabilities.eraseCryptoCapable.toString(10))
        devSecCapsReading.Labels.addLabel("master_passphrase_capable", dev.securityCapabilities.masterPassphraseCapable.toString(10))
        results[i] = MetricReading(devSecCapsReading)
    }
    return results, nil
}

func GetDeviceCapabilitiesInfo() ([]MetricReading, error) {
    dimmID      := nvmUID("")
    opstat, count, err := GetNumberOfDevices()
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, 1)
        devCapsReading := *newDeviceCapabilitiesReading(dimmID, 0)
        results[0] = MetricReading(devCapsReading)
        return results, err
    }
    devices := make([]deviceDiscovery, count)
    opstat, devices, err = GetDevices(count)
    if nvmStatusCodeEnum.nvmSuccess != opstat {
        results := make([]MetricReading, count)
        for i:=0; i < int(count); i++ {
            devCapsReading := *newDeviceCapabilitiesReading(dimmID, 0)
            results[0] = MetricReading(devCapsReading)
            return results, err
        }
    }
    results := make([]MetricReading, count)
    for i, dev := range devices {
        devCapsReading := *newDeviceCapabilitiesReading(dimmID, 1)
        devCapsReading.Labels.addLabel("uid", string(dev.uid))
        devCapsReading.Labels.addLabel("package_sparing_capable", dev.deviceCapabilities.packageSparingCapable.toString(10))
        devCapsReading.Labels.addLabel("memory_mode_capable", dev.deviceCapabilities.memoryModeCapable.toString(10))
        devCapsReading.Labels.addLabel("app_direct_mode_capable", dev.deviceCapabilities.appDirectModeCapable.toString(10))
        results[i] = MetricReading(devCapsReading)
    }
    return results, nil
}

func GetIpmctlExporterInfo() ([]MetricReading, error) {
    dimmID      := nvmUID("")
    const ieInfoLabelsCount = 1
    results := make([]MetricReading, ieInfoLabelsCount)
    var err error = nil
    ipmctlExpReading := *newIpmctlExporterReading(dimmID, 0)
    if len(Version) == 0 {
        err = fmt.Errorf("Unable to get IPMCTL exporter info")
    }
    ipmctlExpReading.MetricValue = 1
    ipmctlExpReading.Labels.addLabel("version", Version)
    results[0] = MetricReading(ipmctlExpReading)
    return results, err
}
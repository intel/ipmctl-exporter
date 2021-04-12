// +build cgo

/**
 * Copyright (c) 2020-2021, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduces wrapper for ipmctl library written in C.
 * api_lib.go file exposes exactly the same API as C library, however C types were
 * replace by proper go-like wrapers (see api_types.go). This API should be used
 * only inside the nvm package to write function used by exporter to collect
 * specific metrics (only these functions should be marked as public). It's a good
 * place to add wrapper for ipmctl API function needed to expose more metrics,
 * so far the biggest part of these functions were stubbed.
 */

package nvm

// #cgo pkg-config: libipmctl
// #include <include/nvm_management.h>
import "C"
import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Init - libipmctl initialiation method
func Init() (bool, error) {
	if isLibInitialized {
		log.Info("libipmctl was already initialized, nothing to be done")
		return true, nil
	}
	opstat := C.nvm_init()
	if C.NVM_SUCCESS != opstat {
		log.Error("libipmctl initialization failed with status:", opstat)
		return false, fmt.Errorf("libipmctl initialization failed with status: %d", opstat)
	}
	isLibInitialized = true
	return true, nil
}

// Uninit - libipmctl un-initialization method
func Uninit() (bool, error) {
	if !isLibInitialized {
		log.Warn("libipmctl was not initialized, nothing to be done")
		return true, nil
	}
	C.nvm_uninit()
	isLibInitialized = false
	return true, nil
}

// GetNumberOfDevices retrieves the number of devices installed in the system whether they are
// fully compatible with the current native API library version or not.
// @pre The caller must have administrative privileges.
// @remarks This method should be called before #nvm_get_devices.
// @remarks The number of devices can be 0.
// @return status, devices counter, error obj
func GetNumberOfDevices() (nvmStatusCodeEnumAttr, nvmUint8, error) {
	cCount := C.uint(0)
	cOpstat := C.nvm_get_number_of_devices(&cCount)
	opstat := nvmStatusCodeEnumAttr(cOpstat)
	count := nvmUint8(cCount)
	if C.NVM_SUCCESS != cOpstat {
		return opstat, count, fmt.Errorf("Unable to get number of NVM devices")
	}
	return opstat, count, nil
}

// ConfFileInit - stubbed - implement if needed
func ConfFileInit() {
}

// ConfFileFlush - stubbed - implement if needed
func ConfFileFlush() {
}

// GetDimmID - stubbed - implement if needed
func GetDimmID(iniFileName nvmUID) (nvmStatusCodeEnumAttr, uint, uint, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	dimmID := uint(0)
	dimmHandle := uint(0)
	return opstat, dimmID, dimmHandle, fmt.Errorf("Method is not implemented")
}

// GetConfigInt - stubbed - implement if needed
func GetConfigInt(paramName string, defaultVal int) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetHostName - stubbed - implement if needed
func GetHostName() (nvmStatusCodeEnumAttr, string, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	hostname := "localhost"
	return opstat, hostname, fmt.Errorf("Method is not implemented")
}

// GetHost - stubbed - implement if needed
func GetHost() (nvmStatusCodeEnumAttr, host, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, host{}, fmt.Errorf("Method is not implemented")
}

// GetSwInventory - stubbed - implement if needed
func GetSwInventory() (nvmStatusCodeEnumAttr, host, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, host{}, fmt.Errorf("Method is not implemented")
}

// GetNumberOfSockets - stubbed - implement if needed
func GetNumberOfSockets() (nvmStatusCodeEnumAttr, int, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := int(0)
	return opstat, count, fmt.Errorf("Method is not implemented")
}

// GetSockets - stubbed - implement if needed
func GetSockets() (nvmStatusCodeEnumAttr, socket, nvmUint16, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := nvmUint16(0)
	return opstat, socket{}, count, fmt.Errorf("Method is not implemented")
}

// GetNumberOfMemoryTopologyDevices - stubbed - implement if needed
func GetNumberOfMemoryTopologyDevices() (nvmStatusCodeEnumAttr, uint, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := uint(0)
	return opstat, count, fmt.Errorf("Method is not implemented")
}

// GetMemoryTopology - stubbed - implement if needed
func GetMemoryTopology(count uint8) (nvmStatusCodeEnumAttr, memoryTopology, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, memoryTopology{}, fmt.Errorf("Method is not implemented")
}

// @brief Retrieves #device_discovery information
// about each device in the system whether they are fully compatible
// with the current native API library version or not.
// @param[in] count: The number of elements in array.
// @pre The caller must have administrative privileges.
// @remarks To allocate the array of #device_discovery structures,
// call #nvm_get_device_count before calling this method.
// @return status, An array of #DeviceDiscovery structures, error obj
// ::NVM_SUCCESS @n
// ::NVM_ERR_INVALID_PARAMETER @n
// ::NVM_ERR_UNKNOWN @n
// ::NVM_ERR_BAD_SIZE @n
func GetDevices(count nvmUint8) (nvmStatusCodeEnumAttr, []deviceDiscovery, error) {
	cDevices := make([]C.struct_device_discovery, count)
	if count <= 0 {
		return nvmStatusCodeEnum.nvmErrDIMMNotFound, []deviceDiscovery{},
			fmt.Errorf("No NVM devices detected")
	}
	cOpstat := C.nvm_get_devices(&cDevices[0], C.uchar(count))
	opstat := nvmStatusCodeEnumAttr(cOpstat)
	if C.NVM_SUCCESS != cOpstat {
		return opstat, []deviceDiscovery{},
			fmt.Errorf("Unable to get all NVM devices")
	}
	devices := make([]deviceDiscovery, count)
	for i, cDev := range cDevices {
		devices[i] = *newDeviceDiscovery(cDev)
	}
	return opstat, devices, nil
}

// GetDevicesNfit - stubbed - implement if needed
func GetDevicesNfit(count nvmUint8) (nvmStatusCodeEnumAttr, deviceDiscovery, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, deviceDiscovery{}, fmt.Errorf("Method is not implemented")
}

// GetDeviceDiscovery - stubbed - implement if needed
func GetDeviceDiscovery(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceDiscovery, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, deviceDiscovery{}, fmt.Errorf("Method is not implemented")
}

// GetDeviceStatus - stubbed - implement if needed
func GetDeviceStatus(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceStatus, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, deviceStatus{}, fmt.Errorf("Method is not implemented")
}

// GetPMOMRegister - stubbed - implement if needed
func GetPMOMRegister(deviceUID nvmUID,
	smartDataMask uint8) (nvmStatusCodeEnumAttr, pmonRegisters, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, pmonRegisters{}, fmt.Errorf("Method is not implemented")
}

// SetPMONRegisters - stubbed - implement if needed
func SetPMONRegisters(deviceUID nvmUID, pmonGroupEnable uint8) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetDeviceSettings - stubbed - implement if needed
func GetDeviceSettings(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceSettings, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, deviceSettings{}, fmt.Errorf("Method is not implemented")
}

// GetDeviceDetails - stubbed - implement if needed
func GetDeviceDetails(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceDetails, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, deviceDetails{}, fmt.Errorf("Method is not implemented")
}

// @brief Retrieves a current snapshot of the performance metrics
// for the specified device.
// @param[in] device_uid: The device identifier.
// @pre The caller must have administrative privileges.
// @pre The device is manageable.
// @return #DevicePerformance structure, operation status:
// ::NVM_SUCCESS @n
// ::NVM_ERR_INVALID_PARAMETER @n
// ::NVM_ERR_UNKNOWN @n
func GetDevicePerformance(deviceUID nvmUID) (nvmStatusCodeEnumAttr,
	devicePerformance,
	error) {
	cResult := C.struct_device_performance{}
	cDeviceUID := deviceUID.toCharArray()
	cOpstat := C.nvm_get_device_performance(&cDeviceUID[0], &cResult)
	if C.NVM_SUCCESS != cOpstat {
		opstat := nvmStatusCodeEnumAttr(cOpstat)
		return opstat, devicePerformance{},
			fmt.Errorf("Unable to get performance readings from DIMM: %s", deviceUID)
	}
	result := *newDevicePerformance(cResult)
	opstat := nvmStatusCodeEnumAttr(cOpstat)
	return opstat, result, nil
}

// GetDeviceFwImageInfo - stubbed - implement if needed
func GetDeviceFwImageInfo(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceFWInfo, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, deviceFWInfo{}, fmt.Errorf("Method is not implemented")
}

// UpdateDeviceFw - stubbed - implement if needed
func UpdateDeviceFw(deviceUID nvmUID,
	path nvmPath,
	pathLen nvmSize,
	force bool) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// ExamineDeviceFw - stubbed - implement if needed
func ExamineDeviceFw(deviceUID nvmUID,
	path nvmPath,
	pathLen nvmSize,
	imageVersion nvmVersion,
	imageVersionLen nvmSize) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetNVMCapabilities - stubbed - implement if needed
func GetNVMCapabilities() (nvmStatusCodeEnumAttr, nvmCapabilities, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, nvmCapabilities{}, fmt.Errorf("Method is not implemented")
}

// GetNVMCapacities - stubbed - implement if needed
func GetNVMCapacities() (nvmStatusCodeEnumAttr, deviceCapacities, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, deviceCapacities{}, fmt.Errorf("Method is not implemented")
}

// GetSensors - stubbed - implement if needed
func GetSensors(deviceUID nvmUID,
	count nvmUint16) (nvmStatusCodeEnumAttr, []sensor, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, []sensor{}, fmt.Errorf("Method is not implemented")
}

// @brief Retrieves a specific health sensor from the specified DCPMM.
// @param[in] deviceUid: The device identifier.
// @param[in] stype: The specific #SensorType to retrieve.
// @pre The caller has administrative privileges.
// @pre The device is manageable.
// @return #Sensor structure allocated by the caller, status:
// ::NVM_SUCCESS @n
// ::NVM_ERR_INVALID_PARAMETER @n
// ::NVM_ERR_UNKNOWN @n
func GetSensor(deviceUID nvmUID,
	stype sensorTypeEnumAttr) (nvmStatusCodeEnumAttr, sensor, error) {
	cResult := C.struct_sensor{}
	cDeviceUID := deviceUID.toCharArray()
	cSensorType := uint32(stype)
	cOpstat := C.nvm_get_sensor(&cDeviceUID[0], cSensorType, &cResult)
	if C.NVM_SUCCESS != cOpstat {
		opstat := nvmStatusCodeEnumAttr(cOpstat)
		return opstat, sensor{},
			fmt.Errorf("Unable to get readings from sensor number: %d, DIMM: %s", stype, cDeviceUID)
	}
	result := *newSensor(cResult)
	opstat := nvmStatusCodeEnumAttr(cOpstat)
	return opstat, result, nil
}

// SetSensorSettings - stubbed - implement if needed
func SetSensorSettings(deviceUID nvmUID,
	stype sensorType,
	settings sensorSettings) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetNumberOfEvents - stubbed - implement if needed
func GetNumberOfEvents(filter eventFilter) (nvmStatusCodeEnumAttr, int, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := int(0)
	return opstat, count, fmt.Errorf("Method is not implemented")
}

// GetEvents - stubbed - implement if needed
func GetEvents(filter eventFilter, count nvmUint16) (nvmStatusCodeEnumAttr, []event, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, []event{}, fmt.Errorf("Method is not implemented")
}

// PurgeEvents - stubbed - implement if needed
func PurgeEvents(filter eventFilter) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// AcknowledgeEvent - stubbed - implement if needed
func AcknowledgeEvent(eventID nvmUint32) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetNumberOfRegions - stubbed - implement if needed
func GetNumberOfRegions() (nvmStatusCodeEnumAttr, nvmUint8, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := nvmUint8(0)
	return opstat, count, fmt.Errorf("Method is not implemented")
}

// GetNumberOfRegionsEx - - stubbed - implement if needed
func GetNumberOfRegionsEx(useNfit nvmBool) (nvmStatusCodeEnumAttr, nvmUint8, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := nvmUint8(0)
	return opstat, count, fmt.Errorf("Method is not implemented")
}

// GetRegions - stubbed - implement if needed
func GetRegions() (nvmStatusCodeEnumAttr, []region, nvmUint8, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := nvmUint8(0)
	return opstat, []region{}, count, fmt.Errorf("Method is not implemented")
}

// GetRegionsEx - stubbed - implement if needed
func GetRegionsEx(useNfit nvmBool) (nvmStatusCodeEnumAttr, []region, nvmUint8, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := nvmUint8(0)
	return opstat, []region{}, count, fmt.Errorf("Method is not implemented")
}

// CreateConfigGoal - stubbed - implement if needed
func CreateConfigGoal(deviceUIDs []nvmUID,
	deviceUIDsCount nvmUint32,
	goal configGoalInput) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetConfigGoal - stubbed - implement if needed
func GetConfigGoal(deviceUIDs []nvmUID,
	deviceUIDsCount nvmUint32) (nvmStatusCodeEnumAttr, configGoal, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, configGoal{}, fmt.Errorf("Method is not implemented")
}

// DeleteConfigGoal - stubbed - implement if needed
func DeleteConfigGoal(deviceUIDs []nvmUID, deviceUIDsCount nvmUint32) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// DumpGoalConfig - stubbed - implement if needed
func DumpGoalConfig(file nvmPath, fileLen nvmSize) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// LoadGoalConfig - stubbed - implement if needed
func LoadGoalConfig(file nvmPath, fileLen nvmSize) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetMajorVersion - stubbed - implement if needed
func GetMajorVersion() int {
	majorVersion := int(0)
	return majorVersion
}

// GetMinorVersion - stubbed - implement if needed
func GetMinorVersion() int {
	minorVersion := int(0)
	return minorVersion
}

// GetHotfixNumber - stubbed - implement if needed
func GetHotfixNumber() int {
	hotfixVersionNumber := int(0)
	return hotfixVersionNumber
}

// GetBuildNumber - stubbed - implement if needed
func GetBuildNumber() int {
	buildVersionNumber := int(0)
	return buildVersionNumber
}

// GetVersion - stubbed - implement if needed
func GetVersion(strLen nvmSize) (nvmStatusCodeEnumAttr, nvmVersion, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	versionStr := nvmVersion("")
	return opstat, versionStr, fmt.Errorf("Method is not implemented")
}

// GatherSupport - stubbed - implement if needed
func GatherSupport(supportFile nvmPath, supportFileLen nvmSize) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// InjectDeviceError - stubbed - implement if needed
func InjectDeviceError(deviceUID nvmUID, derror deviceError) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// ClearInjectedDeviceError - stubbed - implement if needed
func ClearInjectedDeviceError(deviceUID nvmUID,
	derror deviceError) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// RunDiagnostic - stubbed - implement if needed
func RunDiagnostic(deviceUID nvmUID,
	diagnostic diagnostic) (nvmStatusCodeEnumAttr, nvmUint32, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	results := nvmUint32(0)
	return opstat, results, fmt.Errorf("Method is not implemented")
}

// SetUserPreferences - stubbed - implement if needed
func SetUserPreferences(key nvmPreferenceKey,
	value nvmPreferenceValue) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// DebugLoggingEnabled - stubbed - implement if needed
func DebugLoggingEnabled() (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// ToggleDebugLogging - stubbed - implement if needed
func ToggleDebugLogging(enabled nvmBool) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetJobs - stubbed - implement if needed
func GetJobs(count nvmUint32) (nvmStatusCodeEnumAttr, job, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, job{}, fmt.Errorf("Method is not implemented")
}

// CreateContext - stubbed - implement if needed
func CreateContext() (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// FreeContext - stubbed - implement if needed
func FreeContext(foce nvmBool) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// SendDevicePassthroughCmd - stubbed - implement if needed
func SendDevicePassthroughCmd(deviceUID nvmUID, cmd devicePTCmd) (nvmStatusCodeEnumAttr, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, fmt.Errorf("Method is not implemented")
}

// GetFwErrorLogEntryCmd - stubbed - implement if needed
func GetFwErrorLogEntryCmd(deviceUID nvmUID,
	seqNum uint,
	logLevel uint8,
	logType uint8) (nvmStatusCodeEnumAttr, errorLog, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, errorLog{}, fmt.Errorf("method is not implemented")
}

// GetFwErrLogStats - stubbed - implement if needed
func GetFwErrLogStats(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceErrorLogStatus, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, deviceErrorLogStatus{}, fmt.Errorf("method is not implemented")
}

// GetNumberOfEffectLogEntries - stubbed - implement if needed
func GetNumberOfEffectLogEntries(deviceUID nvmUID) (nvmStatusCodeEnumAttr, nvmUint32, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := nvmUint32(0)
	return opstat, count, fmt.Errorf("method is not implemented")
}

// GetCommandEffectLog - stubbed - implement if needed
func GetCommandEffectLog(deviceUID nvmUID) (nvmStatusCodeEnumAttr, []commandEffectLog, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	return opstat, []commandEffectLog{}, fmt.Errorf("method is not implemented")
}

// GetCommandAccessPolicy - stubbed - implement if needed
func GetCommandAccessPolicy(deviceUID nvmUID) (nvmStatusCodeEnumAttr, []commandAccessPolicy, nvmUint32, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := nvmUint32(0)
	return opstat, []commandAccessPolicy{}, count, fmt.Errorf("method is not implemented")
}

// GetNumberOfCapEntries - stubbed - implement if needed
func GetNumberOfCapEntries(deviceUID nvmUID) (nvmStatusCodeEnumAttr, nvmUint32, error) {
	opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
	count := nvmUint32(0)
	return opstat, count, fmt.Errorf("method is not implemented")
}

// SyncLockAPI - stubbed - implement if needed
func SyncLockAPI() {
}

//SyncUnlockAPI - stubbed - implement if needed
func SyncUnlockAPI() {
}

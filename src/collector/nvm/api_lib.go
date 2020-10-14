// +build cgo

/**
 * Copyright (c) 2020, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduce wrapper for ipmctl library written in C.
 * api_lib.go file expose exactly the same API as C library however or C types were
 * replace by proper go-like wrapers (see api_types.go). This API should be used
 * only inside the nvm package to write function used by exporter to collect
 * specific metrics (only these functions should be marked as public). It's a good
 * place to add wrapper for ipmictl API function needed to expose more metrics,
 * so far the biggest part of these functions were stubbed.
 */
package nvm

// #cgo pkg-config: libipmctl
// #include <include/nvm_management.h>
import "C"
import (
    "fmt"
)

// libipmctl initialiation method
func Init() (bool, error) {
    if isLibInitialized {
        fmt.Printf("libipmctl was already initialized, nothing to be done\n")
        return true, nil
    }
    opstat := C.nvm_init()
    if C.NVM_SUCCESS != opstat {
        fmt.Printf("libipmctl initialization failed with status: %d\n", opstat)
        return false, fmt.Errorf("libipmctl initialization failed with status: %d", opstat)
    }
    isLibInitialized = true
    return true, nil
}

// libipmctl un-initialization method
func Uninit() (bool, error) {
    if !isLibInitialized {
        fmt.Printf("libipmctl was not initialized, nothing to be done\n")
        return true, nil
    }
    C.nvm_uninit()
    isLibInitialized = false
    return true, nil
}

// Retrieves the number of devices installed in the system whether they are
// fully compatible with the current native API library version or not.
// @pre The caller must have administrative privileges.
// @remarks This method should be called before #nvm_get_devices.
// @remarks The number of devices can be 0.
// @return status, devices counter, error obj
func GetNumberOfDevices() (nvmStatusCodeEnumAttr, nvmUint8, error) {
    cCount  := C.uint(0)
    cOpstat := C.nvm_get_number_of_devices(&cCount)
    opstat  := nvmStatusCodeEnumAttr(cOpstat)
    count   := nvmUint8(cCount)
    if C.NVM_SUCCESS != cOpstat {
        return opstat, count, fmt.Errorf("Unable to get number of NVM devices")
    }
    return opstat, count, nil
}

func ConfFileInit() () {
    // stubbed - implement if needed
}

func ConfFileFlush() () {
    // stubbed - implement if needed
}

func GetDimmId(iniFileName nvmUID) (nvmStatusCodeEnumAttr, uint, uint, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    dimmId := uint(0)
    dimmHandle := uint(0)
    return opstat, dimmId, dimmHandle, fmt.Errorf("Method is not implemented")
}

func GetConfigInt(paramName string, defaultVal int) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetHostName() (nvmStatusCodeEnumAttr, string, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    hostname := "localhost"
    return opstat, hostname, fmt.Errorf("Method is not implemented")
}

func GetHost() (nvmStatusCodeEnumAttr, host, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, host{}, fmt.Errorf("Method is not implemented")
}

func GetSwInventory() (nvmStatusCodeEnumAttr, host, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, host{}, fmt.Errorf("Method is not implemented")
}

func GetNumberOfSockets() (nvmStatusCodeEnumAttr, int, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    count := int(0)
    return opstat, count, fmt.Errorf("Method is not implemented")
}

func GetSockets() (nvmStatusCodeEnumAttr, socket, nvmUint16, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    count  := nvmUint16(0)
    return opstat, socket{}, count, fmt.Errorf("Method is not implemented")
}

func GetNumberOfMemoryTopologyDevices() (nvmStatusCodeEnumAttr, uint, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    count  := uint(0)
    return opstat, count, fmt.Errorf("Method is not implemented")
}

func GetMemoryTopology(count uint8) (nvmStatusCodeEnumAttr, memoryTopology, error) {
    // stubbed - implement if needed
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

func GetDevicesNfit(count nvmUint8) (nvmStatusCodeEnumAttr, deviceDiscovery, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, deviceDiscovery{}, fmt.Errorf("Method is not implemented")
}

func GetDeviceDiscovery(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceDiscovery, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, deviceDiscovery{}, fmt.Errorf("Method is not implemented")
}

func GetDeviceStatus(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceStatus, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, deviceStatus{}, fmt.Errorf("Method is not implemented")
}

func GetPMOMRegister(deviceUID nvmUID,
                     smartDataMask uint8) (nvmStatusCodeEnumAttr, pmonRegisters, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, pmonRegisters{}, fmt.Errorf("Method is not implemented")
}

func SetPMONRegisters(deviceUID nvmUID, pmonGroupEnable uint8) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetDeviceSettings(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceSettings, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, deviceSettings{}, fmt.Errorf("Method is not implemented")
}

func GetDeviceDetails(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceDetails, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, deviceDetails{}, fmt.Errorf("Method is not implemented")
}

// @brief Retrieve a current snapshot of the performance metrics for the device
// specified.
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
    cDeviceUid := deviceUID.toCharArray()
    cOpstat := C.nvm_get_device_performance(&cDeviceUid[0], &cResult)
    if C.NVM_SUCCESS != cOpstat {
        opstat := nvmStatusCodeEnumAttr(cOpstat)
        return opstat, devicePerformance{},
               fmt.Errorf("Unable to get performance readings from DIMM: %s", deviceUID)
    }
    result := *newDevicePerformance(cResult)
    opstat := nvmStatusCodeEnumAttr(cOpstat)
    return opstat, result, nil
}


func GetDeviceFwImageInfo(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceFWInfo, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, deviceFWInfo{}, fmt.Errorf("Method is not implemented")
}

func UpdateDeviceFw(deviceUID nvmUID,
                    path nvmPath,
                    pathLen nvmSize,
                    force bool) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func ExamineDeviceFw(deviceUID nvmUID,
                     path nvmPath,
                     pathLen nvmSize,
                     imageVersion nvmVersion,
                     imageVersionLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetNVMCapabilities() (nvmStatusCodeEnumAttr, nvmCapabilities, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, nvmCapabilities{}, fmt.Errorf("Method is not implemented")
}

func GetNVMCapacities() (nvmStatusCodeEnumAttr, deviceCapacities, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, deviceCapacities{}, fmt.Errorf("Method is not implemented")
}

func GetSensors(deviceUID nvmUID,
                count nvmUint16) (nvmStatusCodeEnumAttr, []sensor, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, []sensor{}, fmt.Errorf("Method is not implemented")
}

// @brief Retrieve a specific health sensor from the specified DCPMM.
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

func SetSensorSettings(deviceUID nvmUID,
                       stype sensorType,
                       settings sensorSettings) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func SetPassphrase(deviceUID nvmUID,
                   oldPassphrase nvmPassphrase,
                   oldPassphraseLen nvmSize,
                   newPassphrase nvmPassphrase,
                   newPassphraseLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func RemovePassphrase(deviceUID nvmUID,
                      passphrase nvmPassphrase,
                      passphraseLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func UnlockDevice(deviceUID nvmUID,
                  passphrase nvmPassphrase,
                  passphraseLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func FreezelockDevice(deviceUID nvmUID) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func EraseDevice(deviceUID nvmUID,
                 passphrase nvmPassphrase,
                 passphraseLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func SetMasterPassphrase(deviceUID nvmUID,
                         oldMasterPassphrase nvmPassphrase,
                         oldMasterPassphraseLen nvmSize,
                         newMasterPassphrase nvmPassphrase,
                         newMasterPassphraseLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetNumberOfEvents(filter eventFilter) (nvmStatusCodeEnumAttr, int, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    count  := int(0)
    return opstat, count, fmt.Errorf("Method is not implemented")
}

func GetEvents(filter eventFilter, count nvmUint16) (nvmStatusCodeEnumAttr, []event, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, []event{}, fmt.Errorf("Method is not implemented")
}

func PurgeEvents(filter eventFilter) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func AcknowledgeEvent(eventID nvmUint32) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetNumberOfRegions() (nvmStatusCodeEnumAttr, nvmUint8, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    count  := nvmUint8(0)
    return opstat, count, fmt.Errorf("Method is not implemented")
}

func GetNumberOfRegionsEx(useNfit nvmBool) (nvmStatusCodeEnumAttr, nvmUint8, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    count  := nvmUint8(0)
    return opstat, count, fmt.Errorf("Method is not implemented")
}

func GetRegions() (nvmStatusCodeEnumAttr, []region, nvmUint8, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    count  := nvmUint8(0)
    return opstat, []region{}, count, fmt.Errorf("Method is not implemented")
}

func GetRegionsEx(useNfit nvmBool) (nvmStatusCodeEnumAttr, []region, nvmUint8, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    count := nvmUint8(0)
    return opstat, []region{}, count, fmt.Errorf("Method is not implemented")
}

func CreateConfigGoal(deviceUIDs []nvmUID,
                      deviceUIDsCount nvmUint32,
                      goal configGoalInput) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetConfigGoal(deviceUIDs []nvmUID,
                   deviceUIDsCount nvmUint32) (nvmStatusCodeEnumAttr, configGoal, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, configGoal{}, fmt.Errorf("Method is not implemented")
}

func DeleteConfigGoal(deviceUIDs []nvmUID, deviceUIDsCount nvmUint32) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func DumpGoalConfig(file nvmPath, fileLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func LoadGoalConfig(file nvmPath, fileLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetMajorVersion() (int) {
    // stubbed - implement if needed
    majorVersion := int(0)
    return majorVersion
}

func GetMinorVersion() (int) {
    // stubbed - implement if needed
    minorVersion := int(0)
    return minorVersion
}

func GetHotfixNumber() (int) {
    // stubbed - implement if needed
    hotfixVersionNumber := int(0)
    return hotfixVersionNumber
}

func GetBuildNumber() (int) {
    // stubbed - implement if needed
    buildVersionNumber := int(0)
    return buildVersionNumber
}

func GetVersion(strLen nvmSize) (nvmStatusCodeEnumAttr, nvmVersion, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    versionStr := nvmVersion("")
    return opstat, versionStr, fmt.Errorf("Method is not implemented")
}

func GatherSupport(supportFile nvmPath, supportFileLen nvmSize) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func InjectDeviceError(deviceUID nvmUID, derror deviceError) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func ClearInjectedDeviceError(deviceUID nvmUID,
                              derror deviceError) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func RunDiagnostic(deviceUID nvmUID,
                   diagnostic diagnostic) (nvmStatusCodeEnumAttr, nvmUint32, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    results := nvmUint32(0)
    return opstat, results, fmt.Errorf("Method is not implemented")
}

func SetUserPreferences(key nvmPreferenceKey,
                        value nvmPreferenceValue) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func ClearDimmLsa(deviceUid nvmUID) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func DebugLoggingEnabled() (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func ToggleDebugLogging(enabled nvmBool) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetJobs(count nvmUint32) (nvmStatusCodeEnumAttr, job, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, job{}, fmt.Errorf("Method is not implemented")
}

func CreateContext() (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func FreeContext(foce nvmBool) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func SendDevicePassthroughCmd(deviceUID nvmUID, cmd devicePTCmd) (nvmStatusCodeEnumAttr, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, fmt.Errorf("Method is not implemented")
}

func GetFwErrorLogEntryCmd(deviceUID nvmUID,
                           seqNum uint,
                           logLevel uint8,
                           logType uint8) (nvmStatusCodeEnumAttr, errorLog, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, errorLog{}, fmt.Errorf("method is not implemented")
}

func GetFwErrLogStats(deviceUID nvmUID) (nvmStatusCodeEnumAttr, deviceErrorLogStatus, error) {
    // stubbed - implement if needed
    opstat := nvmStatusCodeEnum.nvmErrAPINotSupported
    return opstat, deviceErrorLogStatus{}, fmt.Errorf("method is not implemented")
}

func SyncLockApi() () {
    // stubbed
}

func SyncUnlockApi() () {
    // stubbed
}

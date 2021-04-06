// +build cgo

/**
 * Copyright (c) 2020-2021, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduces wrapper for ipmctl library written in C.
 * api_types.go is a wrapper for all types used by ipmctl library written in C
 * These types and only these should be used by lib API functions wrapper.
 */
package nvm

// #cgo pkg-config: libipmctl
// #include <include/nvm_management.h>
import "C"

type Labels interface {
	GetLabelValues() []string
	GetLabelNames() []string
	addLabel(name string, value string)
}

type MetricReading struct {
	DIMMUID     string
	ReadStatus  int
	MetricType  uint8
	MetricValue float64
	Labels      Labels
}

type MetricLabels struct {
	labels map[string]string
}

type (
	nvmUID                  string
	nvmUint8                uint8
	nvmUint16               uint16
	nvmUint32               uint32
	nvmUint64               uint64
	nvmBool                 bool
	timeT                   uint64
	enumAttr                int
	nvmPath                 string
	nvmSize                 uint64
	nvmPassphrase           string
	nvmEventMsg             string
	nvmEventArg             string
	diagnosticThresholdType nvmUint64
	nvmStatusCodeEnumAttr   enumAttr
	nvmStatusCode           struct {
		nvmSuccess                                   nvmStatusCodeEnumAttr
		nvmSuccessFWResetRequired                    nvmStatusCodeEnumAttr
		nvmErrOperationNotStarted                    nvmStatusCodeEnumAttr
		nvmErrOperationFailed                        nvmStatusCodeEnumAttr
		nvmErrForceRequired                          nvmStatusCodeEnumAttr
		nvmErrInvalidParameter                       nvmStatusCodeEnumAttr
		nvmErrCommandNotSupportedByThisSKU           nvmStatusCodeEnumAttr
		nvmErrDIMMNotFound                           nvmStatusCodeEnumAttr
		nvmErrDIMMIDDuplicated                       nvmStatusCodeEnumAttr
		nvmErrSocketIDNotValid                       nvmStatusCodeEnumAttr
		nvmErrSocketIDIncompatiblewDIMMID            nvmStatusCodeEnumAttr
		nvmErrSocketIDDuplicated                     nvmStatusCodeEnumAttr
		nvmErrConfigNotSupportedByCurrentSKU         nvmStatusCodeEnumAttr
		nvmErrManageableDIMMNotFound                 nvmStatusCodeEnumAttr
		nvmErrNoUsableDIMMs                          nvmStatusCodeEnumAttr
		nvmErrPassphraseNotProvided                  nvmStatusCodeEnumAttr
		nvmErrNewPassphraseNotProvided               nvmStatusCodeEnumAttr
		nvmErrPassphrasesDoNotMatch                  nvmStatusCodeEnumAttr
		nvmErrPassphraseTooLong                      nvmStatusCodeEnumAttr
		nvmErrEnableSecurityNotAllowed               nvmStatusCodeEnumAttr
		nvmErrCreateGoalNotAllowed                   nvmStatusCodeEnumAttr
		nvmErrInvalidSecurityState                   nvmStatusCodeEnumAttr
		nvmErrInvalidSecurityOperation               nvmStatusCodeEnumAttr
		nvmErrUnableToGetSecurityState               nvmStatusCodeEnumAttr
		nvmErrInconsistentSecurityState              nvmStatusCodeEnumAttr
		nvmErrInvalidPassphrase                      nvmStatusCodeEnumAttr
		nvmErrSecurityUserPPCountExpired             nvmStatusCodeEnumAttr
		nvmErrRecoveryAccessNotEnabled               nvmStatusCodeEnumAttr
		nvmErrSecureEraseNamespaceExists             nvmStatusCodeEnumAttr
		nvmErrSecurityMasterPPCountExpired           nvmStatusCodeEnumAttr
		nvmErrImageFileNotCompatibleToCTLRStepping   nvmStatusCodeEnumAttr
		nvmErrFilenameNotProvided                    nvmStatusCodeEnumAttr
		nvmSuccessImageExamineOK                     nvmStatusCodeEnumAttr
		nvmErrImageFileNotValid                      nvmStatusCodeEnumAttr
		nvmErrImageExamineLowerVersion               nvmStatusCodeEnumAttr
		nvmErrImageExamineInvalid                    nvmStatusCodeEnumAttr
		nvmErrFirmwareAPINotValid                    nvmStatusCodeEnumAttr
		nvmErrFirmwareVersionNotValid                nvmStatusCodeEnumAttr
		nvmErrFirmwareTooLowForceRequired            nvmStatusCodeEnumAttr
		nvmErrFirmwareAlreadyLoaded                  nvmStatusCodeEnumAttr
		nvmErrFirmwareFailedToStage                  nvmStatusCodeEnumAttr
		nvmErrSensorNotValid                         nvmStatusCodeEnumAttr
		nvmErrSensorMediaTempOutOfRange              nvmStatusCodeEnumAttr
		nvmErrSensorControllerTempOutOfRange         nvmStatusCodeEnumAttr
		nvmErrSensorCapacityOutOfRange               nvmStatusCodeEnumAttr
		nvmErrSensorEnabledStateInvalidValue         nvmStatusCodeEnumAttr
		nvmErrErrorInjectionBIOSKNOBNotEnabled       nvmStatusCodeEnumAttr
		nvmErrMediaDisabled                          nvmStatusCodeEnumAttr
		nvmWarnGoalCreationSecurityUnlocked          nvmStatusCodeEnumAttr
		nvmWarnRegionMaxPMInterleaveSetsExceeded     nvmStatusCodeEnumAttr
		nvmWarnRegionMaxADPMInterleaveSetsExceeded   nvmStatusCodeEnumAttr
		nvmWarnRegionMaxADNIPMInterleaveSetsExceeded nvmStatusCodeEnumAttr
		nvmWarnRegionADNIPMInterleaveSetsReduced     nvmStatusCodeEnumAttr
		nvmErrRegionMaxPMInterleaveSetsExceeded      nvmStatusCodeEnumAttr
		nvmWarn2LMModeOFF                            nvmStatusCodeEnumAttr
		nvmWarnIMCDDRPMMNotPaired                    nvmStatusCodeEnumAttr
		nvmErrPCDBadDeviceConfig                     nvmStatusCodeEnumAttr
		nvmErrRegionGoalConfAffectsUnspecDIMM        nvmStatusCodeEnumAttr
		nvmErrRegionCURRConfAffectsUnspecDIMM        nvmStatusCodeEnumAttr
		nvmErrRegionGoalCURRConfAffectsUnspecDIMM    nvmStatusCodeEnumAttr
		nvmErrRegionConfApplyingFailed               nvmStatusCodeEnumAttr
		nvmErrRegionConfUnsupportedConfig            nvmStatusCodeEnumAttr
		nvmErrRegionNotFound                         nvmStatusCodeEnumAttr
		nvmErrPlatformNotSupportManagementSoft       nvmStatusCodeEnumAttr
		nvmErrPlatformNotSupport2LMMode              nvmStatusCodeEnumAttr
		nvmErrPlatformNotSupportPMMode               nvmStatusCodeEnumAttr
		nvmErrRegionCurrConfExists                   nvmStatusCodeEnumAttr
		nvmErrRegionSizeTooSmallForIntSetAlignment   nvmStatusCodeEnumAttr
		nvmErrPlatformNotSupportSpecifiedIntSizes    nvmStatusCodeEnumAttr
		nvmErrPlatformNotSupportDefaultIntSizes      nvmStatusCodeEnumAttr
		nvmErrRegionNotHealthy                       nvmStatusCodeEnumAttr
		nvmErrRegionNotEnoughSpaceForPMNamespace     nvmStatusCodeEnumAttr
		nvmErrRegionNoGoalExistsOnDIMM               nvmStatusCodeEnumAttr
		nvmErrReserveDIMMRequiresAtLeastTwoDIMMs     nvmStatusCodeEnumAttr
		nvmErrRegionGoalNamespaceExists              nvmStatusCodeEnumAttr
		nvmErrRegionRemainingSizeNotInLastProperty   nvmStatusCodeEnumAttr
		nvmErrPersMemMustBeAppliedToAllDIMMs         nvmStatusCodeEnumAttr
		nvmWarnMappedMemReducedDueToCPUSKU           nvmStatusCodeEnumAttr
		nvmErrRegionGoalAutoProvEnabled              nvmStatusCodeEnumAttr
		nvmErrCreateNamespaceNotAllowed              nvmStatusCodeEnumAttr
		nvmErrOpenFileWithWriteModeFailed            nvmStatusCodeEnumAttr
		nvmErrDumpNoConfiguredDIMMs                  nvmStatusCodeEnumAttr
		nvmErrDumpFileOperationFailed                nvmStatusCodeEnumAttr
		nvmErrLoadVersion                            nvmStatusCodeEnumAttr
		nvmErrLoadInvalidDataInFile                  nvmStatusCodeEnumAttr
		nvmErrLoadImproperConfigInFile               nvmStatusCodeEnumAttr
		nvmErrLoadDIMMCountMismatch                  nvmStatusCodeEnumAttr
		nvmErrDIMMSKUModeMismatch                    nvmStatusCodeEnumAttr
		nvmErrDIMMSKUSecurityMismatch                nvmStatusCodeEnumAttr
		nvmErrNoneDIMMFulfillsCriteria               nvmStatusCodeEnumAttr
		nvmErrUnsupportedBlockSize                   nvmStatusCodeEnumAttr
		nvmErrInvalidNamespaceCapacity               nvmStatusCodeEnumAttr
		nvmErrNotEnoughFreeSpace                     nvmStatusCodeEnumAttr
		nvmErrNamespaceConfigurationBroken           nvmStatusCodeEnumAttr
		nvmErrNamespaceDoesNotExist                  nvmStatusCodeEnumAttr
		nvmErrNamespaceCouldNotUninstall             nvmStatusCodeEnumAttr
		nvmErrNamespaceCouldNotInstall               nvmStatusCodeEnumAttr
		nvmErrNamespaceReadOnly                      nvmStatusCodeEnumAttr
		nvmErrPlatformNotSupportBlockMode            nvmStatusCodeEnumAttr
		nvmWarnBlockModeDisabled                     nvmStatusCodeEnumAttr
		nvmErrNamespaceTooSmallForBTT                nvmStatusCodeEnumAttr
		nvmErrNotEnoughFreeSpaceBTT                  nvmStatusCodeEnumAttr
		nvmErrFailedToUpdateBTT                      nvmStatusCodeEnumAttr
		nvmErrBadalignment                           nvmStatusCodeEnumAttr
		nvmErrRenameNamespaceNotSupported            nvmStatusCodeEnumAttr
		nvmErrFailedToInitNSLabels                   nvmStatusCodeEnumAttr
		nvmErrFWDBGLogFailedToGetSize                nvmStatusCodeEnumAttr
		nvmErrFWDBGSetLogLevelFailed                 nvmStatusCodeEnumAttr
		nvmInfoFWDBGLogNOLogsToFetch                 nvmStatusCodeEnumAttr
		nvmErrFailedToFetchErrorLog                  nvmStatusCodeEnumAttr
		nvmSuccessNoErrorLogEntry                    nvmStatusCodeEnumAttr
		nvmErrSmartFailedToGetSmartInfo              nvmStatusCodeEnumAttr
		nvmWarnSmartNoncriticalHealthIssue           nvmStatusCodeEnumAttr
		nvmErrSmartCriticalHealthIssue               nvmStatusCodeEnumAttr
		nvmErrSmartFatalHealthIssue                  nvmStatusCodeEnumAttr
		nvmErrSmartReadOnlyHealthIssue               nvmStatusCodeEnumAttr
		nvmErrSmartUnknownHealthIssue                nvmStatusCodeEnumAttr
		nvmErrFWSetOptionalDataPolicyFailed          nvmStatusCodeEnumAttr
		nvmErrInvalidOptionalDataPolicyState         nvmStatusCodeEnumAttr
		nvmErrFailedToGetDIMMInfo                    nvmStatusCodeEnumAttr
		nvmErrFailedToGetDIMMRegisters               nvmStatusCodeEnumAttr
		nvmErrSMBIOSDIMMEntryNotFoundInNFIT          nvmStatusCodeEnumAttr
		nvmOperationInProgress                       nvmStatusCodeEnumAttr
		nvmErrGetPCDFailed                           nvmStatusCodeEnumAttr
		nvmErrARSInProgress                          nvmStatusCodeEnumAttr
		nvmErrAPPDirectInSystem                      nvmStatusCodeEnumAttr
		nvmErrOperationNotSupportedByMixedSKU        nvmStatusCodeEnumAttr
		nvmErrFWGetFAUnsupported                     nvmStatusCodeEnumAttr
		nvmErrFWGetFADataFailed                      nvmStatusCodeEnumAttr
		nvmErrAPINotSupported                        nvmStatusCodeEnumAttr
		nvmErrUnknown                                nvmStatusCodeEnumAttr
		nvmErrInvalidPermissions                     nvmStatusCodeEnumAttr
		nvmErrBadDevice                              nvmStatusCodeEnumAttr
		nvmErrBusyDevice                             nvmStatusCodeEnumAttr
		nvmErrGeneralOSDriverFailure                 nvmStatusCodeEnumAttr
		nvmErrNoMem                                  nvmStatusCodeEnumAttr
		nvmErrBadSize                                nvmStatusCodeEnumAttr
		nvmErrTimeout                                nvmStatusCodeEnumAttr
		nvmErrDataTransfer                           nvmStatusCodeEnumAttr
		nvmErrGeneralDevFailure                      nvmStatusCodeEnumAttr
		nvmErrBadFW                                  nvmStatusCodeEnumAttr
		nvmErrDriverFailed                           nvmStatusCodeEnumAttr
		nvmErrDriverfailed                           nvmStatusCodeEnumAttr
		nvmErrInvalidparameter                       nvmStatusCodeEnumAttr
		nvmErrOperationNotSupported                  nvmStatusCodeEnumAttr
		nvmErrRetrySuggested                         nvmStatusCodeEnumAttr
		nvmErrSPDNotAccessible                       nvmStatusCodeEnumAttr
		nvmErrIncompatibleHardwareRevision           nvmStatusCodeEnumAttr
		nvmSuccessNoEventFound                       nvmStatusCodeEnumAttr
		nvmErrFileNotFound                           nvmStatusCodeEnumAttr
		nvmErrOverwriteDIMMInProgress                nvmStatusCodeEnumAttr
		nvmErrFWupdateInProgress                     nvmStatusCodeEnumAttr
		nvmErrUnknownLongOPInProgress                nvmStatusCodeEnumAttr
		nvmErrLongOPAbortedOrRevisionFailure         nvmStatusCodeEnumAttr
		nvmErrFWUpdateAuthFailure                    nvmStatusCodeEnumAttr
		nvmErrUnsupportedCommand                     nvmStatusCodeEnumAttr
		nvmErrDeviceError                            nvmStatusCodeEnumAttr
		nvmErrTransferError                          nvmStatusCodeEnumAttr
		nvmErrUnableToStageNoLongop                  nvmStatusCodeEnumAttr
		nvmErrLongOPUnknown                          nvmStatusCodeEnumAttr
		nvmErrPCDDeleteDenied                        nvmStatusCodeEnumAttr
		nvmErrMixedGenerationsNotSupported           nvmStatusCodeEnumAttr
		nvmErrDimmHealthyFWNotRecoverable            nvmStatusCodeEnumAttr
		nvmLastStatusValue                           nvmStatusCodeEnumAttr
	}
	osTypeEnumAttr enumAttr
	osType         struct {
		osTypeUnknown osTypeEnumAttr
		osTypeWindows osTypeEnumAttr
		osTypeLinux   osTypeEnumAttr
		osTypeEsx     osTypeEnumAttr
	}
	memoryTypeEnumAttr enumAttr
	memoryType         struct {
		memoryTypeUnknown memoryTypeEnumAttr
		memoryTypeDDR4    memoryTypeEnumAttr
		memoryTypeNVMDIMM memoryTypeEnumAttr
		memoryTypeDDR5    memoryTypeEnumAttr
	}
	host struct {
		name         string
		osType       osTypeEnumAttr
		osName       string
		osVersion    string
		mixedSku     nvmBool
		skuViolation nvmBool
		reserved     nvmUint8
	}
	socket struct {
		id                nvmUint16
		mappedMemoryLimit nvmUint64
		totalMappedMemory nvmUint64
		reserved          [64]nvmUint8
	}
	memoryTopology struct {
		physicalID    nvmUint16
		memoryType    memoryTypeEnumAttr
		deviceLocator string
		bankLabel     string
		reserved      [58]nvmUint8
	}
	deviceHandleParts struct {
		memChannelDIMMNum  nvmUint32
		memChannelID       nvmUint32
		memoryControllerID nvmUint32
		socketID           nvmUint32
		nodeControllerID   nvmUint32
		rsvd               nvmUint32
	}
	nvmNfitDeviceHandle []byte
	nvmManufacturer     []nvmUint8
	nvmSerialNumber     []nvmUint8
	nvmVersion          string
	// Structure that describes the security capabilities of a device
	deviceSecurityCapabilities struct {
		passphraseCapable       nvmBool // DCPMM supports the nvm_(set|remove)_passphrase command
		unlockDeviceCapable     nvmBool // DCPMM supports the nvm_unlock_device command
		eraseCryptoCapable      nvmBool // DCPMM supports nvm_erase command with the CRYPTO
		masterPassphraseCapable nvmBool // DCPMM supports set master passphrase command
		reserved                [4]nvmUint8
	}
	// Structure that describes the capabilities supported by a DCPMM
	deviceCapabilities struct {
		packageSparingCapable nvmBool // DCPMM supports package sparing
		memoryModeCapable     nvmBool // DCPMM supports memory mode
		appDirectModeCapable  nvmBool // DCPMM supports app direct mode
		reserved              [5]nvmUint8
	}
	lockStateEnumAttr enumAttr
	lockState         struct {
		lockStateUnknown         lockStateEnumAttr
		lockStateDisable         lockStateEnumAttr
		lockStateUnlocked        lockStateEnumAttr
		lockStateLocked          lockStateEnumAttr
		lockStateFrozen          lockStateEnumAttr
		lockStatePassphraseLimit lockStateEnumAttr
		lockStateNotSupported    lockStateEnumAttr
	}
	manageabilityStateEnumAttr enumAttr
	manageabilityState         struct {
		managementUnknown       manageabilityStateEnumAttr
		managementValidConfig   manageabilityStateEnumAttr
		managementInvalidConfig manageabilityStateEnumAttr
		managementNonFunctional manageabilityStateEnumAttr
	}
	// The device_discovery structure describes an enterprise-level view of
	// a device with enough information to allow callers to uniquely identify
	// a device and determine its status. The UID in this structure is used
	// for all other device management calls to uniquely identify a device.
	// It is intended that this structure will not change over time to allow
	// the native API library to communicate with older and newer revisions of
	// devices
	deviceDiscovery struct {
		// ACPI
		allPropertiesPopulated nvmBool
		deviceHandle           nvmNfitDeviceHandle // The unique device handle of the memory module
		physicalID             nvmUint16           // The unique physical ID of the memory module
		vendorID               nvmUint16           // The vendor identifier - Little Endian
		deviceID               nvmUint16           // The device identifier - Little Endian
		revisionID             nvmUint16           // The revision identifier.
		channelPos             nvmUint16           // The memory module's position in the memory channel
		channelID              nvmUint16           // The memory channel number
		memoryControllerID     nvmUint16           // The ID of the associated memory controller
		socketID               nvmUint16           // The processor socket identifier.
		nodeControllerID       nvmUint16           // The node controller ID.
		// SMBIOS
		memoryType             memoryTypeEnumAttr // The type of memory used by the DCPMM
		dimmSKU                nvmUint32          // Identify Intel DCPMM Gen
		manufacturer           nvmManufacturer    // The manufacturer ID code determined by JEDEC JEP-106 - Little Endian
		serialNumber           nvmSerialNumber    // Serial number assigned by the vendor - Little Endian
		subsystemVendorID      nvmUint16          // vendor identifier of the DCPMM non-volatile memory subsystem controller - Little Endian
		subsystemDeviceID      nvmUint16          // device identifier of the DCPMM non-volatile memory subsystem controller
		subsystemRevisionID    nvmUint16          // revision identifier of the DCPMM non-volatile memory subsystem controller from NFIT
		manufacturingInfoValid nvmBool            // manufacturing location and date validity
		manufacturingLocation  nvmUint8           // DCPMM manufacturing location assigned by vendor only valid if manufacturing_info_valid=1
		manufacturingDate      nvmUint16          // Date the DCPMM was manufactured, assigned by vendor only valid if manufacturing_info_valid=1
		partNumber             string             // The manufacturer's model part number
		fwRevision             nvmVersion         // The current active firmware revision.
		fwAPIVersion           nvmVersion         // API version of the currently running FW
		capacity               nvmUint64          // Raw capacity in bytes.
		interfaceFormatCodes   [9]nvmUint16       // NVM_MAX_IFCS_PER_DIMM
		securityCapabilities   deviceSecurityCapabilities
		deviceCapabilities     deviceCapabilities
		// Calculated by MGMT from NFIT table properties
		uid nvmUID
		// Indicates if the DCPMM is in a locked security state
		lockState lockStateEnumAttr
		// Whether the DCPMM is manageable or not is derived based on what
		// calls are made to populate this struct. If partial properties are
		// requested, then only those properties are used to derive this value.
		// If all properties are requested, then the partial properties plus
		// the firmware API version (requires a DSM call) are used to set this
		// value.
		manageability           manageabilityStateEnumAttr
		controllerRevisionID    nvmUint16
		masterPassphraseEnabled nvmBool
		reserved                [47]nvmUint8
	}
	configStatusEnumAttr enumAttr
	configStatus         struct {
		configStatusNotConfigured       configStatusEnumAttr
		configStatusValid               configStatusEnumAttr
		configStatusErrCorrupt          configStatusEnumAttr
		configStatusErrBrokenInterleave configStatusEnumAttr
		configStatusErrReverted         configStatusEnumAttr
		configStatusErrNotSupported     configStatusEnumAttr
		configStatusUnknown             configStatusEnumAttr
	}
	deviceARSStatusEnumAttr enumAttr
	deviceARSStatus         struct {
		deviceARSStatusUnknown    deviceARSStatusEnumAttr
		deviceARSStatusNotStarted deviceARSStatusEnumAttr
		deviceARSStatusInprogress deviceARSStatusEnumAttr
		deviceARSStatusComplete   deviceARSStatusEnumAttr
		deviceARSStatusAborted    deviceARSStatusEnumAttr
	}
	deviceOverwriteDIMMStatusEnumAttr enumAttr
	deviceOverwriteDIMMStatus         struct {
		deviceOverwriteDIMMStatusUnknown    deviceOverwriteDIMMStatusEnumAttr
		deviceOverwriteDIMMStatusNotstarted deviceOverwriteDIMMStatusEnumAttr
		deviceOverwriteDIMMStatusInprogress deviceOverwriteDIMMStatusEnumAttr
		deviceOverwriteDIMMStatusComplete   deviceOverwriteDIMMStatusEnumAttr
	}
	deviceStatus struct {
		health                             nvmUint8
		isNew                              nvmBool
		isConfigured                       nvmBool
		isMissing                          nvmBool
		packageSparesAvailable             nvmUint8
		lastShutdownStatusDetails          nvmUint32
		configStatus                       configStatusEnumAttr
		lastShutdownTime                   nvmUint64
		mixedSKU                           nvmBool
		skuViolation                       nvmBool
		viralState                         nvmBool
		arsStatus                          deviceARSStatusEnumAttr
		overwritedimmStatus                deviceOverwriteDIMMStatusEnumAttr
		aitDRAMEnabled                     nvmBool
		bootStatus                         nvmUint64
		injectedMediaErrors                nvmUint32
		injectedNonMediaErrors             nvmUint32
		unlachedLastShutdownStatusDetails  nvmUint32
		thermalThrottlePerformanceLossPCNT nvmUint8
		reserved                           [64]nvmUint8
	}
	// enumAttrSharedDefs.h
	pmonRegisters struct {
		// This will specify whether or not to return the extra smart data
		// along with the PMON Counter data.
		smartDataMask uint8
		reserved1     [3]uint8
		// This will specify which group that is currently enabled. If no
		// groups are enabled Group F will be returned.
		groupEnabled uint8
		reserved2    [19]uint8
		pmon4Counter uint
		pmon5Counter uint
		reserved3    [4]uint8
		pmon7Counter uint
		pmon8Counter uint
		pmon9Counter uint
		reserved4    [16]uint8
		reserved5    [4]uint8
		// DDRT Reads for current power cycle
		ddrtrd uint64
		// DDRT Writes for current power cycle
		ddrtwr uint64
		// Media Reads for current power cycle
		merd uint64
		// Media Writes for current power cycle
		mewr uint64
		// Current Media temp
		mtp uint16
		// Current Controller temp
		ctp      uint16
		reserved [20]uint8
	}
	deviceSettings struct {
		viralPolicy nvmBool
		viralStatus nvmBool
		reserved    [6]nvmUint8
	}
	fwUpdateStatusEnumAttr enumAttr
	fwUpdateStatus         struct {
		fwUpdateUnknown fwUpdateStatusEnumAttr
		fwUpdateStaged  fwUpdateStatusEnumAttr
		fwUpdateSuccess fwUpdateStatusEnumAttr
		fwUpdateFailed  fwUpdateStatusEnumAttr
	}
	// Detailed information about firmware image log information of a device.
	deviceFWInfo struct {
		activeFwRevision nvmVersion
		stagedFwRevision nvmVersion
		fwImageMaxSize   nvmUint32
		fwUpdateStatus   fwUpdateStatusEnumAttr
		reserved         [4]nvmUint8
	}
	devicePerformance struct {
		time         timeT
		bytesRead    nvmUint64
		hostReads    nvmUint64
		bytesWritten nvmUint64
		hostWrites   nvmUint64
		blockReads   nvmUint64
		blockWrites  nvmUint64
		reserved     [8]nvmUint8
	}
	sensorTypeEnumAttr enumAttr
	sensorType         struct {
		sensorHealth                     sensorTypeEnumAttr
		sensorMediaTemperature           sensorTypeEnumAttr
		sensorControllerTemperature      sensorTypeEnumAttr
		sensorPercentageRemaining        sensorTypeEnumAttr
		sensorLatchedDirtyShutdownCount  sensorTypeEnumAttr
		sensorPowerontime                sensorTypeEnumAttr
		sensorUptime                     sensorTypeEnumAttr
		sensorPowerCycles                sensorTypeEnumAttr
		sensorFWerrorlogcount            sensorTypeEnumAttr
		sensorUnlachedDirtyShutdownCount sensorTypeEnumAttr
	}
	sensorUnitsEnumAttr enumAttr
	sensorUnits         struct {
		unitCount   sensorUnitsEnumAttr
		unitCelsius sensorUnitsEnumAttr
		unitSeconds sensorUnitsEnumAttr
		unitMinutes sensorUnitsEnumAttr
		unitHours   sensorUnitsEnumAttr
		unitCycles  sensorUnitsEnumAttr
		unitPercent sensorUnitsEnumAttr
	}
	// The current status of a sensor
	sensorStatusEnumAttr enumAttr
	sensorStatus         struct {
		sensorNotInitialized sensorStatusEnumAttr
		sensorNormal         sensorStatusEnumAttr
		sensorNoncritical    sensorStatusEnumAttr
		sensorCritical       sensorStatusEnumAttr
		sensorFatal          sensorStatusEnumAttr
		sensorUnknown        sensorStatusEnumAttr
	}
	sensorSettings struct {
		enabled                   nvmBool
		upperCriticalThreshold    nvmUint64
		lowerCriticalThreshold    nvmUint64
		upperFatalThreshold       nvmUint64
		lowerFatalThreshold       nvmUint64
		upperNoncriticalThreshold nvmUint64
		lowerNoncriticalThreshold nvmUint64
		reserved                  [8]nvmUint8
	}
	sensor struct {
		stype                    sensorTypeEnumAttr   // The type of sensor.
		units                    sensorUnitsEnumAttr  // The units of measurement for the sensor.
		currentState             sensorStatusEnumAttr // The current state of the sensor.
		reading                  nvmUint64            // The current value of the sensor.
		settings                 sensorSettings       // The settings for the sensor.
		lowerCriticalSettable    nvmBool              // If the lower_critical_threshold value is modifiable.
		upperCriticalSettable    nvmBool              // If the upper_critical_threshold value is modifiable
		lowerCriticalSupport     nvmBool              // If the lower_critical_threshold value is supported.
		upperCriticalSupport     nvmBool              // If the upper_critical_threshold value is supported.
		lowerFatalSettable       nvmBool              // If the lower_fatal_threshold value is modifiable.
		upperFatalSettable       nvmBool              // If the upper_fatal_threshold value is modifiable.
		lowerFatalSupport        nvmBool              // If the lower_fatal_threshold value is supported.
		upperFatalSupport        nvmBool              // If the upper_fatal_threshold value is supported.
		lowerNoncriticalSettable nvmBool              // If the lower_noncritical_threshold value is modifiable.
		upperNoncriticalSettable nvmBool              // If the upper_noncritical_threshold value is modifiable.
		lowerNoncriticalSupport  nvmBool              // If the lower_noncritical_threshold value is supported.
		upperNoncriticalSupport  nvmBool              // If the upper_noncritical_threshold value is supported.
		reserved                 [24]nvmUint8
	}
	deviceCapacities struct {
		capacity                  nvmUint64
		memoryCapacity            nvmUint64
		appDirectoryCapacity      nvmUint64
		reserved1                 nvmUint64
		unconfiguredCapacity      nvmUint64
		inaccessibleCapacity      nvmUint64
		reservedCapacity          nvmUint64
		reserved                  [64]nvmUint8
	}
	deviceFromFactorEnumAttr enumAttr
	deviceFromFactor         struct {
		deviceFromFactorUnknown deviceFromFactorEnumAttr
		deviceFromFactorDIMM    deviceFromFactorEnumAttr
		deviceFromFactorSODIMM  deviceFromFactorEnumAttr
	}
	// Detailed information about a device.
	deviceDetails struct {
		discovery   deviceDiscovery
		status      deviceStatus
		fwInfo      deviceFWInfo
		padding     [2]nvmUint8
		performance devicePerformance
		sensors     [11]sensor // NVM_MAX_DEVICE_SENSORS
		capacities  deviceCapacities
		// from SMBIOS Type 17 Table
		formFactor            deviceFromFactorEnumAttr
		dataWidth             nvmUint64
		totalWidth            nvmUint64
		speed                 nvmUint64
		deviceLocator         string
		bankLabel             string
		peakPowerBudget       nvmUint16
		avgPowerBudget        nvmUint16
		packageSparingEnabled nvmBool
		settings              deviceSettings
		reserved              [8]nvmUint8
	}
	nvmFeatures struct {
		getPlatformCapabilities  nvmBool
		getDevices               nvmBool
		getDeviceSMBIOS          nvmBool
		getDeviceHealth          nvmBool
		getDeviceSettings        nvmBool
		modifyDeviceSettings     nvmBool
		getDeviceSecurity        nvmBool
		getDevicePerformance     nvmBool
		getDeviceDeviceFirmware  nvmBool
		updateDeviceFirmware     nvmBool
		getSensors               nvmBool
		modifySensors            nvmBool
		getDeviceCapacity        nvmBool
		modifyDeviceCapacity     nvmBool
		getRegions               nvmBool
		getNamespaces            nvmBool
		getNamespaceDetails      nvmBool
		createNamespace          nvmBool
		enableNamespace          nvmBool
		disableNamespace         nvmBool
		deleteNamespace          nvmBool
		getAddressScrubData      nvmBool
		startAddressScrub        nvmBool
		quickDiagnostic          nvmBool
		platformConfigDiagnostic nvmBool
		pmMetadataDiagnostic     nvmBool
		securityDiagnostic       nvmBool
		fwConsistencyDiagnostic  nvmBool
		memoryMode               nvmBool
		appDirectMode            nvmBool
		errorInjection           nvmBool
		reserved                 [33]nvmUint8
	}
	swCapabilities struct {
		minNamespaceSize                     nvmUint64
		namespaceMemoryPageAllocationCapable nvmBool
		reserved                             [48]nvmUint8
	}
	interleaveSizeEnumAttr enumAttr
	interleaveSize         struct {
		interleaveSizeNone interleaveSizeEnumAttr
		interleaveSize64b  interleaveSizeEnumAttr
		interleaveSize128b interleaveSizeEnumAttr
		interleaveSize256b interleaveSizeEnumAttr
		interleaveSize4kb  interleaveSizeEnumAttr
		interleaveSize1gb  interleaveSizeEnumAttr
	}
	interleaveWaysEnumAttr enumAttr
	interleaveWays         struct {
		interleaveWays0  interleaveWaysEnumAttr
		interleaveWays1  interleaveWaysEnumAttr
		interleaveWays2  interleaveWaysEnumAttr
		interleaveWays3  interleaveWaysEnumAttr
		interleaveWays4  interleaveWaysEnumAttr
		interleaveWays6  interleaveWaysEnumAttr
		interleaveWays8  interleaveWaysEnumAttr
		interleaveWays12 interleaveWaysEnumAttr
		interleaveWays16 interleaveWaysEnumAttr
		interleaveWays24 interleaveWaysEnumAttr
	}
	interleaveFormat struct {
		recommended nvmBool
		channel     interleaveSizeEnumAttr
		imc         interleaveSizeEnumAttr
		ways        interleaveWaysEnumAttr
	}
	memoryCapabilities struct {
		supported               nvmBool
		interleaveAlignmentSize nvmUint16
		interleaveFormatsCount  nvmUint16
		interleaveFormats       [32]interleaveFormat // NVM_INTERLEAVE_FORMATS
		reserved                [56]nvmUint8
	}
	// The volatile memory mode currently selected by the BIOS.
	volatileModeEnumAttr enumAttr
	volatileMode         struct {
		volatileMode1LM     volatileModeEnumAttr
		volatileModeMemory  volatileModeEnumAttr
		volatileModeAuto    volatileModeEnumAttr
		volatileModeUnknown volatileModeEnumAttr
	}
	// The App Direct mode currently selected by the BIOS.
	appDirectModeEnumAttr enumAttr
	appDirectMode         struct {
		appDirectModeDisabled appDirectModeEnumAttr
		appDirectModeEnabled  appDirectModeEnumAttr
		appDirectModeUnknown  appDirectModeEnumAttr
	}
	// Supported features and capabilities BIOS supports
	platformCapabilities struct {
		biosConfigSupport        nvmBool
		biosRuntimeSupport       nvmBool
		memoryMirrorSupported    nvmBool
		memoryMigrationSupported nvmBool
		oneLMMode                memoryCapabilities
		memoryMode               memoryCapabilities
		appDirectMode            memoryCapabilities
		currentVolatileMode      volatileModeEnumAttr
		currentAppDirectMode     appDirectModeEnumAttr
		reserved                 [48]nvmUint8
	}
	skuCapabilities struct {
		mixedSKU     nvmBool
		skuViolation nvmBool
		memorySKU    nvmBool
		appDirectSKU nvmBool
		reserved     [4]nvmUint8
	}
	nvmCapabilities struct {
		nvmFeatures          nvmFeatures
		swCapabilities       swCapabilities
		platformCapabilities platformCapabilities
		dimmSKUCapabilities  skuCapabilities
		reserved             [56]nvmUint8
	}
	eventTypeEnumAttr enumAttr
	eventType         struct {
		eventTypeAll                eventTypeEnumAttr
		eventTypeConfig             eventTypeEnumAttr
		eventTypeHealth             eventTypeEnumAttr
		eventTypeMgmt               eventTypeEnumAttr
		eventTypeDiag               eventTypeEnumAttr
		eventTypeDiagQuick          eventTypeEnumAttr
		eventTypeDiagPlatformConfig eventTypeEnumAttr
		eventTypeDiagSecurity       eventTypeEnumAttr
		eventTypeDiagFWConsistency  eventTypeEnumAttr
	}
	eventSeverityEnumAttr enumAttr
	eventSeverity         struct {
		eventSeverityInfo     eventSeverityEnumAttr
		eventSeverityWarn     eventSeverityEnumAttr
		eventSeverityCritical eventSeverityEnumAttr
		eventSeverityFatal    eventSeverityEnumAttr
	}
	eventFilter struct {
		filterMask nvmUint8
		etype      eventTypeEnumAttr
		severity   eventSeverityEnumAttr
		uid        nvmUID
		eventID    int
		reserved   [21]nvmUint8
	}
	diagnosticResultEnumAttr enumAttr
	diagnosticResult         struct {
		diagnosticResultUnknown diagnosticResultEnumAttr
		diagnosticResultOK      diagnosticResultEnumAttr
		diagnosticResultWarning diagnosticResultEnumAttr
		diagnosticResultFailed  diagnosticResultEnumAttr
		diagnosticResultAborted diagnosticResultEnumAttr
	}
	event struct {
		eventID    nvmUint32
		etype      eventTypeEnumAttr
		severity   eventSeverityEnumAttr
		code       nvmUint16
		reserved1  nvmBool
		uid        nvmUID
		time       timeT
		message    nvmEventMsg
		args       [3]nvmEventArg // NVM_MAX_EVENT_ARGS
		diagResult diagnosticResultEnumAttr
		reserved2  [8]nvmUint8
	}
	regionTypeEnumAttr enumAttr
	regionType         struct {
		regionTypeUnknown          regionTypeEnumAttr
		regionTypePersistent       regionTypeEnumAttr
		regionTypeVolatile         regionTypeEnumAttr
		regionTypePersistentMirror regionTypeEnumAttr
	}
	regionHealthEnumAttr enumAttr
	regionHealth         struct {
		regionHealthNormal  regionHealthEnumAttr
		regionHealthError   regionHealthEnumAttr
		regionHealthUnknown regionHealthEnumAttr
		regionHealthPending regionHealthEnumAttr
		regionHealthLocked  regionHealthEnumAttr
	}
	region struct {
		isetId       nvmUint64
		rtype        regionTypeEnumAttr
		capacity     nvmUint64
		freeCapacity nvmUint64
		socketID     nvmUint16
		dimmCount    nvmUint16
		dimms        [24]nvmUint16 // NVM_MAX_DEVICES_PER_SOCKET
		health       regionHealthEnumAttr
		reserved     [40]nvmUint8
	}
	// Describes the configuration goal for a particular DCPMM.
	configGoalInput struct {
		persistentMemType   nvmUint8
		volatilePercent     nvmUint32
		reservedPercent     nvmUint32
		reserveDIMM         nvmUint32
		namespaceLabelMajor nvmUint16
		namespaceLabelMinor nvmUint16
		reserved            [44]nvmUint8
	}
	interleaveTypeEnumAttr enumAttr
	interleaveType         struct {
		interleaveTypeDefault        interleaveTypeEnumAttr
		interleaveTypeInterleaved    interleaveTypeEnumAttr
		interleaveTypeNotInterleaved interleaveTypeEnumAttr
	}
	configGoalStatusEnumAttr enumAttr
	configGoalStatus         struct {
		configGoalStatusNoGoalOrSuccess          configGoalStatusEnumAttr
		configGoalStatusUnknown                  configGoalStatusEnumAttr
		configGoalStatusNew                      configGoalStatusEnumAttr
		configGoalStatusErrBadrequest            configGoalStatusEnumAttr
		configGoalStatusErrInsufficientresources configGoalStatusEnumAttr
		configGoalStatusErrFW                    configGoalStatusEnumAttr
		configGoalStatusErrUnknown               configGoalStatusEnumAttr
	}
	configGoal struct {
		dimmUID             nvmUID
		socketID            nvmUint16
		persistentRegions   nvmUint32
		volatileSize        nvmUint64
		storageCapacity     nvmUint64
		interleaveSetType   [2]interleaveTypeEnumAttr // MAX_IS_PER_DIMM
		appdirectSize       [2]nvmUint64              // MAX_IS_PER_DIMM
		imcInterleaving     [2]interleaveSizeEnumAttr // MAX_IS_PER_DIMM
		channelInterleaving [2]interleaveSizeEnumAttr // MAX_IS_PER_DIMM
		appdirectIndex      nvmUint8
		status              configGoalStatusEnumAttr
	}
	errorTypeEnumAttr enumAttr
	errorType         struct {
		errorTypePoison          errorTypeEnumAttr
		errorTypeTemperature     errorTypeEnumAttr
		errorTypePackageSparing  errorTypeEnumAttr
		errorTypeSpareCapacity   errorTypeEnumAttr
		errorTypeMediaFatalError errorTypeEnumAttr
		errorTypeDirtyShutdown   errorTypeEnumAttr
	}
	poisonMemoryTypeEnumAttr enumAttr
	poisonMemoryType         struct {
		poisonMemoryTypeMemorymode  poisonMemoryTypeEnumAttr
		poisonMemoryTypeAPPdirect   poisonMemoryTypeEnumAttr
		poisonMemoryTypePatrolscrub poisonMemoryTypeEnumAttr
	}
	deviceError struct {
		etype               errorTypeEnumAttr
		memoryType          poisonMemoryTypeEnumAttr
		dpa                 nvmUint64
		temperature         nvmUint64
		percentageRemaining nvmUint64
		Reserved            [32]nvmUint8
	}
	diagnosticTestEnumAttr enumAttr
	diagnosticTest         struct {
		diagTypeQuick          diagnosticTestEnumAttr
		diagTypePlatformConfig diagnosticTestEnumAttr
		diagTypeSecurity       diagnosticTestEnumAttr
		diagTypeFWConsistency  diagnosticTestEnumAttr
	}
	diagnosticThreshold struct {
		dtype        diagnosticThresholdType
		threshold    nvmUint64
		thresholdStr string
		reserved     [48]nvmUint8
	}
	diagnostic struct {
		test         diagnosticTestEnumAttr
		excludes     nvmUint64
		pOverrides   diagnosticThreshold
		overridesLen nvmUint32
		reserved     [32]nvmUint8
	}
	nvmPreferenceKey     string
	nvmPreferenceValue   string
	nvmJobStatusEnumAttr enumAttr
	nvmJobStatus         struct {
		nvmJobStatusUnknown    nvmJobStatusEnumAttr
		nvmJobStatusNotStarted nvmJobStatusEnumAttr
		nvmJobStatusRunning    nvmJobStatusEnumAttr
		nvmJobStatusComplete   nvmJobStatusEnumAttr
	}
	nvmJobTypeEnumAttr enumAttr
	nvmJobType         struct {
		nvmJobTypeSanitize nvmJobTypeEnumAttr
		nvmJobTypeARS      nvmJobTypeEnumAttr
		nvmJobTypeFWUpdate nvmJobTypeEnumAttr
		nvmJobTypeUnknown  nvmJobTypeEnumAttr
	}
	job struct {
		uid             nvmUID
		percentComplete nvmUint8
		status          nvmJobStatusEnumAttr
		jType           nvmJobType
		result          interface{}
		reserved        [64]nvmUint8
	}
	commandEffectLog struct {
		opcode          nvmUint32
		effects         nvmUint32
	  
	}
	commandAccessPolicy struct {
        opcode                nvmUint8
        sub_opcode            nvmUint8
        restriction           nvmUint8
	}
	devicePTCmd struct {
		opcode                 nvmUint8
		subOpcode              nvmUint8
		inputPayloadSize       nvmUint32
		inputPayload           interface{}
		outputPayloadSize      nvmUint32
		outputPayload          interface{}
		largeInputPayloadSize  nvmUint32
		largeInputPayload      interface{}
		largeOutputPayloadSize nvmUint32
		largeOutputPayload     interface{}
		result                 int
	}
	errorLog struct {
		dimmID          nvmUint16
		systemTimestamp nvmUint64
		errorType       nvmUint8
		outputData      [64]nvmUint8 // MAX_ERROR_LOG_SZ
	}
	fwErrorLogSequenceNumbers struct {
		oldest   nvmUint16
		current  nvmUint16
		reserved [4]nvmUint8
	}
	deviceErrorLogStatus struct {
		thermLow  fwErrorLogSequenceNumbers
		thermHigh fwErrorLogSequenceNumbers
		mediaLow  fwErrorLogSequenceNumbers
		mediaHigh fwErrorLogSequenceNumbers
		reserved  [32]nvmUint8
	}
)

var (
	isLibInitialized       = false
	nvmMaxUIDLen      uint = 22
	nvmStatusCodeEnum      = &nvmStatusCode{
		nvmSuccess:                                   0,
		nvmSuccessFWResetRequired:                    1,
		nvmErrOperationNotStarted:                    2,
		nvmErrOperationFailed:                        3,
		nvmErrForceRequired:                          4,
		nvmErrInvalidParameter:                       5,
		nvmErrCommandNotSupportedByThisSKU:           9,
		nvmErrDIMMNotFound:                           11,
		nvmErrDIMMIDDuplicated:                       12,
		nvmErrSocketIDNotValid:                       13,
		nvmErrSocketIDIncompatiblewDIMMID:            14,
		nvmErrSocketIDDuplicated:                     15,
		nvmErrConfigNotSupportedByCurrentSKU:         16,
		nvmErrManageableDIMMNotFound:                 17,
		nvmErrNoUsableDIMMs:                          18,
		nvmErrPassphraseNotProvided:                  30,
		nvmErrNewPassphraseNotProvided:               31,
		nvmErrPassphrasesDoNotMatch:                  32,
		nvmErrPassphraseTooLong:                      34,
		nvmErrEnableSecurityNotAllowed:               35,
		nvmErrCreateGoalNotAllowed:                   36,
		nvmErrInvalidSecurityState:                   37,
		nvmErrInvalidSecurityOperation:               38,
		nvmErrUnableToGetSecurityState:               39,
		nvmErrInconsistentSecurityState:              40,
		nvmErrInvalidPassphrase:                      41,
		nvmErrSecurityUserPPCountExpired:             42,
		nvmErrRecoveryAccessNotEnabled:               43,
		nvmErrSecureEraseNamespaceExists:             44,
		nvmErrSecurityMasterPPCountExpired:           45,
		nvmErrImageFileNotCompatibleToCTLRStepping:   59,
		nvmErrFilenameNotProvided:                    60,
		nvmSuccessImageExamineOK:                     61,
		nvmErrImageFileNotValid:                      62,
		nvmErrImageExamineLowerVersion:               63,
		nvmErrImageExamineInvalid:                    64,
		nvmErrFirmwareAPINotValid:                    65,
		nvmErrFirmwareVersionNotValid:                66,
		nvmErrFirmwareTooLowForceRequired:            67,
		nvmErrFirmwareAlreadyLoaded:                  68,
		nvmErrFirmwareFailedToStage:                  69,
		nvmErrSensorNotValid:                         70,
		nvmErrSensorMediaTempOutOfRange:              71,
		nvmErrSensorControllerTempOutOfRange:         72,
		nvmErrSensorCapacityOutOfRange:               73,
		nvmErrSensorEnabledStateInvalidValue:         74,
		nvmErrErrorInjectionBIOSKNOBNotEnabled:       75,
		nvmErrMediaDisabled:                          90,
		nvmWarnGoalCreationSecurityUnlocked:          97,
		nvmWarnRegionMaxPMInterleaveSetsExceeded:     98,
		nvmWarnRegionMaxADPMInterleaveSetsExceeded:   99,
		nvmWarnRegionMaxADNIPMInterleaveSetsExceeded: 100,
		nvmWarnRegionADNIPMInterleaveSetsReduced:     101,
		nvmErrRegionMaxPMInterleaveSetsExceeded:      102,
		nvmWarn2LMModeOFF:                            103,
		nvmWarnIMCDDRPMMNotPaired:                    104,
		nvmErrPCDBadDeviceConfig:                     105,
		nvmErrRegionGoalConfAffectsUnspecDIMM:        106,
		nvmErrRegionCURRConfAffectsUnspecDIMM:        107,
		nvmErrRegionGoalCURRConfAffectsUnspecDIMM:    108,
		nvmErrRegionConfApplyingFailed:               109,
		nvmErrRegionConfUnsupportedConfig:            110,
		nvmErrRegionNotFound:                         111,
		nvmErrPlatformNotSupportManagementSoft:       112,
		nvmErrPlatformNotSupport2LMMode:              113,
		nvmErrPlatformNotSupportPMMode:               114,
		nvmErrRegionCurrConfExists:                   115,
		nvmErrRegionSizeTooSmallForIntSetAlignment:   116,
		nvmErrPlatformNotSupportSpecifiedIntSizes:    117,
		nvmErrPlatformNotSupportDefaultIntSizes:      118,
		nvmErrRegionNotHealthy:                       119,
		nvmErrRegionNotEnoughSpaceForPMNamespace:     121,
		nvmErrRegionNoGoalExistsOnDIMM:               122,
		nvmErrReserveDIMMRequiresAtLeastTwoDIMMs:     123,
		nvmErrRegionGoalNamespaceExists:              124,
		nvmErrRegionRemainingSizeNotInLastProperty:   125,
		nvmErrPersMemMustBeAppliedToAllDIMMs:         126,
		nvmWarnMappedMemReducedDueToCPUSKU:           127,
		nvmErrRegionGoalAutoProvEnabled:              128,
		nvmErrCreateNamespaceNotAllowed:              129,
		nvmErrOpenFileWithWriteModeFailed:            130,
		nvmErrDumpNoConfiguredDIMMs:                  131,
		nvmErrDumpFileOperationFailed:                132,
		nvmErrLoadVersion:                            140,
		nvmErrLoadInvalidDataInFile:                  141,
		nvmErrLoadImproperConfigInFile:               142,
		nvmErrLoadDIMMCountMismatch:                  148,
		nvmErrDIMMSKUModeMismatch:                    151,
		nvmErrDIMMSKUSecurityMismatch:                152,
		nvmErrNoneDIMMFulfillsCriteria:               168,
		nvmErrUnsupportedBlockSize:                   171,
		nvmErrInvalidNamespaceCapacity:               174,
		nvmErrNotEnoughFreeSpace:                     175,
		nvmErrNamespaceConfigurationBroken:           176,
		nvmErrNamespaceDoesNotExist:                  177,
		nvmErrNamespaceCouldNotUninstall:             178,
		nvmErrNamespaceCouldNotInstall:               179,
		nvmErrNamespaceReadOnly:                      180,
		nvmErrPlatformNotSupportBlockMode:            181,
		nvmWarnBlockModeDisabled:                     182,
		nvmErrNamespaceTooSmallForBTT:                183,
		nvmErrNotEnoughFreeSpaceBTT:                  184,
		nvmErrFailedToUpdateBTT:                      185,
		nvmErrBadalignment:                           186,
		nvmErrRenameNamespaceNotSupported:            187,
		nvmErrFailedToInitNSLabels:                   188,
		nvmErrFWDBGLogFailedToGetSize:                195,
		nvmErrFWDBGSetLogLevelFailed:                 196,
		nvmInfoFWDBGLogNOLogsToFetch:                 197,
		nvmErrFailedToFetchErrorLog:                  200,
		nvmSuccessNoErrorLogEntry:                    201,
		nvmErrSmartFailedToGetSmartInfo:              220,
		nvmWarnSmartNoncriticalHealthIssue:           221,
		nvmErrSmartCriticalHealthIssue:               222,
		nvmErrSmartFatalHealthIssue:                  223,
		nvmErrSmartReadOnlyHealthIssue:               224,
		nvmErrSmartUnknownHealthIssue:                225,
		nvmErrFWSetOptionalDataPolicyFailed:          230,
		nvmErrInvalidOptionalDataPolicyState:         231,
		nvmErrFailedToGetDIMMInfo:                    235,
		nvmErrFailedToGetDIMMRegisters:               240,
		nvmErrSMBIOSDIMMEntryNotFoundInNFIT:          241,
		nvmOperationInProgress:                       250,
		nvmErrGetPCDFailed:                           260,
		nvmErrARSInProgress:                          261,
		nvmErrAPPDirectInSystem:                      262,
		nvmErrOperationNotSupportedByMixedSKU:        263,
		nvmErrFWGetFAUnsupported:                     264,
		nvmErrFWGetFADataFailed:                      265,
		nvmErrAPINotSupported:                        266,
		nvmErrUnknown:                                267,
		nvmErrInvalidPermissions:                     268,
		nvmErrBadDevice:                              269,
		nvmErrBusyDevice:                             270,
		nvmErrGeneralOSDriverFailure:                 271,
		nvmErrNoMem:                                  272,
		nvmErrBadSize:                                273,
		nvmErrTimeout:                                274,
		nvmErrDataTransfer:                           275,
		nvmErrGeneralDevFailure:                      276,
		nvmErrBadFW:                                  277,
		nvmErrDriverFailed:                           288,
		nvmErrInvalidparameter:                       289,
		nvmErrOperationNotSupported:                  290,
		nvmErrRetrySuggested:                         291,
		nvmErrSPDNotAccessible:                       300,
		nvmErrIncompatibleHardwareRevision:           301,
		nvmSuccessNoEventFound:                       302,
		nvmErrFileNotFound:                           303,
		nvmErrOverwriteDIMMInProgress:                304,
		nvmErrFWupdateInProgress:                     305,
		nvmErrUnknownLongOPInProgress:                306,
		nvmErrLongOPAbortedOrRevisionFailure:         307,
		nvmErrFWUpdateAuthFailure:                    308,
		nvmErrUnsupportedCommand:                     309,
		nvmErrDeviceError:                            310,
		nvmErrTransferError:                          311,
		nvmErrUnableToStageNoLongop:                  312,
		nvmErrLongOPUnknown:                          313,
		nvmErrPCDDeleteDenied:                        314,
		nvmErrMixedGenerationsNotSupported:           315,
		nvmErrDimmHealthyFWNotRecoverable:            316,
	}
	nvmJobStatusEnum = &nvmJobStatus{
		nvmJobStatusUnknown:    0,
		nvmJobStatusNotStarted: 1,
		nvmJobStatusRunning:    2,
		nvmJobStatusComplete:   3,
	}
	nvmJobTypeEnum = &nvmJobType{
		nvmJobTypeSanitize: 0,
		nvmJobTypeARS:      1,
		nvmJobTypeFWUpdate: 3,
		nvmJobTypeUnknown:  4,
	}
	diagnosticTestEnum = &diagnosticTest{
		diagTypeQuick:          0,
		diagTypePlatformConfig: 1,
		diagTypeSecurity:       2,
		diagTypeFWConsistency:  3,
	}
	errorTypeEnum = &errorType{
		errorTypePoison:          1,
		errorTypeTemperature:     2,
		errorTypePackageSparing:  3,
		errorTypeSpareCapacity:   4,
		errorTypeMediaFatalError: 5,
		errorTypeDirtyShutdown:   6,
	}
	poisonMemoryTypeEnum = &poisonMemoryType{
		poisonMemoryTypeMemorymode:  1,
		poisonMemoryTypeAPPdirect:   2,
		poisonMemoryTypePatrolscrub: 4,
	}
	configGoalStatusEnum = &configGoalStatus{
		configGoalStatusNoGoalOrSuccess:          0,
		configGoalStatusUnknown:                  1,
		configGoalStatusNew:                      2,
		configGoalStatusErrBadrequest:            3,
		configGoalStatusErrInsufficientresources: 4,
		configGoalStatusErrFW:                    5,
		configGoalStatusErrUnknown:               6,
	}
	interleaveTypeEnum = &interleaveType{
		interleaveTypeDefault:        0,
		interleaveTypeInterleaved:    1,
		interleaveTypeNotInterleaved: 2,
	}
	regionHealthEnum = &regionHealth{
		regionHealthNormal:  1,
		regionHealthError:   2,
		regionHealthUnknown: 3,
		regionHealthPending: 4,
		regionHealthLocked:  5,
	}
	regionTypeEnum = &regionType{
		regionTypeUnknown:          0,
		regionTypePersistent:       1,
		regionTypeVolatile:         2,
	}
	diagnosticResultEnum = &diagnosticResult{
		diagnosticResultUnknown: 0,
		diagnosticResultOK:      2,
		diagnosticResultWarning: 3,
		diagnosticResultFailed:  5,
		diagnosticResultAborted: 6,
	}
	eventSeverityEnum = &eventSeverity{
		eventSeverityInfo:     2,
		eventSeverityWarn:     3,
		eventSeverityCritical: 6,
		eventSeverityFatal:    7,
	}
	eventTypeEnum = &eventType{
		eventTypeAll:                0,
		eventTypeConfig:             1,
		eventTypeHealth:             2,
		eventTypeMgmt:               3,
		eventTypeDiag:               4,
		eventTypeDiagQuick:          5,
		eventTypeDiagPlatformConfig: 6,
		eventTypeDiagSecurity:       7,
		eventTypeDiagFWConsistency:  8,
	}
	appDirectModeEnum = &appDirectMode{
		appDirectModeDisabled: 0,
		appDirectModeEnabled:  1,
		appDirectModeUnknown:  2,
	}
	volatileModeEnum = &volatileMode{
		volatileMode1LM:     0,
		volatileModeMemory:  1,
		volatileModeAuto:    2,
		volatileModeUnknown: 3,
	}
	interleaveWaysEnum = &interleaveWays{
		interleaveWays0:  0x00,
		interleaveWays1:  0x01,
		interleaveWays2:  0x02,
		interleaveWays3:  0x04,
		interleaveWays4:  0x08,
		interleaveWays6:  0x10,
		interleaveWays8:  0x20,
		interleaveWays12: 0x40,
		interleaveWays16: 0x80,
		interleaveWays24: 0x100,
	}
	interleaveSizeEnum = &interleaveSize{
		interleaveSizeNone: 0x00,
		interleaveSize64b:  0x01,
		interleaveSize128b: 0x02,
		interleaveSize256b: 0x04,
		interleaveSize4kb:  0x40,
		interleaveSize1gb:  0x80,
	}
	deviceFromFactorEnum = &deviceFromFactor{
		deviceFromFactorUnknown: 0,
		deviceFromFactorDIMM:    8,
		deviceFromFactorSODIMM:  12,
	}
	sensorStatusEnum = &sensorStatus{
		sensorNotInitialized: -1,
		sensorNormal:         0,
		sensorNoncritical:    1,
		sensorCritical:       2,
		sensorFatal:          3,
		sensorUnknown:        4,
	}
	sensorUnitsEnum = &sensorUnits{
		unitCount:   1,
		unitCelsius: 2,
		unitSeconds: 21,
		unitMinutes: 22,
		unitHours:   23,
		unitCycles:  39,
		unitPercent: 65,
	}
	sensorTypeEnum = &sensorType{
		sensorHealth:                     0,
		sensorMediaTemperature:           1,
		sensorControllerTemperature:      2,
		sensorPercentageRemaining:        3,
		sensorLatchedDirtyShutdownCount:  4,
		sensorPowerontime:                5,
		sensorUptime:                     6,
		sensorPowerCycles:                7,
		sensorFWerrorlogcount:            8,
		sensorUnlachedDirtyShutdownCount: 9,
	}
	fwUpdateStatusEnum = &fwUpdateStatus{
		fwUpdateUnknown: 0,
		fwUpdateStaged:  1,
		fwUpdateSuccess: 2,
		fwUpdateFailed:  3,
	}
	deviceOverwriteDIMMStatusEnum = &deviceOverwriteDIMMStatus{
		deviceOverwriteDIMMStatusUnknown:    0,
		deviceOverwriteDIMMStatusNotstarted: 1,
		deviceOverwriteDIMMStatusInprogress: 2,
		deviceOverwriteDIMMStatusComplete:   3,
	}
	deviceARSStatusEnum = &deviceARSStatus{
		deviceARSStatusUnknown:    0,
		deviceARSStatusNotStarted: 1,
		deviceARSStatusInprogress: 2,
		deviceARSStatusComplete:   3,
		deviceARSStatusAborted:    4,
	}
	configStatusEnum = &configStatus{
		configStatusNotConfigured:       0,
		configStatusValid:               1,
		configStatusErrCorrupt:          2,
		configStatusErrBrokenInterleave: 3,
		configStatusErrReverted:         4,
		configStatusErrNotSupported:     5,
		configStatusUnknown:             6,
	}
	lockStateEnum = &lockState{
		lockStateUnknown:         0,
		lockStateDisable:         1,
		lockStateUnlocked:        2,
		lockStateLocked:          3,
		lockStateFrozen:          4,
		lockStatePassphraseLimit: 5,
		lockStateNotSupported:    6,
	}
	manageabilityStateEnum = &manageabilityState{
		managementUnknown:       0,
		managementValidConfig:   1,
		managementInvalidConfig: 2,
		managementNonFunctional: 3,
	}
	memoryTypeEnum = &memoryType{
		memoryTypeUnknown: 0,
		memoryTypeDDR4:    1,
		memoryTypeNVMDIMM: 2,
		memoryTypeDDR5:    3,
	}
	osTypeEnum = &osType{
		osTypeUnknown: 0,
		osTypeWindows: 1,
		osTypeLinux:   2,
		osTypeEsx:     3,
	}
)

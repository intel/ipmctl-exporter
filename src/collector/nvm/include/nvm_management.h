/*
 * Copyright (c) 2018, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 */

/**
 * @file nvm_management.h
 * @brief The file describes the entry points of the Native Management API.
 * It is intended to be used by clients of the Native Management API
 * in order to perform management actions.
 *
 * @mainpage Intel(R) Barlow Pass Software Management API
 *
 * @license This project is licensed under the BSD-3-Clause License.
 *
 * @section Introduction
 * The native management API is provided as a convenience for the developers of management utilities.
 * The library serves as an abstraction layer above the underlying driver and operating system.
 * The intent of the abstraction is to simplify the interface, unify the API across operating systems
 * and drivers and to reduce programming errors in the applications utilizing the library.
 *
 * @subsection Compiling
 * The following header files are required to compile applications using the native management library:
 *
 *      - nvm_management.h: The native management API interface definition.
 *      - nvm_types.h: Common types used by the native management API.
 *      - NvmSharedDefs.h: Return code definitions.
 *      - export_api.h: Export definitions for libararies.
 *
 * Be sure to link with the -lipmctl option when compiling.
 *
 * @subsection Versioning
 * The Management Library is versioned in two ways.  First, standard shared library versioning techniques are used so that the OS run-time linkers can combine applications with the appropriate version of the library if possible.  Second, C macros are provided to allow an application to determine and react to different versions of the library in different run-time environments.
 * The version is formatted as MM.mm.hh.bbbb where MM is the 2-digit major version (00-99), mm is the 2-digit minor version (00-99), hh is the 2-digit hot fix number (00-99), and bbbb is the 4-digit build number (0000-9999).
 * The following C macros and interfaces are provided to retrieve the native API version information.
 *
 * @subsection Concurrency
 * The Management Library is not thread-safe.
 *
 * <table>
 * <tr><td>Synopsis</td><td><strong>int nvm_get_major_version</strong>();</td></tr>
 * <tr><td>Description</td><td>Retrieve the native API library major version number (00-99).</td></tr>
 * <tr><td>Arguments</td><td>None</td></tr>
 * <tr><td>Conditions</td><td>No limiting conditions apply to this function.</td></tr>
 * <tr><td>Remarks</td><td>Applications and the native API library are not compatible if they were written against different major versions.&nbsp; For this reason, it is recommended that every application that uses the native API library performs the following check:
 * if (nvm_get_major_version() != NVM_VERSION_MAJOR)
 * // The application cannot continue with this version of the library
 * </td></tr>
 * <tr><td>Returns</td><td>Returns the major version number.</td></tr>
 * </table>
 *
 * <table>
 * <tr><td>Synopsis</td><td><strong>int nvm_get_minor_version</strong>();</td></tr>
 * <tr><td>Description</td><td>Retrieve the native API library minor version number (00-99).</td></tr>
 * <tr><td>Arguments</td><td>None</td></tr>
 * <tr><td>Conditions</td><td>No limiting conditions apply to this function.</td></tr>
 * <tr><td>Remarks</td><td>Unless otherwise stated, every data structure, function, and description described in this document has existed with those exact semantics since version 1.0 of the native API library.  In cases where functions have been added, the appropriate section in this document will describe the version that introduced the new feature.  Applications wishing to check for features that were added may do so by comparing the return value from nvm_get_minor_version() against the minor number in this specification associated with the introduction of the new feature.
 * if (nvm_get_minor_version() != NVM_VERSION_MINOR)
 * // Specific APIs may not be supported
 * </td></tr>
 * <tr><td>Returns</td><td>Returns the minor version number.</td></tr>
 * </table>
 *
 * <table>
 * <tr><td>Synopsis</td><td><strong>int nvm_get_hotfix_number</strong>();</td></tr>
 * <tr><td>Description</td><td>Retrieve the native API library hot fix version number (00-99).</td></tr>
 * <tr><td>Arguments</td><td>None</td></tr>
 * <tr><td>Conditions</td><td>No limiting conditions apply to this function.</td></tr>
 * <tr><td>Remarks</td><td>The hotfix number is used when reporting incidents but has no significance with respect to library compatibility.
 * </td></tr>
 * <tr><td>Returns</td><td>Returns the hot fix version number.</td></tr>
 * </table>
 *
 * <table>
 * <tr><td>Synopsis</td><td><strong>int nvm_get_build_number</strong>();</td></tr>
 * <tr><td>Description</td><td>Retrieve the native API library build version number (0000-9999).</td></tr>
 * <tr><td>Arguments</td><td>None</td></tr>
 * <tr><td>Conditions</td><td>No limiting conditions apply to this function.</td></tr>
 * <tr><td>Remarks</td><td>The build number is used when reporting incidents but has no significance with respect to library compatibility.
 * </td></tr>
 * <tr><td>Returns</td><td>Returns the build version number.</td></tr>
 * </table>
 *
 * @subsection Caller Privileges
 * Unless otherwise specified, all interfaces require the caller to have administrative/root privileges. The library will return NVM_ERR_INVALID_PERMISSIONS if not.
 *
 * @subsection Return Codes
 * Each interface returns a code indicating the status of the operation as defined in ::return_code. Use nvm_get_error to convert the code into a textual description. Specific codes that may be returned by a particular interface are defined in the "Returns" section of each interface.
 *
 * @subsection Microsoft Windows* Notes and Limitations
 * The Windows driver that enables ipmctl communication to Intel's DCPMMs prevents
 * executing commands that change configuration of any DCPMM when there is a related
 * logical disk (namespace) associated with that DCPMM. This is done to protect user
 * data. If a logical disk (namespace) is associated with the target DCPMM, the
 * command will return an error. The logical disk (namespace) must first be deleted
 * before attempting to execute commands that change configuration.
 *
 * Generally, all commands that retrieve status will succeed regardless of logical
 * disk presence.
 */

#ifndef _NVM_MANAGEMENT_H_
#define _NVM_MANAGEMENT_H_

#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>
#include <time.h>

#include "nvm_types.h"
#include "export_api.h"

#ifdef __cplusplus
extern "C"
{
#endif

#define NVM_VERSION_MAJOR   __VERSION_MAJOR__           // Major version number
#define NVM_VERSION_MINOR   __VERSION_MINOR__           // Minor version number
#define NVM_VERSION_HOTFIX   __VERSION_HOTFIX__         // Hot fix version number
#define NVM_VERSION_BUILDNUM    __VERSION_BUILDNUM__    // Build version number

/**
 * Convert a BCD value to the one byte hex value
 */
#define BCD_TO_BYTE(bcd) (bcd > 0x255 ? MAX_UINT8_VALUE : (((bcd & 0xF00) >> 8) * 100) + (((bcd & 0xF0) >> 4) * 10) + (bcd & 0xF))

// the following defines and inline functions should no longer be needed but are
// included for backward compilation compatibility

/**
 * Convert an array of 8 unsigned chars into an unsigned 64 bit value
 * @remarks While it doesn't seem right to be casting 8 bit chars to unsigned long
 * long, this is an issue with gcc - see http:// gcc.gnu.org/bugzilla/show_bug.cgi?id=47821.
 */
#define NVM_8_BYTE_ARRAY_TO_64_BIT_VALUE(arr, val) \
  val = ((unsigned long long)(arr[7] & 0xFF) << 56) + \
        ((unsigned long long)(arr[6] & 0xFF) << 48) + \
        ((unsigned long long)(arr[5] & 0xFF) << 40) + \
        ((unsigned long long)(arr[4] & 0xFF) << 32) + \
        ((unsigned long long)(arr[3] & 0xFF) << 24) + \
        ((unsigned long long)(arr[2] & 0xFF) << 16) + \
        ((unsigned long long)(arr[1] & 0xFF) << 8) + \
        (unsigned long long)(arr[0] & 0xFF);

/**
 * Convert an unsigned 64 bit integer to an array of 8 unsigned chars
 */
#define NVM_64_BIT_VALUE_TO_8_BYTE_ARRAY(val, arr) \
  arr[7] = (unsigned char)((val >> 56) & 0xFF); \
  arr[6] = (unsigned char)((val >> 48) & 0xFF); \
  arr[5] = (unsigned char)((val >> 40) & 0xFF); \
  arr[4] = (unsigned char)((val >> 32) & 0xFF); \
  arr[3] = (unsigned char)((val >> 24) & 0xFF); \
  arr[2] = (unsigned char)((val >> 16) & 0xFF); \
  arr[1] = (unsigned char)((val >> 8) & 0xFF); \
  arr[0] = (unsigned char)(val & 0xFF);

/**
 * ****************************************************************************
 * ENUMS
 * ****************************************************************************
 */

/**
 * The operating system type.
 */
enum os_type {
  OS_TYPE_UNKNOWN = 0,    ///< The OS type can not be determined
  OS_TYPE_WINDOWS = 1,    ///< Windows
  OS_TYPE_LINUX	= 2,    ///< Linux
  OS_TYPE_ESX	= 3     ///< ESX
};

/**
 * Compatibility of the device, FW and configuration with the management software.
 */
enum manageability_state {
  MANAGEMENT_UNKNOWN		= 0,        ///< Device is not recognized or manageability cannot be determined.
  MANAGEMENT_VALIDCONFIG		= 1,    ///< Device is fully manageable.
  MANAGEMENT_INVALIDCONFIG	= 2,    ///< Device is recognized but cannot be managed.
  MANAGEMENT_NON_FUNCTIONAL	= 3     ///< Device is disabled per NFIT
};

/**
 * Security and Sanitize state of the DCPMM.
 */
enum lock_state {
  LOCK_STATE_UNKNOWN		= 0,    ///< Device lock state can not be determined.
  LOCK_STATE_DISABLED		= 1,    ///< Security is not enabled on the device.
  LOCK_STATE_UNLOCKED		= 2,    ///< Security is enabled and unlocked and un-frozen.
  LOCK_STATE_LOCKED		= 3,    ///< Security is enabled and locked and un-frozen.
  LOCK_STATE_FROZEN		= 4,    ///< Security is enabled, unlocked and frozen.
  LOCK_STATE_PASSPHRASE_LIMIT	= 5,    ///< The passphrase limit has been reached, reset required.
  LOCK_STATE_NOT_SUPPORTED	= 6     ////< Security is not supported
};

/**
 * The device type.
 */
enum memory_type {
  MEMORY_TYPE_UNKNOWN	= 0,    ///< The type of DCPMM cannot be determined.
  MEMORY_TYPE_DDR4	= 1,      ///< DDR4.
  MEMORY_TYPE_NVMDIMM	= 2     ///< NGNVM.
};

/**
 * The device format factor.
 */
enum device_form_factor {
  DEVICE_FORM_FACTOR_UNKNOWN	= 0,  ///< The form factor cannot be determined.
  DEVICE_FORM_FACTOR_DIMM		= 8,    ///< DCPMM.
  DEVICE_FORM_FACTOR_SODIMM	= 12,   ///< SODIMM.
};

/**
 * The address range scrub (ARS) operation status for the DCPMM
 */
enum device_ars_status {
  DEVICE_ARS_STATUS_UNKNOWN,      ///< ARS status unknown
  DEVICE_ARS_STATUS_NOTSTARTED,   ///< ARS not started
  DEVICE_ARS_STATUS_INPROGRESS,   ///< ARS in-progress
  DEVICE_ARS_STATUS_COMPLETE,     ///< ARS complete
  DEVICE_ARS_STATUS_ABORTED       ///< ARS aborted
};

/**
 * The overwrite DCPMM operation status for the DCPMM
 */
enum device_overwritedimm_status {
  DEVICE_OVERWRITEDIMM_STATUS_UNKNOWN,      ///< Overwrite DCPMM status unknown
  DEVICE_OVERWRITEDIMM_STATUS_NOTSTARTED,   ///< Overwrite DCPMM not started
  DEVICE_OVERWRITEDIMM_STATUS_INPROGRESS,   ///< Overwrite DCPMM in-progress
  DEVICE_OVERWRITEDIMM_STATUS_COMPLETE      ///< Overwrite DCPMM complete
};

/**
 * The type of sensor.
 * @internal
 * These enums are also used as indexes in the device.sensors array.  It is important to
 * keep them in order and with valid values (0 - 17)
 * @endinternal
 */
enum sensor_type {
  SENSOR_HEALTH = 0,    ///< DCPMM health as reported in the SMART log
  SENSOR_MEDIA_TEMPERATURE = 1,    ///< Device media temperature in degrees Celsius.
  SENSOR_CONTROLLER_TEMPERATURE = 2,    ///< Device media temperature in degrees Celsius.
  SENSOR_PERCENTAGE_REMAINING = 3,    ///< Amount of percentage remaining as a percentage.
  SENSOR_LATCHED_DIRTY_SHUTDOWN_COUNT = 4,    ///< Device shutdowns without notification.
  SENSOR_POWERONTIME = 5,    ///< Total power-on time over the lifetime of the device.
  SENSOR_UPTIME = 6,    ///< Total power-on time since the last power cycle of the device.
  SENSOR_POWERCYCLES = 7,    ///< Number of power cycles over the lifetime of the device.
  SENSOR_FWERRORLOGCOUNT = 8,    ///< The total number of firmware error log entries.
  SENSOR_UNLATCHED_DIRTY_SHUTDOWN_COUNT = 9,    ///!< Number of times that the FW received an unexpected power loss
};

#define SENSOR_COUNT                10

typedef NVM_UINT64 NVM_SENSOR_CATEGORY_BITMASK;

/*
 * The bitmask for sensor type.
 */
enum sensor_category {
  SENSOR_CAT_SMART_HEALTH = 0x1,    ///< SMART and Health
  SENSOR_CAT_POWER = 0x2,           ///< Power related
  SENSOR_CAT_FW_ERROR = 0x4,        ///< FW Error related
  SENSOR_CAT_ALL = SENSOR_CAT_SMART_HEALTH | SENSOR_CAT_POWER | SENSOR_CAT_FW_ERROR ///< All sensor types
};

/**
 * The units of measurement for a sensor.
 */
enum sensor_units {
  UNIT_COUNT = 1,     ///< In numbers of something (0,1,2 ... n).
  UNIT_CELSIUS = 2,   ///< In units of Celsius degrees.
  UNIT_SECONDS = 21,  ///< In seconds of time.
  UNIT_MINUTES = 22,  ///< In minutes of time.
  UNIT_HOURS = 23,    ///< In hours of time.
  UNIT_CYCLES = 39,   ///< Cycles
  UNIT_PERCENT = 65   ///< In units of percentage.
};

/**
 * The current status of a sensor
 */
enum sensor_status {
  SENSOR_NOT_INITIALIZED = -1,    ///< no attempt to read sensor value yet.
  SENSOR_NORMAL = 0,              ///< Current value of the sensor is in the normal range.
  SENSOR_NONCRITICAL = 1,         ///< Current value of the sensor is in non critical range.
  SENSOR_CRITICAL = 2,            ///< Current value of the sensor is in the critical error range.
  SENSOR_FATAL = 3,               ///< Current value of the sensor is in the fatal error range.
  SENSOR_UNKNOWN = 4,             ///< Sensor status cannot be determined.
};

/**
 *      The type of the event that occurred.  Can be used to filter subscriptions.
 */
enum event_type {
  EVENT_TYPE_ALL = 0,                     ///< Subscribe or filter on all event types
  EVENT_TYPE_CONFIG = 1,                  ///< Device configuration status
  EVENT_TYPE_HEALTH = 2,                  ///< Device health event.
  EVENT_TYPE_MGMT = 3,                    ///< Management software generated event.
  EVENT_TYPE_DIAG = 4,                    ///< Subscribe or filter on all diagnostic event types
  EVENT_TYPE_DIAG_QUICK = 5,              ///< Quick diagnostic test event.
  EVENT_TYPE_DIAG_PLATFORM_CONFIG = 6,    ///< Platform config diagnostic test event.
  EVENT_TYPE_DIAG_SECURITY = 7,           ///< Security diagnostic test event.
  EVENT_TYPE_DIAG_FW_CONSISTENCY = 8      ///< FW consistency diagnostic test event.
};

/**
 * Perceived severity of the event
 */
enum event_severity {
  EVENT_SEVERITY_INFO = 2,        ///< Informational event.
  EVENT_SEVERITY_WARN = 3,        ///< Warning or degraded.
  EVENT_SEVERITY_CRITICAL = 6,    ///< Critical.
  EVENT_SEVERITY_FATAL = 7        ///< Fatal or nonrecoverable.
};

enum diagnostic_result {
  DIAGNOSTIC_RESULT_UNKNOWN = 0,      ///< Diagnostic result unknown
  DIAGNOSTIC_RESULT_OK = 2,           ///< Diagnostic result OK
  DIAGNOSTIC_RESULT_WARNING = 3,      ///< Diagnostic result warning
  DIAGNOSTIC_RESULT_FAILED = 5,       ///< Diagnostic result failed
  DIAGNOSTIC_RESULT_ABORTED = 6       ///< Diagnostic result aborted
};

/**
 * Logging level used with the library logging functions.
 */
enum log_level {
  LOG_LEVEL_ERROR = 0,    ///< Error message
  LOG_LEVEL_WARN  = 1,    ///< Warning message
  LOG_LEVEL_INFO  = 2,    ///< Informational message
  LOG_LEVEL_DEBUG = 3     ///< Debug message
};

/**
 * Injected error type - should match the #defines in types.h
 */
enum error_type {
  ERROR_TYPE_POISON             = 1,    ///< Inject a poison error.
  ERROR_TYPE_TEMPERATURE        = 2,    ///< Inject a media temperature error.
  ERROR_TYPE_PACKAGE_SPARING    = 3,    ///< Trigger or revert an artificial package sparing.
  ERROR_TYPE_SPARE_CAPACITY     = 4,    ///< Trigger or clear a percentage remaining threshold alarm.
  ERROR_TYPE_MEDIA_FATAL_ERROR  = 5,    ///< Inject or clear a fake media fatal error.
  ERROR_TYPE_DIRTY_SHUTDOWN     = 6,    ///< Inject or clear a dirty shutdown error.
};

/*
 * Inject a poison error at specific dpa
 */
enum poison_memory_type {
  POISON_MEMORY_TYPE_MEMORYMODE   = 1,    ///< currently allocated in Memory mode
  POISON_MEMORY_TYPE_APPDIRECT    = 2,    ///< currently allocated in AppDirect
  POISON_MEMORY_TYPE_PATROLSCRUB  = 4,    ///< simulating an error found during a patrol scrub operation indifferent to how the memory is currently allocated
};

/**
 * Diagnostic test type
 */
enum diagnostic_test {
  DIAG_TYPE_QUICK           = 0,    ///< verifies manageable DCPMM host mailbox is accessible and basic health
  DIAG_TYPE_PLATFORM_CONFIG = 1,    ///< verifies BIOS config matches installed HW
  DIAG_TYPE_SECURITY        = 2,    ///< verifies all manageable DCPMMs have consistent security state
  DIAG_TYPE_FW_CONSISTENCY  = 3     ///< verifies all DCPMMs have consistent FW and attributes
};

/**
* Health status type
*/
enum health_status {
  HEALTH_STATUS_UNKNOWN             =  0,    ///< Unknown health status
  HEALTH_STATUS_HEALTHY             =  1,    ///< DIMM Healthy
  HEALTH_STATUS_NON_CRITICAL_FAILURE=  2,    ///< Non-Critical (maintenance required)
  HEALTH_STATUS_CRITICAL_FAILURE    =  3,    ///< Critical (feature or performance degraded due to failure)
  HEALTH_STATUS_FATAL_FAILURE       =  4,    ///< Fatal (data loss has occurred or is imminent)
  HEALTH_STATUS_UNMANAGEABLE        =  5,    ///< DIMM is unmanagable
  HEALTH_STATUS_NON_FUNCTIONAL      =  6
};
/**
 * Diagnostic threshold type.
 */
typedef NVM_UINT64 diagnostic_threshold_type;

#define DIAG_THRESHOLD_QUICK_HEALTH                         (1 << 0)
#define DIAG_THRESHOLD_QUICK_MEDIA_TEMP                     (1 << 1)
#define DIAG_THRESHOLD_QUICK_CONTROLLER_TEMP                (1 << 2)
#define DIAG_THRESHOLD_QUICK_AVAIL_SPARE                    (1 << 3)
#define DIAG_THRESHOLD_QUICK_PERC_USED                      (1 << 4)
#define DIAG_THRESHOLD_QUICK_SPARE_DIE                      (1 << 5)
#define DIAG_THRESHOLD_QUICK_UNCORRECT_ERRORS               (1 << 6)
#define DIAG_THRESHOLD_QUICK_CORRECTED_ERRORS               (1 << 7)
#define DIAG_THRESHOLD_QUICK_ERASURE_CODED_CORRECTED_ERRORS (1 << 8)
#define DIAG_THRESHOLD_QUICK_VALID_VENDOR_ID                (1 << 9)
#define DIAG_THRESHOLD_QUICK_VALID_MANUFACTURER             (1 << 10)
#define DIAG_THRESHOLD_QUICK_VALID_PART_NUMBER              (1 << 11)
#define DIAG_THRESHOLD_QUICK_VIRAL                          (1 << 12)
#define DIAG_THRESHOLD_SECURITY_CONSISTENT                  (1 << 13)
#define DIAG_THRESHOLD_SECURITY_ALL_DISABLED                (1 << 14)
#define DIAG_THRESHOLD_SECURITY_ALL_NOTSUPPORTED            (1 << 15)
#define DIAG_THRESHOLD_FW_CONSISTENT                        (1 << 16)
#define DIAG_THRESHOLD_FW_MEDIA_TEMP                        (1 << 17)
#define DIAG_THRESHOLD_FW_CORE_TEMP                         (1 << 18)
#define DIAG_THRESHOLD_FW_SPARE                             (1 << 19)
#define DIAG_THRESHOLD_FW_POW_MGMT_POLICY                   (1 << 20)
#define DIAG_THRESHOLD_FW_PEAK_POW_BUDGET_MIN               (1 << 21)
#define DIAG_THRESHOLD_FW_PEAK_POW_BUDGET_MAX               (1 << 22)
#define DIAG_THRESHOLD_FW_AVG_POW_BUDGET_MIN                (1 << 23)
#define DIAG_THRESHOLD_FW_AVG_POW_BUDGET_MAX                (1 << 24)
#define DIAG_THRESHOLD_FW_DIE_SPARING_POLICY                (1 << 25)
#define DIAG_THRESHOLD_FW_DIE_SPARING_LEVEL                 (1 << 26)
#define DIAG_THRESHOLD_FW_TIME                              (1 << 27)
#define DIAG_THRESHOLD_FW_DEBUGLOG                          (1 << 28)
#define DIAG_THRESHOLD_PCONFIG_NFIT                         (1 << 29)
#define DIAG_THRESHOLD_PCONFIG_PCAT                         (1 << 30)
#define DIAG_THRESHOLD_PCONFIG_PCD                          (1llu << 31)
#define DIAG_THRESHOLD_PCONFIG_CURRENT_PCD                  (1llu << 32)
#define DIAG_THRESHOLD_PCONFIG_UNCONFIGURED                 (1llu << 33)
#define DIAG_THRESHOLD_PCONFIG_BROKEN_ISET                  (1llu << 34)
#define DIAG_THRESHOLD_PCONFIG_MAPPED_CAPACITY              (1llu << 35)
#define DIAG_THRESHOLD_PCONFIG_BEST_PRACTICES               (1llu << 36)

///< The volatile memory mode currently selected by the BIOS.
enum volatile_mode {
  VOLATILE_MODE_1LM       = 0,    ///< 1LM Mode
  VOLATILE_MODE_MEMORY    = 1,    ///< Memory Mode
  VOLATILE_MODE_AUTO      = 2,    ///< Memory Mode if DDR4 + PMM present, 1LM otherwise
  VOLATILE_MODE_UNKNOWN   = 3,    ///< The current volatile memory mode cannot be determined.
};

///< Interface format code as reported by NFIT
enum nvm_format {
  FORMAT_NONE = 0,                  ///< No format indicated
  FORMAT_BLOCK_STANDARD = 0x201,    ///< Block format
  FORMAT_BYTE_STANDARD = 0x301      ///< Byte format
};

///< The App Direct mode currently selected by the BIOS.
enum app_direct_mode {
  APP_DIRECT_MODE_DISABLED    = 0,    ///< App Direct mode disabled.
  APP_DIRECT_MODE_ENABLED     = 1,    ///< App Direct mode enabled.
  APP_DIRECT_MODE_UNKNOWN     = 2,    ///< The current App Direct mode cannot be determined.
};

/**
 * Detailed status of last DCPMM shutdown
 */
enum shutdown_status {
  SHUTDOWN_STATUS_UNKNOWN = 0,                ///< The last shutdown status cannot be determined.
  SHUTDOWN_STATUS_PM_ADR = 1 << 0,            ///< Async DCPMM Refresh command received
  SHUTDOWN_STATUS_PM_S3 = 1 << 1,             ///< PM S3 received
  SHUTDOWN_STATUS_PM_S5 = 1 << 2,             ///< PM S5 received
  SHUTDOWN_STATUS_DDRT_POWER_FAIL = 1 << 3,   ///< DDRT power fail command received
  SHUTDOWN_STATUS_PMIC_POWER_LOSS = 1 << 4,   ///< PMIC Power Loss received
  SHUTDOWN_STATUS_WARM_RESET = 1 << 5,        ///< PM warm reset received
  SHUTDOWN_STATUS_FORCED_THERMAL = 1 << 6,    ///< Thermal shutdown received
  SHUTDOWN_STATUS_CLEAN = 1 << 7              ///< Denotes a proper clean shutdown
};

/**
 * Extended detailed status of last DCPMM shutdown
 */

enum shutdown_status_extended {
  SHUTDOWN_STATUS_VIRAL_INT_RCVD              = 1 << 0,   ///< Virtal interrupt received
  SHUTDOWN_STATUS_SURPRISE_CLK_STOP_INT_RCVD  = 1 << 1,   ///< Surprise clock stop interrupt received
  SHUTDOWN_STATUS_WR_DATA_FLUSH_RCVD          = 1 << 2,   ///< Write Data Flush Complete
  SHUTDOWN_STATUS_S4_PWR_STATE_RCVD           = 1 << 3,   ///< S4 Power State received
  SHUTDOWN_STATUS_PM_IDLE_RCVD                = 1 << 4,   ///< PM Idle Power State received
  SHUTDOWN_STATUS_SURPRISE_RESET_RCVD         = 1 << 5,   ///< Surprise Reset received
};

/**
 * Status of the device current configuration
 */
enum config_status {
  CONFIG_STATUS_NOT_CONFIGURED        = 0,    ///< The device is not configured.
  CONFIG_STATUS_VALID                 = 1,    ///< The device has a valid configuration.
  CONFIG_STATUS_ERR_CORRUPT           = 2,    ///< The device configuration is corrupt.
  CONFIG_STATUS_ERR_BROKEN_INTERLEAVE = 3,    ///< The interleave set is broken.
  CONFIG_STATUS_ERR_REVERTED          = 4,    ///< The configuration failed and was reverted.
  CONFIG_STATUS_ERR_NOT_SUPPORTED     = 5,    ///< The configuration is not supported by the BIOS.
  CONFIG_STATUS_UNKNOWN               = 6,    ///< The configuration status cannot be determined
};

/**
 * Status of current configuration goal
 */
enum config_goal_status {
  CONFIG_GOAL_STATUS_NO_GOAL_OR_SUCCESS		= 0,    ///< The configuration goal status cannot be determined.
  CONFIG_GOAL_STATUS_UNKNOWN			= 1,    ///< The configuration goal has not yet been applied.
  CONFIG_GOAL_STATUS_NEW				= 2,    ///< The configuration goal was applied successfully.
  CONFIG_GOAL_STATUS_ERR_BADREQUEST		= 3,    ///< The configuration goal was invalid.
  CONFIG_GOAL_STATUS_ERR_INSUFFICIENTRESOURCES	= 4,    ///< Not enough resources to apply the goal.
  CONFIG_GOAL_STATUS_ERR_FW			= 5,    ///< Failed to apply the goal due to a FW error.
  CONFIG_GOAL_STATUS_ERR_UNKNOWN			= 6,    ///< Failed to apply the goal for an unknown reason.
};

/**
 *  * Status of NVM jobs
 */
enum nvm_job_status {
  NVM_JOB_STATUS_UNKNOWN      = 0,  ///< Job status unknown
  NVM_JOB_STATUS_NOT_STARTED  = 1,  ///< Job status not started
  NVM_JOB_STATUS_RUNNING      = 2,  ///< Job status in-progress
  NVM_JOB_STATUS_COMPLETE     = 3   ///< Job status complete
};

/**
 * Type of job
 */
enum nvm_job_type {
  NVM_JOB_TYPE_SANITIZE   = 0,  ///< Sanitize
  NVM_JOB_TYPE_ARS        = 1,  ///< ARS
  NVM_JOB_TYPE_FW_UPDATE  = 3,  ///< FW Update
  NVM_JOB_TYPE_UNKNOWN          ///< Unknown
};

/**
 * firmware type
 */
enum device_fw_type {
  DEVICE_FW_TYPE_UNKNOWN      = 0, ///< fw image type cannot be determined
  DEVICE_FW_TYPE_PRODUCTION   = 1, ///< Production image
  DEVICE_FW_TYPE_DFX          = 2, ///< DFX image
  DEVICE_FW_TYPE_DEBUG        = 3  ///< Debug image
};

/**
 * status of last firmware update operation
 */
enum fw_update_status {
  FW_UPDATE_UNKNOWN = 0, ///< status of the last FW update cannot be retrieved
  FW_UPDATE_STAGED  = 1, ///< FW Update Staged
  FW_UPDATE_SUCCESS = 2, ///< FW Update Success
  FW_UPDATE_FAILED  = 3  ///< FW Update Failed
};

/**
 * ****************************************************************************
 * STRUCTURES
 * ****************************************************************************
 */

/**
 * The host server that the native API library is running on.
 */
struct host {
  char		name[NVM_COMPUTERNAME_LEN];     ///<The host computer name.
  enum os_type	os_type;                        ///<OS type.
  char		os_name[NVM_OSNAME_LEN];        ///< OS name string.
  char		os_version[NVM_OSVERSION_LEN];  ///< OS version string.
  NVM_BOOL	mixed_sku;                      ///< One or more DCPMMs have different SKUs.
  NVM_BOOL	sku_violation;                  ///< Configuration of DCPMMs are unsupported due to a license issue.
  NVM_UINT8     reserved[56];                   ///< reserved
};

/**
 * Software versions (one per server).
 */
struct sw_inventory {
  NVM_VERSION	mgmt_sw_revision;               ///< Host software version.
  NVM_VERSION	vendor_driver_revision;         ///< Vendor specific NVDIMM driver version.
  NVM_BOOL	vendor_driver_compatible;       ///< Is vendor driver compatible with MGMT SW?
  NVM_UINT8     reserved[13];                   ///< reserved
};

/**
 * Structure that describes a memory device in the system.
 * This data is harvested from the SMBIOS table Type 17 structures.
 */
struct memory_topology {
  NVM_UINT16		physical_id;                            ///< Memory device's physical identifier (SMBIOS handle)
  enum memory_type	memory_type;                            ///< Type of memory device
  char			device_locator[NVM_DEVICE_LOCATOR_LEN]; ///< Physically-labeled socket of device location
  char			bank_label[NVM_BANK_LABEL_LEN];         ///< Physically-labeled bank of device location
  NVM_UINT8     reserved[58];                                   ///< reserved
};

/**
 * Structure that describes the security capabilities of a device
 */
struct device_security_capabilities {
  NVM_BOOL	passphrase_capable;         ///< DCPMM supports the nvm_(set|remove)_passphrase command
  NVM_BOOL	unlock_device_capable;      ///< DCPMM supports the nvm_unlock_device command
  NVM_BOOL	erase_crypto_capable;       ///< DCPMM supports nvm_erase command with the CRYPTO
  NVM_BOOL      master_passphrase_capable;  ///< DCPMM supports set master passphrase command
  NVM_UINT8     reserved[4];                ///< reserved
};

/**
 * Structure that describes the capabilities supported by a DCPMM
 */
struct device_capabilities {
  NVM_BOOL	package_sparing_capable;        ///< DCPMM supports package sparing
  NVM_BOOL	memory_mode_capable;            ///< DCPMM supports memory mode
  NVM_BOOL	app_direct_mode_capable;        ///< DCPMM supports app direct mode
  NVM_UINT8     reserved[5];                    ///< reserved
};

/**
 * The device_discovery structure describes an enterprise-level view of a device with
 * enough information to allow callers to uniquely identify a device and determine its status.
 * The UID in this structure is used for all other device management calls to uniquely
 * identify a device.  It is intended that this structure will not change over time to
 * allow the native API library to communicate with older and newer revisions of devices.
 * @internal
 * Keep this structure to data from the Identify DCPMM command and calculated data.
 * @endinternal
 */
struct device_discovery {
  // Properties that are fast to access
  ///////////////////////////////////////////////////////////////////////////
  // Indicate whether the struct was populated with the full set of
  // properties (nvm_get_devices()) or just a minimal set (NFIT + SMBIOS)
  // The calls originate at populate_devices() and use the
  // parameter populate_all_properties to distinguish each
  NVM_BOOL		all_properties_populated;

  // ACPI
  NVM_NFIT_DEVICE_HANDLE	device_handle;          ///< The unique device handle of the memory module
  NVM_UINT16		physical_id;            ///< The unique physical ID of the memory module
  NVM_UINT16		vendor_id;              ///< The vendor identifier - Little Endian
  NVM_UINT16		device_id;              ///< The device identifier - Little Endian
  NVM_UINT16		revision_id;            ///< The revision identifier.
  NVM_UINT16		channel_pos;            ///< The memory module's position in the memory channel
  NVM_UINT16		channel_id;             ///< The memory channel number
  NVM_UINT16		memory_controller_id;   ///< The ID of the associated memory controller
  NVM_UINT16		socket_id;              ///< The processor socket identifier.
  NVM_UINT16		node_controller_id;     ///< The node controller ID.

  // SMBIOS
  enum memory_type	memory_type; ///<	The type of memory used by the DCPMM.

  ///////////////////////////////////////////////////////////////////////////



  // Slow (>15ms per passthrough ioctl) properties stored on each DCPMM
  ///////////////////////////////////////////////////////////////////////////
  // Identify Intel DCPMM Gen 1
  // add_identify_dimm_properties_to_device() in device.c
  NVM_UINT32				dimm_sku;
  NVM_MANUFACTURER			manufacturer;                ///< The manufacturer ID code determined by JEDEC JEP-106 - Little Endian
  NVM_SERIAL_NUMBER			serial_number;               ///< Serial number assigned by the vendor - Little Endian
  NVM_UINT16				subsystem_vendor_id;             ///< vendor identifier of the DCPMM non-volatile memory subsystem controller - Little Endian
  NVM_UINT16				subsystem_device_id;            ///< device identifier of the DCPMM non-volatile memory subsystem controller
  NVM_UINT16				subsystem_revision_id;          ///< revision identifier of the DCPMM non-volatile memory subsystem controller from NFIT
  NVM_BOOL				manufacturing_info_valid;       ///< manufacturing location and date validity
  NVM_UINT8				manufacturing_location;         ///< DCPMM manufacturing location assigned by vendor only valid if manufacturing_info_valid=1
  NVM_UINT16				manufacturing_date;             ///< Date the DCPMM was manufactured, assigned by vendor only valid if manufacturing_info_valid=1
  char					part_number[NVM_PART_NUM_LEN];  ///< The manufacturer's model part number
  NVM_VERSION				fw_revision;                    ///< The current active firmware revision.
  NVM_VERSION				fw_api_version;                 ///< API version of the currently running FW
  NVM_UINT64				capacity;                       ///< Raw capacity in bytes.
  NVM_UINT16				interface_format_codes[NVM_MAX_IFCS_PER_DIMM]; ///< calculate_capabilities_for_populated_devices() in device.c
  struct device_security_capabilities	security_capabilities; ///< Security capabilities
  struct device_capabilities		device_capabilities; ///< Capabilities supported by the device

  ///< Calculated by MGMT from NFIT table properties
  NVM_UID					uid; ///< Unique identifier of the device.


  // Get Security State
  // add_security_state_to_device() in device.c
  enum lock_state				lock_state; // Indicates if the DCPMM is in a locked security state
  ///////////////////////////////////////////////////////////////////////////

  // Whether the DCPMM is manageable or not is derived based on what calls are
  // made to populate this struct. If partial properties are requested, then
  // only those properties are used to derive this value. If all properties are
  // requested, then the partial properties plus the firmware API version
  // (requires a DSM call) are used to set this value.
  enum manageability_state manageability;
  NVM_UINT16				controller_revision_id;          ///< revision identifier of the DCPMM non-volatile memory subsystem controller from FIS
  NVM_BOOL				master_passphrase_enabled;	 ///< If 1, master passphrase is enabled on the DCPMM
  NVM_UINT8                             reserved[47];                    ///< reserved
};

struct fw_error_log_sequence_numbers {
  NVM_UINT16	oldest;
  NVM_UINT16	current;
  NVM_UINT8     reserved[4];                    ///< reserved
};

struct device_error_log_status {
  struct fw_error_log_sequence_numbers	therm_low;
  struct fw_error_log_sequence_numbers	therm_high;
  struct fw_error_log_sequence_numbers	media_low;
  struct fw_error_log_sequence_numbers	media_high;
  NVM_UINT8                             reserved[32];   ///< reserved
};

/**
 * The status of a particular device
 */

struct device_status {
  NVM_UINT8			health;                                 ///< Overall device health.
  NVM_BOOL			is_new;                                 ///< Unincorporated with the rest of the devices.
  NVM_BOOL			is_configured;                          ///< only the values 1(Success) and 6 (old config used) from CCUR are considered configured
  NVM_BOOL			is_missing;                             ///< If the device is missing.
  NVM_UINT8			package_spares_available;               ///< Number of package spares on the DCPMM that are available.
  NVM_UINT32		last_shutdown_status_details;           ///< Extended fields as per FIS 1.6 (Latched LSS Details/Extended Details)
  enum config_status		config_status;                  ///< Status of last configuration request.
  NVM_UINT64			last_shutdown_time;                   ///< Time of the last shutdown - seconds since 1 January 1970
  NVM_BOOL			mixed_sku;                              ///< One or more DCPMMs have different SKUs.
  NVM_BOOL			sku_violation;                          ///< The DCPMM configuration is unsupported due to a license issue.
  NVM_BOOL			viral_state;                            ///< Current viral status of DCPMM.
  enum device_ars_status		ars_status;                 ///< Address range scrub operation status for the DCPMM
  enum device_overwritedimm_status	overwritedimm_status;         ///< Overwrite DCPMM operation status for the DCPMM
  NVM_BOOL			ait_dram_enabled;                       ///< Whether or not the AIT DRAM is enabled.
  NVM_UINT64			boot_status;                            ///< The status of the DCPMM as reported by the firmware in the BSR
  NVM_UINT32			injected_media_errors;                  ///< The number of injected media errors on DCPMM
  NVM_UINT32			injected_non_media_errors;              ///< The number of injected non-media errors on DCPMM
  NVM_UINT32    unlatched_last_shutdown_status_details;   ///< Extended fields valid per FIS 1.13+ (Unlatched LSS Details/Extended Details)
  NVM_UINT8     thermal_throttle_performance_loss_pcnt;   ///< the average percentage loss (0..100) due to thermal throttling since last read in current boot (FIS 2.1+)
  NVM_UINT8                             reserved[64];                   ///< reserved
};

/**
 * A snapshot of the performance metrics for a specific device.
 * @remarks All data is cumulative over the life the device.
 */
struct device_performance {
  time_t		time; ///< The time the performance snapshot was gathered.
  // These next fields are 16 bytes in the fw spec, but it would take 100 years
  // of over 31 million reads/writes per second to reach the limit, so we
  // are just using 8 bytes here.
  NVM_UINT64	bytes_read;     ///< Lifetime number of 64 byte reads from media on the DCPMM
  NVM_UINT64	host_reads;     ///< Lifetime number of DDRT read transactions the DCPMM has serviced
  NVM_UINT64	bytes_written;  ///< Lifetime number of 64 byte writes to media on the DCPMM
  NVM_UINT64	host_writes;    ///< Lifetime number of DDRT write transactions the DCPMM has serviced
  NVM_UINT64	block_reads;    ///< Invalid field. "Lifetime number of BW read requests the DCPMM has serviced"
  NVM_UINT64	block_writes;   ///< Invalid field. "Lifetime number of BW write requests the DCPMM has serviced"
  NVM_UINT8     reserved[8];   ///< reserved
};

/**
 * The threshold settings for a particular sensor
 */
struct sensor_settings {
  NVM_BOOL	enabled;                        ///< If firmware notifications are enabled when sensor value is critical.
  NVM_UINT64	upper_critical_threshold;       ///< The upper critical threshold.
  NVM_UINT64	lower_critical_threshold;       ///< The lower critical threshold.
  NVM_UINT64	upper_fatal_threshold;          ///< The upper fatal threshold.
  NVM_UINT64	lower_fatal_threshold;          ///< The lower fatal threshold.
  NVM_UINT64	upper_noncritical_threshold;    ///< The upper noncritical threshold.
  NVM_UINT64	lower_noncritical_threshold;    ///< The lower noncritical threshold.
  NVM_UINT8     reserved[8];                    ///< reserved
};

/**
 * The current state and settings of a particular sensor
 */
struct sensor {
  enum sensor_type	type;                           ///< The type of sensor.
  enum sensor_units	units;                          ///< The units of measurement for the sensor.
  enum sensor_status	current_state;                  ///< The current state of the sensor.
  NVM_UINT64		reading;                        ///< The current value of the sensor.
  struct sensor_settings	settings;                       ///< The settings for the sensor.
  NVM_BOOL		lower_critical_settable;        ///< If the lower_critical_threshold value is modifiable.
  NVM_BOOL		upper_critical_settable;        ///< If the upper_critical_threshold value is modifiable.
  NVM_BOOL		lower_critical_support;         ///< If the lower_critical_threshold value is supported.
  NVM_BOOL		upper_critical_support;         ///< If the upper_critical_threshold value is supported.
  NVM_BOOL		lower_fatal_settable;           ///< If the lower_fatal_threshold value is modifiable.
  NVM_BOOL		upper_fatal_settable;           ///< If the upper_fatal_threshold value is modifiable.
  NVM_BOOL		lower_fatal_support;            ///< If the lower_fatal_threshold value is supported.
  NVM_BOOL		upper_fatal_support;            ///< If the upper_fatal_threshold value is supported.
  NVM_BOOL		lower_noncritical_settable;     ///< If the lower_noncritical_threshold value is modifiable.
  NVM_BOOL		upper_noncritical_settable;     ///< If the upper_noncritical_threshold value is modifiable.
  NVM_BOOL		lower_noncritical_support;      ///< If the lower_noncritical_threshold value is supported.
  NVM_BOOL		upper_noncritical_support;      ///< If the upper_noncritical_threshold value is supported.
  NVM_UINT8             reserved[24];                    ///< reserved
};

/**
 * Device partition capacities (in bytes) used for a single device or aggregated across the server.
 */
struct device_capacities {
  NVM_UINT64  capacity;                       ///< The total DCPMM capacity in bytes.
  NVM_UINT64  memory_capacity;                ///< The total DCPMM capacity in bytes for memory mode.
  NVM_UINT64  app_direct_capacity;            ///< The total DCPMM capacity in bytes for app direct mode.
  NVM_UINT64  mirrored_app_direct_capacity;   ///< The total DCPMM mirrored app direct capacity.
  NVM_UINT64  unconfigured_capacity;          ///< Unconfigured DCPMM capacity. Can be used as storage.
  NVM_UINT64  inaccessible_capacity;          ///< DCPMM capacity that is not acccessible.
  NVM_UINT64  reserved_capacity;              ///< DCPMM app direct capacity reserved and unmapped to SPA.
  NVM_UINT8   reserved[64];                   ///< reserved
};

/**
 * Modifiable settings of a device.
 */
struct device_settings {
  NVM_BOOL  viral_policy;           ///< Viral Policy Enabled/Disabled
  NVM_BOOL  viral_status;           ///< Viral Policy Status
  NVM_UINT8 reserved[6];            ///< reserved
};

/**
 * Detailed information about firmware image log information of a device.
 */
struct device_fw_info {
  /**
   * BCD-formatted revision of the active firmware in the format MM.mm.hh.bbbb
   * MM = 2-digit major version
   * mm = 2-digit minor version
   * hh = 2-digit hot fix version
   * bbbb = 4-digit build version
   */
  NVM_VERSION active_fw_revision;
  NVM_VERSION staged_fw_revision;               ///<  BCD formatted revision of the staged FW.
  NVM_UINT32    FWImageMaxSize;     ///<  The size of FW Image in bytes.
  enum fw_update_status fw_update_status;       ///< status of last FW update operation.
  NVM_UINT8 reserved[4];            ///< reserved
};

/**
 * Detailed information about a device.
 */
struct device_details {
  struct device_discovery     discovery;                                ///< Basic device identifying information.
  struct device_status		status;                                 ///< Device health and status.
  struct device_fw_info       fw_info;                                  ///< The firmware image information for the PMem DCPMM.
  NVM_UINT8			padding[2];                             ///< struct alignment
  struct device_performance	performance;                            ///< A snapshot of the performance metrics.
  struct sensor			sensors[NVM_MAX_DEVICE_SENSORS];        ///< Device sensors.
  struct device_capacities	capacities;                             ///< Partition information

  // from SMBIOS Type 17 Table
  enum device_form_factor		form_factor;                            ///< The type of DCPMM.
  NVM_UINT64                  data_width;                               ///< The width in bits used to store user data.
  NVM_UINT64                  total_width;                              ///< The width in bits for data and ECC and/or redundancy.
  NVM_UINT64			speed;                                  ///< The speed in nanoseconds.
  char				device_locator[NVM_DEVICE_LOCATOR_LEN]; ///< The socket or board position label
  char				bank_label[NVM_BANK_LABEL_LEN];         ///< The bank label

  NVM_UINT16			peak_power_budget;                      ///< instantaneous power budget in mW (100-20000 mW).
  NVM_UINT16			avg_power_budget;                       ///< average power budget in mW (100-18000 mW).
  NVM_BOOL			package_sparing_enabled;                    ///< Enable or disable package sparing.
  struct device_settings		settings;                               ///< Modifiable features of the device.
  NVM_UINT8			reserved[8];				///< reserved
};

/**
 * Supported capabilities of a specific memory mode
 */
struct memory_capabilities {
  NVM_BOOL			supported;                                      ///< is the memory mode supported by the BIOS
  NVM_UINT16			interleave_alignment_size;                      ///< interleave alignment size in 2^n bytes.
  NVM_UINT16			interleave_formats_count;                       ///< Number of interleave formats supported by BIOS
  struct interleave_format	interleave_formats[NVM_INTERLEAVE_FORMATS];     ///< interleave formats
  NVM_UINT8			reserved[56];					///< reserved
};

/**
 * Supported features and capabilities BIOS supports
 */
struct platform_capabilities {
  NVM_BOOL			bios_config_support;            ///< available BIOS support for DCPMM config changes
  NVM_BOOL			bios_runtime_support;           ///< runtime interface used to validate management configuration
  NVM_BOOL			memory_mirror_supported;        ///< indicates if DCPMM mirror is supported
  NVM_BOOL			memory_spare_supported;         ///< pm spare is supported
  NVM_BOOL			memory_migration_supported;     ///< pm memory migration is supported
  struct memory_capabilities	one_lm_mode;                    ///< capabilities for 1LM mode
  struct memory_capabilities	memory_mode;                    ///< capabilities for Memory mode
  struct memory_capabilities	app_direct_mode;                ///< capabilities for App Direct mode
  enum volatile_mode		current_volatile_mode;          ///< The volatile memory mode selected by the BIOS.
  enum app_direct_mode		current_app_direct_mode;        ///< The App Direct mode selected by the BIOS.
  NVM_UINT8			reserved[48];			///< reserved
};

/**
 * DCPMM software-supported features
 */
struct nvm_features {
  NVM_BOOL	get_platform_capabilities;      ///< get platform supported capabilities
  NVM_BOOL	get_devices;                    ///< retrieve the list of DCPMMs installed on the server
  NVM_BOOL	get_device_smbios;              ///< retrieve the SMBIOS information for DCPMMs
  NVM_BOOL	get_device_health;              ///< retrieve the health status for DCPMMs
  NVM_BOOL	get_device_settings;            ///< retrieve DCPMM settings
  NVM_BOOL	modify_device_settings;         ///< modify DCPMM settings
  NVM_BOOL	get_device_security;            ///< retrieve DCPMM security state
  NVM_BOOL	modify_device_security;         ///< modify DCPMM security settings
  NVM_BOOL	get_device_performance;         ///< retrieve DCPMM performance metrics
  NVM_BOOL	get_device_firmware;            ///< retrieve DCPMM firmware version
  NVM_BOOL	update_device_firmware;         ///< update the firmware version on DCPMMs
  NVM_BOOL	get_sensors;                    ///< get health sensors on DCPMMs
  NVM_BOOL	modify_sensors;                 ///< modify the DCPMM health sensor settings
  NVM_BOOL	get_device_capacity;            ///< retrieve how DCPMM capacity is mapped by BIOS
  NVM_BOOL	modify_device_capacity;         ///< modify how the DCPMM capacity is provisioned
  NVM_BOOL	get_regions;                    ///< retrieve regions of DCPMM capacity
  NVM_BOOL	get_namespaces;                 ///< retrieve the list of namespaces allocated from regions
  NVM_BOOL	get_namespace_details;          ///< retrieve detailed info about each namespace
  NVM_BOOL	create_namespace;               ///< create a new namespace
  NVM_BOOL	enable_namespace;               ///< enable a namespace
  NVM_BOOL	disable_namespace;              ///< disable a namespace
  NVM_BOOL	delete_namespace;               ///< delete a namespace
  NVM_BOOL	get_address_scrub_data;         ///< retrieve address range scrub data
  NVM_BOOL	start_address_scrub;            ///< initiate an address range scrub
  NVM_BOOL	quick_diagnostic;               ///< quick health diagnostic
  NVM_BOOL	platform_config_diagnostic;     ///< platform configuration diagnostic
  NVM_BOOL	pm_metadata_diagnostic;         ///< persistent memory metadata diagnostic
  NVM_BOOL	security_diagnostic;            ///< security diagnostic
  NVM_BOOL	fw_consistency_diagnostic;      ///< firmware consistency diagnostic
  NVM_BOOL	memory_mode;                    ///< access DCPMM capacity as memory
  NVM_BOOL	app_direct_mode;                ///< access DCPMM persistent memory in App Direct Mode
  NVM_BOOL	error_injection;                ///< error injection on DCPMMs
  NVM_UINT8	reserved[32];			///< reserved
};

/**
 * Supported features and capabilities the driver/software supports
 */
struct sw_capabilities {
  NVM_UINT64	min_namespace_size; ///< smallest namespace supported by the driver, in bytes
  NVM_BOOL	namespace_memory_page_allocation_capable; ///< namespace memory page allocation capable
  NVM_UINT8	reserved[48];			///< reserved
};

/**
 * Aggregation of DCPMM SKU capabilities across all manageable DCPMMs in the system.
 */
struct dimm_sku_capabilities {
  NVM_BOOL	mixed_sku;      ///< One or more DCPMMs have different SKUs.
  NVM_BOOL	sku_violation;  ///< One or more DCPMMs are in violation of their SKU.
  NVM_BOOL	memory_sku;     ///< One or more DCPMMs support memory mode.
  NVM_BOOL	app_direct_sku; ///< One or more DCPMMs support app direct mode.
  NVM_UINT8	reserved[4];	///< reserved
};

/**
 * Combined DCPMM capabilities
 */
struct nvm_capabilities {
  struct nvm_features		nvm_features;           ///< supported features of the PMM software
  struct sw_capabilities	sw_capabilities;        ///< driver supported capabilities
  struct platform_capabilities	platform_capabilities;  ///< platform-supported capabilities
  struct dimm_sku_capabilities	sku_capabilities;       ///< aggregated DCPMM SKU capabilities
  NVM_UINT8			reserved[56];		///< reserved
};

/*
 * Interleave set information
 */
struct interleave_set {
  NVM_UINT32			set_index;      ///< unique identifier from the PCD
  NVM_UINT32			driver_id;      ///< unique identifier from the driver
  NVM_UINT64			size;           ///< size in bytes
  NVM_UINT64			available_size; ///< free size in bytes
  struct interleave_format	settings; ///< interleave format settings
  NVM_UINT8			socket_id;        ///< socket ID
  NVM_UINT8			dimm_count;       ///< number of dimms in member dimms
  NVM_UID				dimms[NVM_MAX_DEVICES_PER_SOCKET]; ///< UID of DCPMM
  NVM_BOOL			mirrored;         ///< Is mirrored
  enum interleave_set_health	health; ///< health status
  enum encryption_status		encryption;  ///< on if lockstates of all dimms is enabled
  NVM_BOOL			erase_capable;          ///< true if all dimms in the set support erase
  NVM_UINT8			reserved[56];		///< reserved
};

/**
 * Information about a persistent memory region
 */
struct region {
  NVM_UINT64 isetId;       ///< Unique identifier of the region.
  enum region_type		type;           ///< The type of region.
  NVM_UINT64		capacity;       ///< Size of the region in bytes.
  NVM_UINT64		free_capacity;  ///< Available size of the region in bytes.
  NVM_INT16		socket_id;        ///< socket ID
  NVM_UINT16		dimm_count;     ///< The number of dimms in this region.
  NVM_UINT16		dimms[NVM_MAX_DEVICES_PER_SOCKET]; ///< Unique ID's of underlying DCPMMs.
  enum region_health	health; ///< Rolled up health of the underlying DCPMMs.
  NVM_UINT8		reserved[40];		///< reserved
};

/**
 * Describes the configuration goal for a particular DCPMM.
 */
struct config_goal_input {
  NVM_UINT8	persistent_mem_type;      ///< Persistent memory type: 0x1 - AppDirect, 0x2 - AppDirect Non-Interleaved
  NVM_UINT32	volatile_percent;       ///< Volatile region size in percents
  NVM_UINT32	reserved_percent;       ///< Amount of AppDirect memory to not map in percents
  NVM_UINT32	reserve_dimm;           ///< Reserve one DCPMM for use as not interleaved AppDirect memory: 0x0 - RESERVE_DIMM_NONE, 0x1 - STORAGE (NOT SUPPORTED), 0x2 - RESERVE_DIMM_AD_NOT_INTERLEAVED
  NVM_UINT16	namespace_label_major;  ///< Major version of label to init: 0x1 (only supported major version)
  NVM_UINT16	namespace_label_minor;  ///< Minor version of label to init: 0x1 or 0x2 (only supported minor versions)
  NVM_UINT8	reserved[44];		///< reserved
};

struct config_goal {
  NVM_UID			dimm_uid;                                        ///< DCPMM UID
  NVM_UINT16		socket_id;                                     ///< Socket ID
  NVM_UINT32		persistent_regions;                            ///< count of persistent regions
  NVM_UINT64		volatile_size;                                 ///< Gibibytes of memory mode capacity on the DCPMM.
  NVM_UINT64		storage_capacity;                              ///< Gibibytes of storage capacity on the DCPMM.
  enum interleave_type	interleave_set_type[MAX_IS_PER_DIMM];  ///< type of interleave set
  NVM_UINT64		appdirect_size[MAX_IS_PER_DIMM];               ///< appdirect size
  enum interleave_size	imc_interleaving[MAX_IS_PER_DIMM];     ///< IMC interleaving
  enum interleave_size	channel_interleaving[MAX_IS_PER_DIMM]; ///< Channel interleaving
  NVM_UINT8		appdirect_index[MAX_IS_PER_DIMM];                ///< appdirect Index
  enum config_goal_status status;                              ///< Status for the config goal. Ignored for input.
  NVM_UINT8		reserved[32];				///< reserved
};

/*
 * The details of a specific device event that can be subscribed to
 * using #nvm_add_event_notify.
 */
struct event {
  NVM_UINT32		event_id;                       ///< Unique ID of the event.
  enum event_type		type;                           ///< The type of the event that occurred.
  enum event_severity	severity;                       ///< The severity of the event.
  NVM_UINT16		code;                           ///< A numerical code for the specific event that occurred.
  NVM_BOOL		Reserved;                ///< Reserved for future use
  NVM_UID			uid;                            ///< The unique ID of the item that had the event.
  time_t			time;                           ///< The time the event occurred.
  NVM_EVENT_MSG		message;                        ///< A detailed description of the event type that occurred in English.
  NVM_EVENT_ARG		args[NVM_MAX_EVENT_ARGS];       ///< The message arguments.
  enum diagnostic_result	diag_result;                    ///< The diagnostic completion state (only for diag events).
  NVM_UINT8		reserved[8];				///< reserved
};

/**
 * Limits the events returned by the #nvm_get_events method to
 * those that meet the conditions specified.
 */
struct event_filter {
  /**
   * A bit mask specifying the values in this structure used to limit the results.
   * Any combination of the following or 0 to return all events.
   * NVM_FILTER_ON_TYPE
   * NVM_FITLER_ON_SEVERITY
   * NVM_FILTER_ON_CODE
   * NVM_FILTER_ON_UID
   * NVM_FILTER_ON_AFTER
   * NVM_FILTER_ON_BEFORE
   * NVM_FILTER_ON_EVENT
   */
  NVM_UINT8		filter_mask;

  /**
   * The type of events to retrieve. Only used if
   * NVM_FILTER_ON_TYPE is set in the #filter_mask.
   */
  enum event_type		type;

  /**
   * The type of events to retrieve. Only used if
   * NVM_FILTER_ON_SEVERITY is set in the #filter_mask.
   */
  enum event_severity	severity;

  /**
   * The identifier to retrieve events for.
   * Only used if NVM_FILTER_ON_UID is set in the #filter_mask.
   */
  NVM_UID			uid; ///< filter on specific item

  /**
   * Event ID number (row ID)
   * Only used if NVM_FILTER_ON_EVENT is set in the #filter mask.
   */
  int			event_id; ///< filter of specified event

  NVM_UINT8		reserved[21];	///< reserved
};

/**
 * An entry in the native API trace log.
 */
struct nvm_log {
  char		message[NVM_LOG_MESSAGE_LEN];   ///< The log message
  NVM_UINT8		reserved[64];	///< reserved
};

/**
 * An injected device error.
 */
struct device_error {
  enum error_type		type;           ///< The type of error to inject.
  enum poison_memory_type memory_type;    ///< Poison type
  NVM_UINT64		dpa;            ///< Inject poison address - only valid if injecting poison error
  NVM_UINT64		temperature;    ///< Inject temperature - only valid if injecting temperature error
  NVM_UINT64		percentageRemaining;  ///< only valid if injecting percentage remaining error
  NVM_UINT8		reserved[32];	///< reserved
};

/**
 * A structure to hold a diagnostic threshold.
 * Primarily for allowing caller to override default thresholds.
 */
struct diagnostic_threshold {
  diagnostic_threshold_type	type;                                   ///< A diagnostic threshold indicator
  NVM_UINT64			threshold;                              ///< numeric threshold
  char				threshold_str[NVM_THRESHOLD_STR_LEN];   ///< text value used as a "threshold"
  NVM_UINT8			reserved[48];	///< reserved
};

/**
 * A diagnostic test.
 */
struct diagnostic {
  enum diagnostic_test		test;           ///< The type of diagnostic test to run
  NVM_UINT64			excludes;       ///< Bitmask - zero or more diagnostic_threshold_type enums
  struct diagnostic_threshold *	p_overrides;    ///< override default thresholds that trigger failure
  NVM_UINT32			overrides_len;  ///< size of p_overrides array
  NVM_UINT8			reserved[32];	///< reserved
};

/**
 * Describes the identity of a system's physical processor in a NUMA context.
 */
struct socket {
  NVM_UINT16	id;                                             ///< Zero-indexed NUMA node number
  NVM_UINT64	mapped_memory_limit;                            ///< Maximum allowed memory (via PCAT)
  NVM_UINT64	total_mapped_memory;                            ///< Current occupied memory (via PCAT)
  NVM_UINT8	reserved[64];					///< reserved
};

/** Describes the status of a job */
struct job {
  NVM_UID			uid;                ///< UID of the DCPMM
  NVM_UINT8		percent_complete;   ///< Percent complete
  enum nvm_job_status	status;     ///< Job status
  enum nvm_job_type	type;         ///< Job type
  NVM_UID			affected_element;   ///< Affected element
  void *			result;             ///< Result
  NVM_UINT8		reserved[64];		///< reserved
};

#define TEMP_POSITIVE           0
#define TEMP_NEGATIVE           1
#define TEMP_USER_ALARM         0
#define TEMP_LOW                        1
#define TEMP_HIGH                       2
#define TEMP_CRIT                       4
#define TEMP_TYPE_MEDIA         0
#define TEMP_TYPE_CORE          1
/*
 * ****************************************************************************
 * ENTRY POINT METHODS
 * ****************************************************************************
 */

/**
* @brief  Initialize the library.
* @return
*  ::NVM_SUCCESS @n
*/
NVM_API int nvm_init();

/**
 * @brief  Clean up the library.
 */
NVM_API void nvm_uninit();

/**
* @brief    Initialize the config file. Only the first call to the
* function changes the conf file configuration, the following
* function calls have no effect and the conf file configuration
* remains unchanged up to next application execution.
*
* @param    p_ini_file_name Pointer to the name of the ini file to read
* @return  void
*/
NVM_API void nvm_conf_file_init(const char *p_ini_file_name);

/**
* @brief    Flush the config structre to the config file, the previous config
* file content is being overwritten
*
* @return  void
*/
NVM_API void nvm_conf_file_flush();

/*
 * system.c
 */

/**
* @brief Convert DCPMM UID to DCPMM ID and/or DCPMM Handle
*
* @param[in] device_uid UID of the DCPMM
* @param[out] dimm_id optional. pointer to get DCPMM ID.
* @param[out] dimm_handle optional. pointer to get DCPMM Handle.
*
* @return
* ::NVM_SUCCESS @n
* ::NVM_ERR_UNKNOWN @n
*/
NVM_API int nvm_get_dimm_id(const NVM_UID device_uid, unsigned int *dimm_id, unsigned int *dimm_handle);

/**
* @brief Get configuration parameter as integer. If not found, default_val will
* be returned.
*
* @param[in] param_name name of configuration parameter
* @param[in] default_val value to be returned if param_name is not found
*
* @returnint value found in configuration or default_val if not found.
*/
NVM_API int nvm_get_config_int(const char *param_name, int default_val);
/**
 * @brief  Retrieve just the host server name that the native API is running on.
 * @param[in, out] host_name
 *              A caller supplied buffer to hold the host server name
 * @param[in] host_name_len
 *              The length of the host_name buffer. Should be = NVM_COMPUTERNAME_LEN.
 * @return
 *            ::NVM_SUCCESS @n
 *  ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_host_name(char *host_name, const NVM_SIZE host_name_len);

/**
 * @brief Retrieve basic information about the host server the native API library is running on.
 * @param[in,out] p_host
 *              A pointer to a #host structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_host(struct host *p_host);

/**
 * @brief Retrieve a list of installed software versions related to DCPMM management.
 * @param[in,out] p_inventory
 *              A pointer to a #sw_inventory structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @remarks If a version cannot be retrieved, the version is returned as all zeros.
 * @remarks DCPMM firmware revisions are not included.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_sw_inventory(struct sw_inventory *p_inventory);

/**
 * @brief Retrieves the number of physical processors (NUMA nodes) in the system.
 * @pre
 *              The OS must support its respective NUMA implementation.
 * @remarks
 *              This method should be called before #nvm_get_socket or #nvm_get_sockets
 * @remarks
 *              This method should never return a value less than 1.
 * @param[in,out] count
 *              A pointer to an integer which contain the number of sockets on return.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_number_of_sockets(int *count);
/**
 * @brief Retrieves #socket information about each processor socket in the system.
 *
 * @param[in,out] p_sockets
 *              An array of #socket structures allocated by the caller.
 * @param[in] count
 *              The number of elements in the array.
 * @remarks To allocate the array of #socket structures,
 * call #nvm_get_number_of_sockets before calling this method.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_BAD_SIZE @n
 */
NVM_API int nvm_get_sockets(struct socket *p_sockets, const NVM_UINT16 count);

/**
 * @brief Retrieves #socket information about a given processor socket.
 * @pre
 *              The OS must support its respective NUMA implementation.
 * @param[in] socket_id
 *              The NUMA node identifier
 * @param[in,out] p_socket
 *              A pointer to a #socket structure allocated by the caller.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_socket(const NVM_UINT16 socket_id, struct socket *p_socket);

/**
* @brief Retrieve the number of memory devices installed in the system. This count includes
* both DCPMMs and other memory devices, such as DRAM.
* @pre The caller must have administrative privileges.
* @remarks This method should be called before #nvm_get_memory_topology.
* @param[out] count pointer to number of memory devices
* @return
*       ::NVM_SUCCESS @n
*       ::NVM_ERR_INVALID_PARAMETER @n
*       ::NVM_ERR_UNKNOWN @n
*/
NVM_API int nvm_get_number_of_memory_topology_devices(unsigned int *count);

/**
 * @brief Retrieves basic topology information about all memory devices installed in the
 * system, including both PMMs and other memory devices, such as DRAM.
 * @pre The caller must have administrative privileges.
 * @param[out] p_devices pointer to #memory_topology array of size count
 * @param[in] count number of elements in p_devices array
 * @remarks To allocate the array of #memory_topology structures,
 * call #nvm_get_number_of_memory_topology_devices before calling this method.
 * @return
 *              ::NVM_SUCCESS @n
 *              ::NVM_ERR_INVALID_PARAMETER @n
 *              ::NVM_ERR_UNKNOWN @n
 *              ::NVM_ERR_BAD_SIZE @n
 */
NVM_API int nvm_get_memory_topology(struct memory_topology *p_devices, const NVM_UINT8 count);

/*
* @brief Retrieves the number of devices installed in the system whether they are
* fully compatible with the current native API library version or not.
* @pre The caller must have administrative privileges.
* @remarks This method should be called before #nvm_get_devices.
* @remarks The number of devices can be 0.
* @param[out] count pointer to count of devices
* @return
*              ::NVM_SUCCESS @n
*              ::NVM_ERR_INVALID_PARAMETER @n
*              ::NVM_ERR_UNKNOWN @n
*/
NVM_API int nvm_get_number_of_devices(unsigned int *count);

/**
 * @brief Retrieves #device_discovery information
 * about each device in the system whether they are fully compatible
 * with the current native API library version or not.
 * @param[in,out] p_devices
 *              An array of #device_discovery structures allocated by the caller.
 * @param[in] count
 *              The number of elements in array.
 * @pre The caller must have administrative privileges.
 * @remarks To allocate the array of #device_discovery structures,
 * call #nvm_get_device_count before calling this method.
 * @return
 *              ::NVM_SUCCESS @n
 *              ::NVM_ERR_INVALID_PARAMETER @n
 *              ::NVM_ERR_UNKNOWN @n
 *              ::NVM_ERR_BAD_SIZE @n
 */
NVM_API int nvm_get_devices(struct device_discovery *p_devices, const NVM_UINT8 count);

/**
* @brief Retrieves -PARTIAL- #device_discovery information
* about each device in the system whether they are fully compatible
* with the current native API library version or not.
* @remarks Only attributes that can be found from NFIT will be populated on #device_discovery.
* @param[in,out] p_devices
*              An array of #device_discovery structures allocated by the caller.
* @param[in] count
*              The number of elements in the array.
* @pre The caller must have administrative privileges.
* @remarks To allocate the array of #device_discovery structures,
* call #nvm_get_device_count before calling this method.
* @return
*              ::NVM_SUCCESS @n
*              ::NVM_ERR_UNKNOWN @n
*              ::NVM_ERR_OPERATION_FAILED @n
*              ::NVM_ERR_NOT_ENOUGH_FREE_SPACE @n
*              ::NVM_ERR_BAD_SIZE @n
*/
NVM_API int nvm_get_devices_nfit(struct device_discovery *p_devices, const NVM_UINT8 count);

/**
 * @brief Retrieve #device_discovery information about the device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in,out] p_discovery
 *              A pointer to a #device_discovery structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_DIMM_NOT_FOUND @n
 */
NVM_API int nvm_get_device_discovery(const NVM_UID device_uid, struct device_discovery *p_discovery);

/**
 * @brief Retrieve the #device_status of the device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in,out] p_status
 *              A pointer to a #device_status structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @pre The device is manageable.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_DIMM_NOT_FOUND @n
 */
NVM_API int nvm_get_device_status(const NVM_UID device_uid, struct device_status *p_status);

/**
 * @brief Retrieve the PMON Registers of device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] SmartDataMask
 *              This will specify whether or not to return the extra smart data along with the PMON
 * Counter data
 * @param[out] p_output_payload
 *               A pointer to the output payload PMON registers
 * @pre The caller must have administrative privileges.
 * @pre The device is manageable.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_DIMM_NOT_FOUND @n
 */
NVM_API int nvm_get_pmon_registers(const NVM_UID device_uid, const NVM_UINT8 SmartDataMask, PMON_REGISTERS *p_output_payload);

/**
 * @brief Set the PMON Registers of device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] PMONGroupEnable
 *              Specifies which PMON Group to enable
 * @pre The caller must have administrative privileges.
 * @pre The device is manageable.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_DIMM_NOT_FOUND @n
 */
NVM_API int nvm_set_pmon_registers(const NVM_UID device_uid, NVM_UINT8 PMONGroupEnable);


/**
 * @brief Retrieve #device_settings information about the device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[out] p_settings
 *              A pointer to a #device_settings structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @pre The device is manageable.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_device_settings(const NVM_UID device_uid, struct device_settings *p_settings);

/**
 * @brief Retrieve #device_details information about the device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in,out] p_details
 *              A pointer to a #device_details structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @pre The device is manageable.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_DIMM_NOT_FOUND @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_device_details(const NVM_UID device_uid, struct device_details *p_details);

/**
 * @brief Retrieve a current snapshot of the performance metrics for the device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in,out] p_performance
 *              A pointer to a #device_performance structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @pre The device is manageable.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_device_performance(const NVM_UID device_uid, struct device_performance *p_performance);

/**
 * @brief Retrieve the firmware image log information from the device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in, out] p_fw_info
 *              A pointer to a #device_fw_info structure allocated by the caller.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_device_fw_image_info(const NVM_UID device_uid, struct device_fw_info *p_fw_info);

/**
 * @brief Push a new FW image to the device specified.
 *
 * @remarks If Address Range Scrub (ARS) is in progress on any target DIMM,
 * an attempt will be made to abort ARS and the proceed with the firmware update.
 *
 * @remarks A reboot is required to activate the updated firmware image and is
 * recommended to ensure ARS runs to completion.
 *
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] path
 *              Absolute file path to the new firmware image.
 * @param[in] path_len
 *              String length of path, should be < NVM_PATH_LEN.
 * @param[in] force
 *              If attempting to downgrade the minor version, force must be true.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @remarks A FW update may require similar changes to related devices to
 * represent a consistent correct configuration.
 *
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_NO_MEM @n
 *            ::NVM_ERR_BAD_DEVICE @n
 *            ::NVM_ERR_INVALID_PERMISSIONS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_MANAGEABLE_DIMM_NOT_FOUND @n
 *            ::NVM_ERR_DRIVER_FAILED @n
 *            ::NVM_ERR_IMAGE_FILE_NOT_VALID @n
 *            ::NVM_ERR_DATA_TRANSFER @n
 *            ::NVM_ERR_GENERAL_DEV_FAILURE @n
 *            ::NVM_ERR_BUSY_DEVICE @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_BAD_FW @n
 *            ::NVM_ERR_DUMP_FILE_OPERATION_FAILED @n
 *            ::NVM_ERR_GENERAL_OS_DRIVER_FAILURE @n
 *            ::NVM_ERR_IMAGE_EXAMINE_INVALID @n
 */
NVM_API int nvm_update_device_fw(const NVM_UID device_uid, const NVM_PATH path, const NVM_SIZE path_len, const NVM_BOOL force);

/**
 * @brief Examine the FW image to determine if it is valid for the device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] path
 *              Absolute file path to the new firmware image.
 * @param[in] path_len
 *              String length of path, should be < NVM_PATH_LEN.
 * @param image_version
 *              Firmware image version returned after examination
 * @param image_version_len
 *              Buffer size for the image version
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @remarks A FW update may require similar changes to related devices to
 * represent a consistent correct configuration.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_BAD_FW @n
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_NO_MEM @n
 *            ::NVM_ERR_BAD_DEVICE @n
 *            ::NVM_ERR_INVALID_PERMISSIONS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_MANAGEABLE_DIMM_NOT_FOUND @n
 *            ::NVM_ERR_DRIVER_FAILED @n
 *            ::NVM_ERR_IMAGE_FILE_NOT_VALID @n
 *            ::NVM_ERR_DATA_TRANSFER @n
 *            ::NVM_ERR_GENERAL_DEV_FAILURE @n
 *            ::NVM_ERR_BUSY_DEVICE @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_GENERAL_OS_DRIVER_FAILURE @n
 */
NVM_API int nvm_examine_device_fw(const NVM_UID device_uid, const NVM_PATH path, const NVM_SIZE path_len, NVM_VERSION image_version, const NVM_SIZE image_version_len);

/**
 * @brief Retrieve the supported capabilities for all devices in aggregate.
 * @param[in,out] p_capabilties
 *              A pointer to an #nvm_capabilities structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_nvm_capabilities(struct nvm_capabilities *p_capabilties);

/**
 * @brief Retrieve the aggregate capacities across all manageable DCPMMs in the system.
 * @param[in,out] p_capacities
 *              A pointer to an #device_capacities structure allocated by the caller.
 * @pre The caller must have administrative privileges.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_nvm_capacities(struct device_capacities *p_capacities);

/**
* @brief Retrieve all the health sensors for the specified DCPMM.
* @param[in] device_uid
*              The device identifier.
* @param[in,out] p_sensors
*              An array of #sensor structures allocated by the caller.
* @param[in] count
*              The number of elements in the array. Should be NVM_MAX_DEVICE_SENSORS.
* @pre The caller has administrative privileges.
* @pre The device is manageable.
* @remarks Sensors are used to monitor a particular aspect of a device by
* settings thresholds against a current value.
* @remarks The number of sensors for a device is defined as NVM_MAX_DEVICE_SENSORS.
* @remarks Sensor information is returned as part of the #device_details structure.
* @return
*            ::NVM_SUCCESS @n
*            ::NVM_ERR_INVALID_PARAMETER @n
*            ::NVM_ERR_UNKNOWN @n
*/
NVM_API int nvm_get_sensors(const NVM_UID device_uid, struct sensor *p_sensors, const NVM_UINT16 count);

/**
* @brief Retrieve a specific health sensor from the specified DCPMM.
* @param[in] device_uid
*              The device identifier.
* @param[in] type
*              The specific #sensor_type to retrieve.
* @param[in,out] p_sensor
*              A pointer to a #sensor structure allocated by the caller.
* @pre The caller has administrative privileges.
* @pre The device is manageable.
* @return
*            ::NVM_SUCCESS @n
*            ::NVM_ERR_INVALID_PARAMETER @n
*            ::NVM_ERR_UNKNOWN @n
*/
NVM_API int nvm_get_sensor(const NVM_UID device_uid, const enum sensor_type type, struct sensor *p_sensor);

/**
* @brief Change the critical threshold on the specified health sensor for the specified
* DCPMM.
* @param[in] device_uid
*              The device identifier.
* @param[in] type
*              The specific #sensor_type to modify.
* @param[in] p_settings
*              The modified settings.
* @pre The caller has administrative privileges.
* @pre The device is manageable.
* @return
*            ::NVM_SUCCESS @n
*            ::NVM_ERR_INVALID_PARAMETER @n
*            ::NVM_ERR_UNKNOWN @n
*/
NVM_API int nvm_set_sensor_settings(const NVM_UID device_uid, const enum sensor_type type, const struct sensor_settings *p_settings);

/**
 * @}
 * @defgroup Security
 * These functions manage the security state of DCPMMs.
 * @{
 */

/**
 * @brief If data at rest security is not enabled, this method enables it and
 * sets the passphrase. If data at rest security was previously enabled, this method changes
 * the passphrase to the new passphrase specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] old_passphrase
 *              The current passphrase or NULL if security is disabled.
 * @param[in] old_passphrase_len
 *              String length of old_passphrase,
 *              should be <= NVM_PASSPHRASE_LEN or 0 if security is disabled.
 * @param[in] new_passphrase
 *              The new passphrase.
 * @param[in] new_passphrase_len
 *              String length of new_passphrase, should be <= NVM_PASSPHRASE_LEN.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @pre Device security is not frozen.
 * @pre The device passphrase limit has not been reached.
 * @post The device will be unlocked and frozen.
 * @post The device will be locked on the next reset.
 * @return
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 */
NVM_API int nvm_set_passphrase(const NVM_UID device_uid, const NVM_PASSPHRASE old_passphrase, const NVM_SIZE old_passphrase_len, const NVM_PASSPHRASE new_passphrase, const NVM_SIZE new_passphrase_len);

/**
 * @brief Disables data at rest security and removes the passphrase.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] passphrase
 *              The current passphrase.
 * @param[in] passphrase_len
 *              String length of passphrase, should be <= NVM_PASSPHRASE_LEN.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @pre Device security is enabled and the passphrase has been set using #nvm_set_passphrase.
 * @pre Device security is not frozen.
 * @pre The device passphrase limit has not been reached.
 * @post The device will be unlocked if it is currently locked.
 * @post Device security will be disabled.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_NO_MEM @n
 *            ::NVM_ERR_BAD_DEVICE @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_INVALID_PERMISSIONS @n
 *            ::NVM_ERR_MANAGEABLE_DCPMM_NOT_FOUND @n
 *            ::NVM_ERR_DRIVER_FAILED @n
 *            ::NVM_ERR_INVALID_SECURITY_OPERATION @n
 *            ::NVM_ERR_INVALID_PASSPHRASE @n
 *            ::NVM_ERR_PASSPHRASES_DO_NOT_MATCH @n
 *            ::NVM_ERR_DATA_TRANSFER @n
 *            ::NVM_ERR_GENERAL_DEV_FAILURE @n
 *            ::NVM_ERR_BUSY_DEVICE @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_GENERAL_OS_DRIVER_FAILURE @n
 */
NVM_API int nvm_remove_passphrase(const NVM_UID device_uid, const NVM_PASSPHRASE passphrase, const NVM_SIZE passphrase_len);

/**
 * @brief Unlocks the device with the passphrase specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] passphrase
 *              The current passphrase.
 * @param[in] passphrase_len
 *              String length of passphrase, should be <= NVM_PASSPHRASE_LEN.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @pre Device security is enabled and the passphrase has been set using #nvm_set_passphrase.
 * @pre Device security is not frozen.
 * @pre The device passphrase limit has not been reached.
 * @post The device will be unlocked and frozen.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_API_NOT_SUPPORTED @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_unlock_device(const NVM_UID device_uid, const NVM_PASSPHRASE passphrase, const NVM_SIZE passphrase_len);

/**
 * @brief Prevent security lock state changes to the DCPMM until the next reboot
 * @param[in] device_uid
 *              The device identifier.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @pre The device supports unlocking a device.
 * @pre Current DCPMM security state is unlocked.
 * @post DCPMM security state will be frozen.
 * @post Device security will be changed.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_API_NOT_SUPPORTED @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_freezelock_device(const NVM_UID device_uid);

/**
 * @brief Erases data on the device specified by zeroing the device encryption key.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] passphrase
 *              The current passphrase.
 * @param[in] passphrase_len
 *              String length of passphrase, should be <= NVM_PASSPHRASE_LEN.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @pre The device supports overwriting a device.
 * @pre Device security is disabled or sanitize antifreeze.
 * @post All user data is inaccessible.
 * @post Device security will be changed.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_NO_MEM @n
 *            ::NVM_ERR_BAD_DEVICE @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_INVALID_PERMISSIONS @n
 *            ::NVM_ERR_MANAGEABLE_DIMM_NOT_FOUND @n
 *            ::NVM_ERR_DRIVER_FAILED @n
 *            ::NVM_ERR_INVALID_SECURITY_OPERATION @n
 *            ::NVM_ERR_PASSPHRASES_DO_NOT_MATCH @n
 *            ::NVM_ERR_DATA_TRANSFER @n
 *            ::NVM_ERR_GENERAL_DEV_FAILURE @n
 *            ::NVM_ERR_BUSY_DEVICE @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_GENERAL_OS_DRIVER_FAILURE @n
 */
NVM_API int nvm_erase_device(const NVM_UID device_uid, const NVM_PASSPHRASE passphrase, const NVM_SIZE passphrase_len);

/**
 * @brief If data at rest security is not enabled and master passphrase is enabled
 * in the DCPMM security state, this method modifies the master passphrase. On 
 * Microsoft(R) Windows(TM) this functionality may be prohibited if there are any
 * namespaces present.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] old_master_passphrase
 *              The current master passphrase. For default Master Passphrase (0's) use a zero length, null terminated string.
 * @param[in] old_master_passphrase_len
 *              String length of old_master_passphrase,
 *              should be <= NVM_PASSPHRASE_LEN.
 * @param[in] new_master_passphrase
 *              The new master passphrase.
 * @param[in] new_master_passphrase_len
 *              String length of new_master_passphrase, should be <= NVM_PASSPHRASE_LEN.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @pre The device master passphrase is enabled.
 * @pre Device security is not enabled.
 * @pre The device master passphrase limit has not been reached.
 * @pre The device master passphrase has not been changed on this boot.
 * @return
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_SECURITY_COUNT_EXPIRED @n
 *            ::NVM_ERR_INVALID_SECURITY_STATE@n
 *            ::NVM_ERR_PASSPHRASE_NOT_PROVIDED@n
 */
NVM_API int nvm_set_master_passphrase(const NVM_UID device_uid,
                                      const NVM_PASSPHRASE old_master_passphrase,
                                      const NVM_SIZE old_master_passphrase_len,
                                      const NVM_PASSPHRASE new_master_passphrase,
                                      const NVM_SIZE new_master_passphrase_len);

/**
 * @}
 * @defgroup Events
 * These functions provide access to various events generated from
 * DCPMMs.
 * @{
 */

/**
 * @brief Retrieve the number of events in the native API library event database.
 * @param[in] p_filter
 *              A pointer to an event_filter structure allocated by the caller to
 *              optionally filter the event count.
 * @param[in,out] count
 *              A pointer an integer that will contain the number of events
 * @pre The caller must have administrative privileges.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_API_NOT_SUPPORTED @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_number_of_events(const struct event_filter *p_filter, int *count);

/**
 * @brief Retrieve a list of stored events from the native API library database and
 * optionally filter the results.
 * @param[in] p_filter
 *              A pointer to an event_filter structure to optionally
 *              limit the results.  NULL to return all the events.
 * @param[in,out] p_events
 *              An array of #event structures allocated by the caller.
 * @param[in] count
 *              The number of elements in the array.
 * @pre The caller must have administrative privileges.
 * @remarks The native API library stores a maximum of 10,000 events in the table,
 * rolling the table once the maximum is reached. However, the maximum number of events
 * is configurable by modifying the EVENT_LOG_MAX_ROWS value in the configuration database.
 * @remarks To allocate the array of #event structures,
 * call #nvm_get_number_of_events before calling this method.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_NOT_ENOUGH_FREE_SPACE @n
 *            ::NVM_ERR_BAD_SIZE @n
 */
NVM_API int nvm_get_events(const struct event_filter *p_filter, struct event *p_events, const NVM_UINT16 count);

/**
 * @brief Purge stored events from the native API database.
 * @param[in] p_filter
 *              A pointer to an event_filter structure to optionally
 *              purge only specific events.
 * @pre The caller must have administrative privileges.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_purge_events(const struct event_filter *p_filter);

/**
 * @brief Acknowledge an event from the native API database.
 * (i.e. setting action required field from true to false)
 * @param[in] event_id
 *              The event id of the event to be acknowledged.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_acknowledge_event(NVM_UINT32 event_id);

/**
 * @brief Retrieve the number of configured persistent memory regions in the host server.
 * @pre The caller has administrative privileges.
 * @remarks This method should be called before #nvm_get_regions.
 * @param[in,out] count
 *              A pointer an integer that will contain the number of region count on return
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_number_of_regions(NVM_UINT8 *count);

/**
 * @brief Retrieve the number of configured persistent memory regions in the host server.
 * @pre The caller has administrative privileges.
 * @remarks This method should be called before #nvm_get_regions.
 * @param[in] use_nfit
 *              0: Use PCD data to get region information.
 *              1: Use NFIT table to get region information.
 * @param[in,out] count
 *              A pointer an integer that will contain the number of region count on return
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_number_of_regions_ex(const NVM_BOOL use_nfit, NVM_UINT8 *count);

/**
 * @brief Retrieve a list of the configured persistent memory regions in host server.
 * @param[in,out] p_regions
 *              An array of #region structures allocated by the caller.
 * @param[in,out] count
 *              The number of elements in the array allocated by the caller and returns the count of regions that were returned.
 * @pre The caller has administrative privileges.
 * @remarks To allocate the array of #region structures,
 * call #nvm_get_region_count before calling this method.
 * @return
 *            ::NVM_SUCCESS
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_NO_MEM @n
 */
NVM_API int nvm_get_regions(struct region *p_regions, NVM_UINT8 *count);

/**
 * @brief Retrieve a list of the configured persistent memory regions in host server.
 * @param[in,out] p_regions
 *              An array of #region structures allocated by the caller.
 * @param[in] use_nfit
 *              0: Use PCD data to get region information.
 *              1: Use NFIR table to get region information.
 * @param[in,out] count
 *              The number of elements in the array allocated by the caller and returns the count of regions that were returned.
 * @pre The caller has administrative privileges.
 * @remarks To allocate the array of #region structures,
 * call #nvm_get_region_count before calling this method.
 * @return
 *            ::NVM_SUCCESS
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_NO_MEM @n
 */
NVM_API int nvm_get_regions_ex(const NVM_BOOL use_nfit, struct region *p_regions, NVM_UINT8 *count);

/**
 * @brief Modify how the DCPMM capacity is provisioned by the BIOS on the next reboot.
 * @param p_device_uids
 *              Pointer to list of device uids to configure.
 *              If NULL, all devices on platform will be configured.
 * @param device_uids_count
 *              Number of devices in p_device_uids list.
 * @param p_goal
 *              Values that defines how regions are created.
 * @pre The caller has administrative privileges.
 * @pre The specified DCPMM is manageable by the host software.
 * @pre Any existing namespaces created from capacity on the
 *              DCPMM must be deleted first.
 * @remarks This operation stores the specified configuration goal on the DCPMM
 *              for the BIOS to read on the next reboot.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_create_config_goal(NVM_UID *p_device_uids, NVM_UINT32 device_uids_count, struct config_goal_input *p_goal);

/**
 * @brief Retrieve the configuration goal from the specified DCPMM.
 * @param p_device_uids
 *              Pointer to list of device uids to retrieve config goal from.
 *              If NULL, retrieve goal configs from all devices on platform.
 * @param device_uids_count
 *              Number of devices in p_device_uids list.
 * @param p_goal
 *              A pointer to a list of config_goal structures allocated by the caller.
 * @pre The caller has administrative privileges.
 * @pre The specified DCPMM is manageable by the host software.
 * @remarks A configuration goal is stored on the DCPMM until the
 *              BIOS successfully processes it on reboot.
 *              Use @link nvm_delete_config_goal @endlink to erase a
 *              configuration goal from a DCPMM.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @nNVM_SUCCESS @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_get_config_goal(NVM_UID *p_device_uids, NVM_UINT32 device_uids_count, struct config_goal *p_goal);

/**
 * @brief Erase the region configuration goal from the specified DCPMM.
 * @param p_device_uids
 *              Pointer to list of device uids to erase the region config goal.
 *              If NULL, all devices on platform will have their region config goal erased.
 * @param device_uids_count
 *              Number of devices in p_device_uids list.
 * @pre The caller has administrative privileges.
 * @pre The specified DCPMM is manageable by the host software.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_delete_config_goal(NVM_UID *p_device_uids, NVM_UINT32 device_uids_count);

/**
 * @brief Store the configuration settings of how the DCPMM capacity
 * is currently provisioned to a file in order to duplicate the
 * configuration elsewhere.
 * @param file
 *              The absolute file path in which to store the configuration data.
 * @param file_len
 *              String length of file, should be < #NVM_PATH_LEN.
 * @pre The caller has administrative privileges.
 * @pre The specified DCPMM is manageable by the host software.
 * @pre The specified DCPMM is currently configured.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_DUMP_FILE_OPERATION_FAILED @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_dump_goal_config(const NVM_PATH file, const NVM_SIZE file_len);

/**
 * @brief Modify how the DCPMM capacity is provisioned by the BIOS on the
 * next reboot by applying the configuration goal previously stored in the
 * specified file with @link nvm_dump_config @endlink.
 * @param file
 *              The absolute file path containing the region configuration goal to load.
 * @param file_len
 *              String length of file, should be < NVM_PATH_LEN.
 * @pre The caller has administrative privileges.
 * @pre The specified DCPMM is manageable by the host software.
 * @pre Any existing namespaces created from capacity on the
 *              DCPMM must be deleted first.
 * @pre If the configuration goal contains any app direct memory,
 *              all DCPMMs that are part of the interleave set must be included in the file.
 * @pre The specified DCPMM must be >= the total capacity of the DCPMM
 *              specified in the file.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_CREATE_GOAL_NOT_ALLOWED @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_load_goal_config(const NVM_PATH file, const NVM_SIZE file_len);

/**
 * @}
 * @defgroup Support
 * These functions provide various support functionality of DCPMMs.
 * @{
 */

/**
 * @brief Retrieve the native API library major version number.
 * @remarks Applications and the native API Library are not compatible if they were
 *              written against different major versions of the native API definition.
 *              For this reason, it is recommended that every application that uses the
 *              native API Library to perform the following check:
 *              if (#nvm_get_major_version() != NVM_VERSION_MAJOR)
 * @returnThe major version number of the library.
 */
NVM_API int nvm_get_major_version();

/**
 * @brief Retrieve the native API library minor version number.
 * @remarks Unless otherwise stated, every data structure, function, and description
 *              described in this document has existed with those exact semantics since version 1.0
 *              of the library.  In cases where functions have been added,
 *              the appropriate section in this document will describe the version that introduced
 *              the new feature.  Applications wishing to check for features that were added
 *		may do so by comparing the return value from #nvm_get_minor_version() against the
 *              minor number in this specification associated with the introduction of the new feature.
 * @returnThe minor version number of the library.
 */
NVM_API int nvm_get_minor_version();

/**
 * @brief Retrieve the native API library hot fix version number.
 * @returnThe hot fix version number of the library.
 */
NVM_API int nvm_get_hotfix_number();

/**
 * @brief Retrieve the native API library build version number.
 * @returnThe build version number of the library.
 */
NVM_API int nvm_get_build_number();

/**
 * @brief Retrieve native API library version as a string in the format MM.mm.hh.bbbb,
 * where MM is the major version, mm is the minor version, hh is the hotfix number
 * and bbbb is the build number.
 * @param[in,out] version_str
 *              A buffer for the version string allocated by the caller.
 * @param[in] str_len
 *              Size of the version_str buffer.  Should be NVM_VERSION_LEN.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 */
NVM_API int nvm_get_version(NVM_VERSION version_str, const NVM_SIZE str_len);


/**
 * @brief Collect support data into a single file to document the context of a problem
 * for offline analysis by support or development personnel.
 * @param[in] support_file
 *              Absolute file path where the support file will be stored.
 * @param[in] support_file_len
 *              String length of the file path, should be < NVM_PATH_LEN.
 * @pre The caller must have administrative privileges.
 * @post A support file exists at the path specified for debug by
 * support or development personnel.
 * @remarks The support file contains a current snapshot of the system, events logs, current
 * performance data, basic #host server information, SW version, memoryresources, system
 * capabilities, topology, sensor values and diagnostic data.
 * @remarks This operation will be attempt to gather as much information as possible about
 * the state of the system.  Therefore, it will ignore errors during the information
 * gathering process and only generate errors for invalid input parameters
 * or if the support file is not able to be generated.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_gather_support(const NVM_PATH support_file, const NVM_SIZE support_file_len);


/**
 * @brief Inject an error into the device specified for debugging purposes.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] p_error
 *              A pointer to a #device_error structure containing the injected
 *              error information allocated by the caller.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @pre This interface is only supported by the underlying DCPMM firmware when it is in a
 * debug state.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_inject_device_error(const NVM_UID device_uid, const struct device_error *p_error);

/**
 * @brief Clear an injected error into the device specified for debugging purposes.
 *        From a FIS perspective, it is setting the enable/disable field to disable for
 *        the specified injected error type.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] p_error
 *              A pointer to a #device_error structure containing the injected
 *              error information allocated by the caller.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @pre This interface is only supported by the underlying DCPMM firmware when it is in a
 * debug state.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_clear_injected_device_error(const NVM_UID device_uid, const struct device_error *p_error);

/**
 * @brief Run a diagnostic test on the device specified.
 * @param[in] device_uid
 *              The device identifier.
 * @param[in] p_diagnostic
 *              A pointer to a #diagnostic structure containing the
 *              diagnostic to run allocated by the caller.
 * @param[in,out] p_results
 *              The number of diagnostic failures. To see full results use #nvm_get_events.
 * @pre The caller has administrative privileges.
 * @pre The device is manageable.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_run_diagnostic(const NVM_UID device_uid, const struct diagnostic *p_diagnostic, NVM_UINT32 *p_results);

/**
 * @brief Set the user preference config value in DCPMM software.  See the Change Preferences section of the CLI
 * specification for a list of supported preferences and values.  Note, this API does not verify if the property key
 * is supported, or if the value is supported per the CLI specification.
 * @param[in] key
 *              The preference name.
 * @param[in] value
 *              The preference value.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_set_user_preference(const NVM_PREFERENCE_KEY key, const NVM_PREFERENCE_VALUE value);

/**
 * @brief Clear namespace label storage area in PCD on the specified DCPMM.
 * @param[in] device_uid
 *              The device identifier.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_DIMM_NOT_FOUND @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_clear_dimm_lsa(const NVM_UID device_uid);

/**
 * @}
 * @defgroup Logging
 * These functions manage the logging features of
 * DCPMM software.
 * @{
 */

/**
 * @brief Determine if the native API debug logging is enabled.
 * @pre The caller must have administrative privileges.
 * @returnReturns true (1) if debug logging is enabled and false (0) if not,
 * or
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_debug_logging_enabled();

/**
 * @brief Toggle whether the native API library performs debug logging.
 * @param[in] enabled @n
 *              0: Debug logger disabled. @n
 *              1: Log warning and error debug traces to the file. @n
 * @pre The caller must have administrative privileges.
 * @remarks By default, the native API library starts logging errors only.
 * @remarks Debug logging may impact native API library performance depending
 * on the workload of the library.  It's recommended that debug logging is only
 * turned on during troubleshooting or debugging.
 * @remarks Changing the debug log level is NOT persistent.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_UNKNOWN @n
 */
NVM_API int nvm_toggle_debug_logging(const NVM_BOOL enabled);

/**
 * @brief Retrieves #job information about each device in the system
 * @param[in,out] p_jobs
 *              An array of #job structures allocated by the caller.
 *              One for each device in the system.
 * @param[in] count
 *              The number of elements in the array.
 * @pre The caller must have administrative privileges.
 * @remarks To allocate the array of #job structures,
 * call #nvm_get_number_of_devices before calling this method.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_NOT_ENOUGH_FREE_SPACE @n
 *            ::NVM_ERR_OPERATION_FAILED @n
 *            ::NVM_ERR_BAD_SIZE @n
 */
NVM_API int nvm_get_jobs(struct job *p_jobs, const NVM_UINT32 count);

/**
 * @brief Initialize a new context
 */
NVM_API int nvm_create_context();

/**
 * @brief Clean up the current context
 */
NVM_API int nvm_free_context(const NVM_BOOL force);

/**
 * A device pass-through command. Refer to the FW specification
 * for specific details about the individual fields.
 */
struct device_pt_cmd {
  NVM_UINT8	opcode;                         ///< Command opcode.
  NVM_UINT8	sub_opcode;                     ///<  Command sub-opcode.
  NVM_UINT32	input_payload_size;             ///<  Size of the input payload.
  void *		input_payload;                  ///< A pointer to the input payload buffer.
  NVM_UINT32	output_payload_size;            ///< Size of the output payload.
  void *		output_payload;                 ///< A pointer to the output payload buffer.
  NVM_UINT32	large_input_payload_size;       ///< Size of the large input payload.
  void *		large_input_payload;            ///< A pointer to the large input payload buffer.
  NVM_UINT32	large_output_payload_size;      ///< Size of the large output payload.
  void *		large_output_payload;           ///< A pointer to the large output payload buffer.
  int		result;                         ///< Return code from the pass through command
};

/**
 * @brief Send a firmware command directly to the specified device without
 * checking for valid input.
 * @param device_uid
 *              The device identifier.
 * @param p_cmd
 *              A pointer to a @link #device_pt_command @endlink structure defining the command to send.
 * @return
 *            ::NVM_SUCCESS @n
 *            ::NVM_ERR_INVALID_PARAMETER @n
 *            ::NVM_ERR_INVALID_PERMISSIONS @n
 *            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
 *            ::NVM_ERR_NO_MEM @n
 *            ::NVM_ERR_UNKNOWN @n
 *            ::NVM_ERR_BAD_DEVICE @n
 *            ::NVM_ERR_DRIVER_FAILED @n
 *            ::NVM_ERR_DATA_TRANSFER @n
 *            ::NVM_ERR_GENERAL_DEV_FAILURE @n
 *            ::NVM_ERR_BUSY_DEVICE @n
 */
NVM_API int nvm_send_device_passthrough_cmd(const NVM_UID device_uid, struct device_pt_cmd *p_cmd);

/**
* @brief Retrieve a FW error log entry
* @param[in] device_uid The device identifier
* @param[in] seq_num Log entry sequence number
* @param[in] log_level Log entry log level (0: Low, 1: High)
* @param[in] log_type Log entry log type (0: Media, 1: Thermal)
* @param[out] error_entry pointer to buffer to store a single FW error log entry
* @return
*            ::NVM_SUCCESS @n
*            ::NVM_SUCCESS_NO_ERROR_LOG_ENTRY @n
*            ::NVM_ERR_INVALID_PARAMETER @n
*            ::NVM_ERR_INVALID_PERMISSIONS @n
*            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
*            ::NVM_ERR_NO_MEM @n
*            ::NVM_ERR_UNKNOWN @n
*            ::NVM_ERR_BAD_DEVICE @n
*            ::NVM_ERR_DRIVER_FAILED @n
*            ::NVM_ERR_GENERAL_DEV_FAILURE @n
*            ::NVM_ERR_BUSY_DEVICE @n
*/
NVM_API int nvm_get_fw_error_log_entry_cmd(const NVM_UID   device_uid, const unsigned short  seq_num, const unsigned char log_level, const unsigned char log_type, ERROR_LOG * error_entry);

/**
* @brief Retrieve a FW error log counters: current and oldest sequence number for each log type.
* @param[in] device_uid The device identifier
* @param[out] error_log_stats Pointer to #device_error_log_status.
* @return
*            ::NVM_SUCCESS @n
*            ::NVM_ERR_INVALID_PARAMETER @n
*            ::NVM_ERR_INVALID_PERMISSIONS @n
*            ::NVM_ERR_OPERATION_NOT_SUPPORTED @n
*            ::NVM_ERR_NO_MEM @n
*            ::NVM_ERR_UNKNOWN @n
*            ::NVM_ERR_BAD_DEVICE @n
*            ::NVM_ERR_DRIVER_FAILED @n
*            ::NVM_ERR_GENERAL_DEV_FAILURE @n
*            ::NVM_ERR_BUSY_DEVICE @n
*/

NVM_API int nvm_get_fw_err_log_stats(const NVM_UID device_uid, struct device_error_log_status *error_log_stats);

/**
* @brief Lock API
*/
NVM_API void nvm_sync_lock_api();

/**
* @brief Unlock API
*/
NVM_API void nvm_sync_unlock_api();

#ifdef __cplusplus
}
#endif

#endif  /* _NVM_MANAGEMENT_H_ */

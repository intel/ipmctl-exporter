# Intel® Optane™ Persistent Memory Controller Exporter
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

Intel® Optane™ PMCE is a utility for exposing health and performance
metrics from Intel Optane DC Persistent memory modules (DCPMM) for
[Prometheus](https://prometheus.io/) (an open source monitoring system).
Exporter is linking to libipmctl and consume its API. Library is a part of
[IPMCTL project](https://github.com/intel/ipmctl).

# Exported Metrics

basic set: `# sudo ./ipmctl_exporter`


Name                                                      | Description
----------------------------------------------------------|-------------
ipmctl_health                                             | DCPMM health as reported in the SMART log
ipmctl_media_temperature_degrees_c                        | Device media temperature in degrees Celsius
ipmctl_controller_temperature_degrees_c                   | Device media temperature in degrees Celsius
ipmctl_percentage_remaining                               | Amount of percentage remaining as a percentage
ipmctl_latched_dirty_shutdown_count_total                 | Device shutdowns without notification
ipmctl_power_on_time_total                                | Total power-on time over the lifetime of the device
ipmctl_up_time_seconds_total                              | Total power-on time since the last power cycle of the device
ipmctl_power_cycles_total                                 | Number of power cycles over the lifetime of the device
ipmctl_fw_error_count_total                               | The total number of firmware error log entries
ipmctl_unlatched_dirty_shutdown_count_total               | Number of times that the FW received an unexpected power loss
ipmctl_total_media_reads_total                            | Lifetime number of 64 byte reads from media on the DCPMM
ipmctl_total_media_writes_total                           | Lifetime number of 64 byte writes to media on the DCPMM
ipmctl_total_read_requests_total                          | Lifetime number of DDRT read transactions the DCPMM has serviced
ipmctl_total_write_requests_total                         | Lifetime number of DDRT write transactions the DCPMM has serviced
ipmctl_device_discovery_info                              | Describes the capabilities supported by a DCPMM
ipmctl_device_security_capabilities_info                  | Describes the security capabilities of a device
ipmctl_device_discovery_info                              | Describes an enterprise-level view of a device
ipmctl_exporter_info                                      | Describes ipmctl-exporter information


If you would like to add some alerts in Prometheus to get notification after
reaching some configured thresholds, you may enable it as well (these are
disabled by default) to do it try:
`# sudo ./ipmctl_exporter --enable-thresholds`


Name                                                      | Description
----------------------------------------------------------|-------------
ipmctl_media_temperature_enabled                          | Indictes if firmware notifications are enabled when media temperature value is critical
ipmctl_media_temperature_upper_critical_threshold         | The upper media temperature critical threshold
ipmctl_media_temperature_lower_critical_threshold         | The lower media temperature critical threshold
ipmctl_media_temperature_upper_fatal_threshold            | The upper media temperature fatal threshold
ipmctl_media_temperature_lower_fatal_threshold            | The lower media temperature fatal threshold
ipmctl_media_temperature_upper_noncritical_threshold      | The upper media temperature noncritical threshold
ipmctl_media_temperature_lower_noncritical_threshold      | The lower media temperature noncritical threshold
ipmctl_controller_temperature_enabled                     | Indictes if firmware notifications are enabled when controller temperature value is critical
ipmctl_controller_temperature_upper_critical_threshold    | The upper controller temperature critical threshold
ipmctl_controller_temperature_lower_critical_threshold    | The lower controller temperature critical threshold
ipmctl_controller_temperature_upper_fatal_threshold       | The upper controller temperature fatal threshold
ipmctl_controller_temperature_lower_fatal_threshold       | The lower controller temperature fatal threshold
ipmctl_controller_temperature_upper_noncritical_threshold | The upper controller temperature noncritical threshold
ipmctl_controller_temperature_lower_noncritical_threshold | The lower controller temperature noncritical threshold
ipmctl_percentage_remaining_enabled                       | Indictes if firmware notifications are enabled when percentage remaining value is critical
ipmctl_percentage_remaining_upper_critical_threshold      | The upper percentage remaining critical threshold
ipmctl_percentage_remaining_lower_critical_threshold      | The lower percentage remaining critical threshold
ipmctl_percentage_remaining_upper_fatal_threshold         | The upper percentage remaining fatal threshold"
ipmctl_percentage_remaining_lower_fatal_threshold         | The lower percentage remaining fatal threshold"
ipmctl_percentage_remaining_upper_noncritical_threshold   | The upper percentage remaining noncritical threshold
ipmctl_percentage_remaining_lower_noncritical_threshold   | The lower percentage remaining noncritical threshold

# Build

As far as IPMCTL exporter utilize libipmctl as well as libndctl (both are
external libraries) supported systems depends on availability of these libraries
under different Operating Systems.


For Linux we highly recommend:

**Fedora 29/30 (Workstation Edition) x64** with installed
latest [golang](https://golang.org/) compiler,
latest [pkg-config](http://pkg-config.freedesktop.org/),
latest [GCC](https://gcc.gnu.org/),
latest [cmake](https://cmake.org/download/),
and latest [ipmctl](https://github.com/intel/ipmctl/releases) +
[ndctl](https://github.com/pmem/ndctl) libraries,
follow the steps below to prepare you environment for builds:
```shell
dnf install -y git cmake pkg-config gcc golang ndctl-libs libipmctl
git clone https://sv-gitlab.igk.intel.com/validation-analytics-platform/ipmctl_exporter.git
cd ./ipmctl_exporter/src
cmake -S . -B output
```


To proceed with build:
```shell
export PKG_CONFIG_PATH=`pwd`/output/
make -C output
```


For Windows we highly recommend:

**Windows Server 2016 Standard x64** with installed latest
[golang](https://golang.org/) compiler,
latest [pkg-config](http://pkg-config.freedesktop.org/),
latest [TDM64-GCC](https://jmeubank.github.io/tdm-gcc/),
latest [cmake](https://cmake.org/download/),
and latest [ipmctl](https://github.com/intel/ipmctl/releases) library,
follow the steps below to prepare you environment for builds:
- Install **golang**
from [here](https://golang.org/doc/install?download=go1.14.4.windows-amd64.msi)
to `C:\Go` directory
- Install **tdm64-gcc**
from [here](https://jmeubank.github.io/tdm-gcc/)
to `C:\TDM-GCC-64` directory
- Install **pkgconfiglite**
from [here](https://sourceforge.net/projects/pkgconfiglite/files/)
to `C:\TDM-GCC-64\bin` directory
- Install **cmake**
from [here](https://cmake.org/download/)
to `C:\Program Files\CMake` directory
- Install **ipmctl library**
from [here](https://github.com/intel/ipmctl/releases)
choose latest build for Windows OS
- From cmd.exe:
```powershell
git clone https://sv-gitlab.igk.intel.com/validation-analytics-platform/ipmctl_exporter.git
cd ipmctl_exporter\src
cmake -S . -B output -G "MinGW Makefiles"
```

To proceed with build:
```powershell
set PKG_CONFIG_PATH=%cd%\output\
mingw32-make -C output
```


# Run

Referring to the
[list of default ports](https://github.com/prometheus/prometheus/wiki/Default-port-allocations)
by default ipmclt-exporter serves on port `0.0.0.0:9757` at endpoint `/metrics`,
for more details about the usage type:

```
sudo ./ipmctl_exporter --help
```

ipmctl_exporter as well as ipmctl tool has to be run as root user, otherwise
you should receive error code 268 (INVALID PERMISSIONS) trying to collect some
data.


# Code of Conduct
We are following rules defined by
[Contributor Covenant Code of Coduct](CODE_OF_CONDUCT.md) version 2.0

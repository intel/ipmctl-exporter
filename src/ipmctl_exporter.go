 /**
 * Copyright (c) 2020, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package main

import (
    "fmt"
    "flag"
    "os"
    "os/signal"
    "github.com/intel/ipmctl_exporter/collector"
)

var Version string

func parseCmdArgs() (string, bool, bool) {
    port := flag.String("port", "9165",
        "Listening port number used by exporter")
    enableThresholds := flag.Bool("thresholds-enable", false,
        "Enable media and controller temperature, plus percentage remaining thresholds collection")
    showVersion := flag.Bool("version", false,
        "Shows ipmctl_exporter version")
    flag.Parse()
    return *port, *enableThresholds, *showVersion
}

func handleSIGINT() {
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)
    go func() {
        for sig := range signals {
            collector.Stop()
            fmt.Printf("ipmctl exporter - catch %s, stopping service\n", sig)
            os.Exit(1)
        }
    }()
}

func main() {
    port, enableThresholds, showVersion := parseCmdArgs()
    if showVersion {
        fmt.Printf("%s\n", Version)
        os.Exit(0)
    }
    handleSIGINT()
    fmt.Printf("ipmctl exporter listening on port :%s\n", port)
    collector.Version = Version
    collector.Run(port, enableThresholds)
}
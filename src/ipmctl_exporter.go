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

func parseCmdArgs() (string, bool){
    port := flag.String("port", "9165",
        "Listening port number used by exporter")
    thresholds_enable := flag.Bool("thresholds-enable", false,
        "Enable media and controller temperature, plus percentage remaining thresholds collection")
    flag.Parse()
    return *port, *thresholds_enable
}

func handleSIGINT() () {
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)
    go func(){
        for sig := range signals {
            collector.Stop()
            fmt.Printf("ipmctl exporter - catch %s, stopping service\n", sig)
            os.Exit(1)
        }
    }()
}

func main() {
    port, thresholds_enable := parseCmdArgs()
    handleSIGINT()
    fmt.Printf("ipmctl exporter listening on port :%s\n", port)
    collector.Run(port, thresholds_enable)
}

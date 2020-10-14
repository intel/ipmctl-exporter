/**
 * Copyright (c) 2020, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package main

import (
    "fmt"
    "flag"
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

func main() {
    port, thresholds_enable := parseCmdArgs()
    fmt.Printf("ipmctl exporter listening on port :%s\n", port)
    collector.Run(port, thresholds_enable)
}

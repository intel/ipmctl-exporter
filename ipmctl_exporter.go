 /**
 * Copyright (c) 2020-2021, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 */

package main
import (
    "fmt"
    "flag"
    "os"
    "os/signal"
    "github.com/intel/ipmctl_exporter/collector"
    "github.com/natefinch/lumberjack"
    "github.com/elastic/go-elasticsearch/v7"
    "gopkg.in/go-extras/elogrus.v7"
    log "github.com/sirupsen/logrus"
)

var Version string

func setupLogger(loggingLevel string, logOnConsole bool, elasticUsed bool, elasticAddress string, indexName string) (){
    log.SetFormatter(&log.JSONFormatter {
        DisableTimestamp: false,
        PrettyPrint: false,
    })
    if !logOnConsole {
        /*
        Lumberjack is used for easy log rotation.
        When MaxSize(MB) is exceeded current file is closed and renamed(original name + timestamp)
        then a new file is created (with original name) where new logs are written.
        */
        lumberjackLogger := &lumberjack.Logger {
            Filename:   "ipmctl-exporter.log",
            MaxSize:    10,
            MaxBackups: 1,
            MaxAge:     7,
            LocalTime:  true,
        }
        log.SetOutput(lumberjackLogger)
    }
    
    // Logrus does not support a silent mode - a workaround
    if loggingLevel == "Silent" {
        loggingLevel = "Panic"
    }
    logLevel, err := log.ParseLevel(loggingLevel)
    if err != nil {
        fmt.Printf("Given logging level does not exist\n")
        os.Exit(1)
    }
    fmt.Printf("Logging level = %s\n", logLevel)
    log.SetLevel(logLevel)
    log.Debug("Logger Initialized")
    if elasticUsed {
        client, err := elasticsearch.NewClient(elasticsearch.Config {
            Addresses: []string{elasticAddress},
        })
        if err != nil {
            log.Fatal(err)
        }
        hostname, err := os.Hostname()
        if err != nil {
            log.Fatal(err)
        }
        hook, err := elogrus.NewAsyncElasticHook(client, hostname, logLevel, indexName)
        if err != nil {
            log.Fatal(err)
        }
        log.AddHook(hook)
        log.Debug("Logging to elasticsearch")
    }
}

func parseCmdArgs() (string, bool, bool, string, bool, bool, string, string) {
    port := flag.String("port", "9757",
        "Listening port number used by exporter")
    enableThresholds := flag.Bool("thresholds-enable", false,
        "Enable media and controller temperature, plus percentage remaining thresholds collection")
    showVersion := flag.Bool("version", false,
        "Shows ipmctl_exporter version")
    loggingLevel := flag.String("log-level", "Info",
        "Level of logging done by the application. Higher level means less log messages.\n" +
        "Set logging to desired level:" +
        "\n\t0. Debug\n\t1. Info\n\t2. Warn\n\t3. Error\n\t4. Silent - no logging output\n")
    useOnConsole := flag.Bool("console", false,
        "Output logs to console instead of file")
    useElastic := flag.Bool("elastic", false,
        "Enable additional logging output to elasticsearch")
    elasticAddress := flag.String("elastic-address", "http://localhost:9200",
        "URL used for elasticsearch connection")
    elasticIndexName := flag.String("index-name", "cr-telemetry-ipmctl-exporter",
        "Index name used/created in elasticsearch")
    flag.Parse()
    return *port, *enableThresholds, *showVersion, *loggingLevel, *useOnConsole, *useElastic, *elasticAddress, *elasticIndexName
}

func handleSIGINT() {
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)
    go func() {
        for sig := range signals {
            collector.Stop()
            fmt.Printf("ipmctl exporter - catch %s, stopping service\n", sig)
            log.Debug("ipmctl exporter caught signal ", sig, " stopping service")
            os.Exit(1)
        }
    }()
}

func main() {
    port, enableThresholds, showVersion, loggingLevel, logOnConsole, elasticUsed, elasticAddress, indexName := parseCmdArgs()
    setupLogger(loggingLevel, logOnConsole, elasticUsed, elasticAddress, indexName)
    if showVersion {
        fmt.Printf("%s\n", Version)
        os.Exit(0)
    }
    handleSIGINT()
    log.Debug("Ipmctl_exporter version: ", Version)
    log.Debug("Ipmctl exporter listening port: ", port)
    fmt.Printf("ipmctl exporter listening on port :%s\n", port)
    collector.Version = Version
    collector.Run(port, enableThresholds)
}
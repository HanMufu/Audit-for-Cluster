package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"audit-cluster/logic"
	"audit-cluster/settings"
	"audit-cluster/logger"
	"go.uber.org/zap"
)

var (
	fs = flag.NewFlagSet("audit", flag.ExitOnError)
)

func main() {
	// Load config files
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settings failed, err:%v\n", err)
		return
	}
	// Init uber/zap logger
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("Init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Info("logger init success")
	fs.Parse(os.Args[1:])
	if err := logic.Read(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

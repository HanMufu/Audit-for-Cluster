package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"audit-cluster/logic"
	"audit-cluster/settings"
	"audit-cluster/neo4jdb"
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
	// Init neo4j database
	if err := neo4jdb.Init(settings.Conf.Neo4jConfig); err != nil {
		fmt.Printf("Init neo4j failed, err:%v\n", err)
		return
	}
	// res, err := neo4jdb.TestConnection()
	// if err != nil {
	// 	fmt.Printf("Test connection to neo4j failed, err:%v\n", err)
	// 	return
	// }
	// zap.L().Info("Greetings from neo4j: ", zap.String("greeting", res))
	defer neo4jdb.Close()
	fs.Parse(os.Args[1:])
	if err := logic.Read(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

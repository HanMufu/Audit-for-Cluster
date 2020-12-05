package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"audit-cluster/logic"
	"audit-cluster/settings"
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
	fs.Parse(os.Args[1:])
	if err := logic.Read(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

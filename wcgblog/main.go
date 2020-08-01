package main

import (
	"flag"
	"wcgblog/controller"
	"wcgblog/service"
	"wcgblog/system"

	"github.com/cihub/seelog"
)

func main() {
	configFilePath := flag.String("C", "conf/conf.yaml", "config file path")
	logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")
	flag.Parse()
	logger, err := seelog.LoggerFromConfigAsFile(*logConfigPath)
	if err != nil {
		seelog.Critical("err parsing seelog config file", err)
		return
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()
	if err := system.LoadConfiguration(*configFilePath); err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	service.ConnectDB()
	defer service.DisconnectDB()
	router := controller.MapRoutes()
	router.Run(system.GetConfiguration().Port)
}

package main

import (
	"flag"

	"github.com/jrmConduce/redishappy"
	"github.com/jrmConduce/redishappy/configuration"
	"github.com/jrmConduce/redishappy/services/logger"
)

var configFile string
var logPath string

func init() {
	flag.StringVar(&configFile, "config", "config.json", "Full path of the configuration JSON file.")
	flag.StringVar(&logPath, "log", "log", "Folder for the logging folder.")
}

func main() {

	flag.Parse()

	logger.InitLogging(logPath)

	config, err := configuration.LoadFromFile(configFile)

	if err != nil {
		logger.Error.Panicf("Error opening config file : %s", err.Error())
	}

	sane, errors := config.GetCurrentConfiguration().SanityCheckConfiguration(
		&configuration.ConfigContainsRequiredSections{},
		&configuration.CheckForObviousMisConfiguration{})

	if !sane {

		for _, errorAsStr := range errors {
			logger.Error.Print(errorAsStr)
		}

		logger.Error.Panic("Configuration fails checks")
	}

	flipper := NewConsulFlipperClient(config)

	engine := redishappy.NewRedisHappyEngine(flipper, config)
	engine.Serve()
}

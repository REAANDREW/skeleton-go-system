package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/imdario/mergo"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type ArgumentParser interface {
	Parse() Configuration
}

type ArgumentParserFactory interface {
	Build(configuration Configuration) ArgumentParser
}

type KingpinArgumentParser struct {
	configuration Configuration
}

func (parser KingpinArgumentParser) Parse() Configuration {
	var versionString = fmt.Sprintf(`Version: %s BuildTime: %s CommitHash: %s`, Version, BuildTime, CommitHash)
	var configFile = kingpin.Flag("config", "Configuration File").Short('c').File()
	var configuration = parser.configuration

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel != "" {
		configuration.LogLevel = logLevel
	}

	if configFile != nil {
		var userConfiguration Configuration

		var data, _ = ioutil.ReadAll(bufio.NewReader(*configFile))
		json.Unmarshal(data, &userConfiguration)
		mergo.MergeWithOverwrite(&configuration, userConfiguration)
	}

	kingpin.Version(versionString)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	return configuration
}

type KingpinArgumentParserFactory struct {
}

func (factory KingpinArgumentParserFactory) Build(configuration Configuration) ArgumentParser {
	return KingpinArgumentParser{
		configuration: configuration,
	}
}

func NewKingpinArgumentParser() KingpinArgumentParserFactory {
	return KingpinArgumentParserFactory{}
}

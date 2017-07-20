package main

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger interface {
	Debug(fields ...interface{})
	Info(fields ...interface{}) Warn(fields ...interface{})
	Error(fields ...interface{})
	Fatal(fields ...interface{})
}

type LoggerFactory interface {
	Build(config Configuration) Logger
}

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func (log ZapLogger) Debug(fields ...interface{}) {
	fmt.Println("DEBUGGING")
	log.logger.Debug(fields)
}

func (log ZapLogger) Error(fields ...interface{}) {
	log.logger.Error(fields)
}

func (log ZapLogger) Warn(fields ...interface{}) {
	log.logger.Warn(fields)
}

func (log ZapLogger) Info(fields ...interface{}) {
	log.logger.Info(fields)
}

func (log ZapLogger) Fatal(fields ...interface{}) {
	log.logger.Fatal(fields)
}

type ZapLoggerFactory struct {
}

func (factory ZapLoggerFactory) Build(config Configuration) Logger {
	var loggerConfig = zap.NewProductionConfig()
	var atomicLevel = zap.NewAtomicLevel()

	atomicLevel.SetLevel(zap.DebugLevel)

	loggerConfig.Level = atomicLevel
	var logger, _ = loggerConfig.Build()
	return ZapLogger{
		logger: logger.Sugar(),
	}
}

func NewZapLogger() ZapLoggerFactory {
	return ZapLoggerFactory{}
}

func Build() {

}

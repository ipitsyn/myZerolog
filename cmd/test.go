package main

import (
	"sync"
	"time"

	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
	"go.uber.org/zap"

	"github.com/ipitsyn/myzap"
	log "github.com/ipitsyn/myzerolog"
)

var logger = log.MyNewConsoleLogger()
var zapSugar = myzap.New(zap.DebugLevel).Sugar

func main() {
	log.AdjustCallerWidth(20)
	logger.AdjustCallerWidth(20)
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	log.Trace("Trace test (global): ", 123)
	log.Tracef("Trace test (global) formatted: %06.2f", 1.23)

	logger.Trace("Trace test (logger): ", 123)
	logger.Tracef("Trace test (logger) formatted: %06.2f", 1.23)

	log.Debug("Debug test (global): ", 123)
	log.Debugf("Debug test (global) formatted: %06.2f", 1.23)

	logger.Debug("Debug test (logger): ", 123)
	logger.Debugf("Debug test (logger) formatted: %06.2f", 1.23)

	log.Info("Info test (global): ", 123)
	log.Infof("Info test (global) formatted: %06.2f", 1.23)

	logger.Info("Info test (logger): ", 123)
	logger.Infof("Info test (logger) formatted: %06.2f", 1.23)

	log.Warn("Warn test (global): ", 123)
	log.Warnf("Warn test (global) formatted: %06.2f", 1.23)

	logger.Warn("Warn test (logger): ", 123)
	logger.Warnf("Warn test (logger) formatted: %06.2f", 1.23)

	log.Error("Error test (global): ", 123)
	log.Errorf("Error test (global) formatted: %06.2f", 1.23)

	logger.Error("Error test (logger): ", 123)
	logger.Errorf("Error test (logger) formatted: %06.2f", 1.23)

	log.Fatal("Fatal test (global): ", 123)
	log.Fatalf("Fatal test (global) formatted: %06.2f", 1.23)

	logger.Fatal("Fatal test (logger): ", 123)
	logger.Fatalf("Fatal test (logger) formatted: %06.2f", 1.23)

	log.Panic("Panic test (global): ", 123)
	log.Panicf("Panic test (global) formatted: %06.2f", 1.23)

	logger.Panic("Panic test (logger): ", 123)
	logger.Panicf("Panic test (logger) formatted: %06.2f", 1.23)

	zl.Info().Msg("Test message")

	var resZL float64
	var resZap float64

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			log.Sync()
			wg.Done()
		}()
		resZL = testMyZerolog()
	}()

	wg.Add(1)
	go func() {
		defer func() {
			zapSugar.Sync()
			wg.Done()
		}()
		resZap = testZapSugar()
	}()

	wg.Wait()
	log.Warnf("myZerolog speed: %.3f operations per second", resZL)
	log.Warnf("ZapSugar speed: %.3f operations per second", resZap)
	log.Sync()

	runRestful(7890)
}

func testMyZerolog() float64 {
	nOps := 1000000

	start := time.Now()
	for i := 0; i < nOps; i++ {
		log.Debug("MyZerolog test, line no.: ", i)
	}
	stop := time.Now()
	return (float64(nOps) / float64(stop.Sub(start))) * float64(time.Second)
}

func testZapSugar() float64 {
	nOps := 1000000

	start := time.Now()
	for i := 0; i < nOps; i++ {
		zapSugar.Debug("ZapSugar test, line no.: ", i)
	}
	stop := time.Now()
	return (float64(nOps) / float64(stop.Sub(start))) * float64(time.Second)
}

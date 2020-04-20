package main

import (
	"fmt"
	"github.com/asut-inc/go-rest/client"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	atom := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	logger, err := newLogger(atom)
	if err != nil {
		fmt.Printf("failed to create logger: %s\n", err)
		os.Exit(2)
	}
	logger = logger.Named("app")
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("error when calling logger sync", zap.Error(err))
		}
	}()

	httpClient, err := client.NewClient(client.WithHTTPClient(nil))
	if err != nil {
		logger.Error("http client error", zap.Error(err))
	}
	_ = httpClient

	args := []string{""}
	err = run(args, logger)
}

func newLogger(atom zap.AtomicLevel) (*zap.Logger, error) {
	conf := zap.NewProductionConfig()
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	conf.Level = atom
	logger, err := conf.Build()

	if err != nil {
		return nil, err
	}

	return logger, err
}

func run(args []string, logger *zap.Logger) error {
	return nil
}

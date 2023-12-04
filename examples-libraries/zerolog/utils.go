package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitializeLogger(filePath string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	var multi zerolog.LevelWriter

	if filePath == "" {
		multi = zerolog.MultiLevelWriter(consoleWriter)
	} else {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to open log file")
			multi = zerolog.MultiLevelWriter(consoleWriter)
		} else {
			multi = zerolog.MultiLevelWriter(consoleWriter, file)
		}
	}

	logger := zerolog.New(multi).With().Timestamp().Logger()
	return logger
}

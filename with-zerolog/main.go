package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func init() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

func main() {
}

package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type event struct {
	Type string      `json:"type"`
	Time time.Time   `json:"time"`
	Data interface{} `json:"data"`
}

func main() {
	args := os.Args[1:]
	interval, err := strconv.Atoi(args[0])
	if err != nil {
		interval = 5
	}

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	l, _ := config.Build()
	defer l.Sync()

	// Generating logs indefinitely.
	for range time.Tick(time.Duration(interval) * time.Second) {
		ascii := rand.Intn(26)
		e := event{
			Type: string(ascii + 97), // Event type generated ranges from "a" to "z"
			Time: time.Now().UTC(),
			Data: map[string]interface{}{
				"id":        (ascii % 10) + 1, // Maximum of 10 different entity IDs.
				"uppercase": string(ascii + 65),
				"random":    rand.Intn(100) + 1,
			},
		}

		l.Info("test message", zap.Any("event", e))
	}
}

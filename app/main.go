package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/dchest/uniuri"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	maxRandomNumber     = 100
	maxRandomInProgress = 10
	maxRandomEventTimer = 10
)

func main() {
	rand.Seed(time.Now().UnixNano())

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
		go generateLog(l)
	}
}

func generateLog(l *zap.Logger) {
	id := uuid.New().String()

	e := event{
		Type: "event_started",
		Time: time.Now().UTC(),
		Data: eventStartedData{
			ID:      id,
			Message: uniuri.New(),
		},
	}
	l.Info("event started", zap.Any("event", e))
	time.Sleep(time.Duration(rand.Intn(maxRandomEventTimer)+1) * time.Second)

	numInProgress := rand.Intn(maxRandomInProgress) + 1
	for i := 0; i < numInProgress; i++ {
		e = event{
			Type: "event_in_progress",
			Time: time.Now().UTC(),
			Data: eventInProgressData{
				ID:     id,
				Random: rand.Intn(maxRandomNumber) + 1,
			},
		}
		l.Info("event in progress", zap.Any("event", e))
		time.Sleep(time.Duration(rand.Intn(maxRandomEventTimer)+1) * time.Second)
	}

	e = event{
		Type: "event_finished",
		Time: time.Now().UTC(),
		Data: eventFinishedData{
			ID:      id,
			Message: uniuri.New(),
		},
	}
	l.Info("event finished", zap.Any("event", e))
}

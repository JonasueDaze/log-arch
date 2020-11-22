package main

import "time"

type event struct {
	Type string      `json:"type"`
	Time time.Time   `json:"time"`
	Data interface{} `json:"data"`
}

type eventStartedData struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type eventInProgressData struct {
	ID     string `json:"id"`
	Random int    `json:"random"`
}

type eventFinishedData struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

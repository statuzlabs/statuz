package main

import (
	"context"
	"log"
	"time"

	"github.com/Unfield/Statuz/monitors"
	"github.com/Unfield/Statuz/scheduler"
)

func main() {
	monitor := monitors.HTTPMonitor{
		ID:                  "test-monitor",
		URL:                 "https://www.speedtest.net/",
		HBInterval:          10 * time.Second,
		Retries:             3,
		RetryInterval:       2 * time.Second,
		ReqTimeout:          5,
		MaxRedirects:        0,
		AcceptedStatusCodes: []int{200},
		IPFamily:            "auto",
		HTTPMethod:          "GET",
	}

	monitor2 := monitors.HTTPMonitor{
		ID:                  "test-monitor2",
		URL:                 "https://www.google.com/",
		HBInterval:          60 * time.Second,
		Retries:             3,
		RetryInterval:       2 * time.Second,
		ReqTimeout:          5,
		MaxRedirects:        0,
		AcceptedStatusCodes: []int{200},
		IPFamily:            "auto",
		HTTPMethod:          "GET",
	}

	_ = monitor2

	monitorsList := []monitors.Monitor{&monitor, &monitor2}

	s := scheduler.NewScheduler(context.Background(), monitorsList)
	go s.Run()

	for {
		select {
		case result, ok := <-s.ResultChannel:
			if !ok {
				log.Println("Result channel closed")
				return
			}

			log.Printf("[%s] %s → %s (success=%t) (duration=%d)\n",
				result.MonitorID,
				result.Type,
				result.Status,
				result.Success,
				result.Duration.Milliseconds(),
			)

		case <-s.Context().Done():
			log.Println("Scheduler context canceled, stopping listener.")
			return
		}
	}
}

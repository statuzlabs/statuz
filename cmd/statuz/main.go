package main

import (
	"fmt"
	"time"

	"github.com/Unfield/Statuz/monitors"
)

func main() {
	montior := monitors.HTTPMonitor{
		ID:                  "test-monitor",
		URL:                 "https://www.speedtest.net/",
		HBInterval:          60,
		Retries:             3,
		RetryInterval:       2 * time.Second,
		ReqTimeout:          5,
		MaxRedirects:        0,
		AcceptedStatusCodes: []int{200},
		IPFamily:            "auto",
		HTTPMethod:          "GET",
	}

	res := montior.Check()

	fmt.Println(res.Status)
	fmt.Println(res.Duration)
}

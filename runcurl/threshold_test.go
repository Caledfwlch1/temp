package runcurl

import (
	"testing"
	"time"
)

func TestThreshold(t *testing.T) {
	//RunCurl("http://35.190.167.185/custom?branch=dev&commit=e43c8b290fb575808b8aafb680a45eeeb2221282&roles=threshold&wait=true")
	time.Sleep(time.Minute * 11)
	RunCurl("http://localhost/")
	return
}
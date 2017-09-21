package runcurl

import (
	"testing"
)

func TestBuilds(t *testing.T) {
	RunCurl("http://35.190.167.185/custom?branch=dev&commit=6c910dffc96d6a4d0c5a32fc6f53bc4f9d050570&roles=builds&wait=true")
	return
}
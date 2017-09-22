package runcurl

import (
	"testing"
)

func TestDbdelays(t *testing.T) {
	//RunCurl("http://35.190.167.185/custom?branch=dev&commit=6c910dffc96d6a4d0c5a32fc6f53bc4f9d050570&roles=dbdelays&wait=true")
	RunCurl("http://localhost/")
	return
}
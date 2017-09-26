package runcurl

import (
	"testing"
)

func TestFailed_nodes(t *testing.T) {
	//RunCurl("http://35.190.167.185/custom?branch=dev&commit=e43c8b290fb575808b8aafb680a45eeeb2221282&roles=failed_nodes&wait=true")
	RunCurl("http://localhost/")
	return
}
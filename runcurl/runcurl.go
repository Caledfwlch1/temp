package runcurl

import (
	"flag"
	"net/http"
	"log"
)


func RunCurl(getString string) {

	resp, err := http.Get(getString)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln(resp.Status)
	}

	log.Println(resp.Status)

	return
}


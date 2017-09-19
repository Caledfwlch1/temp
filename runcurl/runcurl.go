package runcurl

import (
	"net/http"
	"log"
	"io"
)


func RunCurl(getString string) {

	resp, err := http.Get(getString)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		printResp(resp.Body)
		log.Fatalln(resp.Status)
	}

	log.Println(resp.Status)

	return
}

func printResp(r io.ReadCloser) {
	p := make([]byte, 10 * 1024)
	for {
		l, err := r.Read(p)
		log.Println(l)
		if err != nil {
			break
		}
	}
	return
}
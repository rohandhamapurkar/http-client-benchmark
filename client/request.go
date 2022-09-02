package client

import (
	"bytes"
	"log"
	"net/http"
)

const endpoint = "http://127.0.0.1:3333"

func GetRequest() *http.Request {
	req, err := http.NewRequest(http.MethodGet, endpoint, bytes.NewReader([]byte{}))

	if err != nil {
		log.Fatalln("Could not get request")
	}

	return req
}

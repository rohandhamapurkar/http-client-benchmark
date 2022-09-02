package client

import (
	"log"
	"net/http"
	"time"
)

func DefaultClientRequest() {
	// using default client
	var httpClient = &http.Client{}

	response, err := httpClient.Do(GetRequest())
	if err != nil {
		log.Println(err)
	}

	defer func() {
		if closeErr := response.Body.Close(); closeErr != nil {
			log.Println(err, "closing response body: "+closeErr.Error())
		}
	}()

}

func OverridingDefaultClientRequest() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	response, err := httpClient.Do(GetRequest())
	if err != nil {
		log.Println(err)
	}

	defer func() {
		if closeErr := response.Body.Close(); closeErr != nil {
			log.Println(err, "closing response body: "+closeErr.Error())
		}
	}()

}

func OverridingDefaultClientRequestWithIdleConnTimeout() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	t.IdleConnTimeout = 30

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	response, err := httpClient.Do(GetRequest())
	if err != nil {
		log.Println(err)
	}

	defer func() {
		if closeErr := response.Body.Close(); closeErr != nil {
			log.Println(err, "closing response body: "+closeErr.Error())
		}
	}()

}

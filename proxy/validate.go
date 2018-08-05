package proxy

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

func IsValidProxy(checkURL, proxy string, expectCode int) bool {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return false
	}

	transport := http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	myClient := &http.Client{Transport: &transport}
	myClient.Timeout = 10 * time.Second
	resp, err := myClient.Get(checkURL)
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode == expectCode {
		return true
	}
	return false
}

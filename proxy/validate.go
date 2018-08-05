package proxy

import (
	"net/http"
	"net/url"
	"time"

	"github.com/weaming/golib/debug"
)

func IsValidProxy(checkURL, proxy string, expectCode, timeout int) bool {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		debug.Debug(err)
		return false
	}

	transport := http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	myClient := &http.Client{Transport: &transport}
	myClient.Timeout = time.Duration(timeout) * time.Second
	resp, err := myClient.Get(checkURL)
	if err != nil {
		debug.Debug(err)
		return false
	}
	if resp.StatusCode == expectCode {
		return true
	}
	return false
}

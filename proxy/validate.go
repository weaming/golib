package proxy

import (
	"net/http"
	"net/url"
	"time"

	"github.com/weaming/golib/debug"
)

func IsValidProxy(checkURLStr, proxy string, expectCode, timeout int) bool {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		debug.Debug(err)
		return false
	}

	checkURL, err := url.Parse(checkURLStr)
	if err != nil {
		debug.Debug(err)
		return false
	}

	transport := http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	myClient := &http.Client{Transport: &transport}
	myClient.Timeout = time.Duration(timeout) * time.Second
	resp, err := myClient.Get(checkURL.String())
	if err != nil {
		debug.Debug(err)
		return false
	}
	if resp.StatusCode == expectCode {
		return true
	}
	return false
}

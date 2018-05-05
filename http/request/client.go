package request

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

func NewHTTPClient(timeout time.Duration) *http.Client {
	// https://golang.org/src/net/http/transport.go
	/*
		// Default transport uses HTTP proxies as directed by the $HTTP_PROXY and $NO_PROXY
		// (or $http_proxy and $no_proxy) environment variables.
		var DefaultTransport RoundTripper = &Transport{
			Proxy: ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
	*/

	tr := &http.Transport{
		// MaxIdleConnsPerHost, if non-zero, controls the maximum idle (keep-alive) connections to keep per-host.
		// If zero, DefaultMaxIdleConnsPerHost is used, whose value is 2.
		MaxIdleConnsPerHost: 1024,
		// MaxIdleConns controls the maximum number of idle (keep-alive) connections across all hosts.
		// Zero means no limit.
		MaxIdleConns: 0,
	}
	return &http.Client{
		Transport: tr,
		Timeout:   timeout * time.Second,
	}
}

func SortedQueryString(m map[string]interface{}) (rv string) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if i > 0 {
			rv += "&"
		}
		rv += fmt.Sprintf("%v=%v", k, m[k])
	}
	return
}

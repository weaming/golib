package http

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

func NewHTTPClient(timeout time.Duration) *http.Client {
	// Default transport uses HTTP proxies as directed by the $HTTP_PROXY and $NO_PROXY
	// (or $http_proxy and $no_proxy) environment variables.
	// var DefaultTransport RoundTripper = &Transport{
	//         Proxy: ProxyFromEnvironment,
	//         DialContext: (&net.Dialer{
	//                 Timeout:   30 * time.Second,
	//                 KeepAlive: 30 * time.Second,
	//                 DualStack: true,
	//         }).DialContext,
	//         MaxIdleConns:          100,
	//         IdleConnTimeout:       90 * time.Second,
	//         TLSHandshakeTimeout:   10 * time.Second,
	//         ExpectContinueTimeout: 1 * time.Second,
	// }

	// DefaultMaxIdleConnsPerHost is the default value of Transport's MaxIdleConnsPerHost:
	// const DefaultMaxIdleConnsPerHost = 2

	tr := &http.Transport{
		MaxIdleConnsPerHost: 1024,
		MaxIdleConns:        1024,
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

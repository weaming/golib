package main

import (
	"flag"
	"os"

	proxylib "github.com/weaming/golib/proxy"
)

func main() {
	url := flag.String("url", "https://www.baidu.com", "target url to test using proxy")
	proxy := flag.String("proxy", "http://localhost:8123", "proxy url")
	code := flag.Int("code", 200, "expected response status code")
	timeout := flag.Int("timeout", 3, "timeout in seconds")
	flag.Parse()

	if proxylib.IsValidProxy(*url, *proxy, *code, *timeout) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

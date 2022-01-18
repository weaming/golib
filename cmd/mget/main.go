package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/weaming/golib/fs"
	"github.com/weaming/golib/http/downloader"
)

func main() {
	url := flag.String("url", "", "url to download")
	flag.Parse()

	if *url == "" {
		fmt.Println("missing -url")
		os.Exit(1)
	}
	now := time.Now()
	data, e := downloader.Download(*url, 8)
	if e != nil {
		log.Println("err:", e)
		os.Exit(2)
	}
	name := filepath.Base(*url)
	ioutil.WriteFile(name, data, 0644)
	fmt.Printf("downloaded as \"%v\" with size %v, cost %v\n", name, fs.HumanSize(uint64(len(data))), time.Now().Sub(now))
}

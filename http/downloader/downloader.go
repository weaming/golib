package downloader

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

// ctx 可以为 nil
// start，end 都是闭区间
// end 可以为 0 表示下载 start 往后所有的
func HTTPGetByRange(ctx context.Context, url string, start, end uint64) ([]byte, error) {
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return nil, e
	}

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range
	if end == 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%v-", start))
	} else {
		req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", start, end))
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		if res != nil {
			defer res.Body.Close()
		}
		return nil, e
	}

	return ioutil.ReadAll(res.Body)
}

func HTTPGetByParts(url string, size uint64, N int) ([]byte, error) {
	partSize := uint64(math.Ceil(float64(size) / float64(N)))

	// context 完成的事件没有被利用起来取消剩余操作
	g, ctx := errgroup.WithContext(context.Background())
	parts := make([][]byte, N)
	lock := sync.Mutex{}
	for i := 0; i < N; i++ {
		start := partSize * uint64(i)
		end := start + partSize
		if end >= size {
			end = size - 1
		}

		go func(i int, start, end uint64) {
			g.Go(func() error {
				bs, e := HTTPGetByRange(ctx, url, start, end)
				if e != nil {
					return e
				}

				lock.Lock()
				defer lock.Unlock()
				parts[i] = bs
				return nil
			})
		}(i, start, end)
	}

	if e := g.Wait(); e != nil {
		return nil, e
	}

	data := []byte{}
	for _, bs := range parts {
		data = append(data, bs...)
	}
	return data, nil
}

func HTTPGetSize(url string) (uint64, error) {
	res, e := http.Head(url)
	if e != nil {
		if res != nil {
			defer res.Body.Close()
		}
		return 0, e
	}
	cl := res.Header.Get("Content-Length")
	if cl == "" {
		return 0, nil
	}
	return strconv.ParseUint(cl, 10, 64)
}

func Download(url string, N int) ([]byte, error) {
	size, e := HTTPGetSize(url)
	if e != nil {
		return nil, e
	}
	if size == 0 {
		return HTTPGetByRange(nil, url, 0, 0)
	}
	return HTTPGetByParts(url, size, N)
}

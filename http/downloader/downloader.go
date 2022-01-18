package downloader

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/weaming/golib/fs"
	"golang.org/x/sync/errgroup"
)

// ctx 可以为 nil
// start，end 都是闭区间
// end 可以为 0 表示下载 start 往后所有的
func HTTPGetByRange(ctx context.Context, url string, start, end uint64, updates chan<- int) ([]byte, error) {
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
		return nil, e
	}
	defer res.Body.Close()

	if end != 0 {
		data := []byte{}
		buf := make([]byte, 1024)
		for {
			n, e := res.Body.Read(buf)

			sizeNeeds := end - start + 1
			// 检查下载的大小是否超出需要下载的大小
			// 数据大小不正常，一般是因为网络环境不好导致，比如用中国电信下载国外文件
			if n > int(sizeNeeds) {
				// 设置数据大小来去掉多余数据
				// 并结束这个协程的下载
				n = int(sizeNeeds)
				e = io.EOF
			}

			if n > 0 {
				// 将缓冲数据写入
				data = append(data, buf[:n]...)
				// 更新已下载大小
				updates <- n
			}

			if e != nil {
				if e == io.EOF {
					// 数据已经下载完毕
					return data, nil
				}
				return nil, e
			}
		}
	}

	// 一次性读取
	bs, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}

	updates <- len(bs)
	return bs, e
}

type Part struct {
	sync.Mutex
	Content []byte
}

func HTTPGetByParts(url string, size uint64, N int, updates chan<- int) ([]byte, error) {
	partSize := uint64(math.Ceil(float64(size) / float64(N)))

	// context 完成的事件没有被利用起来取消剩余操作
	g, ctx := errgroup.WithContext(context.Background())
	parts := make([]Part, N)
	for i := 0; i < N; i++ {
		start := partSize * uint64(i)
		end := start + partSize
		if end >= size {
			end = size - 1
		}

		func(i int, start, end uint64) {
			g.Go(func() error {
				bs, e := HTTPGetByRange(ctx, url, start, end, updates)
				if e != nil {
					return e
				}

				parts[i].Content = bs
				return nil
			})
		}(i, start, end)
	}

	if e := g.Wait(); e != nil {
		return nil, e
	}

	data := []byte{}
	for _, bs := range parts {
		data = append(data, bs.Content...)
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
	fmt.Printf("Total size is %v\n", size)

	updates := make(chan int)
	defer close(updates)
	if size == 0 {
		return HTTPGetByRange(nil, url, 0, 0, updates)
	}
	go func() {
		total := 0
		before := time.Now()
		for {
			change, ok := <-updates
			if ok {
				total += change
				percentChange := float64(uint64(total)*10000/size) / 100 // 2 位小数
				now := time.Now()
				if now.Sub(before).Seconds() >= 1 {
					fmt.Printf("\r=> %.2f%% %v   ", percentChange, fs.HumanSize(uint64(total)))
					before = now
				}
			} else {
				fmt.Printf("\n")
				return
			}
		}
	}()
	return HTTPGetByParts(url, size, N, updates)
}

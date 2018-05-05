package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/weaming/golib/database/influx"
)

const (
	dbhost = "http://localhost:8086"
	dbname = "machine"
)

func main() {
	c := influx.New(influx.Config{
		Addr: dbhost,
	})
	c.CreateDB(dbname)
	// c.Use(dbname) // optional after create

	hi, err := host.Info()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(1000 * time.Millisecond)
	for {
		bp, _ := c.NewBatch()

		n := 1
		for _ = range ticker.C {
			n++

			v, _ := mem.VirtualMemory()
			bp.Add(influx.Point{
				Name: "memory",
				Tags: influx.Tags{
					"os": hi.OS,
				},
				Fields: influx.Fields{
					"free":    float64(v.Free),
					"percent": float64(v.UsedPercent),
				},
			})
			if n%10 == 0 {
				if err = bp.Write(c); err != nil {
					panic(err)
				}
				fmt.Println("writing...")
				break
			}
		}
	}
}

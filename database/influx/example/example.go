package main

import (
	"github.com/weaming/golib/database/influx"
)

const (
	dbhost = "http://localhost:8086"
	dbname = "test"
)

func main() {
	c := influx.New(influx.Config{
		Addr: dbhost,
	})
	c.CreateDB(dbname)
	// c.Use(dbname) // optional after create

	bp, err := c.NewBatch()
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		bp.Add(influx.Point{
			Name: "hi",
			Tags: influx.Tags{
				"type": "A",
			},
			Fields: influx.Fields{
				"cpu": 29,
			},
		})
	}
	bp.Write(c)
}

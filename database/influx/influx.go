package influx

import (
	"fmt"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
)

type Client struct {
	influx.Client
	addr   string
	dbname string
}

func New(config influx.HTTPConfig) *Client {
	client, err := influx.NewHTTPClient(config)

	if err != nil {
		panic(err)
	}

	return &Client{
		Client: client,
		addr:   config.Addr,
	}
}

func (c *Client) CreateDB(dbname string) error {
	q := influx.NewQuery("CREATE DATABASE "+dbname, "", "")
	_, err := c.Query(q)
	if err == nil {
		c.Use(dbname)
	}
	return err
}

func (c *Client) NewBatch() (*Batch, error) {
	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database: c.dbname,
		//Precision: "s", // timestamp precision
	})

	if err != nil {
		return nil, err
	}

	return &Batch{bp}, err
}

func (c *Client) Use(dbname string) {
	c.dbname = dbname
}

func (c *Client) String() string {
	// stringer interface
	return fmt.Sprintf("<%v %v>", c.addr, c.dbname)
}

// Batch of points.
type Batch struct {
	influx.BatchPoints
}

// Tags of a point.
type Tags map[string]string

// Fields of a point.
type Fields map[string]interface{}

// Point is a single point.
type Point struct {
	Name   string
	Tags   Tags
	Fields Fields
}

// Add point.
func (b *Batch) Add(p Point) error {
	point, err := influx.NewPoint(p.Name, p.Tags, p.Fields, time.Now())
	if err != nil {
		return err
	}

	b.AddPoint(point)
	return nil
}

func (b *Batch) Write(client *Client) error {
	return client.Write(b)
}

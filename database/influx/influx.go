package influx

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
)

type Client struct {
	influx.Client
	addr   string
	dbname string
}

type Config struct {
	// Addr should be of the form "http://host:port"
	// or "http://[ipv6-host%zone]:port".
	Addr string

	// Username is the influxdb username, optional.
	Username string

	// Password is the influxdb password, optional.
	Password string

	// UserAgent is the http User Agent, defaults to "InfluxDBClient".
	UserAgent string

	// Timeout for influxdb writes, defaults to no timeout.
	Timeout time.Duration

	// InsecureSkipVerify gets passed to the http client, if true, it will
	// skip https certificate verification. Defaults to false.
	InsecureSkipVerify bool

	// TLSConfig allows the user to set their own TLS config for the HTTP
	// Client. If set, this option overrides InsecureSkipVerify.
	TLSConfig *tls.Config

	// Proxy configures the Proxy function on the HTTP client.
	Proxy func(req *http.Request) (*url.URL, error)
}

func New(config Config) *Client {
	client, err := influx.NewHTTPClient(influx.HTTPConfig(config))

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

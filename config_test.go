package main

import (
	"testing"
)

var data = `
http:
  host: "127.0.0.1"
  port: 8080
carbon:
  host: "127.0.0.1"
  port: 2003
  rollups:
    - name: application metrics
      pattern: ^app\..*
      retentions: 
        - 10s:6h
        - 60s:7d
        - 1h:180d
        - 1d:5y
    - name: user status metrics
      pattern: ^app\..*\.users\..*\.status\..*
      retentions:
        - 10s:7d
        - 60s:30d
  aggregations:
    - name: min
      pattern: \.min$
      aggregationMethod: min
    - name: max
      pattern: \.max$
      aggregationMethod: max
    - name: avg
      pattern: \.avg$
      aggregationMethod: average
    - name: last
      pattern: \.current$
      aggregationMethod: last
    - name: status
      pattern: \.status$
      aggregationMethod: last
    - name: default_sum
      pattern: .*
      aggregationMethod: sum
logging:
  level: info
  console: true
  files:
    - "/var/log/gneiss/gneiss.log"
    - "/tmp/gneiss.log"
store:
  cluster: 'localhost'
  keyspace: 'metric'
index:
  index: "my_paths"
  url: "http://myes.host.com:9200"
`

func TestParseConfig(t *testing.T) {
	c := Config{}
	c.Parse(data)

	if c.Http.Host != "127.0.0.1" {
		t.Fatal("Parse failed at Http.Host!")
	}

	if c.Carbon.Port != 2003 {
		t.Fatal("Parse failed at Carbon.Port!")
	}

	if c.Carbon.Rollups[0].Pattern != "^app\\..*" {
		t.Fatal("Parse failed at Carbon.Rollups[0].Pattern!")
	}

	if c.Carbon.Rollups[1].Retentions[1] != "60s:30d" {
		t.Fatal("Parse failed at Carbon.Rollups[1].Retentions[1]")
	}

	if c.Carbon.Aggregations[2].AggregationMethod != "average" {
		t.Fatal("Parse failed at Carbon.Aggregations[2].AggregationMethod")
	}

	if c.Logging.Console != true {
		t.Fatal("Parse failed at c.Logging.Console")
	}

	if c.Logging.Files[1] != "/tmp/gneiss.log" {
		t.Fatal("Parse failed at c.Logging.Files[1]")
	}

	if c.Store.Cluster != "localhost" {
		t.Fatal("Parse failed at c.Store.Cluster")
	}

	if c.Index.Url != "http://myes.host.com:9200" {
		t.Fatal("Parse failed at c.Index.Url")
	}
}

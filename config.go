package main

import (
	"gopkg.in/yaml.v1"
	"log"
)

type RollupConfig struct {
	Name       string
	Pattern    string
	Retentions []string
}

type AggregationConfig struct {
	Name              string
	Pattern           string
	AggregationMethod string "aggregationMethod"
}

type CarbonConfig struct {
	Host         string
	Port         int
	Rollups      []RollupConfig
	Aggregations []AggregationConfig
}

type HttpConfig struct {
	Host string
	Port int
}

type LoggingConfig struct {
	Level   string
	Console bool
	Files   []string
}

type StoreConfig struct {
	Cluster  string
	Keyspace string
}

type IndexConfig struct {
	Index string
	Url   string
}

type Config struct {
	Carbon  CarbonConfig
	Http    HttpConfig
	Logging LoggingConfig
	Store   StoreConfig
	Index   IndexConfig
}

func (c *Config) Parse(data string) {
	err := yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

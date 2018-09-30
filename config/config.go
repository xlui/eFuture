package config

import (
	"eFuture/common"
	"encoding/json"
	"io/ioutil"
	"os"
)

type configuration struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Smtp          string `json:"smtp"`
	SmtpPort      string `json:"smtp_port"`
	RedisAddress  string `json:"redis_address"`
	RedisPassword string `json:"redis_password"`
	RedisDB       int    `json:"redis_db"`
}

var Configuration configuration

// Read configurations from file. If specified `EFUTURE_CONFIG` environment variable,
// will read from the specified path, and if not specified, will try to read from
// `/data/eFuture/config.json`, make sure this file exist if you don't want to specified
// a environment variable!
func init() {
	config := os.Getenv("EFUTURE_CONFIG")
	if config == "" {
		config = "/data/eFuture/config.json"
	}
	bytes, e := ioutil.ReadFile(config)
	common.FailOnError(e, "Cannot read from file")
	json.Unmarshal(bytes, &Configuration)
}

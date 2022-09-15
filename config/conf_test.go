package config

import "testing"

func TestNewConfig(t *testing.T) {
	config := NewConfig("./test.yaml")
	if config.Mysql.User != "testdb" {
		t.Error("parse config error")
	}
}

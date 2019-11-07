package config

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	conf, err := LoadTom("./config.example.toml")
	if err != nil {
		t.Fatal(err.Error())
	}
	var e *Config

	if reflect.TypeOf(conf) != reflect.TypeOf(e) {
		t.Fatal("error type")
	}
	server := conf.Services["httpserver"]
	if server.Proto != "http" {
		t.Fatal("content error inside Toml")
	}
	if server.Hosts[0] != "localhost:8081" {
		t.Fatal("error parsing Hosts array")
	}

}

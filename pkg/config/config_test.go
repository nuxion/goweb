package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	conf, err := LoadTom("./config.example.toml")
	fmt.Println("hola test")
	if err != nil {
		t.Fatal(err.Error())
	}
	var e *Config

	if reflect.TypeOf(conf) != reflect.TypeOf(e) {
		t.Fatal("error type")
	}
	server := conf.Service[0]
	if server.Proto != "http" {
		t.Fatal("content error inside Toml")
	}
	if server.Hosts[0] != "localhost:8080" {
		t.Fatal("error parsing Hosts array")
	}
	if server.Name != "test" {
		t.Fatal("error parsing Name service")
	}

	general := conf.General
	if general.Port != "9090" {
		t.Fatal("error parsing general port")
	}

}

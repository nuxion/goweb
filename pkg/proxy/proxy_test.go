package proxy

import (
	"reflect"
	"testing"

	"github.com/nuxion/goweb/pkg/config"
)

func TestNewClient(t *testing.T) {
	hash := "g9XI5pH0Q-sGoaCZf0byjX9BGPk="
	c := NewClient("127.0.0.1", "Firefox")
	var expected *ClientID
	if hash != c.hash {
		t.Fatalf("Client hash doesn't match: Generated: %s, Expected: %s", c.hash, hash)
	}
	if reflect.TypeOf(c) != reflect.TypeOf(expected) {
		t.Fatalf("Wrong type of Client. Generated: %s, Expected: %s",
			reflect.TypeOf(c), reflect.TypeOf(expected))
	}

}

func TestPrepareUrls(t *testing.T) {
	conf, err := config.LoadTom("../config/config.example.toml")
	if err != nil {
		t.Fatal(err.Error())
	}
	service := conf.Services["httpserver"]
	urls := prepareUrls(&service)
	if len(urls) != 2 {
		t.Fatal("Bad parsed services")
	}
	h := urls[0].Host
	if h != "localhost:8081" {
		t.Error("Host doesn't match. Result: ", h)
	}

}

package proxy

import (
	"reflect"
	"testing"
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

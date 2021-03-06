package main

import (
	"fmt"
	"testing"

	"github.com/benschw/dns-clb-go/clb"
	"github.com/benschw/opin-go/rando"
)

func TestDiscover(t *testing.T) {
	//given
	expected := fmt.Sprintf("http://%s:8080", rando.MyIp())
	lb := clb.NewClb("localhost", "8600", clb.Random)

	// when
	found := Discover(lb, "test")

	// then
	if found != expected {
		t.Errorf("%s not equal to %s", found, expected)
	}
}

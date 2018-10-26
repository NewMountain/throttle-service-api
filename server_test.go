package main

import "testing"

func TestMkRt(t *testing.T) {
	newRoute := mkRt("/funk")
	if newRoute != "/api/throttling-service/v1/funk" {
		t.Errorf("mkRt working outside of specification")
	}
}

func TestMakeServer(t *testing.T) {
	server := makeServer()

	routes := server.Routes()

	// There should only be one route
	if len(routes) != 1 {
		t.Errorf("The number of routes is greater than anticipated")
	}

	// The only route should be
	onlyRoutePath := "/api/throttling-service/v1/throttle"
	if routes[0].Path != onlyRoutePath {
		t.Errorf("The only route is not matching the anticipated path.")
	}
}

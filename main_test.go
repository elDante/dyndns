package main

import (
	"testing"
)

var tests = map[string]map[string]string{
	"ohjahz7Loh0daid5Chiedeejiex8ahcu": map[string]string{
		"dns":  "foo.bar.com",
		"name": "foo",
	},
	"Uth0phuchai0itheish3Oogheyiethae": map[string]string{
		"dns":  "bar.foo.com",
		"name": "bar",
	},
}

func TestParseConfig(t *testing.T) {
	config := parseConfig("config.example.toml")

	for key := range tests {
		if tests[key]["dns"] != config.Clients[key].DNS {
			t.Error(
				"For", "config.Clients[key].DNS",
				"expected", tests[key]["dns"],
				"got", config.Clients[key].DNS,
			)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	for _, test := range []struct {
		fn          string
		expectError bool
	}{
		{".config.toml", false},
		{".nonsuch", true},
		{".", true},
	} {
		t.Run(test.fn, func(t *testing.T) {
			_, err := LoadConfig(test.fn)
			if err == nil && test.expectError {
				t.Errorf("expected error, received none")
			} else if err != nil && !test.expectError {
				t.Errorf("unexpected error %#v", err)
			}
		})
	}
}

func TestConfig_Topics(t *testing.T) {
	expect := []string{"atmosphere", "precipitation", "weather"}

	c, err := LoadConfig(".config.toml")
	if err != nil {
		t.Fatal(err)
	}

	rcvd := c.Topics()

	if !reflect.DeepEqual(expect, rcvd) {
		t.Errorf("expected %#v, received %#v", expect, rcvd)
	}
}

func TestConfig_Type(t *testing.T) {
	for _, test := range []struct {
		topic       string
		expectError bool
	}{
		{"atmosphere", false},
		{"precipitation", false},
		{"weather", false},
		{"", true},
		{"typo", true},
	} {
		t.Run(test.topic, func(t *testing.T) {
			c, err := LoadConfig(".config.toml")
			if err != nil {
				t.Fatal(err)
			}

			_, err = c.Type(test.topic)
			if err == nil && test.expectError {
				t.Errorf("expected error, received none")
			} else if err != nil && !test.expectError {
				t.Errorf("unexpected error %#v", err)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_fail(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Fail()
			if err == nil {
				t.Log(err)
				t.Fail()
			}
		})
	}
}

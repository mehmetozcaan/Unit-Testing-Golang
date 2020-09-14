package main

import (
	"testing"
	"time"
)

//takes about 4 seconds
func TestParallel(t *testing.T) {
	tc := []struct {
		dur time.Duration
	}{
		{time.Second},
		{2 * time.Second},
		{3 * time.Second},
		{4 * time.Second},
	}
	for _, tt := range tc {
		t.Run("", func(ts *testing.T) {
			ts.Parallel()
			time.Sleep(tt.dur)
		})
	}
}

//takes about 10 seconds
func TestNormal(t *testing.T) {
	tc := []struct {
		dur time.Duration
	}{
		{time.Second},
		{2 * time.Second},
		{3 * time.Second},
		{4 * time.Second},
	}
	for _, tt := range tc {
		t.Run("", func(ts *testing.T) {
			time.Sleep(tt.dur)
		})
	}
}

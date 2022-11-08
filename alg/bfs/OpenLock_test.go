package bfs_test

import (
	"algorithm/alg/bfs"
	"testing"
)

func TestOpenLock(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	expected := 6
	result := bfs.OpenLock(deadends, target)
	if expected != result {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}

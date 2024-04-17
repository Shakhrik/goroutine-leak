package main

import "testing"

func TestProcessRecords(t *testing.T) {
	recs := []string{"rec1", "rec2", "rec3"}
	processRecords(recs, 2)
}

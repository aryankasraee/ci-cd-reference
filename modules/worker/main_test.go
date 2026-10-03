package main

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	got := Process([]int{1, 2, 3})
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

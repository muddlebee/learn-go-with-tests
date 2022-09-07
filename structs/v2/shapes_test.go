package main

import "testing"

func TestPerimeter(t *testing.T) {
	object := Object{10.0, 10.0}
	got := Perimeter(object)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	object := Object{10.0, 10.0}
	got := Area(object)
	want := 100.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

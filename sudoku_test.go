package main

import (

	"testing"
)

func TestDupicated(t *testing.T) {
	if duplicate([10]int{
		0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	}) {
		t.Fatal("failed")
	}
}

func TestVerify(t *testing.T) {


	b := Borad{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	if !verify(b) {
		t.Fatal("to be false")
	}
}

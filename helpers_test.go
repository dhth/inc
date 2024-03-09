package main

import (
	"fmt"
	"testing"
)

func TestIsTagValid(t *testing.T) {

	tests := []struct {
		tag      string
		expected bool
	}{
		{"v0.0.0", true},
		{"v0.1.99", true},
		{"v0.10.1", true},
		{"v100.0.1", true},
		{"v999.999.999", true},
		{"999.999.999", true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%v", tt.tag, tt.expected)
		t.Run(testname, func(t *testing.T) {
			// given
			// when
			got, _ := isTagValid(tt.tag)
			// then
			if got != tt.expected {
				t.Errorf("got %v, expected %v", got, tt.expected)
			}
		})

	}
}
func TestGetSemVer(t *testing.T) {

	tests := []struct {
		v        string
		expected SemVer
	}{
		{"v0.0.1", SemVer{true, 0, 0, 1}},
		{"v0.1.2", SemVer{true, 0, 1, 2}},
		{"v1.2.3", SemVer{true, 1, 2, 3}},
		{"1.2.3", SemVer{false, 1, 2, 3}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%v", tt.v, tt.expected)
		t.Run(testname, func(t *testing.T) {
			// given
			// when
			got, _ := getSemVer(tt.v)
			// then
			if got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}
		})

	}
}

func TestIncSemVer(t *testing.T) {

	tests := []struct {
		s        SemVer
		i        IncType
		expected SemVer
	}{
		{SemVer{true, 0, 0, 1}, IncPatch, SemVer{true, 0, 0, 2}},
		{SemVer{true, 0, 0, 10}, IncPatch, SemVer{true, 0, 0, 11}},
		{SemVer{true, 0, 1, 10}, IncMinor, SemVer{true, 0, 2, 0}},
		{SemVer{true, 0, 1, 10}, IncMajor, SemVer{true, 1, 0, 0}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%v,%v", tt.s, tt.i, tt.expected)
		t.Run(testname, func(t *testing.T) {
			// given
			// when
			got := incSemVer(&tt.s, tt.i)
			// then
			if *got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}
		})

	}

}

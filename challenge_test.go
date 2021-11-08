package main

import (
	"testing"
)

func Test_checkIfValidFile(t *testing.T) {

	tests := []struct {
		name     string
		filepath string
		want     bool
	}{
		{"valid files test", "/home/alexdripchak/go/src/praccy/hands_on.go", true},
		{"csv file test", "/home/alexdripchak/Downloads/targets.csv", false},
		{"not a file test", "bepis", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := go_file_check(tt.filepath)

			if got != tt.want {
				t.Errorf("checkIfValidFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllGoFiles(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		want []string
	}{
		{"go dir", "/home/alexdripchak/go/src/praccy/", []string{"/home/alexdripchak/go/src/praccy/handson_level1.go"}},
		{"no go files in dir", "/home/alexdripchak/Downloads/", make([]string, 0)},
		//{"misc", "bepis", make([]string, 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := recurse_through_repo(tt.dir)
			for i, f := range got {
				if tt.want[i] == f {
					continue
				} else {
					t.Errorf("checkIfValidFile() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_countLines(t *testing.T) {
	tests := []struct {
		name string
		dir  string
		want int
	}{
		{"go dir", "/home/alexdripchak/go/src/praccy/", 41},
		{"go dir", "/home/alexdripchak/go/src/derby/", 76},
		{"no go files in dir", "/home/alexdripchak/Downloads/", 0},
		//{"misc", "bepis", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := recurse_through_repo(tt.dir)
			got := count_lines_in_files(list)

			if got != tt.want {
				t.Errorf("checkIfValidFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

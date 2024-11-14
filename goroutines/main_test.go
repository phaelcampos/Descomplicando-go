package main

import "testing"

func Test_getPokemonInfo(t *testing.T) {
	type args struct {
		url string
		c   chan Pokemon
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getPokemonInfo(tt.args.url, tt.args.c)
		})
	}
}

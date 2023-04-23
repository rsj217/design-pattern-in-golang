package proxy

import "testing"

func Test_client(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"normal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client1()
		})
	}
}

func Test_client2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"proxy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client2()
		})
	}
}

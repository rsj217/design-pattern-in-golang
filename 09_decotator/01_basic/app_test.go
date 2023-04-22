package decorator

import "testing"

func Test_client1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"fetch from db"},
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
		{"fetch with redis wrapper"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client2()
		})
	}
}

func Test_client3(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"multi wrapper cache hit"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client3()
		})
	}
}

func Test_client4(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"multi wrapper cache missing"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client4()
		})
	}
}

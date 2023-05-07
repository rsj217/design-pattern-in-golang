package prototype

import "testing"

func TestCopy(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"copy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Copy()
		})
	}
}

func TestDeepCopy(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"deep copy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeepCopy()
		})
	}
}

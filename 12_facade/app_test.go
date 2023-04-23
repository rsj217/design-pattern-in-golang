package facade

import "testing"

func Test_client(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"facade"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client()
		})
	}
}

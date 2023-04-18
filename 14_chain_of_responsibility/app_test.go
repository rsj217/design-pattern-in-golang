package chain_of_responsibility

import "testing"

func Test_Client(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"chain of responsibility"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Client()
		})
	}
}

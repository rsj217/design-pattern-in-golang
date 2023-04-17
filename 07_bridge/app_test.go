package bridge

import "testing"

func TestApp(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"bridge"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			App()
		})
	}
}

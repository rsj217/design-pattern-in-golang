package template_method

import "testing"

func TestClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"template method"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Client()
		})
	}
}

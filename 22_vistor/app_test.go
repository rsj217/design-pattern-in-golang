package visitor

import "testing"

func TestClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"visitor"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Client()
		})
	}
}
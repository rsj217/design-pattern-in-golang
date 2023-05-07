package anti_ocp

import "testing"

func Test_client(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"anti ocp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client()
		})
	}
}

package abstract_factory

import "testing"

func Test_client(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"abstract-factory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client()
		})
	}
}

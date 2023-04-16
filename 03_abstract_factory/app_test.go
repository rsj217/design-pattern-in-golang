package abstract_factory

import "testing"

func TestApp(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"abstract-factory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			App()
		})
	}
}

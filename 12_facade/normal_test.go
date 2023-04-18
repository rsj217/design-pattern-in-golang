package facade

import "testing"

func TestApp(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"normal",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := App(); got != tt.want {
				t.Errorf("App() = %v, want %v", got, tt.want)
			}
		})
	}
}

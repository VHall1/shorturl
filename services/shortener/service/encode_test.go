package service

import (
	"fmt"
	"testing"
)

func TestBase62(t *testing.T) {
	tests := []struct {
		id   int64
		want string
	}{
		{id: 0, want: "0"},
		{id: 1, want: "1"},
		{id: 61, want: "z"},
		{id: 62, want: "10"},
		{id: 12345, want: "3D7"},
		{id: 987654321, want: "14q60P"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.id), func(t *testing.T) {
			got := Base62(tt.id)
			if got != tt.want {
				t.Errorf("%d got %s want %s", tt.id, got, tt.want)
			}
		})
	}
}

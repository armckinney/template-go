package server

import (
	"os"
	"testing"
)

func TestNewServer_PortSelection(t *testing.T) {
	tests := []struct {
		name     string
		envPort  string
		wantAddr string
	}{
		{
			name:     "Default Port",
			envPort:  "",
			wantAddr: ":8080",
		},
		{
			name:     "Custom Port",
			envPort:  "4000",
			wantAddr: ":4000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save current env and defer restore
			oldEnv := os.Getenv("PORT")
			defer os.Setenv("PORT", oldEnv)

			os.Setenv("PORT", tt.envPort)

			srv := NewServer()

			if srv.Addr != tt.wantAddr {
				t.Errorf("NewServer().Addr = %v, want %v", srv.Addr, tt.wantAddr)
			}
		})
	}
}

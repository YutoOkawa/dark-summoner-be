package server

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	// Test the NewServer function
	port := ":8080"
	shutdownTimeout := 5 * time.Second
	server := NewServer(port, shutdownTimeout)

	if server.port != port {
		t.Errorf("Expected port %s, got %s", port, server.port)
	}

	if server.shutdownTimeout != shutdownTimeout {
		t.Errorf("Expected shutdownTimeout %v, got %v", shutdownTimeout, server.shutdownTimeout)
	}
}

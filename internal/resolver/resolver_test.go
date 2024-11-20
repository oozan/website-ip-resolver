package resolver

import (
	"testing"
)

func TestResolveValid(t *testing.T) {
	resolver := NewResolver()
	ips, err := resolver.Resolve("google.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ips) == 0 {
		t.Fatal("expected at least one IP address, got none")
	}
}

func TestResolveInvalid(t *testing.T) {
	resolver := NewResolver()
	_, err := resolver.Resolve("invalid.domain")
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestResolveEmpty(t *testing.T) {
	resolver := NewResolver()
	_, err := resolver.Resolve("")
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

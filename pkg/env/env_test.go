package env

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	t.Helper()
	t.Run("Check int", func(t *testing.T) {
		os.Setenv("INT", "42")
		got := GetEnv[int]("INT")
		if got != 42 {
			t.Fatalf("expected 42, got: %v", got)
		}
	})

	t.Run("Check bool", func(t *testing.T) {
		os.Setenv("BOOL", "true")
		got := GetEnv[bool]("BOOL")
		if !got {
			t.Fatalf("expected true, got: %v", got)
		}
	})

	t.Run("Check string", func(t *testing.T) {
		os.Setenv("STRING", "lorem ipsum")
		got := GetEnv[string]("STRING")
		if got != "lorem ipsum" {
			t.Fatalf("expected 'lorem ipsum', got: %v", got)
		}
	})

	t.Run("Check bool invalid defaults to false", func(t *testing.T) {
		t.Setenv("BOOL_INVALID", "notaboolean")
		got := GetEnv[bool]("BOOL_INVALID")
		if got != false {
			t.Fatalf("expected false, got %v", got)
		}
	})

	t.Run("Not set should return zero value", func(t *testing.T) {
		gotInt := GetEnv[int]("NONEXISTENT_INT")
		if gotInt != 0 {
			t.Fatalf("expected 0, got %d", gotInt)
		}

		gotStr := GetEnv[string]("NONEXISTENT_STR")
		if gotStr != "" {
			t.Fatalf("expected empty string, got %q", gotStr)
		}

		gotBool := GetEnv[bool]("NONEXISTENT_BOOL")
		if gotBool != false {
			t.Fatalf("expected false, got %v", gotBool)
		}
	})
}

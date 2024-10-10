package auth

import (
	"testing"
)


func TestGetAPIKey(t *testing.T) {
	got := GetAPIKey("ApiKey")
	want := (string, error)"ApiKey", nil
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}


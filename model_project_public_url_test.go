package openapi

import (
	"encoding/json"
	"testing"
)

func TestProjectPublicUrl(t *testing.T) {
	var p Project
	if err := json.Unmarshal([]byte(`{"id":1,"slug":"textproxyz","public_url":"https://nitrozen.io/c/textproxyz"}`), &p); err != nil {
		t.Fatal(err)
	}
	if !p.HasPublicUrl() || p.GetPublicUrl() != "https://nitrozen.io/c/textproxyz" {
		t.Fatalf("public_url not decoded: %+v", p)
	}
}

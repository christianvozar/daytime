// Copyright © 2015-2016 Christian R. Vozar
// MIT Licensed.

package daytime

import (
	"testing"
)

// TestDaytimeServer tests a simple Daytime protocol server.
func TestDaytimeServer(t *testing.T) {
	err := ListenAndServe(":13")
	if err != nil {
		t.Fatal(err)
	}
}

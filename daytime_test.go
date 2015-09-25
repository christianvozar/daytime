// Copyright © 2015 Rogue Ethic
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE.markdown file.

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

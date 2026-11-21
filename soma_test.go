package aprovae2e

import "testing"

func TestSomar(t *testing.T) {
	if Somar(2, 3) != 5 {
		t.Fatal("2 + 3 deveria ser 5")
	}
}

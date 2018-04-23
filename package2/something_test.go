package package2

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	panic("Attention, we got a panic here.")
	os.Exit(m.Run())
}

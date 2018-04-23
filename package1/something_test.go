package package1

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test1(t *testing.T) {
	t.Log("Some fancy test")
}

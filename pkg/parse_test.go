package pkg

import (
	"testing"
)

func TestParse(t *testing.T) {
	deps, err := Parse("../test/repo/go.mod")
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(deps) > 0 {
		t.Errorf("should be zero deps")
	}
}

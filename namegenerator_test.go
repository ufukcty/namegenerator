package namegenerator

import (
	"strings"
	"testing"
)

func TestNew_CheckName(t *testing.T) {
	got := New()
	if arr := strings.Split(got, " "); len(arr) == 2 {
		t.Skip()
	} else {
		t.Fail()
	}
}

package bablmodule

import (
	"strings"
	"testing"
)

func TestModuleGrpcModuleName(t *testing.T) {
	mod := New("larskluge/http-forwarder")
	expected := "HttpForwarder"
	actual := mod.GrpcModuleName()
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected '%s', but result was '%s'", expected, actual)
	}
}

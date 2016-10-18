package bablmodule

import (
	"strings"
	"testing"
)

func TestModuleGrpcModuleName(t *testing.T) {
	mod := New("larskluge/http-forward")
	expected := "HttpForward"
	actual := mod.GrpcModuleName()
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected '%s', but result was '%s'", expected, actual)
	}
}

func TestDockerServiceName(t *testing.T) {
	mod := New("larskluge/http-forward")
	expected := "larskluge--http-forward"
	actual := mod.DockerServiceName()
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected '%s', but result was '%s'", expected, actual)
	}
}

func TestDockerFullName(t *testing.T) {
	mod := New("larskluge/http-forward")
	expected := "larskluge/http-forward"

	actual := mod.Fullname()
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected '%s', but result was '%s'", expected, actual)
	}
}

func TestDockerFullNameWithTag(t *testing.T) {
	mod := New("larskluge/http-forward:babl")
	expected := "larskluge/http-forward:babl"
	actual := mod.Fullname()
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected '%s', but result was '%s'", expected, actual)
	}
}

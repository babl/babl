package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestMain(t *testing.T) { TestingT(t) }

type MainSuite struct{}

var _ = Suite(&MainSuite{})

func (s *MainSuite) TestBablEndpoint(c *C) {
	req := ModuleRequest{BablEndpoint: "a,b"}
	c.Assert(req.bablEndpoint(), Equals, "a,b")
}

func (s *MainSuite) TestNextEndpoint(c *C) {
	req := ModuleRequest{BablEndpoint: "a,b"}

	c.Assert(req.nextEndpoint(), Equals, "a")
	c.Assert(req.nextEndpoint(), Equals, "b")
}

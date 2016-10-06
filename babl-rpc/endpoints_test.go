package main

import (
	"testing"

	rr "github.com/larskluge/babl/babl-rpc/roundrobin"
	. "gopkg.in/check.v1"
)

func TestEndpoints(t *testing.T) { TestingT(t) }

type EndpointSuite struct{}

var _ = Suite(&EndpointSuite{})

func (s *EndpointSuite) TestNoEndpoints(c *C) {
	c.Assert(ParseEndpoints(""), DeepEquals, []rr.Endpoint{})
}

func (s *EndpointSuite) TestOneEndpoint(c *C) {
	c.Assert(ParseEndpoints("babl.sh"), DeepEquals, []rr.Endpoint{rr.Endpoint{"babl.sh", 1}})
}

func (s *EndpointSuite) TestTwoEndpoints(c *C) {
	c.Assert(ParseEndpoints("babl.sh, v5.babl.sh"), DeepEquals, []rr.Endpoint{rr.Endpoint{"babl.sh", 1}, rr.Endpoint{"v5.babl.sh", 1}})
}

func (s *EndpointSuite) TestWeightedEndpoints(c *C) {
	c.Assert(ParseEndpoints("babl.sh;q=9, v5.babl.sh;q=1"), DeepEquals, []rr.Endpoint{rr.Endpoint{"babl.sh", 9}, rr.Endpoint{"v5.babl.sh", 1}})
}

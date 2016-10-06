package roundrobin

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestRR(t *testing.T) { TestingT(t) }

type RRSuite struct{}

var _ = Suite(&RRSuite{})

func (s *RRSuite) TestNoEndpoints(c *C) {
	_, err := New([]Endpoint{})
	c.Assert(err, NotNil)
}

func (s *RRSuite) TestOneEndpoint(c *C) {
	lb, err := New([]Endpoint{Endpoint{Endpoint: "a", Weight: 1}})
	c.Assert(err, IsNil)

	c.Assert(seq(c, lb, 3), DeepEquals, []string{"a", "a", "a"})
}

func (s *RRSuite) TestSimple(c *C) {
	lb, err := New([]Endpoint{{"a", 1}, {"b", 1}})
	c.Assert(err, IsNil)

	c.Assert(seq(c, lb, 3), DeepEquals, []string{"a", "b", "a"})
}

func (s *RRSuite) TestWeighted(c *C) {
	lb, err := New([]Endpoint{{"a", 3}, {"b", 2}})
	c.Assert(err, IsNil)

	c.Assert(seq(c, lb, 6), DeepEquals, []string{"a", "a", "b", "a", "b", "a"})
}

func (s *RRSuite) TestWeighted2(c *C) {
	lb, err := New([]Endpoint{{"a", 8}, {"b", 2}})
	c.Assert(err, IsNil)

	c.Assert(seq(c, lb, 11), DeepEquals, []string{"a", "a", "a", "a", "b", "a", "a", "a", "a", "b", "a"})
}

func seq(c *C, lb *RoundRobin, repeat int) []string {
	out := []string{}
	for i := 0; i < repeat; i++ {
		out = append(out, lb.NextEndpoint())
	}
	return out
}

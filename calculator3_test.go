package calc

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestAdd(c *C) {
	c.Assert(Add(1, 2), Equals, 3)
}

func (s *MySuite) TestSubtract(c *C) {
	c.Assert(Subtract(4, 7), Equals, -3)
}

func (s *MySuite) TestMultiply(c *C) {
	c.Assert(Multiply(5, 6), Equals, 30)
}

func (s *MySuite) TestDivide(c *C) {
	c.Assert(Divide(10, 2), Equals, 5.0)
}

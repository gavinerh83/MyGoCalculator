package calc

import (
	"testing"

	. "gopkg.in/check.v1"
)

type tabledata struct {
	x        int
	y        int
	expected int
}

type tableDivide struct {
	x        int
	y        int
	expected float64
}

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

//TestAdd test addition function
func (s *MySuite) TestAdd(c *C) {
	testData := []tabledata{
		{2, 3, 5},
		{6, 7, 13},
		{9, 10, 19},
		{100, 201, 301},
	}
	for _, v := range testData {
		c.Assert(Add(v.x, v.y), Equals, v.expected)
	}
}

//TestSubtract tests subtraction function
func (s *MySuite) TestSubtract(c *C) {
	testData := []tabledata{
		{2, 3, -1},
		{6, 8, -2},
		{30, 10, 20},
		{100, 201, -101},
	}
	for _, v := range testData {
		c.Assert(Subtract(v.x, v.y), Equals, v.expected)
	}
}

//TestMultiply tests the multiply function
func (s *MySuite) TestMultiply(c *C) {
	testData := []tabledata{
		{2, 3, 6},
		{6, 8, 48},
		{30, 10, 300},
		{100, 201, 20100},
	}
	for _, v := range testData {
		c.Assert(Multiply(v.x, v.y), Equals, v.expected)
	}
}

//TestDivide test the division function
func (s *MySuite) TestDivide(c *C) {
	testData := []tableDivide{
		{5, 2, 2.5},
		{6, 8, 0.75},
		{30, 10, 3.0},
		{100, 25, 4.0},
	}
	for _, v := range testData {
		c.Assert(Divide(v.x, v.y), Equals, v.expected)
	}
}

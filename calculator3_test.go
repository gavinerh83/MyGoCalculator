package calc

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Numbers", func() {
		// Passing Test
		g.It("Should add two numbers ", func() {
			g.Assert(Add(3, 1)).Equal(4)
		})
		g.It("Should do subtract operations", func() {
			g.Assert(Subtract(20, 5)).Equal(15)
		})
		g.It("Should Multiply two numbers", func() {
			g.Assert(Multiply(6, 7)).Equal(42)
		})
		// Excluded Test
		g.It("Should divide two numbers ", func() {
			g.Assert(Divide(9, 3)).Equal(3.0)
		})
	})
}

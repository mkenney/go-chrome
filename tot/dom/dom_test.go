package dom

import (
	"testing"
)

// Regression protection for https://github.com/mkenney/go-chrome/pull/89
func TestDOMQuadType(t *testing.T) {
	t.Logf("%+v", Quad{0, 1.1})
	t.Logf("%+v", Quad([2]float64{0, 1.1}))
}

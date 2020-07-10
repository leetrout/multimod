package foo_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leetrout/multimod/pkg/foo"
)

type S struct {
	Bar string
}

func TestFoo(t *testing.T) {
	s1 := S{Bar: foo.Foo()}
	s2 := S{Bar: foo.Bar()}
	t.Fatalf(cmp.Diff(s1, s2))
}

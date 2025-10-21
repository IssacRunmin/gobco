package foo

import "testing"

func TestFooUseBar_Negative(t *testing.T) {
	got := FooUseBar(-2)
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}

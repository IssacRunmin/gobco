package bar

import "testing"

func TestDoSomething_Positive(t *testing.T) {
	if got := DoSomething(1); got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
}

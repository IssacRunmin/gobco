package foo

import "demo/bar"

func FooUseBar(x int) int {
	return bar.DoSomething(x)
}

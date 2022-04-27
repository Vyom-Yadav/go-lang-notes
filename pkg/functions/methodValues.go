package functions

import "fmt"

type S struct {
	*T
}

type T int

func (t T) M() {
	fmt.Print(t)
}

func MethodValues() {
	// This is a "func" but we are discussing method
	// values, beware of that.
	t := new(T)
	s := S{T: t}
	f := t.M // same as *t.M, see selector expression in go
	g := s.M // same as (*(s.T)).M, see https://stackoverflow.com/questions/13533681/when-do-gos-pointers-dereference-themselves
	*t = 41
	f()
	g()
	// They will give out the anwser 00 as for a method value
	// x.M, the expression x is evaluated and saved during the
	// evalutaion of method value; the saved copy is then used
	// as receiver in any calls, which may be executed later.
}

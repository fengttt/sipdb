package errors

import (
	"bytes"
	"fmt"
	"runtime"
)

type Error struct {
	Kind Kind
	Msg string

	stack []uintptr
}

type Kind int

const (
	Internal	Kind = iota		// Internal error.
)

func (k Kind) String() string {
	switch k {
	case Internal:
		return "internal error"
	}
	return "unknown error kind"
}

func E(k Kind, msg string) error { 
	e := &Error{}
	e.Kind = k
	e.Msg = msg

	e.populateStack()
	return e
}

func (e *Error) Error() string {
	b := new(bytes.Buffer)
	e.printStack(b)

	b.WriteString(e.Kind.String())
	b.WriteString(": " + e.Msg)
	return b.String()
}

func (e *Error) populateStack() {
	var stk [64]uintptr
	// skip 2, to get stack frame where the error is constructed.
	n := runtime.Callers(2, stk[:])
	e.stack = stk[:n]
}

func (e *Error) printStack(b *bytes.Buffer) {
	frames := runtime.CallersFrames(e.stack)
	for i := 0; i < len(e.stack); i++ {
		f, ok := frames.Next()
		if !ok {
			break
		}
		fmt.Fprintf(b, "\n\t%s %v:%d",  f.Func.Name(), f.File, f.Line)
	}
	b.WriteString("\n")
}


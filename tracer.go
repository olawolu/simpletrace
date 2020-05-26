package simpletrace

import (
	"fmt"
	"io"
)

// Tracer is an interface that describes an object  capable
// of tracing events throughout code.
// Describes a single method: Trace
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

func (tr *tracer) Trace(a ...interface{}) {
	fmt.Fprintln(tr.out, a...)
}

// New returns the address to a tracer struct
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func (nt *nilTracer) Trace(a ...interface{}) {}

// Off returns a new nilTracer
func Off() Tracer {
	return &nilTracer{}
}

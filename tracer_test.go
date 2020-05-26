package simpletrace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) { // Capture tracer output in a buffer.
	var buf bytes.Buffer
	tracer := New(&buf)
	// Ensure the string in the buffer matches the expected value.
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}

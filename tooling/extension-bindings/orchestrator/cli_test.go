package orchestrator

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestRunAcceptsArgumentsWithAndWithoutBindingsPrefix(t *testing.T) {
	for _, args := range [][]string{{"bindings", "resolve", "--help"}, {"resolve", "--help"}} {
		var out, errout bytes.Buffer
		if err := Run(context.Background(), args, &out, &errout); err != nil {
			t.Fatalf("%v: %v", args, err)
		}
		if !strings.Contains(errout.String(), "-packages") {
			t.Fatalf("%v: missing common flags", args)
		}
	}
}

func TestRunRejectsMissingOrUnknownCommand(t *testing.T) {
	for _, args := range [][]string{{"bindings"}, {"bindings", "nope"}, {"nope"}} {
		var out, errout bytes.Buffer
		if err := Run(context.Background(), args, &out, &errout); err == nil {
			t.Fatalf("%v accepted", args)
		}
	}
}

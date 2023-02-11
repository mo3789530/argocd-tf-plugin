package cmd

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Run("will show version", func(t *testing.T) {
		args := []string{}
		cmd := NewVersionCommand()

		c := bytes.NewBufferString("")
		cmd.SetArgs(args)
		cmd.SetOut(c)
		cmd.Execute()
		out, err := io.ReadAll(c)
		if err != nil {
			t.Fatal(err)
		}

		expected := "argocd-tf-plugin 0.0.1-rc"
		if !strings.Contains(string(out), expected) {
			t.Fatalf("expected to contain: %s but got %s", expected, out)
		}
	})
}

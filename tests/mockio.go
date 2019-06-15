package tests

import (
	"bytes"
	"github.com/zondax/cobra"
	"strings"
)

// ApplyMockIO replaces stdin/out/err with buffers that can be used during testing
func ApplyMockIO(c *cobra.Command) (*strings.Reader, *bytes.Buffer, *bytes.Buffer) {
	mockIn := strings.NewReader("")
	mockOut := bytes.NewBufferString("")
	mockErr := bytes.NewBufferString("")
	c.SetIn(mockIn)
	c.SetOut(mockOut)
	c.SetErr(mockErr)
	return mockIn, mockOut, mockErr
}

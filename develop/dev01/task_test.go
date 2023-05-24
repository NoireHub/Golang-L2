package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CurrentTime(t *testing.T) {
	output := captureOutput(CurrentTime)
	assert.NotEmpty(t,output)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

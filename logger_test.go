package logmatic

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockPrint(buf *bytes.Buffer) {
	printf = func(format string, a ...interface{}) (int, error) {
		buf.WriteString(fmt.Sprintf(format, a...))
		return 0, nil
	}
}

func TestLogger_Trace_Print(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.SetLevel(TRACE)
	l.Trace("Trace")
	assert.NotEmpty(t, buf)
}

func TestLogger_Trace_NoPrint(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.Trace("Trace")
	assert.Empty(t, buf)
}

func TestLogger_Debug_Print(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.SetLevel(DEBUG)
	l.Debug("Debug")
	assert.NotEmpty(t, buf)
}

func TestLogger_Debug_NoPrint(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.Debug("Debug")
	assert.Empty(t, buf)
}

func TestLogger_Info_Print(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.Info("Info")
	assert.NotEmpty(t, buf)
}

func TestLogger_Info_NoPrint(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.SetLevel(WARN)
	l.Info("Info")
	assert.Empty(t, buf)
}

func TestLogger_Warn_Print(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.Warn("Warn")
	assert.NotEmpty(t, buf)
}

func TestLogger_Warn_NoPrint(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.SetLevel(ERROR)
	l.Warn("Warn")
	assert.Empty(t, buf)
}

func TestLogger_Error_Print(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.Error("Error")
	assert.NotEmpty(t, buf)
}

func TestLogger_Error_NoPrint(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)

	l := NewLogger()
	l.SetLevel(FATAL)
	l.Error("Error")
	assert.Empty(t, buf)
}

func TestLogger_Fatal_Print(t *testing.T) {
	var buf bytes.Buffer
	mockPrint(&buf)
	exit = func(code int) {}

	l := NewLogger()
	l.Fatal("Fatal")
	assert.NotEmpty(t, buf)
}

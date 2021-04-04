package handlers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
)

func Test_checkErrorWithError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	err := CheckError("error description", errors.New("error message"))
	assert.True(t, err, "Should return true")
	assert.Contains(t, buf.String(), "error description", "Missing error description")
	assert.Contains(t, buf.String(), "error message", "Missing error message")
}

func Test_checkErrorWithoutError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	err := CheckError("test", nil)
	assert.False(t, err, "Should return false")
	assert.Empty(t, buf, "Unexpected error logging")
}

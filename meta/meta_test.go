package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestVersion unit test.
func TestVersion(t *testing.T) {
	assert.NotNil(t, Version)
	assert.NotEmpty(t, Version)
	assert.NotContains(t, Version, " ")
}

// TestAppname unit test.
func TestAppname(t *testing.T) {
	assert.NotNil(t, ApplicationName)
	assert.NotEmpty(t, ApplicationName)
	assert.NotContains(t, ApplicationName, " ")
}

// TestConfigfilename unit test.
func TestConfigfilename(t *testing.T) {
	assert.NotNil(t, ConfigFilename)
	assert.NotEmpty(t, ConfigFilename)
	assert.NotContains(t, ConfigFilename, " ")
}

// TestConfigfile unit test.
func TestConfigfile(t *testing.T) {
	f, err := ConfigFile()
	assert.NoError(t, err)
	assert.NotNil(t, f)
	assert.NotEmpty(t, f)
}

// TestConfigfileWithSpecialDir unit test.
func TestConfigfileWithSpecialDir(t *testing.T) {
	ConfigDir = "./"
	f, err := ConfigFile()
	assert.NoError(t, err)
	assert.NotNil(t, f)
	assert.NotEmpty(t, f)
}

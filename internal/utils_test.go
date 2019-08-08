package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCanonicalAlias(t *testing.T) {
	tests := []struct {
		hostname string
		alias    string
		expected string
	}{{"host.com", "alias", "#alias:host.com"}}

	for _, test := range tests {
		got := GetCanonicalAlias(test.hostname, test.alias)
		assert.Equal(t, test.expected, got)
	}
}

func TestStripAlias(t *testing.T) {
	tests := []struct {
		hostname       string
		canonicalAlias string
		expected       string
	}{{"host.com", "#alias:host.com", "alias"}}

	for _, test := range tests {
		got := StripAlias(test.hostname, test.canonicalAlias)
		assert.Equal(t, test.expected, got)
	}

}

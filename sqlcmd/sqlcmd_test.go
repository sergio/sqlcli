package sqlcmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func transform(command string, args []string) ([]string, error) {
	return []string{"-S1.1.1.1"}, nil
}

type StringMap map[string]string

func (m StringMap) With(key string, value string) StringMap {
	result := StringMap(make(map[string]string))
	for k, v := range m {
		result[k] = v
	}
	result[key] = value
	return result
}

func TestRun(t *testing.T) {
	const serverAddress = "1.1.1.1"
	config := StringMap(map[string]string{"server": serverAddress})
	testCases := []struct {
		config   StringMap
		command  string
		args     []string
		expected []string
	}{
		{
			config:   config,
			command:  "query",
			args:     []string{""},
			expected: []string{"-S" + serverAddress},
		},
		{
			config:   config.With("server", "hostxyz"),
			command:  "query",
			args:     []string{""},
			expected: []string{"-S" + "hostxyz"},
		},
	}

	for _, c := range testCases {
		actual, err := transform(c.command, c.args)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, c.expected, actual)
	}
}

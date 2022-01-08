package tree

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Skipf("Configs to test are not ready.")

	tests := []string{
		"correct_config_path",
	}

	validate := validator.New()

	for i, test := range tests {
		path, ok := os.LookupEnv(test)

		require.True(t, ok)

		_, err := NewConfig(*validate, path)

		require.NoErrorf(t, err, "Error not equals (%d). ", i)
	}
}

func TestNewConfigReturnsError(t *testing.T) {
	t.Skipf("Configs to test are not ready. ")

	tests := []string{
		"incorrect_root_config_path",
	}

	validate := validator.New()

	for i, test := range tests {
		path, ok := os.LookupEnv(test)

		require.True(t, ok)

		_, err := NewConfig(*validate, path)

		require.Error(t, err, "Error not equals (%d). ", i)
	}
}

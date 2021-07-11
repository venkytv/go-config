package config

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoad(t *testing.T) {
	var intval int64 = 9223372036854775807

	f := flag.NewFlagSet("test", flag.ExitOnError)
	f.String("simple", "", "simple flag")
	f.String("defval", "my-default", "flag with default value")
	f.String("env-val", "overridden-default", "flag with env default")
	f.Int64("intflag", intval, "int64 flag")

	os.Setenv("CONFIG_LOAD_TEST_ENV_VAL", "env-default")

	cfg := Load(f, "CONFIG_LOAD_TEST")

	assert.ElementsMatch(t, []string{"simple", "defval", "env-val", "intflag"},
		cfg.AllKeys())
	assert.Empty(t, cfg.GetString("simple"))
	assert.Equal(t, "my-default", cfg.GetString("defval"))
	assert.Equal(t, "env-default", cfg.GetString("env-val"))
	assert.Equal(t, intval, cfg.GetInt64("intflag"))

	os.Setenv("CONFIG_LOAD_TEST_INTFLAG", "13")
	assert.Equal(t, int64(13), cfg.GetInt64("intflag"))
}

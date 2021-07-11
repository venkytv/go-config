package config

import (
	"flag"
	"os"
	"path"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
	Pflag *pflag.FlagSet
}

func (cfg *Config) Args() []string {
	if !cfg.Pflag.Parsed() {
		panic("Config not loaded")
	}
	return cfg.Pflag.Args()
}

func Load(flagset *flag.FlagSet, prefix string) *Config {
	if flagset == nil {
		flagset = flag.CommandLine
	}
	pf := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
	pf.AddGoFlagSet(flagset)
	pf.Parse(os.Args[1:])
	v := viper.New()
	cfg := &Config{v, pf}
	v.BindPFlags(pf)
	if len(prefix) > 0 {
		v.SetEnvPrefix(prefix)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	bin := path.Base(os.Args[0])
	v.AddConfigPath("/etc/" + bin)
	v.AddConfigPath("$HOME/.config/" + bin)

	if err := v.ReadInConfig(); err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			panic(err)
		}
	}

	return cfg
}

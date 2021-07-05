package config

import (
	"flag"
	"os"
	"path"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config = viper.Viper

func Load(flagset *flag.FlagSet, prefix string) *viper.Viper {
	if flagset == nil {
		flagset = flag.CommandLine
	}
	pflag.CommandLine.AddGoFlagSet(flagset)
	pflag.Parse()
	v := viper.New()
	v.BindPFlags(pflag.CommandLine)
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

	return v
}

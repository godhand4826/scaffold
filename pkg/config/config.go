package config

import (
	"flag"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/itzg/go-flagsfiller"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"scaffold/pkg/logger"
)

type Config struct {
	Config     string        `mapstructure:"config"`
	FxVerbose  bool          `mapstructure:"fx_verbose"`
	Logger     logger.Config `mapstructure:"logger"`
	ServerAddr string        `mapstructure:"server_addr"`
}

func Load() (*Config, error) {
	var config Config

	_ = godotenv.Load()

	if err := flagsfiller.New(
		flagsfiller.WithFieldRenamer(camelSplitByDashToSnakeSplitByDot),
		flagsfiller.WithEnvRenamer(func(s string) string {
			envKey := strcase.ToScreamingSnake(s)
			viper.MustBindEnv(camelSplitByDashToSnakeSplitByDot(s), envKey)
			return envKey
		}),
		flagsfiller.NoSetFromEnv(),
	).Fill(flag.CommandLine, &config); err != nil {
		return nil, err
	}
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return nil, err
	}

	viper.SetDefault("config", "config.yaml")
	viper.SetConfigFile(viper.GetString("config"))
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func camelSplitByDashToSnakeSplitByDot(name string) string {
	return strings.ReplaceAll(strcase.ToSnakeWithIgnore(name, "-"), "-", ".")
}

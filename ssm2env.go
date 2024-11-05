package ssm2env

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/shopsmart/ssm2env/pkg/service"
	"github.com/shopsmart/ssm2env/pkg/utils"
)

// Config represents the various configurations for ssm2env
type Config struct {
	SearchPath       string
	Recursive        bool
	MultilineSupport bool
	Export           bool
}

// ConfigFromEnv returns a config that pulls configurations from environment variables
func ConfigFromEnv() *Config {
	v := viper.New()
	v.SetEnvPrefix("ssm2env")

	_ = v.BindEnv("path")
	_ = v.BindEnv("multiline")
	_ = v.BindEnv("recursive")
	_ = v.BindEnv("export")

	return &Config{
		SearchPath:       v.GetString("path"),
		Recursive:        v.GetBool("recursive"),
		MultilineSupport: v.GetBool("multiline"),
		Export:           v.GetBool("export"),
	}
}

// Get SSM parameters and then pass to the do function
func getParametersAndDo(svc service.Service, searchPath string, recursive bool, do func(map[string]string) error) error {
	recursively := ""
	if recursive {
		recursively = " recursively"
	}

	log.Debugf("Getting parameters for search path: %s%s", searchPath, recursively)
	params, err := svc.GetParameters(searchPath, recursive)
	if err != nil {
		return err
	}

	log.Debugf("Found %d parameters", len(params))
	return do(params)
}

// WriteEnv retrieves the SSM parameters for the given search path and writes to the writer in env format
func WriteEnv(w io.Writer, cfg *Config) error {
	svc, err := service.New()
	if err != nil {
		return err
	}

	return WriteEnvWithService(svc, w, cfg)
}

// WriteEnvWithService retrieves the SSM parameters for the given search path and writes to the writer in env format
func WriteEnvWithService(svc service.Service, w io.Writer, cfg *Config) error {
	return getParametersAndDo(svc, cfg.SearchPath, cfg.Recursive, func(params map[string]string) error {
		return utils.WriteEnv(w, params, cfg.MultilineSupport, cfg.Export)
	})
}

// LoadEnv loads SSM parameters directly into the environment
func LoadEnv(cfg *Config) error {
	svc, err := service.New()
	if err != nil {
		return err
	}

	return LoadEnvWithService(svc, cfg)
}

// LoadEnvWithService loads SSM parameters directly into the environment
func LoadEnvWithService(svc service.Service, cfg *Config) error {
	return getParametersAndDo(svc, cfg.SearchPath, cfg.Recursive, func(params map[string]string) error {
		for key, value := range params {
			os.Setenv(utils.EnvKey(key), value)
		}

		return nil
	})
}

// LoadViper loads SSM parameters into a viper instance
func LoadViper(v *viper.Viper, cfg *Config) error {
	svc, err := service.New()
	if err != nil {
		return err
	}

	return LoadViperWithService(svc, v, cfg)
}

// LoadViperWithService loads SSM parameters into a viper instance
func LoadViperWithService(svc service.Service, v *viper.Viper, cfg *Config) error {
	return getParametersAndDo(svc, cfg.SearchPath, cfg.Recursive, func(params map[string]string) error {
		for key, value := range params {
			v.Set(utils.EscapeKey(key), value)
		}

		return nil
	})
}

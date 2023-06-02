package ssm2env

import (
	"io"

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

// Collect retrieves the SSM parameters for the given search path and
// writes to the writer in env format
func Collect(svc service.Service, w io.Writer, cfg *Config) error {
	var err error
	if svc == nil {
		log.Debug("Initializing session")
		svc, err = service.New()
		if err != nil {
			return err
		}
	}

	recursively := ""
	if cfg.Recursive {
		recursively = " recursively"
	}

	log.Debugf("Getting parameters for search path: %s%s", cfg.SearchPath, recursively)
	params, err := svc.GetParameters(cfg.SearchPath, cfg.Recursive)
	if err != nil {
		return err
	}

	log.Debugf("Found %d parameters", len(params))
	return utils.EnvFormat(w, params, cfg.MultilineSupport, cfg.Export)
}

// LoadViper loads SSM parameters into a viper instance
func LoadViper(svc service.Service, v *viper.Viper, searchPath string, recursive bool) error {
	var err error
	if svc == nil {
		log.Debug("Initializing session")
		svc, err = service.New()
		if err != nil {
			return err
		}
	}

	params, err := svc.GetParameters(searchPath, recursive)
	if err != nil {
		return err
	}

	for key, value := range params {
		key = utils.EscapeKey(key)
		v.Set(key, value)
	}

	return nil
}

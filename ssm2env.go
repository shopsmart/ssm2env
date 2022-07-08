package ssm2env

import (
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/shopsmart/ssm2env/pkg/service"
	"github.com/shopsmart/ssm2env/pkg/utils"
)

// Collect retrieves the SSM parameters for the given search path and
// writes to the writer in env format
func Collect(svc service.Service, w io.Writer, searchPath string, recursive bool) error {
	var err error
	if svc == nil {
		log.Debug("Initializing session")
		svc, err = service.New()
		if err != nil {
			return err
		}
	}

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
	return utils.EnvFormat(w, params)
}

package utils

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"gopkg.in/alessio/shellescape.v1"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`[^a-zA-Z0-9_]`)
}

// SSMParameterToEnvVar converts the name of an SSM Parameter to an env var name
func SSMParameterToEnvVar(ssmParam string, prefix string) string {
	s := strings.TrimPrefix(ssmParam, fmt.Sprintf("%s/", prefix))
	s = re.ReplaceAllString(s, "_")
	return strings.ToUpper(s)
}

// EnvFormat will write the map of parameters into the write in env format
func EnvFormat(w io.Writer, m map[string]string) error {
	for key, value := range m {
		_, err := w.Write([]byte(fmt.Sprintf("%s=%s\n", key, shellescape.Quote(value))))
		if err != nil {
			return err
		}
	}

	return nil
}

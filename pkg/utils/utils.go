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

// EnvFormat will write the map of parameters into the write in env format
func EnvFormat(w io.Writer, m map[string]string) error {
	for key, value := range m {
		_, err := w.Write([]byte(fmt.Sprintf("%s=%s\n", strings.ToUpper(re.ReplaceAllString(key, "_")), shellescape.Quote(value))))
		if err != nil {
			return err
		}
	}

	return nil
}

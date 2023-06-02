package utils

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"gopkg.in/alessio/shellescape.v1"
)

var (
	keyRe     *regexp.Regexp
	newlineRe *regexp.Regexp
)

func init() {
	keyRe = regexp.MustCompile(`[^a-zA-Z0-9_]`)
	newlineRe = regexp.MustCompile(`\r?\n`)
}

// EscapeKey will return an escaped string with only letters, numbers, and anything else an underscore
func EscapeKey(key string) string {
	return keyRe.ReplaceAllString(key, "_")
}

// EnvFormat will write the map of parameters into the write in env format
func EnvFormat(key string, value string) (string, string) {
	return strings.ToUpper(EscapeKey(key)), shellescape.Quote(value)
}

// WriteEnv will write the environment variables in env format to the writer
func WriteEnv(w io.Writer, m map[string]string, multilineSupport bool, export bool) error {
	for key, value := range m {
		escapedKey, escapedValue := EnvFormat(key, value)
		if !multilineSupport {
			escapedValue = newlineRe.ReplaceAllString(escapedValue, "\\n")
		}

		exportPrefix := ""
		if export {
			exportPrefix = "export "
		}

		_, err := w.Write([]byte(fmt.Sprintf("%s%s=%s\n", exportPrefix, escapedKey, escapedValue)))
		if err != nil {
			return err
		}
	}

	return nil
}

//nolint:gochecknoinits,lll,gochecknoglobals // use init function
package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Environments "split_words" means underscores will be inserted to env name before uppercase's letters except the first letter.
// ex. AccessLog => ACCESS_LOG
// IF the env name contains an abbreviation (e.g. ID, URL, or API), please use "envconfig" rather than "split_words".
type Environments struct {
	Port        string `default:"7712" split_words:"true"`
	Debug       bool   `default:"false" split_words:"true"`
	AccessLog   bool   `default:"false" split_words:"true"`
	SpecEnabled bool   `default:"false" split_words:"true"`
	SpecFiles   string `default:"./swagger-ui" split_words:"true"`
	LogLevel    string `default:"warn" split_words:"true"`
}

var (
	// Env is worker's environment variable struct.
	Env Environments
)

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		log.Fatal("Failed to parse environments: " + err.Error())
	}
}

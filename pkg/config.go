package gitils

import "fmt"

type Config struct {
	Address      string
	GoogleApiKey string
	PostgresURL  string
}

func NewConfig() Config {
	return Config{Address: ":8080"}
}

func (config Config) String() string {
	return fmt.Sprintf("Address: %s", config.Address)
}

func (config *Config) Check() {

}

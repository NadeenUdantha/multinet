package multinet

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Look in the sample config file for the docs
type Config struct {
	Algorithm string
	Listen    string
	HashPort  bool
	Paths     []struct {
		Type string
		Addr string
	}
}

func LoadConfig(fp string) (*Config, error) {
	fd, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config: %v\n", err)
	}

	var cfg Config
	err = yaml.Unmarshal(fd, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse config: %v\n", err)
	}

	return &cfg, nil
}

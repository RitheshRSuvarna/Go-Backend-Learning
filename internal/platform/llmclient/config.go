package llmclient

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	BaseURL string
	Timeout time.Duration
}

func (c Config) Validate() error {
	if strings.TrimSpace(c.BaseURL) == "" {
		return fmt.Errorf("LLM base URL is required")
	}

	u, err := url.Parse(c.BaseURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid LLM base URL")
	}

	if c.Timeout <= 0 {
		return fmt.Errorf("LLM timeout must be greater than zero")
	}

	return nil
}

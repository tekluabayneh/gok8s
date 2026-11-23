package config

import (
	"context"
	"io"
	"net/http"
)

func CreatApi(ctx context.Context, method, url string, body io.Reader) (*http.Client, *http.Request) {
	cfg := ExtractKubeConfig()
	client, err := NewHTTPClient(&cfg)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequestWithContext(ctx, method, cfg.BaseURL+url, body)
	if cfg.Token != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.Token)
	}

	return client, req
}

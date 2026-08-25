package llmclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(cfg Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		baseURL: strings.TrimRight(cfg.BaseURL, "/"),
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}, nil
}

func (c *Client) Plan(ctx context.Context, req PlanRequest) (PlanResponse, error) {
	var result PlanResponse

	if err := req.Validate(); err != nil {
		return result, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return result, fmt.Errorf("marshal plan request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/plan", bytes.NewReader(body))
	if err != nil {
		return result, fmt.Errorf("create plan request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return result, fmt.Errorf("llm plan request timed out: %w", err)
		}
		return result, fmt.Errorf("call llm plan endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		message, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return result, fmt.Errorf("llm plan request failed with status %d: %s", resp.StatusCode, strings.TrimSpace(string(message)))
	}

	decoder := json.NewDecoder(io.LimitReader(resp.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("decode llm plan response: %w", err)
	}

	if err := result.Validate(); err != nil {
		return result, fmt.Errorf("invalid llm plan response: %w", err)
	}

	return result, nil
}

func (c *Client) Replan(ctx context.Context, req ReplanRequest) (ReplanResponse, error) {
	var result ReplanResponse

	if err := req.Validate(); err != nil {
		return result, err
	}

	body, err := json.Marshal(req)
	if err != nil {
		return result, fmt.Errorf("marshal replan request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/replan", bytes.NewReader(body))
	if err != nil {
		return result, fmt.Errorf("create replan request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return result, fmt.Errorf("call llm replan endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		message, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return result, fmt.Errorf("llm replan request failed with status %d: %s", resp.StatusCode, strings.TrimSpace(string(message)))
	}

	decoder := json.NewDecoder(io.LimitReader(resp.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("decode llm replan response: %w", err)
	}

	return result, nil
}

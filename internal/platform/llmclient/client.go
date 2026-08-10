package llmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *Client) Plan(req PlanRequest) (PlanResponse, error) {
	var result PlanResponse

	body, err := json.Marshal(req)
	if err != nil {
		return result, err
	}

	resp, err := c.client.Post(
		c.baseURL+"/plan",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("llm returned status %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *Client) Replan(req ReplanRequest) (ReplanResponse, error) {
	var result ReplanResponse

	body, err := json.Marshal(req)
	if err != nil {
		return result, err
	}

	resp, err := c.client.Post(
		c.baseURL+"/replan",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("llm returned status %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
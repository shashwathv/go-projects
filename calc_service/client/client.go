package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func New(baseURL string, token string) *Client {
	return &Client{
		baseURL:    baseURL,
		token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) doRequest(
	ctx context.Context,
	method string,
	path string,
	body any,
	out any,
) error {

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return err
		}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL+path,
		&buf,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errResp struct {
			Error string `json:"error"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		return fmt.Errorf("api error: %s", errResp.Error)
	}

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}

	return nil
}

func (c *Client) Add(ctx context.Context, a, b int) (int, error) {
	req := map[string]int{
		"number1": a,
		"number2": b,
	}

	var resp struct {
		Result int `json:"result"`
	}

	err := c.doRequest(ctx, http.MethodPost, "/add", req, &resp)
	return resp.Result, err
}

func (c *Client) Subtract(ctx context.Context, a, b int) (int, error) {
	req := map[string]int{
		"number1": a,
		"number2": b,
	}

	var resp struct {
		Result int `json:"result"`
	}

	err := c.doRequest(ctx, http.MethodPost, "/subtract", req, &resp)
	return resp.Result, err
}

func (c *Client) Multiply(ctx context.Context, a, b int) (int, error) {
	req := map[string]int{
		"number1": a,
		"number2": b,
	}

	var resp struct {
		Result int `json:"result"`
	}

	err := c.doRequest(ctx, http.MethodPost, "/multiply", req, &resp)
	return resp.Result, err
}

func (c *Client) Divide(ctx context.Context, dividend, divisor int) (int, error) {
	req := map[string]int{
		"dividend": dividend,
		"divisor":  divisor,
	}

	var resp struct {
		Result int `json:"result"`
	}

	err := c.doRequest(ctx, http.MethodPost, "/divide", req, &resp)
	return resp.Result, err
}

func (c *Client) Sum(ctx context.Context, numbers []int) (int, error) {
	var resp struct {
		Result int `json:"result"`
	}

	err := c.doRequest(ctx, http.MethodPost, "/sum", numbers, &resp)
	return resp.Result, err
}

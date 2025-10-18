package sabnzbd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"sonarr-sabnzbd-cli/internal/models"
)

// Client represents a Sabnzbd API client
type Client struct {
	baseURL  string
	apiKey   string
	username string
	password string
	client   *http.Client
}

// NewClient creates a new Sabnzbd API client
func NewClient(host string, port int, apiKey, username, password string, timeout time.Duration) *Client {
	baseURL := fmt.Sprintf("http://%s:%d", host, port)
	if username != "" && password != "" {
		baseURL = fmt.Sprintf("http://%s:%s@%s:%d", username, password, host, port)
	}

	return &Client{
		baseURL:  baseURL,
		apiKey:   apiKey,
		username: username,
		password: password,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// GetVersion retrieves the Sabnzbd version
func (c *Client) GetVersion() (string, error) {
	var resp models.VersionResponse
	err := c.get("mode=version", &resp)
	if err != nil {
		return "", err
	}
	// For version endpoint, we can return the version even if status is false
	// as it doesn't require authentication
	if resp.Version != "" {
		return resp.Version, nil
	}
	if !resp.Status {
		return "", fmt.Errorf("API error: %s", resp.Error)
	}
	return resp.Version, nil
}

// GetQueue retrieves the current download queue
func (c *Client) GetQueue() (*models.Queue, error) {
	var resp models.QueueResponse
	err := c.get("mode=queue", &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Queue, nil
}

// GetHistory retrieves the download history
func (c *Client) GetHistory() (*models.History, error) {
	var resp models.HistoryResponse
	err := c.get("mode=history", &resp)
	if err != nil {
		return nil, err
	}
	return &resp.History, nil
}

// AddNZB adds an NZB file or URL to the queue
func (c *Client) AddNZB(nzbURL string, category string) ([]string, error) {
	params := url.Values{}
	params.Add("mode", "addurl")
	params.Add("name", nzbURL)
	if category != "" {
		params.Add("cat", category)
	}

	var resp models.AddResponse
	err := c.getWithParams(params, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Status {
		return nil, fmt.Errorf("API error: %s", resp.Error)
	}
	return resp.NZOIDS, nil
}

// PauseQueue pauses the download queue
func (c *Client) PauseQueue() error {
	return c.simpleCommand("pause")
}

// ResumeQueue resumes the download queue
func (c *Client) ResumeQueue() error {
	return c.simpleCommand("resume")
}

// SetSpeedLimit sets the download speed limit
func (c *Client) SetSpeedLimit(limit string) error {
	params := url.Values{}
	params.Add("value", limit)
	return c.simpleCommandWithParams("speedlimit", params)
}

// GetCategories retrieves available categories
func (c *Client) GetCategories() ([]string, error) {
	var resp models.CategoriesResponse
	err := c.get("mode=get_cats", &resp)
	if err != nil {
		return nil, err
	}
	return resp.Categories, nil
}

// DeleteFromQueue deletes an item from the queue
func (c *Client) DeleteFromQueue(nzoID string) error {
	params := url.Values{}
	params.Add("value", nzoID)
	return c.simpleCommandWithParams("queue", params)
}

// simpleCommand performs a simple command without parameters
func (c *Client) simpleCommand(command string) error {
	params := url.Values{}
	params.Set("mode", command)
	var resp models.SabnzbdResponse
	err := c.getWithParams(params, &resp)
	if err != nil {
		return err
	}
	if !resp.Status {
		return fmt.Errorf("API error: %s", resp.Error)
	}
	return nil
}

// simpleCommandWithParams performs a command with parameters
func (c *Client) simpleCommandWithParams(command string, params url.Values) error {
	params.Set("mode", command)
	var resp models.SabnzbdResponse
	err := c.getWithParams(params, &resp)
	if err != nil {
		return err
	}
	if !resp.Status {
		return fmt.Errorf("API error: %s", resp.Error)
	}
	return nil
}

// get performs a GET request to the Sabnzbd API
func (c *Client) get(endpoint string, result any) error {
	fullURL := fmt.Sprintf("%s/api?apikey=%s&output=json&%s", c.baseURL, c.apiKey, endpoint)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}

// getWithParams performs a GET request with URL parameters
func (c *Client) getWithParams(params url.Values, result any) error {
	params.Set("apikey", c.apiKey)
	params.Set("output", "json")

	fullURL := fmt.Sprintf("%s/api?%s", c.baseURL, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return json.Unmarshal(body, result)
}

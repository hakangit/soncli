package sonarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"sonarr-sabnzbd-cli/internal/models"
)

// Client represents a Sonarr API client
type Client struct {
	baseURL    string
	apiKey     string
	client     *http.Client
	apiVersion string
}

// NewClient creates a new Sonarr API client
func NewClient(host string, port int, apiKey string, timeout time.Duration) *Client {
	c := &Client{
		baseURL: fmt.Sprintf("http://%s:%d", host, port),
		apiKey:  apiKey,
		client: &http.Client{
			Timeout: timeout,
		},
		apiVersion: "v3",
	}

	c.detectAPIVersion()
	return c
}

func (c *Client) detectAPIVersion() {
	req, err := http.NewRequest("GET", c.baseURL+"/api/v3/system/status", nil)
	if err != nil {
		c.apiVersion = ""
		return
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		c.apiVersion = ""
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		c.apiVersion = "v3"
	} else {
		c.apiVersion = ""
	}
}

func (c *Client) endpoint(path string) string {
	if c.apiVersion != "" {
		return fmt.Sprintf("/api/%s%s", c.apiVersion, path)
	}
	return "/api" + path
}

// GetSystemStatus retrieves the system status
func (c *Client) GetSystemStatus() (*models.SystemStatus, error) {
	var status models.SystemStatus
	err := c.get(c.endpoint("/system/status"), &status)
	return &status, err
}

// GetSeries retrieves all series
func (c *Client) GetSeries() ([]models.Series, error) {
	var series []models.Series
	err := c.get(c.endpoint("/series"), &series)
	return series, err
}

// LookupSeries searches for series by term
func (c *Client) LookupSeries(term string) ([]models.Series, error) {
	var series []models.Series
	params := url.Values{}
	params.Add("term", term)
	err := c.get(c.endpoint("/series/lookup")+"?"+params.Encode(), &series)
	return series, err
}

// GetSeriesByID retrieves a specific series by ID
func (c *Client) GetSeriesByID(id int) (*models.Series, error) {
	var series models.Series
	err := c.get(c.endpoint(fmt.Sprintf("/series/%d", id)), &series)
	return &series, err
}

// GetEpisodes retrieves episodes for a series
func (c *Client) GetEpisodes(seriesID int) ([]models.Episode, error) {
	var episodes []models.Episode
	params := url.Values{}
	params.Add("seriesId", fmt.Sprintf("%d", seriesID))
	err := c.get(c.endpoint("/episode")+"?"+params.Encode(), &episodes)
	return episodes, err
}

// GetQualityProfiles retrieves all quality profiles
func (c *Client) GetQualityProfiles() ([]models.QualityProfile, error) {
	var profiles []models.QualityProfile
	profileEndpoint := c.endpoint("/qualityprofile")
	err := c.get(profileEndpoint, &profiles)
	if err != nil && c.apiVersion == "v3" {
		err = c.get(c.endpoint("/profile"), &profiles)
	}
	return profiles, err
}

// GetRootFolders retrieves all root folders
func (c *Client) GetRootFolders() ([]models.RootFolder, error) {
	var folders []models.RootFolder
	err := c.get(c.endpoint("/rootfolder"), &folders)
	return folders, err
}

// AddSeries adds a new series
func (c *Client) AddSeries(series models.Series, rootFolder models.RootFolder, qualityProfile models.QualityProfile) (*models.Series, error) {
	addSeries := map[string]interface{}{
		"tvdbId":           series.TVDBID,
		"title":            series.Title,
		"qualityProfileId": qualityProfile.ID,
		"titleSlug":        series.TitleSlug,
		"rootFolderPath":   rootFolder.Path,
		"monitored":        true,
		"seasonFolder":     true,
		"addOptions": map[string]interface{}{
			"searchForMissingEpisodes": false,
		},
	}

	if c.apiVersion == "v3" {
		addSeries["languageProfileId"] = 1
		addSeries["monitorNewItems"] = "all"
	}

	var result models.Series
	err := c.post(c.endpoint("/series"), addSeries, &result)
	return &result, err
}

// UpdateSeries updates an existing series
func (c *Client) UpdateSeries(series models.Series) (*models.Series, error) {
	var result models.Series
	err := c.put(c.endpoint(fmt.Sprintf("/series/%d", series.ID)), series, &result)
	return &result, err
}

// ImportDownloads scans for downloaded episodes
func (c *Client) ImportDownloads(path string) error {
	command := map[string]interface{}{
		"name": "DownloadedEpisodesScan",
		"path": path,
	}
	return c.post(c.endpoint("/command"), command, nil)
}

// get performs a GET request
func (c *Client) get(endpoint string, result any) error {
	req, err := http.NewRequest("GET", c.baseURL+endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// post performs a POST request
func (c *Client) post(endpoint string, data any, result any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

// put performs a PUT request
func (c *Client) put(endpoint string, data any, result any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}

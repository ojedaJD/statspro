package mlb

import (
	"compress/gzip"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sports_api/db"
	"time"
)

var MLBSession *Client

type MLBResponse struct {
	Status     string
	StatusCode int
	Data       interface{}
	URL        string
	Headers    http.Header
}

func (r *MLBResponse) GetData() (interface{}, error) {
	if r.Data == nil {
		return nil, fmt.Errorf("no data available")
	}
	return r.Data, nil
}

type Client struct {
	HTTPClient     *http.Client
	BaseURL        string
	DefaultHeaders map[string]string
	Proxy          string
}

func NewMLBClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		BaseURL: "https://statsapi.mlb.com", // MLB Base URL
		DefaultHeaders: map[string]string{
			"User-Agent":      "Mozilla/5.0 (compatible; MLBBot/1.0)",
			"Accept":          "application/json",
			"Accept-Encoding": "gzip, deflate, br",
			"Connection":      "keep-alive",
		},
	}
}

func (c *Client) MLBGetRequest(endpoint string, params map[string]string, referer string, customHeaders map[string]string) (*MLBResponse, error) {
	ctx := context.Background()
	fullURL, err := prepareURL(c.BaseURL, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("error constructing request URL: %w", err)
	}

	redisClient := db.GetRedisClient()
	cachedData, err := redisClient.Get(ctx, fullURL)
	if err == nil {
		var cachedResponse MLBResponse
		if json.Unmarshal([]byte(cachedData), &cachedResponse) == nil {
			log.Println("Cache hit:", fullURL)
			return &cachedResponse, nil
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	headers := make(map[string]string)
	for k, v := range c.DefaultHeaders {
		headers[k] = v
	}
	for k, v := range customHeaders {
		headers[k] = v
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	if referer != "" {
		req.Header.Set("Referer", referer)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := parseResponse(resp.Body, resp.Header)
	if err != nil {
		return nil, err
	}

	response := &MLBResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Data:       data,
		URL:        fullURL,
		Headers:    resp.Header,
	}

	responseJSON, err := json.Marshal(response)
	if err == nil {
		if err := redisClient.Save(ctx, fullURL, responseJSON, 24*time.Hour); err != nil {
			log.Println("Failed to cache response:", err)
		}
	}

	return response, nil
}

func parseResponse(body io.ReadCloser, header http.Header) (interface{}, error) {
	defer body.Close()

	var reader io.Reader = body
	switch header.Get("Content-Encoding") {
	case "gzip":
		gzReader, _ := gzip.NewReader(body)
		defer gzReader.Close()
		reader = gzReader
	case "deflate":
		zlibReader, _ := zlib.NewReader(body)
		defer zlibReader.Close()
		reader = zlibReader
	}

	decoder := json.NewDecoder(reader)
	decoder.UseNumber()

	var result interface{}
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return result, nil
}

func prepareURL(baseURL, endpoint string, params map[string]string) (string, error) {
	u, err := url.Parse(baseURL + endpoint)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func init() {
	MLBSession = NewMLBClient()
}

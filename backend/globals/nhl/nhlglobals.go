package nhl

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

// NBASession is a globally accessible instance of NBAClient.
var NHLSession *Client

// NBAResponse represents the API response structure.
type NHLResponse struct {
	Status     string
	StatusCode int
	Data       interface{}
	URL        string
	Headers    http.Header
}

func (r *NHLResponse) GetData() (interface{}, error) {
	if r.Data == nil {
		return nil, fmt.Errorf("No Data")
	}
	return r.Data, nil
}

// NBAClient wraps an HTTP client with default headers for NBA API requests.
type Client struct {
	HTTPClient     *http.Client
	BaseURL        string
	DefaultHeaders map[string]string
	Proxy          string
}

func (c *Client) SetBaseUrl(newUrl string) {
	c.BaseURL = newUrl
}

// NewNHLClient initializes and returns an NHLClient instance.
func NewNHLClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		BaseURL: "https://api.nhle.com/", // Corrected base URL
		DefaultHeaders: map[string]string{
			"Host":            "api.nhle.com",
			"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36",
			"Accept":          "*/*",
			"Accept-Language": "en-US,en;q=0.9",
			"Accept-Encoding": "gzip, deflate, br, zstd",
			"Origin":          "https://www.nhl.com",
			"Referer":         "https://www.nhl.com/",
			"Connection":      "keep-alive",
			"Sec-Fetch-Dest":  "empty",
			"Sec-Fetch-Mode":  "cors",
			"Sec-Fetch-Site":  "cross-site",
		},
	}
}

// NBA Get Request constructs and executes an API request.
func (c *Client) NHLGetRequest(endpoint string, params map[string]string, referer string, customHeaders map[string]string) (*NHLResponse, error) {
	ctx := context.Background()

	fullURL, err := prepareURL(c.BaseURL, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("error constructing request URL: %w", err)
	}

	redisClient := db.GetRedisClient()

	// Check if response is cached in Redis
	cachedData, err := redisClient.Get(ctx, fullURL)
	if err == nil {
		// If cache exists, parse it and return
		var cachedResponse NHLResponse
		err := json.Unmarshal([]byte(cachedData), &cachedResponse)
		if err == nil {
			log.Println("Cache hit for URL:", fullURL)
			return &cachedResponse, nil
		}
		log.Println("Cache data invalid, making new request")
	} else {
		log.Println("Cache miss for URL:", fullURL)
	}

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	headers := make(map[string]string)
	for k, v := range c.DefaultHeaders {
		headers[k] = v // Copy default headers
	}
	if customHeaders != nil {
		for k, v := range customHeaders {
			headers[k] = v // Override with custom headers
		}
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Set Referer if provided
	if referer != "" {
		req.Header.Set("Referer", referer)
	}

	// Execute request
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error making request: %e", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}
	data, err := parseResponse(resp.Body, resp.Header)
	if err != nil {
		return nil, err
	}
	response := &NHLResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Data:       data,
		URL:        fullURL,
		Headers:    resp.Header,
	}

	// Serialize response to JSON for caching
	responseJSON, err := json.Marshal(response)
	if err == nil {
		// Save to Redis with a 1-day expiration
		err = redisClient.Save(ctx, fullURL, responseJSON, 24*time.Hour)
		if err != nil {
			log.Println("Failed to cache response in Redis:", err)
		}
	}

	return response, nil
}

// parseResponse reads and decodes the HTTP response body based on content encoding.
func parseResponse(body io.ReadCloser, header http.Header) (interface{}, error) {
	defer func() {
		if err := body.Close(); err != nil {
			fmt.Println("Warning: failed to close response body:", err)
		}
	}()

	var reader io.Reader = body
	contentEncoding := header.Get("Content-Encoding")

	switch contentEncoding {
	case "gzip":
		gzReader, err := gzip.NewReader(body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer func() {
			if err := gzReader.Close(); err != nil {
				fmt.Println("Warning: failed to close gzip reader:", err)
			}
		}()
		reader = gzReader
	case "deflate":
		zlibReader, err := zlib.NewReader(body)
		if err != nil {
			return nil, fmt.Errorf("failed to create zlib reader: %w", err)
		}
		defer func() {
			if err := zlibReader.Close(); err != nil {
				fmt.Println("Warning: failed to close zlib reader:", err)
			}
		}()
		reader = zlibReader
	}

	if header.Get("Content-Type") == "text/html; charset=utf-8" {
		htmlContent, err := io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("failed to read HTML response: %w", err)
		}
		return string(htmlContent), fmt.Errorf("NBA API returned HTML error page: %s", string(htmlContent))
	}

	// Decode JSON with UseNumber to avoid float precision issues
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()

	var result interface{}
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return result, nil
}

// PrepareURL constructs a full URL with query parameters.
func prepareURL(baseURL string, endpoint string, params map[string]string) (string, error) {
	// Parse the base URL
	u, err := url.Parse(baseURL + endpoint)
	if err != nil {
		return "", fmt.Errorf("error parsing base URL: %w", err)
	}

	if err != nil {
		return "", fmt.Errorf("error parsing base URL: %w", err)
	}

	// Add query parameters
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func init() {
	NHLSession = NewNHLClient()
}

package odds

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sports_api/db"
	"time"
)

var GlobalOddsClient *Client

func init() {
	apiKey := os.Getenv("ODDS_API_KEY") // Fetch API key from environment variable
	if apiKey == "" {
		fmt.Println("Warning: ODDS_API_KEY is not set")
	}
	GlobalOddsClient = NewOddsApiClient(apiKey)
}

// Client struct for The Odds API
type Client struct {
	HTTPClient     *http.Client
	BaseURL        string
	DefaultHeaders map[string]string
	APIKey         string
}

// OddsApiResponse defines the response structure
type OddsApiResponse struct {
	StatusCode int
	Status     string
	Data       interface{}
	URL        string
	Headers    http.Header
}

// NewOddsApiClient initializes a new API client
func NewOddsApiClient(apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		BaseURL: "https://api.the-odds-api.com/v4/sports",
		DefaultHeaders: map[string]string{
			"Accept":        "application/json",
			"User-Agent":    "Go-OddsAPI-Client",
			"Connection":    "keep-alive",
			"Cache-Control": "no-cache",
		},
		APIKey: apiKey,
	}
}

// CacheResponse stores API responses
func CacheResponse(url string, data *OddsApiResponse) error {
	cacheDir := "results"
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating cache directory: %w", err)
	}

	// Generate cache filename based on the URL
	hash := sha256.Sum256([]byte(url))
	filename := hex.EncodeToString(hash[:])
	cachePath := filepath.Join(cacheDir, filename)

	// Serialize and save response
	file, err := os.Create(cachePath)
	if err != nil {
		return fmt.Errorf("error creating cache file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("error writing cache file: %w", err)
	}

	return nil
}

// LoadCachedResponse checks if a response is cached and returns it
func LoadCachedResponse(url string) (*OddsApiResponse, error) {
	cacheDir := "results"
	hash := sha256.Sum256([]byte(url))
	filename := hex.EncodeToString(hash[:])
	cachePath := filepath.Join(cacheDir, filename)

	// Check if the file exists
	file, err := os.Open(cachePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // Cache miss
		}
		return nil, fmt.Errorf("error opening cache file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Read and decode the cached response
	var data OddsApiResponse
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error reading cache file: %w", err)
	}

	return &data, nil
}

// GetOddsRequest makes a request to The Odds API
func (c *Client) GetOddsRequest(fullUrl string, params map[string]string, customHeaders map[string]string) (*OddsApiResponse, error) {
	ctx := context.Background()

	if params == nil {
		params = make(map[string]string)
	}
	params["apiKey"] = c.APIKey
	// Prepare URL with query parameters

	fullURL, err := prepareURL(fullUrl, params)
	if err != nil {
		return nil, fmt.Errorf("error constructing request URL: %w", err)
	}

	redisClient := db.GetRedisClient()

	cachedData, err := redisClient.Get(ctx, fullURL)

	if err == nil {
		// If cache exists, parse it and return
		var cachedResponse OddsApiResponse
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

	// Execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	// Parse response
	data, err := parseResponse(resp.Body)
	if err != nil {
		return nil, err
	}

	// Prepare structured response
	apiResponse := &OddsApiResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Data:       data,
		URL:        fullURL,
		Headers:    resp.Header,
	}

	// Cache the response
	responseJSON, err := json.Marshal(apiResponse)

	if err == nil {
		// Save to Redis with a 1-day expiration
		err = redisClient.Save(ctx, fullURL, responseJSON, 24*time.Hour)
		if err != nil {
			log.Println("Failed to cache response in Redis:", err)
		}
	}
	return apiResponse, nil
}

func (c *Client) AppendSport(sport string) string {
	return fmt.Sprintf("%s/%s", c.BaseURL, sport)
}

// Helper function to construct a URL with query parameters
func prepareURL(fullURL string, params map[string]string) (string, error) {
	u, err := url.Parse(fullURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	query := u.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	u.RawQuery = query.Encode()
	return u.String(), nil
}

// Helper function to parse JSON response
func parseResponse(body io.ReadCloser) (interface{}, error) {
	var data interface{}
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %w", err)
	}
	return data, nil
}

// SetBaseURL allows changing the base URL dynamically
func (c *Client) SetBaseURL(newBaseURL string) {
	c.BaseURL = newBaseURL
}

// GetSportsURL returns the URL to fetch in-season sports
func (c *Client) GetSportsURL() string {
	return fmt.Sprintf("%s", c.BaseURL)
}

// GetSportOddsURL returns the URL to fetch odds for a specific sport
func (c *Client) GetSportOddsURL(sport string) string {
	return fmt.Sprintf("%s/%s/odds", c.BaseURL, sport)
}

// GetSportScoresURL returns the URL to fetch scores for a specific sport
func (c *Client) GetSportScoresURL(sport string) string {
	return fmt.Sprintf("%s/%s/scores", c.BaseURL, sport)
}

// GetSportEventsURL returns the URL to fetch events for a specific sport
func (c *Client) GetSportEventsURL(sport string) string {
	return fmt.Sprintf("%s/%s/events", c.BaseURL, sport)
}

// GetEventOddsURL returns the URL to fetch odds for a specific event
func (c *Client) GetEventOddsURL(sport string, eventId string) string {
	return fmt.Sprintf("%s/%s/events/%s/odds", c.BaseURL, sport, eventId)
}

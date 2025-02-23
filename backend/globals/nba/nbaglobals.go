package nba

import (
	"compress/gzip"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// NBASession is a globally accessible instance of NBAClient.
var NBASession *Client

// NBAResponse represents the API response structure.
type NBAResponse struct {
	*http.Response
	URL string
}

func (r *NBAResponse) isOk() bool {
	return r.StatusCode == 200
}

// ParseResponse parses the response body into an interface{} while checking encoding type.
func (r *NBAResponse) ParseResponse() (interface{}, error) {
	if r.Response == nil {
		return nil, fmt.Errorf("response is nil")
	}
	defer r.Body.Close()

	var reader io.Reader = r.Body

	// Check Content-Encoding for gzip or deflate
	contentEncoding := r.Header.Get("Content-Encoding")
	switch contentEncoding {
	case "gzip":
		gzReader, err := gzip.NewReader(r.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzReader.Close()
		reader = gzReader
	case "deflate":
		zlibReader, err := zlib.NewReader(r.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create zlib reader: %w", err)
		}
		defer zlibReader.Close()
		reader = zlibReader
	}

	// Decode JSON using UseNumber to preserve number formats
	decoder := json.NewDecoder(reader)
	decoder.UseNumber() // Treat numbers as json.Number instead of float64

	var result interface{}
	err := decoder.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return result, nil
}

// GetParameters extracts the parameters from the response.
func (r *NBAResponse) GetParameters() (map[string]interface{}, error) {
	data, err := r.ParseResponse()
	if err != nil {
		return nil, err
	}

	if response, ok := data.(map[string]interface{}); ok {
		if params, ok := response["parameters"].(map[string]interface{}); ok {
			return params, nil
		}
	}

	return nil, fmt.Errorf("parameters not found in response")
}

// GetResource extracts the resource field from the response.
func (r *NBAResponse) GetResource() (string, error) {
	data, err := r.ParseResponse()
	if err != nil {
		return "", err
	}

	if response, ok := data.(map[string]interface{}); ok {
		if resource, ok := response["resource"].(string); ok {
			return resource, nil
		}
	}

	return "", fmt.Errorf("resource not found in response")
}

// GetResultSets extracts the resultSets from the response.
func (r *NBAResponse) GetResultSets() ([]interface{}, error) {
	data, err := r.ParseResponse()
	if err != nil {
		return nil, err
	}

	if response, ok := data.(map[string]interface{}); ok {
		if resultSets, ok := response["resultSets"].([]interface{}); ok {
			return resultSets, nil
		}
	}

	return nil, fmt.Errorf("resultSets not found in response")
}

func (r *NBAResponse) isNil() bool {
	return r == nil
}

// GetNormalizedDict processes all resultSets concurrently and stores them in a map keyed by resultSet names.
func (r *NBAResponse) GetNormalizedDict() (map[string][]map[string]interface{}, error) {
	if r.isNil() {
		return nil, fmt.Errorf("error %d")
	}
	resultSets, err := r.GetResultSets()
	if err != nil {
		return nil, err
	}

	if len(resultSets) == 0 {
		return nil, fmt.Errorf("resultSets are empty")
	}

	// Map to store processed data keyed by resultSet name
	resultMap := make(map[string][]map[string]interface{})
	var resultMutex sync.Mutex // Mutex to prevent race conditions

	// WaitGroup and channel to process each resultSet concurrently
	var wg sync.WaitGroup
	resultChannel := make(chan struct{}, len(resultSets))

	// Process each resultSet concurrently
	for _, resultSet := range resultSets {
		rs, ok := resultSet.(map[string]interface{})
		if !ok {
			continue // Skip invalid resultSets
		}

		// Get resultSet name
		resultSetName, ok := rs["name"].(string)
		if !ok || resultSetName == "" {
			resultSetName = "Unknown" // Default name if missing
		}

		headers, ok := rs["headers"].([]interface{})
		if !ok {
			continue // Skip if headers are missing or invalid
		}

		rowSet, ok := rs["rowSet"].([]interface{})
		if !ok {
			continue // Skip if rowSet is missing or invalid
		}

		// Convert headers to []string
		headerKeys := make([]string, len(headers))
		for i, h := range headers {
			if headerStr, ok := h.(string); ok {
				headerKeys[i] = headerStr
			} else {
				headerKeys[i] = fmt.Sprintf("%v", h) // Convert non-string headers
			}
		}

		wg.Add(1)
		go func(headerKeys []string, rowSet []interface{}, resultSetName string) {
			defer wg.Done()
			var normalizedData []map[string]interface{}

			// Process rowSet concurrently using another set of goroutines
			var rowWG sync.WaitGroup
			rowChannel := make(chan map[string]interface{}, len(rowSet))

			for _, row := range rowSet {
				rowValues, ok := row.([]interface{})
				if !ok {
					continue
				}

				rowWG.Add(1)
				go func(rowValues []interface{}) {
					defer rowWG.Done()
					rowMap := make(map[string]interface{})
					for i, value := range rowValues {
						if i < len(headerKeys) {
							rowMap[headerKeys[i]] = value
						}
					}
					rowChannel <- rowMap
				}(rowValues)
			}

			// Close channel when all row goroutines finish
			go func() {
				rowWG.Wait()
				close(rowChannel)
			}()

			// Collect results from the channel
			for row := range rowChannel {
				normalizedData = append(normalizedData, row)
			}

			// Safely update the map
			resultMutex.Lock()
			resultMap[resultSetName] = normalizedData
			resultMutex.Unlock()

			resultChannel <- struct{}{} // Signal completion
		}(headerKeys, rowSet, resultSetName)
	}

	// Wait for all resultSets to be processed
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Wait for all goroutines to complete
	for range resultChannel {
	}

	return resultMap, nil
}

// GetAllHeaders extracts headers from all resultSets and returns them in a map.
func (r *NBAResponse) GetAllHeaders() (map[string][]string, error) {
	resultSets, err := r.GetResultSets()
	if err != nil {
		return nil, err
	}

	if len(resultSets) == 0 {
		return nil, fmt.Errorf("resultSets are empty")
	}
	// If only one resultSet, process it directly
	if len(resultSets) == 1 {
		rs, ok := resultSets[0].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid resultSet format")
		}
		return processSingleHeaders(rs)
	}

	return processHeadersConcurrently(resultSets)
}

// GetResultSetNames extracts all resultSet names and returns them as a slice.
func (r *NBAResponse) GetResultSetNames() ([]string, error) {
	resultSets, err := r.GetResultSets()
	if err != nil {
		return nil, err
	}

	if len(resultSets) == 0 {
		return nil, fmt.Errorf("resultSets are empty")
	}

	if len(resultSets) == 1 {
		rs, ok := resultSets[0].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid resultSet format")
		}
		name, exists := rs["name"].(string)
		if !exists {
			name = "Unknown"
		}
		return []string{name}, nil
	}

	return processResultSetNamesConcurrently(resultSets)

}

// GetRowSets extracts rowSet data from all resultSets and returns them in a map.
func (r *NBAResponse) GetRowSets() (map[string][][]interface{}, error) {
	resultSets, err := r.GetResultSets()
	if err != nil {
		return nil, err
	}

	if len(resultSets) == 0 {
		return nil, fmt.Errorf("resultSets are empty")
	}

	if len(resultSets) == 1 {
		rs, ok := resultSets[0].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid resultSet format")
		}
		return processSingleRowSet(rs)
	}

	return processRowSetConcurrently(resultSets)

}

// Singles
// GetHeaders extracts the headers from the first resultSet.
func (r *NBAResponse) GetHeaders() ([]interface{}, error) {
	resultSets, err := r.GetResultSets()
	if err != nil {
		return nil, err
	}

	if len(resultSets) > 0 {
		if resultSet, ok := resultSets[0].(map[string]interface{}); ok {
			if headers, ok := resultSet["headers"].([]interface{}); ok {
				return headers, nil
			}
		}
	}

	return nil, fmt.Errorf("headers not found in resultSets")
}

// GetRowSet extracts the rowSet data from the first resultSet.
func (r *NBAResponse) GetRowSet() ([]interface{}, error) {
	resultSets, err := r.GetResultSets()
	if err != nil {
		return nil, err
	}

	if len(resultSets) > 0 {
		if resultSet, ok := resultSets[0].(map[string]interface{}); ok {
			if rowSet, ok := resultSet["rowSet"].([]interface{}); ok {
				return rowSet, nil
			}
		}
	}

	return nil, fmt.Errorf("rowSet not found in resultSets")
}

// GetResponse returns the raw response as a string.
func (r *NBAResponse) GetResponse() *http.Response {
	return r.Response
}

// NBAClient wraps an HTTP client with default headers for NBA API requests.
type Client struct {
	HTTPClient     *http.Client
	BaseURL        string
	DefaultHeaders map[string]string
	Proxy          string
}

// NewNBAClient initializes and returns an NBAClient instance.
func NewNBAClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		BaseURL: "https://stats.nba.com/stats/",
		DefaultHeaders: map[string]string{
			"Host":               "stats.nba.com",
			"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0",
			"Accept":             "application/json, text/plain, */*",
			"Accept-Language":    "en-US,en;q=0.5",
			"Accept-Encoding":    "gzip, deflate, br",
			"x-nba-stats-origin": "stats",
			"x-nba-stats-token":  "true",
			"Connection":         "keep-alive",
			"Referer":            "https://stats.nba.com/",
			"Pragma":             "no-cache",
			"Cache-Control":      "no-cache",
		},
	}
}

// SetProxy sets a proxy for the client.
func (c *Client) SetProxy(proxyURL string) {
	c.Proxy = proxyURL
}

// NBA Get Request constructs and executes an API request.
func (c *Client) NBAGetRequest(endpoint string, params map[string]string, referer string, customHeaders map[string]string) (*NBAResponse, error) {

	fullURL, err := prepareURL(c.BaseURL, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("error constructing request URL: %w", err)
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
	parseResponse(resp.Body, resp.Header)
	// Return structured response
	return &NBAResponse{
		Response: resp,

		URL: fullURL,
	}, nil
}

func parseResponse(body io.ReadCloser, header http.Header) {

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
	NBASession = NewNBAClient()
}

package nba

import (
	"fmt"
	"sync"
)

// processSingleRowSet extracts rowSet for a single resultSet
func processSingleRowSet(rs map[string]interface{}) (map[string][][]interface{}, error) {
	rowSetsMap := make(map[string][][]interface{})

	name, ok := rs["name"].(string)
	if !ok {
		name = "Unknown"
	}

	rowSet, ok := rs["rowSet"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("rowSet missing in resultSet")
	}

	rows := make([][]interface{}, len(rowSet))
	for i, row := range rowSet {
		if rowArr, ok := row.([]interface{}); ok {
			rows[i] = rowArr
		}
	}

	rowSetsMap[name] = rows
	return rowSetsMap, nil
}

// processRowSetConcurrently extracts rowSet concurrently
func processRowSetConcurrently(resultSets []interface{}) (map[string][][]interface{}, error) {
	rowSetsMap := make(map[string][][]interface{})
	var rowSetsMutex sync.Mutex
	var wg sync.WaitGroup

	// Process each resultSet concurrently
	for _, resultSet := range resultSets {
		rs, ok := resultSet.(map[string]interface{})
		if !ok {
			continue
		}

		name, ok := rs["name"].(string)
		if !ok {
			name = "Unknown" // Default name if missing
		}

		rowSet, ok := rs["rowSet"].([]interface{})
		if !ok {
			continue
		}

		// Convert rowSet to [][]interface{}
		rows := make([][]interface{}, len(rowSet))
		for i, row := range rowSet {
			if rowArr, ok := row.([]interface{}); ok {
				rows[i] = rowArr
			}
		}

		wg.Add(1)
		go func(name string, rows [][]interface{}) {
			defer wg.Done()
			rowSetsMutex.Lock()
			rowSetsMap[name] = rows
			rowSetsMutex.Unlock()
		}(name, rows)
	}

	wg.Wait()
	return rowSetsMap, nil
}

// processResultSetNamesConcurrently extracts rowSet concurrently
func processResultSetNamesConcurrently(resultSets []interface{}) ([]string, error) {
	var names []string
	var namesMutex sync.Mutex
	var wg sync.WaitGroup

	// Process each resultSet concurrently
	for _, resultSet := range resultSets {
		rs, ok := resultSet.(map[string]interface{})
		if !ok {
			continue
		}

		name, ok := rs["name"].(string)
		if !ok {
			name = "Unknown" // Default name if missing
		}

		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			namesMutex.Lock()
			names = append(names, name)
			namesMutex.Unlock()
		}(name)
	}

	wg.Wait()
	return names, nil
}

// processSingleHeaders extracts headers for a single resultSet
func processSingleHeaders(rs map[string]interface{}) (map[string][]string, error) {
	headersMap := make(map[string][]string)

	name, ok := rs["name"].(string)
	if !ok {
		name = "Unknown"
	}

	headers, ok := rs["headers"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("headers missing in resultSet")
	}

	headerKeys := make([]string, len(headers))
	for i, h := range headers {
		if headerStr, ok := h.(string); ok {
			headerKeys[i] = headerStr
		} else {
			headerKeys[i] = fmt.Sprintf("%v", h)
		}
	}

	headersMap[name] = headerKeys
	return headersMap, nil
}

// processHeadersConcurrently extracts headers concurrently
func processHeadersConcurrently(resultSets []interface{}) (map[string][]string, error) {
	headersMap := make(map[string][]string)
	var headersMutex sync.Mutex
	var wg sync.WaitGroup

	// Process each resultSet concurrently
	for _, resultSet := range resultSets {
		rs, ok := resultSet.(map[string]interface{})
		if !ok {
			continue
		}

		name, ok := rs["name"].(string)
		if !ok {
			name = "Unknown" // Default name if missing
		}

		headers, ok := rs["headers"].([]interface{})
		if !ok {
			continue
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
		go func(name string, headerKeys []string) {
			defer wg.Done()
			headersMutex.Lock()
			headersMap[name] = headerKeys
			headersMutex.Unlock()
		}(name, headerKeys)
	}

	wg.Wait()
	return headersMap, nil
}

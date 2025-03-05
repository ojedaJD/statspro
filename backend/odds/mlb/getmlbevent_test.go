package mlb

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestGetAllNBAEventProps(t *testing.T) {
	props, err := GetPlayerProps()
	if err != nil {
		t.Fatalf("Failed to get player props: %v", err) // Fail the test if there's an error
	}

	// Iterate over each matchup and print details line by line
	for i, prop := range props {
		fmt.Printf("\nMatchup %d: %s vs. %s\n", i+1, prop.HomeTeam, prop.AwayTeam)
		fmt.Printf("  Commence Time: %s\n", prop.CommenceTime)
		fmt.Printf("  Event ID: %s\n", prop.Id)
		fmt.Printf("  Sport: %s (%s)\n", prop.SportTitle, prop.SportKey)

		// Iterate over bookmakers
		for _, bookmaker := range prop.Bookmakers {
			fmt.Printf("  Bookmaker: %s (Key: %s)\n", bookmaker.Title, bookmaker.Key)

			// Iterate over markets
			for _, market := range bookmaker.Markets {
				fmt.Printf("    Market: %s (Last Update: %s)\n", market.Key, market.LastUpdate)

				// Iterate over outcomes
				for _, outcome := range market.Outcomes {
					fmt.Printf("      Player: %s - %s\n", outcome.Name, outcome.Description)
					fmt.Printf("      Point: %.2f | Price: %d\n", outcome.Point, outcome.Price)
				}
			}
		}
		fmt.Println("--------------------------------------------------")
	}
}

func TestGetAllMLBEvents(t *testing.T) {
	// Replace with your actual API key or use a mock/stub API key
	// Call GetAllNBAEvents
	response, err := GetAllMLBEvents()
	if err != nil {
		t.Fatalf("Failed to fetch NBA events: %v", err)
	}

	// Print the API response (for debugging purposes)
	fmt.Printf("Response: %+v\n", response)

	// Check if status code is 200
	if response.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", response.StatusCode)
	}

	fmt.Printf("%T", response.Data)

	// Ensure the response contains data
	if response.Data == nil {
		t.Errorf("Expected non-nil data, got nil")
	}

	// Optionally, verify that the response contains expected event data
	if len(response.Data.([]interface{})) == 0 {
		t.Errorf("Expected at least one NBA event, got none")
	}

	jsonData, err := json.MarshalIndent(response.Data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling response data to JSON: %v", err)
	}

	// Print formatted JSON output
	fmt.Println("NBA Events JSON Response:")
	fmt.Println(string(jsonData))
}

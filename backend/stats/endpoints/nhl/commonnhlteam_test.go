package nhl

import "testing"

func TestGetNHLTeams(t *testing.T) {
	// Call the function that hits the real NHL API endpoint
	resp, err := GetNHlTeams()
	if err != nil {
		t.Fatalf("Error calling GetNHLTeams: %v", err)
	}
	if resp == nil {
		t.Fatal("Expected a response, got nil")
	}

	// Optionally, you can add more checks here to validate
	// certain fields in resp if you know what shape the data
	// should have (e.g., length of teams, etc.)
	t.Logf("Successfully fetched NHL teams. Response: %+v", resp.Data)

	GetAndParseNHLTeams()
}

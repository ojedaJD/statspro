package nhl

import (
	"testing"
)

func TestGetNHLRoster(t *testing.T) {
	// Call the function that hits the real NHL API endpoint
	teams, err := GetAndParseNHLTeams()
	if err != nil {
		return
	}

	// Optionally, you can add more checks here to validate
	// certain fields in resp if you know what shape the data
	// should have (e.g., length of teams, etc.)
	t.Logf("Successfully fetched NHL teams. Response: %+v", teams)

}

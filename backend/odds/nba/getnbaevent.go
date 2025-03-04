package nba

import (
	"encoding/json"
	"fmt"
	"log"
	"sports_api/globals/odds"
	"time"
)

type OddsEvent struct {
	AwayTeam     string    `json:"away_team"`
	CommenceTime time.Time `json:"commence_time"`
	HomeTeam     string    `json:"home_team"`
	Id           string    `json:"id"`
	SportKey     string    `json:"sport_key"`
	SportTitle   string    `json:"sport_title"`
}

func GetAllNBAEvents() (*odds.OddsApiResponse, error) {
	// Construct the URL for fetching events

	eventsURL := odds.GlobalOddsClient.GetSportEventsURL("basketball_nba")

	// No query parameters needed for this endpoint
	params := map[string]string{}

	// Make API request
	response, err := odds.GlobalOddsClient.GetOddsRequest(eventsURL, params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events for sport %s: %w", "nba", err)
	}

	return response, nil
}

// GetParsedNBAEvents calls GetAllNBAEvents and unmarshals the data into a slice of NBAOddsEvent
func GetandUnmarshallAllNBAEvents() ([]OddsEvent, error) {
	// Fetch NBA events
	response, err := GetAllNBAEvents()
	if err != nil {
		return nil, err
	}

	// Marshal the response data into JSON bytes
	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	// Unmarshal JSON bytes into an array of NBAOddsEvent
	var events []OddsEvent
	if err := json.Unmarshal(jsonData, &events); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NBA events: %w", err)
	}

	return events, nil
}

// GetPlayerProps fetches player props for all NBA events
func GetPlayerProps() (MatchupOddsSlice, error) {
	// Get all NBA events
	events, err := GetandUnmarshallAllNBAEvents()
	if err != nil {
		return nil, fmt.Errorf("failed to get NBA events: %w", err)
	}

	var matchupOdds []MatchupOdds

	// Loop through each event and fetch player props
	for _, event := range events {

		// Construct URL for player props using the event ID
		playerPropsURL := odds.GlobalOddsClient.GetEventOddsURL("basketball_nba", event.Id)

		// No query parameters needed
		params := map[string]string{"regions": "us",
			"markets":    "player_points",
			"oddsFormat": "american"}

		params = map[string]string{"regions": "us",
			"markets":    "player_points,player_rebounds,player_assists,player_steals,player_points_assists,player_points_rebounds_assists,player_points_rebounds,player_turnovers,player_blocks,player_threes",
			"oddsFormat": "american"}

		// Make API request for player props
		response, err := odds.GlobalOddsClient.GetOddsRequest(playerPropsURL, params, nil)
		if err != nil {
			log.Printf("Failed to fetch player props for event %s: %v", event.Id, err)
			continue // Skip this event and continue to the next one
		}

		var m MatchupOdds
		responseBytes, _ := json.Marshal(response.Data)
		if err := json.Unmarshal(responseBytes, &m); err != nil {
			log.Printf("Failed to unmarshal player props for event %s: %v", event.Id, err)
			continue
		}

		// Append parsed response to the results array
		matchupOdds = append(matchupOdds, m)
		// Append response to the results array

	}

	return matchupOdds, nil
}

type MatchupOddsSlice []MatchupOdds

func (r MatchupOddsSlice) GetOddsByHomeAndAwayTeam(homeTeamName, awayTeamName string) *MatchupOdds {
	for i, _ := range r {
		if r[i].HomeTeam == homeTeamName && r[i].AwayTeam == awayTeamName {
			return &r[i]
		}
	}
	return nil
}

type MatchupOdds struct {
	AwayTeam   string `json:"away_team"`
	Bookmakers []struct {
		Key     string `json:"key"`
		Markets []struct {
			Key        string    `json:"key"`
			LastUpdate time.Time `json:"last_update"`
			Outcomes   []struct {
				Description string  `json:"description"`
				Name        string  `json:"name"`
				Point       float64 `json:"point"`
				Price       int     `json:"price"`
			} `json:"outcomes"`
		} `json:"markets"`
		Title string `json:"title"`
	} `json:"bookmakers"`
	CommenceTime time.Time `json:"commence_time"`
	HomeTeam     string    `json:"home_team"`
	Id           string    `json:"id"`
	SportKey     string    `json:"sport_key"`
	SportTitle   string    `json:"sport_title"`
}

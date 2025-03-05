package nhl

import (
	"encoding/json"
	"fmt"
	"sports_api/globals/nhl"
)

type Player struct {
	FirstName struct {
		Default string `json:"default"`
	} `json:"firstName"`
	Id       int `json:"id"`
	LastName struct {
		Default string `json:"default"`
	} `json:"lastName"`
	PositionCode  string `json:"positionCode"`
	ShootsCatches string `json:"shootsCatches"`
	SweaterNumber int    `json:"sweaterNumber"`
	GameLogs      []GameLog
	GoalieLogs    []GoalieGameLog
	Odds          map[string]map[string][]Outcome `json:"odds,omitempty"`
}
type Outcome struct {
	Name  string  `json:"name"`
	Point float64 `json:"point"`
	Price int     `json:"price"`
}

// SetOutcome method to add/update an outcome for a player
func (p *Player) SetOutcome(bookmaker, outcomeType, name string, point float64, price int) *Player {
	// Initialize Odds map if nil
	if p.Odds == nil {
		p.Odds = make(map[string]map[string][]Outcome)
	}

	// Initialize nested map for the bookmaker if nil
	if p.Odds[outcomeType] == nil {
		p.Odds[outcomeType] = make(map[string][]Outcome)
	}

	// Create a new outcome
	outcome := Outcome{
		Name:  name,
		Point: point,
		Price: price,
	}

	// Append the outcome to the existing slice
	p.Odds[outcomeType][bookmaker] = append(p.Odds[outcomeType][bookmaker], outcome)

	return p
}

// GetNHLRoster SeasonID in form 20242025
func (t *NHLTeam) GetRoster(seasonID string) error {
	nhl.NHLSession.SetBaseUrl("https://api-web.nhle.com/v1/")

	resp, err := nhl.NHLSession.NHLGetRequest(fmt.Sprintf("roster/%s/%s", t.Abbreviation, seasonID), nil, "", nil)
	if err != nil {
		return err
	}

	var subRoster []Player
	roster := resp.Data.(map[string]interface{})
	for _, v := range roster {
		marshal, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = json.Unmarshal(marshal, &subRoster)
		if err != nil {
			return err
		}
		t.Roster = append(t.Roster, subRoster...)
	}

	return nil

}

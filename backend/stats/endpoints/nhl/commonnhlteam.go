package nhl

import (
	"fmt"
	"sports_api/globals/nhl"
	"strings"
	"sync"
)

var NHLTeamAbbreviations = map[string]string{
	"Anaheim Ducks":         "ANA",
	"Arizona Coyotes":       "ARI",
	"Boston Bruins":         "BOS",
	"Buffalo Sabres":        "BUF",
	"Calgary Flames":        "CGY",
	"Carolina Hurricanes":   "CAR",
	"Chicago Blackhawks":    "CHI",
	"Colorado Avalanche":    "COL",
	"Columbus Blue Jackets": "CBJ",
	"Dallas Stars":          "DAL",
	"Detroit Red Wings":     "DET",
	"Edmonton Oilers":       "EDM",
	"Florida Panthers":      "FLA",
	"Los Angeles Kings":     "LAK",
	"Minnesota Wild":        "MIN",
	"Montréal Canadiens":    "MTL",
	"Nashville Predators":   "NSH",
	"New Jersey Devils":     "NJD",
	"New York Islanders":    "NYI",
	"New York Rangers":      "NYR",
	"Ottawa Senators":       "OTT",
	"Philadelphia Flyers":   "PHI",
	"Pittsburgh Penguins":   "PIT",
	"San Jose Sharks":       "SJS",
	"Seattle Kraken":        "SEA",
	"St Louis Blues":        "STL",
	"Tampa Bay Lightning":   "TBL",
	"Toronto Maple Leafs":   "TOR",
	"Vancouver Canucks":     "VAN",
	"Vegas Golden Knights":  "VGK",
	"Washington Capitals":   "WSH",
	"Winnipeg Jets":         "WPG",
	"Utah Hockey Club":      "UTA",
}

func (team *NHLTeam) SetAbbreviation() {
	if abbr, exists := NHLTeamAbbreviations[team.FullName]; exists {
		team.Abbreviation = abbr
	}
}

type NHLTeam struct {
	FullName       string `json:"fullName"`
	Id             int    `json:"id"`
	TeamCommonName string `json:"teamCommonName"`
	TeamPlaceName  string `json:"teamPlaceName"`
	Abbreviation   string
	Roster         []Player
}

func GetNHlTeams() (*nhl.NHLResponse, error) {
	nhl.NHLSession.ResetBaseURL()
	params := map[string]string{
		"include": "lastSeason.id",
	}

	return nhl.NHLSession.NHLGetRequest("stats/rest/en/franchise", params, "", nil)
}

type NHLTeams []*NHLTeam

func (t NHLTeams) GetTeamByFullName(fullName string) *NHLTeam {
	for _, team := range t {
		if team.FullName == fullName {
			return team
		}
	}
	return nil
}

func (t *NHLTeam) GetPlayerByFullName(fullName string) *Player {

	for i := range t.Roster {
		if fmt.Sprintf("%s %s", t.Roster[i].FirstName.Default, t.Roster[i].LastName.Default) == fullName {
			return &t.Roster[i]
		}
	}
	return nil
}

func GetAndParseNHLTeams() (NHLTeams, error) {
	teams, err := GetNHlTeams()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var activeTeams []*NHLTeam

	m := teams.Data.(map[string]interface{})["data"].([]interface{})
	for _, i := range m {
		wg.Add(1)

		go func(teamData interface{}) {
			defer wg.Done()

			m2 := teamData.(map[string]interface{})
			if m2["lastSeason"] == nil {

				fullName := fmt.Sprintf("%v", m2["fullName"])
				if strings.Contains(fullName, "St.") {
					fullName = strings.ReplaceAll(fullName, "St.", "St")
				}
				team := &NHLTeam{
					FullName:       fullName,
					Id:             int(m2["id"].(float64)),
					TeamCommonName: fmt.Sprintf("%v", m2["teamCommonName"]),
					TeamPlaceName:  fmt.Sprintf("%v", m2["teamPlaceName"]),
				}
				team.SetAbbreviation()
				err := team.GetRoster("20242025")
				if err != nil {
					fmt.Println("Could not set roster for team", team.TeamCommonName, team.TeamPlaceName)
				}

				for i := range team.Roster {
					err := team.Roster[i].GetGameLog("20242025", 2)
					if err != nil {
						fmt.Println("failed to get Gamelog for player")
					}
				}

				mu.Lock()
				activeTeams = append(activeTeams, team)
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	return activeTeams, nil
}

package nhl

import (
	"fmt"
	nhl2 "sports_api/odds/nhl"
	"sports_api/stats/endpoints/nhl"
)

type NHLMatchup struct {
	Home *nhl.NHLTeam
	Away *nhl.NHLTeam
}

func GetNHLMatchupsWithOdds() []NHLMatchup {
	teams, err := nhl.GetAndParseNHLTeams()
	if err != nil {
		return nil
	}
	props, err := nhl2.GetPlayerProps()
	if err != nil || props == nil {
		return nil
	}
	var nhlMatchups []NHLMatchup
	for _, prop := range props {
		home := teams.GetTeamByFullName(prop.HomeTeam)
		if home == nil {
			fmt.Println("Could not locate team", prop.HomeTeam)
			continue
		}
		away := teams.GetTeamByFullName(prop.AwayTeam)
		if away == nil {
			fmt.Println("Could not locate team", prop.AwayTeam)
			continue
		}
		for _, bookmaker := range prop.Bookmakers {
			for _, market := range bookmaker.Markets {
				for _, outcome := range market.Outcomes {

					if player := home.GetPlayerByFullName(outcome.Description); player != nil {
						player.SetOutcome(bookmaker.Key, market.Key, outcome.Name, outcome.Point, outcome.Price)
					} else if player := away.GetPlayerByFullName(outcome.Description); player != nil {
						player.SetOutcome(bookmaker.Key, market.Key, outcome.Name, outcome.Point, outcome.Price)
					} else {
						fmt.Println("fail to find ", outcome.Description)
					}

				}

			}
		}
		nhlMatchups = append(nhlMatchups, NHLMatchup{
			Home: home,
			Away: away,
		})
	}
	return nhlMatchups
}

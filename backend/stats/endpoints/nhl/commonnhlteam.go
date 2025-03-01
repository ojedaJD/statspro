package nhl

import (
	"fmt"
	"sports_api/globals/nhl"
)

func GetNHlTeams() (*nhl.NHLResponse, error) {

	params := map[string]string{
		"include": "lastSeason.id",
	}

	return nhl.NHLSession.NHLGetRequest("stats/rest/en/franchise", params, "", nil)
}

func GetAndParseNHLTeams() {
	teams, err := GetNHlTeams()
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	m := teams.Data.(map[string]interface{})["data"].([]interface{})
	for _, i := range m {
		m2 := i.(map[string]interface{})
		if m2["lastSeason"] == nil {
			count++
		}
	}
	fmt.Println(count)

}

//r.json()["data"]

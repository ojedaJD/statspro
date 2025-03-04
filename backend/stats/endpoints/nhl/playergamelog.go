package nhl

import (
	"fmt"
	"sports_api/globals/nhl"
)

func GetNHlPlayerGameLog(playerID, seasonType int, seasonYear string) (*nhl.NHLResponse, error) {

	nhl.NHLSession.SetBaseUrl("https://api-web.nhle.com/v1/")

	return nhl.NHLSession.NHLGetRequest(fmt.Sprintf("player/%s/game-log/%s/%s", playerID, seasonYear, seasonType), nil, "", nil)
}

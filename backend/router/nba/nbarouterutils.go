package nba

import "sports_api/stats/endpoints/nba"

// CalculateHitRate computes the percentage of games where the player's stat met/exceeded the threshold.
func CalculateHitRate(gameLogs nba.BaseGameLogSlice, stat string, numGames int, threshold float64) float64 {
	if len(gameLogs) == 0 {
		return 0
	}

	// Take only the last N games
	if len(gameLogs) > numGames {
		gameLogs = gameLogs[len(gameLogs)-numGames:]
	}

	totalGames := len(gameLogs)
	hitCount := 0

	for _, log := range gameLogs {
		var statValue int
		switch stat {
		case "PTS":
			statValue = log.PTS
		case "REB":
			statValue = log.REB
		case "AST":
			statValue = log.AST
		case "FG3M":
			statValue = log.FG3M
		default:
			continue
		}

		if float64(statValue) >= threshold {
			hitCount++
		}
	}

	if totalGames == 0 {
		return 0
	}
	return float64(hitCount) / float64(totalGames) * 100
}

// CalculateStreak determines the current streak of consecutive games meeting the threshold.
func CalculateStreak(gameLogs nba.BaseGameLogSlice, stat string, threshold float64) int {
	if len(gameLogs) == 0 {
		return 0
	}

	streak := 0

	// Iterate from the most recent game to the oldest
	for i := len(gameLogs) - 1; i >= 0; i-- {
		log := gameLogs[i]
		var statValue int

		switch stat {
		case "PTS":
			statValue = log.PTS
		case "REB":
			statValue = log.REB
		case "AST":
			statValue = log.AST
		case "FG3M":
			statValue = log.FG3M
		default:
			continue
		}

		if float64(statValue) >= threshold {
			streak++
		} else {
			break // Streak ends if the player doesn't meet the threshold
		}
	}

	return streak
}

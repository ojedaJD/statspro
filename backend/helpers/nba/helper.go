package nba

import (
	"errors"
	"regexp"
)

// Precompiled regex patterns for validation
var (
	validLeagueIDs           = map[string]bool{"00": true, "10": true, "20": true} // NBA, WNBA, G-League
	validSeasonTypes         = regexp.MustCompile(`^(Regular Season|Pre Season|Playoffs|All Star)$`)
	validPerMode             = regexp.MustCompile(`^(Totals|PerGame|MinutesPer|Per48|Per40|Per36|PerMinute|PerPossession|PerPlay|Per100Possessions|Per100Plays)$`)
	validLocations           = regexp.MustCompile(`^(Home|Road)$`)
	validConferences         = regexp.MustCompile(`^(East|West)$`)
	validDivisions           = regexp.MustCompile(`^(Atlantic|Central|Northwest|Pacific|Southeast|Southwest|East|West)$`)
	validStarterBench        = regexp.MustCompile(`^(Starters|Bench)?$`)
	validSeasonSegment       = regexp.MustCompile(`^(Post All-Star|Pre All-Star)?$`)
	validPlayerPosition      = regexp.MustCompile(`^(F|C|G|C-F|F-C|F-G|G-F)?$`)
	validPlayerExperience    = regexp.MustCompile(`^(Rookie|Sophomore|Veteran)?$`)
	validGameScope           = regexp.MustCompile(`^(Yesterday|Last 10)?$`)
	validGameIDPattern       = regexp.MustCompile(`^(\d{10})(,\d{10})*$`)
	validPlayerOrTeam        = regexp.MustCompile(`^(Player|Team)$`)
	validPlayerScope         = regexp.MustCompile(`^(All Players|Rookies)$`)
	validSeasonYear          = regexp.MustCompile(`^\d{4}$`) // Ensures "YYYY" format (e.g., "2019")
	validSeasonYearOrAllTime = regexp.MustCompile(`^(\d{4}-\d{2})|(All Time)$`)
)

// ValidateLeagueID checks if the given LeagueID is valid.
func ValidateLeagueID(leagueID string) (bool, error) {
	if _, exists := validLeagueIDs[leagueID]; !exists {
		return false, errors.New("invalid LeagueID: must be '00' (NBA), '10' (WNBA), or '20' (G-League')")
	}
	return true, nil
}

// ValidateSeason checks if the given Season format is valid using regex.
func ValidateSeason(season string) (bool, error) {
	if !validSeasonYearOrAllTime.MatchString(season) {
		return false, errors.New("invalid Season format: must be 'YYYY-YY' (e.g., '2023-24')")
	}
	return true, nil
}

// ValidateSeasonType checks if the given season type is valid.
func ValidateSeasonType(seasonType string) (bool, error) {
	if !validSeasonTypes.MatchString(seasonType) {
		return false, errors.New("invalid SeasonType: must be 'Regular Season', 'Pre Season', 'Playoffs', or 'All Star'")
	}
	return true, nil
}

// ValidatePerMode checks if the given perMode is valid.
func ValidatePerMode(perMode string) (bool, error) {
	if !validPerMode.MatchString(perMode) {
		return false, errors.New("invalid PerMode: must be 'Totals' or 'PerGame'")
	}
	return true, nil
}

// ValidateLocation checks if the given Location is valid.
func ValidateLocation(location string) (bool, error) {
	if !validLocations.MatchString(location) {
		return false, errors.New("invalid Location: must be 'Home' or 'Road'")
	}
	return true, nil
}

// ValidateConference checks if the given Conference is valid.
func ValidateConference(conference string) (bool, error) {
	if !validConferences.MatchString(conference) {
		return false, errors.New("invalid Conference: must be 'East' or 'West'")
	}
	return true, nil
}

// ValidateDivision checks if the given Division is valid.
func ValidateDivision(division string) (bool, error) {
	if !validDivisions.MatchString(division) {
		return false, errors.New("invalid Division: must be a valid NBA division")
	}
	return true, nil
}

// ValidateStarterBench checks if the given StarterBench is valid.
func ValidateStarterBench(starterBench string) (bool, error) {
	if !validStarterBench.MatchString(starterBench) {
		return false, errors.New("invalid StarterBench: must be 'Starters' or 'Bench'")
	}
	return true, nil
}

// ValidateOutcome checks if the given Outcome is valid.
func ValidateOutcome(outcome string) (bool, error) {
	if !(outcome == "W" || outcome == "L") {
		return false, errors.New("invalid Outcome: must be 'W' or 'L'")
	}
	return true, nil
}

// ValidateSeasonSegment checks if the given SeasonSegment is valid.
func ValidateSeasonSegment(seasonSegment string) (bool, error) {
	if !validSeasonSegment.MatchString(seasonSegment) {
		return false, errors.New("invalid SeasonSegment: must be 'Post All-Star' or 'Pre All-Star'")
	}
	return true, nil
}

// ValidatePlayerPosition checks if the given PlayerPosition is valid.
func ValidatePlayerPosition(playerPosition string) (bool, error) {
	if !validPlayerPosition.MatchString(playerPosition) {
		return false, errors.New("invalid PlayerPosition: must be 'F', 'C', 'G', 'C-F', 'F-C', 'F-G', or 'G-F'")
	}
	return true, nil
}

// ValidatePlayerExperience checks if the given PlayerExperience is valid.
func ValidatePlayerExperience(playerExperience string) (bool, error) {
	if !validPlayerExperience.MatchString(playerExperience) {
		return false, errors.New("invalid PlayerExperience: must be 'Rookie', 'Sophomore', or 'Veteran'")
	}
	return true, nil
}

// ValidateGameScope checks if the given GameScope is valid.
func ValidateGameScope(gameScope string) (bool, error) {
	if !validGameScope.MatchString(gameScope) {
		return false, errors.New("invalid GameScope: must be 'Yesterday' or 'Last 10'")
	}
	return true, nil
}

// IsPositive checks if the given integer is positive and returns an error if invalid.
func IsPositive(x int) (bool, error) {
	if x <= 0 {
		return false, errors.New("invalid value: must be a positive integer")
	}
	return true, nil
}

// ValidatePlayerID ensures the playerID is not empty.
func ValidatePlayerID(playerID string) (bool, error) {
	if playerID == "" {
		return false, errors.New("PlayerID is required and cannot be empty")
	}
	return true, nil
}

// ValidateTeamID checks if the given TeamID is not empty.
func ValidateTeamID(teamID string) (bool, error) {
	if teamID == "" {
		return false, errors.New("TeamID is required and cannot be empty")
	}
	return true, nil
}

// ValidatePlayerScope checks if the given PlayerScope is valid.
func ValidatePlayerScope(playerScope string) (bool, error) {
	if !validPlayerScope.MatchString(playerScope) {
		return false, errors.New("invalid PlayerScope: must be 'All Players' or 'Rookies'")
	}
	return true, nil
}

// ValidateGameIDs checks if the given GameIDs string is in a valid format.
func ValidateGameIDs(gameIDs string) (bool, error) {
	if !validGameIDPattern.MatchString(gameIDs) {
		return false, errors.New("invalid GameIDs format: must be a 10-digit GameID or multiple GameIDs separated by commas")
	}
	return true, nil
}

// ValidateGameID ensures the GameID follows the correct format (10-digit numeric string).
func ValidateGameID(gameID string) (bool, error) {
	if !validGameIDPattern.MatchString(gameID) {
		return false, errors.New("invalid GameID: must be a 10-digit number (e.g., '0021700807')")
	}
	return true, nil
}

// ValidatePlayerOrTeam checks if the given value is 'Player' or 'Team'.
func ValidatePlayerOrTeam(playerOrTeam string) (bool, error) {
	if !validPlayerOrTeam.MatchString(playerOrTeam) {
		return false, errors.New("invalid PlayerOrTeam: must be 'Player' or 'Team'")
	}
	return true, nil
}

// ValidateSeasonYear checks if the given season year is valid.
func ValidateSeasonYear(seasonYear string) (bool, error) {
	if !validSeasonYear.MatchString(seasonYear) {
		return false, errors.New("invalid SeasonYear: must be a four-digit year (e.g., '2019')")
	}
	return true, nil
}

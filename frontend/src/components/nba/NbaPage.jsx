import React, { useState } from "react";
import { Box, Typography, Button } from "@mui/material";
import Grid from "@mui/material/Grid2"; // Import Grid2
import PlayersTable from "./PlayersTable";
import PlayerStatsPanel from "./PlayerStatsPanel";

const matchups = [
    {
        HomeTeam: {
            id: 1610612761,
            abbreviation: "TOR",
            full_name: "Toronto Raptors",
            year_founded: 1995,
            championship_years: [2019],
            Roster: [
                { DISPLAY_FIRST_LAST: "Scottie Barnes", FROM_YEAR: "2021", TO_YEAR: "2024", PERSON_ID: 1630567, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "RJ Barrett", FROM_YEAR: "2019", TO_YEAR: "2024", PERSON_ID: 1629628, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Chris Boucher", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1628449, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
            ],
        },
        AwayTeam: {
            id: 1610612738,
            abbreviation: "BOS",
            full_name: "Boston Celtics",
            year_founded: 1946,
            championship_years: [1957, 2008],
            Roster: [
                { DISPLAY_FIRST_LAST: "Jaylen Brown", FROM_YEAR: "2016", TO_YEAR: "2024", PERSON_ID: 1627759, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jaylen Brown", FROM_YEAR: "2016", TO_YEAR: "2024", PERSON_ID: 1627759, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jaylen Brown", FROM_YEAR: "2016", TO_YEAR: "2024", PERSON_ID: 1627759, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jaylen Brown", FROM_YEAR: "2016", TO_YEAR: "2024", PERSON_ID: 1627759, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jaylen Brown", FROM_YEAR: "2016", TO_YEAR: "2024", PERSON_ID: 1627759, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jayson Tatum", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1628369, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
            ],
        },

    },
    {
        HomeTeam: {
            id: 1610612761,
            abbreviation: "TOR",
            full_name: "Toronto Raptors",
            year_founded: 1995,
            championship_years: [2019],
            Roster: [
                { DISPLAY_FIRST_LAST: "Scottie Barnes", FROM_YEAR: "2021", TO_YEAR: "2024", PERSON_ID: 1630567, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "RJ Barrett", FROM_YEAR: "2019", TO_YEAR: "2024", PERSON_ID: 1629628, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Chris Boucher", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1628449, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
                { DISPLAY_FIRST_LAST: "Dr", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1642279, TEAM_ABBREVIATION: "TOR", TEAM_NAME: "Raptors" },
            ],
        },
        AwayTeam: {
            id: 1610612738,
            abbreviation: "BOS",
            full_name: "Boston Celtics",
            year_founded: 1946,
            championship_years: [1957, 2008],
            Roster: [
                { DISPLAY_FIRST_LAST: "Jaylen Brown", FROM_YEAR: "2016", TO_YEAR: "2024", PERSON_ID: 1627759, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jayson Tatum", FROM_YEAR: "2017", TO_YEAR: "2024", PERSON_ID: 1628369, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
                { DISPLAY_FIRST_LAST: "Jrue Holiday", FROM_YEAR: "2009", TO_YEAR: "2024", PERSON_ID: 201950, TEAM_ABBREVIATION: "BOS", TEAM_NAME: "Celtics" },
            ],
        },

    }
];

const MainPage = () => {
    const [selectedPlayer, setSelectedPlayer] = useState(null);

    return (
        <Box sx={{ width: '100%', maxWidth: { sm: '100%', md: '1700px' } }}>
            <Typography>TEST</Typography>
            {selectedPlayer ? (
                <PlayerStatsPanel player={selectedPlayer} onBack={() => setSelectedPlayer(null)} />
            ) : (
                <>
                    <Typography component="h2" variant="h6" sx={{ mb: 2 }}>
                        Rosters
                    </Typography>
                    <PlayersTable onPlayerClick={setSelectedPlayer} />
                </>
            )}
        </Box>
    );

}


// const MainPage2 = () => {
//     const [selectedMatchup, setSelectedMatchup] = useState(null);
//
//     return (
//         <Grid
//             container
//             direction="column"
//             alignItems="center"
//             justifyContent="center"
//             spacing={4}
//             sx={{ p: 3, width: "100%" }}
//         >
//             {/* Title */}
//             <Grid2 xs={12}>
//                 <Typography variant="h4" sx={{ textAlign: "center" }}>
//                     NBA Matchups
//                 </Typography>
//             </Grid2>
//
//             {/* Matchup Selector */}
//             <Grid2 container spacing={3} justifyContent="center">
//                 {matchups.map((matchup, index) => (
//                     <Grid2 key={index} xs={12} sm={6} md={4} lg={3}>
//                         <Button
//                             variant="contained"
//                             fullWidth
//                             onClick={() => setSelectedMatchup(matchup)}
//                         >
//                             {matchup.HomeTeam.abbreviation} vs {matchup.AwayTeam.abbreviation}
//                         </Button>
//                     </Grid2>
//                 ))}
//             </Grid2>
//
//             {/* Team Roster Side-by-Side */}
//             {selectedMatchup && (
//                 <Grid2
//                     container
//                     direction={"row"}
//                     justifyContent="center"
//                     alignItems="flex-start"
//                     sx={{ width: "100%", maxWidth: "1200px" }} // Ensuring proper width
//                 >
//                     <Grid2 xs={12} sm={6} md={5} sx={{ display: "flex", justifyContent: "center" }}>
//                         <TeamComponent team={selectedMatchup.HomeTeam} />
//                     </Grid2>
//                     <Grid2 xs={12} sm={6} md={5} sx={{ display: "flex", justifyContent: "center" }}>
//                         <TeamComponent team={selectedMatchup.AwayTeam} />
//                     </Grid2>
//                 </Grid2>
//             )}
//         </Grid>
//     );
// };

export default MainPage;




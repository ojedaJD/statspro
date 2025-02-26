import React from "react";
import {Box} from "@mui/material";
import PlayerCard from "./PlayerCard";
import Grid2 from "@mui/material/Grid2";

const RosterComponent = ({ roster }) => {
    if (!roster || roster.length === 0) return null;

    return (

        <Box
            sx={{
                display: "flex",
                flexWrap: "wrap", // Wrap players onto new lines
                justifyContent: "center",
                gap: 2, // Space between player cards
                p: 2,
            }}
        >


            {roster.map((player) => (
                <PlayerCard key={player.PERSON_ID} player={player} />
            ))}



        </Box>

    );
};

export default RosterComponent;

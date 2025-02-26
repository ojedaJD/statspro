import React from "react";
import { Card, CardContent, Typography, CardMedia } from "@mui/material";

const PlayerCard = ({ player }) => {
    return (
        <Card
            sx={{
                boxShadow: 2,
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                textAlign: "center",
                p: 1, // Reduce padding
                borderRadius: 8, // Slightly less rounded
                width: 120, // Smaller card width
            }}
        >
            <CardMedia
                component="img"
                height="60" // Decrease image height
                image={`https://cdn.nba.com/headshots/nba/latest/1040x760/${player.PERSON_ID}.png`}
                alt={player.DISPLAY_FIRST_LAST}
                onError={(e) => (e.target.src = "https://via.placeholder.com/100")}
                sx={{
                    borderRadius: "50%", // Circular image
                    width: 60, // Smaller image
                    height: 60,
                    objectFit: "cover",
                    border: "1px solid #ddd", // Subtle border around image
                }}
            />
            <CardContent sx={{ p: 1 }}> {/* Reduce padding */}
                <Typography variant="body2" fontWeight="bold">
                    {player.DISPLAY_FIRST_LAST}
                </Typography>
                <Typography variant="caption" color="text.secondary">
                    {player.TEAM_NAME} ({player.TEAM_ABBREVIATION})
                </Typography>
                <Typography variant="caption">
                    Position : TBD via API call
                </Typography>
            </CardContent>
        </Card>
    );
};

export default PlayerCard;

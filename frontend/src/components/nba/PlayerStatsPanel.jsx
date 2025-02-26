import React, { useEffect, useState } from "react";
import axios from "axios";
import { Box, Button, Typography, ButtonGroup, Slider, Alert, Grid } from "@mui/material";
import { BarChart, Bar, XAxis, YAxis, Tooltip, CartesianGrid, ResponsiveContainer, Legend } from "recharts";
import { motion } from "framer-motion"; // Framer Motion for animations

const PlayerStatsPanel = ({ player, onBack }) => {
    const [stats, setStats] = useState([]);
    const [filteredStats, setFilteredStats] = useState([]);
    const [selectedStats, setSelectedStats] = useState(["PTS", "AST"]);
    const [availableStats, setAvailableStats] = useState([]);
    const [errorMessage, setErrorMessage] = useState("");
    const [dateRange, setDateRange] = useState([0, 10]); // Default range
    const currentSeason = "2023-24";
    const seasonType = "Regular Season";

    useEffect(() => {
        if (!player) return;

        const fetchPlayerStats = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/nba/player/gamelog`, {
                    params: {
                        playerID: player.id,
                        season: currentSeason,
                        seasonType: seasonType
                    }
                });

                let gameLogs = response.data.PlayerGameLog.map((game, index) => {
                    let processedGame = { index, date: game.GAME_DATE };

                    Object.keys(game).forEach(key => {
                        if (typeof game[key] === "number") {
                            processedGame[key] = game[key];
                        }
                    });

                    return processedGame;
                }).reverse(); // Keep chronological order (oldest first)

                const allStats = Object.keys(gameLogs[0] || {}).filter(key => key !== "date" && key !== "index");

                setStats(gameLogs);
                setAvailableStats(allStats);
                setDateRange([0, gameLogs.length - 1]); // Full range by default
                setFilteredStats(gameLogs);
            } catch (error) {
                console.error("Error fetching player stats:", error);
            }
        };

        fetchPlayerStats();
    }, [player]);

    // Limit stat selection to 3 stats
    const toggleStatSelection = (stat) => {
        if (selectedStats.includes(stat)) {
            setSelectedStats(selectedStats.filter(s => s !== stat));
            setErrorMessage("");
        } else if (selectedStats.length < 3) {
            setSelectedStats([...selectedStats, stat]);
            setErrorMessage("");
        } else {
            setErrorMessage("You can only select up to 3 statistics.");
        }
    };

    // Handle slider change (Restored original logic)
    const handleSliderChange = (event, newValue) => {
        setDateRange(newValue);
        setFilteredStats(stats.slice(newValue[0], newValue[1] + 1)); // Directly filter data
    };

    // Color palette (Green, Red, Blue, Yellow)
    const colorPalette = ["#28a745", "#dc3545"];

    return (
        <Box sx={{ p: 3, textAlign: "center" }}>
            <Typography variant="h4">{player.DISPLAY_FIRST_LAST} - Stats</Typography>

            {errorMessage && (
                <Alert severity="warning" sx={{ my: 2 }}>{errorMessage}</Alert>
            )}

            {/* Dynamic Button Group for Multi-Stat Selection */}
            <ButtonGroup variant="outlined" sx={{ my: 2, flexWrap: "wrap", justifyContent: "center" }}>
                {availableStats.map((stat) => (
                    <motion.div
                        key={stat}
                        whileTap={{ scale: 0.9 }}
                        whileHover={{ scale: 1.1 }}
                    >
                        <Button
                            onClick={() => toggleStatSelection(stat)}
                            sx={{
                                bgcolor: selectedStats.includes(stat) ? "#007bff" : "black",
                                color: selectedStats.includes(stat) ? "white" : "gray",
                                border: "1px solid gray",
                                transition: "background-color 0.3s ease-in-out",
                                "&:hover": { bgcolor: "#005F8C", color: "white" },
                                m: 0.5,
                            }}
                        >
                            {stat}
                        </Button>
                    </motion.div>
                ))}
            </ButtonGroup>

            {/* Animated Chart Container */}
            <motion.div
                key={selectedStats.join("-")}
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.5 }}
            >
                <Typography variant="h6" sx={{ mt: 3 }}>Comparison: {selectedStats.join(", ")}</Typography>
                <ResponsiveContainer width="100%" height={400}>
                    <BarChart data={filteredStats} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
                        <CartesianGrid strokeDasharray="3 3" />
                        <XAxis dataKey="date" />
                        <YAxis />
                        <Tooltip wrapperStyle={{ backgroundColor: "rgba(255, 255, 255, 0.9)", borderRadius: "5px" }} />
                        <Legend />
                        {selectedStats.map((stat, index) => (
                            <Bar
                                key={stat}
                                dataKey={stat}
                                fill={colorPalette[index]} // Assign predefined colors
                                animationDuration={500}
                            />
                        ))}
                    </BarChart>
                </ResponsiveContainer>
            </motion.div>

            {/* Date Range Slider - Now Below the Chart with Labels */}
            <Box sx={{ width: "80%", margin: "auto", mt: 5 }}>
                <Typography variant="h6">Select Date Range</Typography>
                <Grid container justifyContent="space-between">
                    <Typography variant="body2">Oldest</Typography>
                    <Typography variant="body2">Most Recent</Typography>
                </Grid>
                <Slider
                    value={dateRange}
                    onChange={handleSliderChange}
                    min={0}
                    max={stats.length - 1}
                    valueLabelDisplay="auto"
                    sx={{ color: "#007bff" }} // Blue slider
                />
                <Typography variant="body1">
                    Showing games from <strong>{stats[dateRange[0]]?.date}</strong> to <strong>{stats[dateRange[1]]?.date}</strong>
                </Typography>
            </Box>

            <Button variant="contained" onClick={onBack} sx={{ mt: 2 }}>Back to Player List</Button>
        </Box>
    );
};

export default PlayerStatsPanel;

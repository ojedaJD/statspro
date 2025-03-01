import React from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Paper from '@mui/material/Paper';
import ConstructionIcon from '@mui/icons-material/Construction';

const UnderConstruction = ({ sport }) => {
    return (
        <Paper
            elevation={2}
            sx={{
                p: 5,
                textAlign: 'center',
                maxWidth: 800,
                mx: 'auto',
                mt: 8
            }}
        >
            <ConstructionIcon sx={{ fontSize: 60, color: 'warning.main', mb: 2 }} />
            <Typography variant="h4" gutterBottom>
                {sport} Props Coming Soon
            </Typography>
            <Typography variant="body1" color="text.secondary">
                We're currently working on bringing you the latest {sport} props and odds.
                Please check back later, or explore our NBA props which are available now.
            </Typography>
        </Paper>
    );
};

export default UnderConstruction;
import React from 'react';
import { Outlet, Link, useLocation } from 'react-router-dom';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid2'; // Using Grid2 as specified
import ShowChartIcon from '@mui/icons-material/ShowChart';

const navItems = [
    { label: 'NBA', path: '/nba' },
    { label: 'NFL', path: '/nfl' },
    { label: 'NHL', path: '/nhl' },
    { label: 'MLB', path: '/mlb' },
    { label: 'NCAAF', path: '/ncaaf' },
    { label: 'WNBA', path: '/wnba' },
];

const Layout = () => {
    const location = useLocation();

    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
            <AppBar position="static" color="default" elevation={1}>
                <Toolbar>
                    <Box display="flex" alignItems="center" sx={{ mr: 4 }}>
                        <ShowChartIcon sx={{ color: 'primary.main', mr: 1 }} />
                        <Typography variant="h6" component="div" sx={{ fontWeight: 'bold' }}>
                            Stats Pro
                        </Typography>
                    </Box>

                    <Box sx={{ display: { xs: 'none', md: 'flex' } }}>
                        {navItems.map((item) => (
                            <Button
                                key={item.label}
                                component={Link}
                                to={item.path}
                                sx={{
                                    mx: 1,
                                    color: 'text.primary',
                                    fontWeight: location.pathname.startsWith(item.path) ? 'bold' : 'normal',
                                    borderBottom: location.pathname.startsWith(item.path) ? '2px solid' : 'none',
                                    borderRadius: 0,
                                    '&:hover': {
                                        backgroundColor: 'transparent',
                                        borderBottom: '2px solid',
                                    },
                                }}
                            >
                                {item.label}
                            </Button>
                        ))}
                    </Box>
                </Toolbar>
            </AppBar>

            <Box component="main" sx={{ flexGrow: 1, py: 3 }}>
                <Container maxWidth={false}>
                    <Outlet />
                </Container>
            </Box>
        </Box>
    );
};

export default Layout;
import * as React from 'react';
import axios from 'axios';
import Box from '@mui/material/Box';
import { DataGrid } from '@mui/x-data-grid';

const columns = [
    { field: 'TEAM_ABBREVIATION', headerName: 'Team', width: 90 },
    { field: 'DISPLAY_FIRST_LAST', headerName: 'Player', width: 150 },
];

const PlayersTable = ({ onPlayerClick }) => {
    const [rows, setRows] = React.useState([]);

    React.useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await axios.get('http://localhost:8080/nba/matchups/players');
                const formattedRows = response.data.map((player, index) => ({
                    id: player.PERSON_ID,
                    TEAM_ABBREVIATION: player.TEAM_ABBREVIATION,
                    DISPLAY_FIRST_LAST: player.DISPLAY_FIRST_LAST,
                }));
                setRows(formattedRows);
            } catch (error) {
                console.error('Error fetching players:', error);
            }
        };

        fetchPlayers();
    }, []);

    return (
        <Box sx={{ height: '100vh', width: '100%' }}>
            <DataGrid
                rows={rows}
                columns={columns}
                pageSizeOptions={[5, 10, 20, 50]}
                onRowClick={(params) => onPlayerClick(params.row)}
                autoHeight={false}
                sx={{ '& .MuiDataGrid-root': { overflow: 'auto' } }}
            />
        </Box>
    );
};

export default PlayersTable;

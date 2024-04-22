import { useState } from "react";
import HomeIcon from '@mui/icons-material/Home';
import PersonIcon from '@mui/icons-material/Person';
import LiveTvIcon from '@mui/icons-material/LiveTv';
import Diversity3Icon from '@mui/icons-material/Diversity3';
import { BottomNavigation, BottomNavigationAction, ThemeProvider, createTheme } from "@mui/material";
import { Link } from "react-router-dom";

const theme = createTheme({
    components: {
        MuiBottomNavigationAction: {
            styleOverrides: {
                root: {
                "&.Mui-selected": {
                    color: 'white', // Cambia el color aquí
                },
                },
            },
            },
        },
});

export const Navbar = () => {
    const [value, setValue] = useState(0);

    return (
        <ThemeProvider theme={theme}>
            <BottomNavigation
                showLabels
                value={value}
                onChange={(event, newValue) => {
                    setValue(newValue);
                }}
                sx={{ position: 'fixed', bottom: 0, left: 0, right: 0, width: '100%', backgroundColor: '#5616eb' }}
            >
                    <BottomNavigationAction 
                        component={Link}
                        to='/'
                        sx={{ color: 'rgba(255,255,255,0.5690651260504201)' }} 
                        label="Inicio" 
                        icon={<HomeIcon />} 
                    />
                    <BottomNavigationAction 
                        component={Link}
                        to='/friends'
                        sx={{ color: 'rgba(255,255,255,0.5690651260504201)' }} 
                        label="Amigos" 
                        icon={<Diversity3Icon />} 
                    />
                    <BottomNavigationAction 
                        component={Link}
                        to='/live'
                        sx={{ color: 'rgba(255,255,255,0.5690651260504201)' }} 
                        label="En vivo" 
                        icon={<LiveTvIcon />} 
                    />

                    <BottomNavigationAction 
                        component={Link}
                        to='/profile'
                        sx={{ color: 'rgba(255,255,255,0.5690651260504201)' }} 
                        label="Perfil" 
                        icon={<PersonIcon />} 
                    />

            </BottomNavigation>
        </ThemeProvider>
    )
}

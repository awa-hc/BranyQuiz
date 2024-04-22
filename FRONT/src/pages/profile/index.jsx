
import { Avatar } from "@mui/material"
import { Navbar } from "../../components/Navbar"


export const Profile = () => {

    /* const user = JSON.parse(localStorage.getItem('user')); */

    return (
        <div className="bg center">
            <div className="container-profile">
            <h1 className="title-profile">¡Hola user!</h1>
                <Avatar 
                    sx={{ 
                        width: {md: 150, sm: 125, xs: 120}, 
                        height: {md: 150, sm: 125, xs: 120},
                        position: 'absolute',
                        top: "-90px",
                    }} 
                />
                <div className="container-quiz">
                    <p className="subtitle-profile">Quiz diario</p>
                    <p className="p-profile">Total quiz: 000</p>
                    <p className="p-profile">Puntos: 000</p>
                </div>
                <div className="container-quiz">
                    <p className="subtitle-profile">Quiz en vivo</p>
                    <p className="p-profile">Total quiz: 000</p>
                    <p className="p-profile">Puntos: 000</p>
                </div>
            </div>

        </div>
    )
}

import { useState } from "react";
import axios from "axios";

export const Friends = () => {

    const [query, setQuery] = useState('');

    const handleAddFriend = async () => {
        try {
            await axios.post('https://backend/', { query });
            alert('Solicitud de amistad enviada!');
        } catch (error) {
            console.error('Hubo un error al enviar la solicitud de amistad:', error);
        }
    };

    return (
        <div className="bg">
            <div className="container-s">
                <h1 className="title-friends">¡Agrega amigos!</h1>
                <div className="flex">
                    <input
                        type="text"
                        value={query}
                        onChange={(e) => setQuery(e.target.value)}
                        placeholder="Email o nombre de usuario"
                        className="input-friend"
                    />
                    <button 
                        onClick={handleAddFriend}
                        className="btn-friend">
                            Agregar
                    </button>
                </div>
            </div>
        </div>
    )
}

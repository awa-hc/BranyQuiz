import { useState } from "react"
import { GoogleButton } from "../../components/GoogleButton"
import showPassw from '../../assets/showPassword.svg'
import hidePassword from '../../assets/hidePassword.svg'

export const Login = () => {

    const [showPassword, setShowPassword] = useState(false);

    const [register, setRegister] = useState(true)
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const handleShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const handleSubmit = async (e) => {
        e.preventDefault()
        if (register === false) {
            //inicio de sesion
        } else{
            // registro de usuario
            const response = await fetch('url_del_backend', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({email: email, password: password})
            });
    
            if (response.ok) {
                const data = await response.json();
                console.log(data);

                localStorage.setItem('token', data.token);
            } else {
                console.error('Error:', response.status);
            }
        };
        }
    

    return (
        <>
            
            <div id="form-ui">
                <form id="form" type="submit" onSubmit={handleSubmit}>
                <div id="form-body">
                    <div id="welcome-lines">
                    <div id="welcome-line-1">Brainy Quiz</div>
                    <div id="welcome-line-2">
                        {register ? 'Registrarse' : 'Iniciar sesión'}
                    </div>
                    </div>
                    <div id="input-area">
                    <div className="form-inp">
                        <input 
                            type="email"
                            placeholder="Correo electrónico" 
                            value={email} 
                            onChange={(e) => setEmail(e.target.value)} 
                            autoComplete="off"
                        />
                        
                    </div>
                    <div className="form-inp">
                        <input 
                            type={showPassword ? "text" : "password"} 
                            placeholder="Contraseña" 
                            value={password} 
                            onChange={(e) => setPassword(e.target.value)}
                            autoComplete="off"
                        />
                        <button onClick={handleShowPassword} className="btn-pssw">
                            <img src={showPassword ? hidePassword : showPassw} className="show-password"/>
                        </button>
                    </div>
                    </div>
                    <div id="submit-button-cvr">
                    <button id="submit-button" type="submit">
                        {register ? 'Registrarme' : 'Ingresar'}
                    </button>

                    <GoogleButton />

                    </div>
                    <div id="have-account">
                    <button type="button" onClick={() => setRegister(!register)}>
                        {register ? '¿Ya tenés una cuenta? Inicia sesión' : '¿Todavia no tenés cuenta? Registrate' }
                    </button>
                    </div>
                </div>
                </form>
                </div>
            

        </>
    )

}

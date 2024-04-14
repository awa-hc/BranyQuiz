import { GoogleButton } from "../../components/GoogleButton"

export const Login = () => {
    return (
        <>
            
            <div id="form-ui">
                <form id="form">
                <div id="form-body">
                    <div id="welcome-lines">
                    <div id="welcome-line-1">Brainy Quiz</div>
                    <div id="welcome-line-2">Iniciar Sesión</div>
                    </div>
                    <div id="input-area">
                    <div className="form-inp">
                        <input placeholder="Correo electrónico" type="text"/>
                    </div>
                    <div className="form-inp">
                        <input placeholder="Contraseña" type="password"/>
                    </div>
                    </div>
                    <div id="submit-button-cvr">
                    <button id="submit-button" type="submit">Ingresar</button>
                    <GoogleButton />
                    </div>
                    <div id="forgot-pass">
                    <a href="#">¿Aún no tienes una cuenta? Registrate</a>
                    </div>
                </div>
                </form>
                </div>
            

        </>
    )
}

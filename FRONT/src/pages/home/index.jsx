import { useState } from "react"
import Pagination from '@mui/material/Pagination';
import Stack from '@mui/material/Stack';

export const Home = () => {

    const [form, setForm] = useState(1)

    const handleQuestion = (e) => {
        e.preventDefault()
        setForm(form + 1)
        console.log(form)
    }

    return (
        <div className="bg center">
                <form type="submit" className="form-quiz">
                <div className="body-quiz">
                    <div id="welcome-lines">
                    <div className="title-formq">¡Crea tu Quiz!</div>
                    </div>
                    <div id="input-area">
                    <p className="text-quiz">
                        {
                            form === 1 ? "Primera pregunta" : 
                            form === 2 ? "Segunda pregunta" : 
                            form === 3 ? "Tercera pregunta" : 
                            form === 4 ? "Cuarta pregunta" : 
                            form === 5 ? "Quinta pregunta" : "Primera pregunta"
                        }
                    </p>
                    <div className="form-inp">
                        <input
                            type="text"
                        />
                    </div>
                    <p className="text-quiz text-opt">Opciones</p>
                    <div className="flex options">
                        <div className="form-inp">
                            <input
                            type="text"
                            autoComplete="off"
                            placeholder="Opción 1"
                            />
                        </div>
                        <div className="form-inp">
                            <input
                            type="text"
                            autoComplete="off"
                            placeholder="Opción 2"
                            />
                        </div>
                    </div>
                    <div className="flex options">
                        <div className="form-inp">
                            <input
                            type="text"
                            autoComplete="off"
                            placeholder="Opción 3"
                            />
                        </div>
                        <div className="form-inp">
                            <input
                            type="text"
                            autoComplete="off"
                            placeholder="Opción 4"
                            />
                        </div>
                    </div>
                    </div>
                    <div className="submit-container">
                        <Stack spacing={2}>
                            <Pagination 
                                count={5} 
                                sx={{ 
                                    '.MuiPaginationItem-root': { color: '#ffffff' }, 
                                    '& .Mui-selected': {backgroundColor: '#410fb8', } 
                                }} 
                                onChange={handleQuestion}
                            />
                        </Stack>
                        {
                            form == 5 && <button id="submit-button" className="btn-quiz" type="submit">
                                            Crear quiz
                                        </button>

                        }
                    </div>
                </div>
                </form>
        </div>
    )
}

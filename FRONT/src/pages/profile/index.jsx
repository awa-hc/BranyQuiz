
import { Avatar, Box, Button, Dialog, DialogActions, DialogContent, styled } from "@mui/material"
import editWhite from "../../assets/editw.svg";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import { useState } from "react";
import { toast } from "react-toastify";


export const Profile = () => {
    
    const [openImg, setOpenImg] = useState(false);
    const [selectedFile, setSelectedFile] = useState(null);
    const [file, setFile] = useState(null);
    /* const user = JSON.parse(localStorage.getItem('user')); */

    const VisuallyHiddenInput = styled("input")({
        clip: "rect(0 0 0 0)",
        clipPath: "inset(50%)",
        height: 1,
        overflow: "hidden",
        position: "absolute",
        bottom: 0,
        left: 0,
        whiteSpace: "nowrap",
        width: 1,
    });

    const handleFileChange = (e) => {
        const [file] = e.target.files;
        const uploadedFile = e.target.files[0];
        setFile(uploadedFile);
    
        const SIZE_50MB = 50 * 1024 * 1024;
        const isValidSize = file.size < SIZE_50MB;
        const isNameOfOneImageRegex = /.(jpe?g|gif|png)$/i;
        const isValidType = isNameOfOneImageRegex.test(file.name);
    
        if (!isValidType)
            return toast.error('Solo puedes subir imagenes', {
                position: "bottom-left",
                autoClose: 2000,
                hideProgressBar: true,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
                });
        if (!isValidSize)
            return toast.error('El archivo es demasiado grande', {
                position: "bottom-left",
                autoClose: 2000,
                hideProgressBar: true,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
                });
        
            const reader = new FileReader();
            reader.onloadend = () => {
            setSelectedFile(reader.result);
            };
            reader.readAsDataURL(file);
        };
        
        const handleUpdateImg = async () => {
            if (!selectedFile) {
            return toast.warn('Debes seleciionar una imagen', {
                position: "bottom-left",
                autoClose: 2000,
                hideProgressBar: true,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                theme: "colored",
                });
            }
            try {
                localStorage.setItem("img", selectedFile);

            }catch (error) {
                alert(error);
            }
            setOpenImg(false);
        }

    return (
        <div className="bg center">
            <div className="container-profile">
            <h1 className="title-profile">¡Hola user!</h1>

            <Box sx={{ display: "flex",  }}>
                <Avatar 
                    sx={{ 
                        width: {md: 150, sm: 125, xs: 120}, 
                        height: {md: 150, sm: 125, xs: 120},
                        position: 'absolute',
                        top: "-90px",
                        left: "36%",
                    }} 
                    src={localStorage.getItem("img")}
                />
                <div className="editimg-container">
                <img
                    src={editWhite}
                    alt="edit"
                    className="editimg-icon"
                    onClick={() => setOpenImg(true)}
                />
                </div>
            </Box>
            <Dialog
                open={openImg}
                onClose={() => setOpenImg(false)}
                aria-describedby="dialog-description"
            >
                <DialogContent>
                    <Avatar src={selectedFile} sx={{ width: "10rem", height: "10rem" }} />
                    <Box sx={{ display: "flex" }}>
                        <Button
                            component="label"
                            role={undefined}
                            variant="contained"
                            color="secondary"
                            tabIndex={-1}
                            startIcon={<CloudUploadIcon />}
                        >
                        Subir imagen
                        <VisuallyHiddenInput
                            type="file"
                            accept=".png, .jpg, .jpeg, .gif"
                            onChange={handleFileChange}
                        />
                        </Button>
                    </Box>
                </DialogContent>
                <DialogActions>
                <Button onClick={() => setOpenImg(false)} color="error" size="small">
                    Cancelar
                </Button>
                <Button
                    size="small"
                    sx={{ marginRight: "0.5rem" }}
                    onClick={handleUpdateImg}
                    color="secondary"
                >
                    Actualizar
                </Button>
                </DialogActions>
            </Dialog>
                
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

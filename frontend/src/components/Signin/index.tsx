import React from 'react'
import { Alert, Box, Button, Snackbar } from '@mui/material'

import frame from './../image/frame.svg'

import './signin.css'


export default function Signin() {

    const [success, setSuccess] = React.useState<boolean>(false)
    const [error, setError] = React.useState<boolean>(false)

    const handleClose:any = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === "clickaway") {
          return;
        }
        setSuccess(false);
        setError(false);
    };

    

    return (
        <Box sx={{ display: "flex" }}>
            <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="success">
                    เข้าสู่ระบบสำเร็จ
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    อีเมลหรือรหัสผ่านไม่ถูกต้อง
                </Alert>
            </Snackbar>

            <div className='from-box'>
                <img
                    style={{ maxHeight: "100vh" }}
                    className="img"
                    alt="Banner"
                    src={frame}
                />
            </div>
        </Box>
    )
}

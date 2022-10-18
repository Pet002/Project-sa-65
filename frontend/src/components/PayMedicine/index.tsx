import { Alert, Box, Button, Container, FormControl, Grid, Paper, Select, Snackbar, TextField, Typography } from '@mui/material';
import React, { useEffect, useState } from 'react'
// import { EmployeeInterface } from '../../models/IEmployee';
import { MedicineLabelInterface } from '../../models/IMedicineLabel';
import { PayMedicineInterface } from '../../models/IPayMedicine';
import { PerscriptionInterface } from '../../models/IPerscription';
import { Link as RouterLink } from "react-router-dom";


import './styles.css'

export default function PayMedicine() {
    //main
    const [payMedicine, setPayMedicine] = React.useState<Partial<PayMedicineInterface>>({});
    //relation
    // const [employee, setEmployee] = React.useState<EmployeeInterface[]>([]);
    const [medicineLabel, setMedicineLabel] = React.useState<MedicineLabelInterface[]>([]);
    const [perscription, setPerscription] = React.useState<PerscriptionInterface[]>([]);
    const [selectPerscription, setSelectPerscription] = React.useState<PerscriptionInterface>();
    const [selectmedicineLabel, setSelectmedicineLabel] = React.useState<MedicineLabelInterface>();



    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    const handleClose: any = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    const handleChange: any = (event: React.ChangeEvent<{ name?: string; value: any }>) => {
        const name = event.target.name as keyof typeof payMedicine;

        if (name === "PerscriptionID") {
            setSelectPerscription(perscription.at(event.target.value - 1))
            if (event.target.value === "") {
                setSelectPerscription(perscription.at(perscription.length + 1))
            }
        } else if (name === "MedicineLabelID") {
            setSelectmedicineLabel(medicineLabel.at(event.target.value - 1))
            if (event.target.value === "") {
                setSelectmedicineLabel(medicineLabel.at(medicineLabel.length + 1))
            }
        }
        setPayMedicine({
            ...payMedicine,
            [name]: event.target.value,
        });


    };


    const apiUrl = "http://localhost:8080";


    useEffect(() => {
        const requestOptions = {
            method: "GET",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
        };
        // const getEmployee = async () => {
        //     fetch(`${apiUrl}/phamacist/employees`, requestOptions)
        //         .then((res) => res.json())
        //         .then((res) => {
        //             if (res.data) {
        //                 setEmployee(res.data)
        //             } else {
        //                 console.log("else1")
        //             }
        //         })
        // }

        const getMedicineLable = async () => {
            fetch(`${apiUrl}/phamacist/medicinelabels`, requestOptions)
                .then((res) => res.json())
                .then((res) => {
                    if (res.data) {
                        setMedicineLabel(res.data)
                    } else {
                        console.log("else2")
                    }
                })
        }

        const getPerscription = async () => {
            fetch(`${apiUrl}/phamacist/perscriptions`, requestOptions)
                .then((res) => res.json())
                .then((res) => {
                    if (res.data) {
                        setPerscription(res.data)
                    } else {
                        console.log("else3")
                    }
                })
        }


        // getEmployee();
        getMedicineLable();
        getPerscription();
    }, [])

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val
    }

    function submit() {
        let data = {
            Amount: convertType(payMedicine.Amount),
            Price: convertType(payMedicine.Price),
            EmployeeID: convertType(localStorage.getItem("uid") as string),
            PerscriptionID: convertType(payMedicine.PerscriptionID),
            MedicineLabelID: convertType(payMedicine.MedicineLabelID),
            PayDate: new Date(),
        }

        console.log(data)

        const requestOptionsPost = {
            method: "POST",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        };

        fetch(`${apiUrl}/phamacist/paymedicines`, requestOptionsPost)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    console.log("บันทึกได้")
                    setSuccess(true)
                    setErrorMessage("")
                } else {
                    console.log("บันทึกไม่ได้")
                    setError(true)
                    setErrorMessage(res.error)
                }
            });
    }

    return (
        <Container maxWidth="lg">
            <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="success">
                    บันทึกข้อมูลสำเร็จ
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                <Alert onClose={handleClose} severity="error">
                    บันทึกข้อมูลไม่สำเร็จ: {errorMessage}
                </Alert>
            </Snackbar>

            <Paper className='paper'>
                <Box display="flex">
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            gutterBottom
                            color="primary"
                        >
                            เริ่มทำการจ่ายยา
                        </Typography>
                    </Box>
                </Box>
                <hr />
                <Grid container spacing={3} className="root">
                    <Grid item xs={12} >
                        <FormControl fullWidth variant="outlined">
                            <p>เลือกใบสั่งยา</p>
                            <Select
                                native
                                value={payMedicine.PerscriptionID}
                                onChange={handleChange}
                                inputProps={{
                                    name: "PerscriptionID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกใบสั่งยา
                                </option>
                                {perscription.map((item: PerscriptionInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.Patient}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={4}>
                        <p>ชื่อผู้ป่วย</p>
                        <TextField label={selectPerscription?.Patient} disabled />
                    </Grid>
                    <Grid item xs={4}>
                        <p>ยาที่ต้องจัด</p>
                        <TextField label={selectPerscription?.Medicine} disabled />
                    </Grid>
                    <Grid item xs={4}>
                        <p>อาการ</p>
                        <TextField label={selectPerscription?.Symptom} disabled />
                    </Grid>
                    <Grid item xs={12} >
                        <FormControl fullWidth variant="outlined">
                            <p>เลือกฉลากยา</p>
                            <Select
                                native
                                value={payMedicine.MedicineLabelID}
                                onChange={handleChange}
                                inputProps={{
                                    name: "MedicineLabelID",
                                }}
                            >
                                <option aria-label="None" value="">
                                    กรุณาเลือกฉลากยา
                                </option>
                                {medicineLabel.map((item: MedicineLabelInterface) => (
                                    <option value={item.ID} key={item.ID}>
                                        {item.MedicineUse}
                                    </option>
                                ))}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <p>วิธีการใช้ยา</p>
                        <TextField label={selectmedicineLabel?.MedicineUse} disabled />
                    </Grid>
                    <Grid item xs={6}>
                        <p>คำเตือน</p>
                        <TextField label={selectmedicineLabel?.Warning} disabled />
                    </Grid>
                    <Grid item xs={6}>
                        <p>จำนวนยา</p>
                        <TextField name='Amount' type="number" value={payMedicine.Amount || ""} placeholder="1" InputProps={{ inputProps: { min: 1 } }} onChange={handleChange} />
                    </Grid>
                    <Grid item xs={6}>
                        <p>ราคายา</p>
                        <TextField name='Price' type="number" value={payMedicine.Price || ""} placeholder="1" InputProps={{ inputProps: { min: 1 } }} onChange={handleChange} />
                    </Grid>

                    <Grid item xs={12} >
                    <Button component={RouterLink} to="/medicinepay" variant='outlined'>
                        ย้อนกลับ
                    </Button>
                        <Button
                            style={{ float: "right" }}
                            variant="contained"
                            onClick={submit}
                        >
                            จัดยา
                        </Button>
                    </Grid>
                </Grid>              
            </Paper>

        </Container>
    )
}

import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Button, Grid } from '@mui/material'
import { Container } from '@mui/system'
import React, { useEffect } from 'react'
import { EmployeeInterface } from '../../models/IEmployee'
import { MedicineLabelInterface } from '../../models/IMedicineLabel'
import { PayMedicineInterface } from '../../models/IPayMedicine'
import { PrescriptionInterface } from '../../models/IPrescription'
import { Link as RouterLink } from "react-router-dom";

export default function PayMedicineHistory() {

    const [payMedicine, setPayMedicine] = React.useState<PayMedicineInterface[]>([])
    const [employee, setEmployee] = React.useState<EmployeeInterface[]>([])
    const [perscription, setPerscription] = React.useState<PrescriptionInterface[]>([])
    const [medicineLabel, setMedicineLabel] = React.useState<MedicineLabelInterface[]>([])


    // const [selectEmployee, setSelectEmployee] = React.useState<EmployeeInterface>()
    const apiUrl = "http://localhost:8080"
    useEffect(() => {
        const reqOpt = {
            method: "GET",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
        }
        const getMedicineLabel = async () => {
            fetch(`${apiUrl}/phamacist/medicinelabels`, reqOpt)
                .then((res) => res.json())
                .then((res) => {
                    if (res.data) {
                        //setPayMedicine(res.data)
                        setMedicineLabel(res.data)
                        // console.log(res.data)
                    } else {
                        console.log("else data")
                    }
                })
        }

        const getPerscription = async () => {
            fetch(`${apiUrl}/phamacist/perscriptions`, reqOpt)
                .then((res) => res.json())
                .then((res) => {
                    if (res.data) {
                        //setPayMedicine(res.data)
                        setPerscription(res.data)
                        // console.log(res.data)
                    } else {
                        console.log("else data")
                    }
                })
        }

        const getEmployee = async () => {
            fetch(`${apiUrl}/phamacist/employees`, reqOpt)
                .then((res) => res.json())
                .then((res) => {
                    if (res.data) {
                        //setPayMedicine(res.data)
                        setEmployee(res.data)
                        // console.log(res.data)
                    } else {
                        console.log("else data")
                    }
                })
        }

        const getPayMedicine = async () => {
            fetch(`${apiUrl}/phamacist/paymedicines`, reqOpt)
                .then((res) => res.json())
                .then((res) => {
                    if (res.data) {
                        // console.log(res.body)
                        setPayMedicine(res.data)
                        // console.log(payMedicine)
                    } else {
                        console.log(res)
                        console.log("else")

                    }
                })
        }


        //run Fetch API
        getPayMedicine()
        getEmployee()
        getPerscription()
        getMedicineLabel()
    }, [])

    //use in all interface data
    function filterData(emp: any, item: number) {
        return emp.ID === item
    }

    const convertFormatTime = (date: Date) => {
        // const padL = (nr : any, len = 2, chr = `0`) => `${nr}`.padStart(2, chr);
        const d1 = new Date(date)

        return `${d1.getDate()} / ${d1.getMonth()} / ${d1.getFullYear()} | ${d1.getHours()} : ${d1.getMinutes()}`

    }



    return (
        <Container maxWidth="lg">
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <TableContainer component={Paper} >
                        <Table>
                            <TableHead>
                                <TableRow>
                                    <TableCell>ชื่อ</TableCell>
                                    <TableCell align="right">ฉลากยา</TableCell>
                                    <TableCell align="right">เจ้าหน้าที่ผู้ออกฉลาก</TableCell>
                                    <TableCell align="right">จำนวนยา</TableCell>
                                    <TableCell align="right">ราคายา</TableCell>
                                    <TableCell align="right">วันที่ออกยา</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {
                                    payMedicine.map((item) => (
                                        <TableRow
                                            key={item.ID}
                                            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                        >
                                            <TableCell >{perscription.filter((pers) => filterData(pers, item.PerscriptionID)).at(0)?.Patient}</TableCell>
                                            <TableCell align="right">{medicineLabel.filter((ml) => filterData(ml, item.MedicineLabelID)).at(0)?.MedicineUse}</TableCell>
                                            <TableCell align="right">{employee.filter((emp) => filterData(emp, item.EmployeeID)).at(0)?.Name + " " + employee.filter((emp) => filterData(emp, item.EmployeeID)).at(0)?.Surname}</TableCell>
                                            <TableCell align="right">{item.Amount}</TableCell>
                                            <TableCell align="right">{item.Price}</TableCell>
                                            <TableCell align="right">{convertFormatTime(item.PayDate).toString()}</TableCell>
                                        </TableRow>
                                    ))
                                }
                            </TableBody>
                        </Table>
                    </TableContainer>
                </Grid>
                <Grid item xs={12}>
                    <Button component={RouterLink} to="/medicinepay/create" style={{ float: "right" }} variant="contained">
                        เริ่มทำการจัดยา
                    </Button>
                </Grid>

            </Grid>
        </Container>
    )
}

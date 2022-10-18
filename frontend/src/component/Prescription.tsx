import React, { useEffect } from "react";
import clsx from "clsx";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import {
  createStyles,
  makeStyles,
  useTheme,
  Theme,
} from "@material-ui/core/styles";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { Link as RouterLink } from "react-router-dom";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import { PatientInterface } from "../models/IPatient";
import { MedicineInterface } from "../models/IMedicine";
import { EmployeeInterface } from "../models/IEmployee";
import { PrescriptionInterface } from "../models/IPrescription";

import MenuItem from "@mui/material/MenuItem";

import Select, { SelectChangeEvent } from "@mui/material/Select";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 800,
    },
    tableSpace: {
      marginTop: 20,
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
      width: "130vh",
    },
    root: {
      flexGrow: 1,
    },
  })
);

function Prescription() {
  const classes = useStyles();
  const [date, setDate] = React.useState<Date | null>(null);
  const [prescription, setPrescription] = React.useState<
    Partial<PrescriptionInterface>
  >({});
  const [patient, setPatient] = React.useState<Partial<PatientInterface>>({});
  const [medicine, setMedicine] = React.useState<Partial<MedicineInterface>>(
    {}
  );
  const [employee, setEmployee] = React.useState<Partial<EmployeeInterface>>(
    {}
  );
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof Prescription;
    const { value } = event.target;
    setPrescription({ ...prescription, [id]: value });
  };

  const handleChange = (event: SelectChangeEvent) => {
    const a = event.target.value as keyof typeof Prescription;
    const name = event.target.name as keyof typeof Prescription;
    setPrescription({ ...Prescription, [name]: a });
    console.log(name);
  };

  function submit() {
    let data = {
      PrescriptionID: prescription.Prescription_ID ?? "",
      PHID: prescription.PH_ID ?? "",
      MedicineID: prescription.Medicine ?? "",
      //   Age: typeof patient.Age === "string" ? parseInt(patient.Age) : 0,
      Patient: prescription.Patient,
      Employee: prescription.Employee ?? "",
      Symptom: prescription.Symptom ?? "",
      Case_Time: prescription.Case_Time ?? "",
    };
    console.log(prescription);

    const apiUrl = "http://localhost:8080/patient";
    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify(data),
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
      });
  }

  <h1>Presription</h1>;

  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>

      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>

      <Paper className={classes.paper}>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}>
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom>
              Prescription
            </Typography>
          </Box>
        </Box>

        <Divider />

        <Grid container spacing={1} sx={{ padding: 1 }}>
          <Grid item xs={12}>
            <FormControl sx={{ width: 210 }} variant="outlined">
              <p>PID</p>
              <TextField
                id="PID"
                variant="outlined"
                type="string"
                size="medium"
                value={prescription.Prescription_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid    sx={{ padding: 2 }}>
            <p>Name</p>
            <FormControl sx={{ width: 210 }} variant="outlined">
              <TextField
                id="Name"
                type="string"
                size="medium"
               
                variant="filled"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid sx={{ padding: 2 }}>
            <p>Surname</p>
            <FormControl sx={{ width: 210 }} variant="outlined">
              <TextField
                id="Name"
                
                type="string"
                size="medium"
                
                variant="filled"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid sx={{ padding: 2 }}>
            <p>Age</p>
            <FormControl sx={{ width: 210 }} variant="outlined">
              <TextField
                id="Name"
                type="number"
                size="medium"
             
                variant="filled"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid sx={{ padding: 2 }}>
            <p>Gender</p>
            <FormControl sx={{ width: 210 }} variant="outlined">
              <TextField
                id="Name"
                type="string"
                size="medium"
                
                variant="filled"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid sx={{ padding: 2 }}>
            <p>Allergy</p>
            <FormControl sx={{ width: 210 }} variant="outlined">
              <TextField
                id="Name"
                type="string"
                size="medium"
                
                variant="filled"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid sx={{ padding: 2 }}>
            <p>Symptom</p>
            <FormControl sx={{ width: 200 }} variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid sx={{ padding: 2 }}>
            <p>Medicine</p>
            <FormControl sx={{ width: 200 }} variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid sx={{ padding: 2 }}>
            <p>Type</p>
            <FormControl sx={{ width: 180 }} variant="outlined">
              <TextField
                id="Name"
                type="string"
                size="medium"
                
                variant="filled"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid sx={{ padding: 2 }}>
            <p>Pharmacist</p>
            <FormControl sx={{ width: 180 }} variant="outlined">
              <TextField
                id="Name"
                type="string"
                size="medium"
                disabled
                variant="filled"
                value={prescription.PH_ID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
           
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary">
              Submit
            </Button>
          </Grid>
          
        </Grid>
      </Paper>
    </Container>
  );
}
export default Prescription;

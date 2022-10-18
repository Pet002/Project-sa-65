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
      minWidth: 350,
    },
    tableSpace: {
      marginTop: 20,
    },
    paper: {
       padding: theme.spacing(2),
      color: theme.palette.text.secondary,
      
    },
    root: {
      flexGrow: 1,
    },
  })
);

function PatientCreate() {
  const classes = useStyles();
  const [date, setDate] = React.useState<Date | null>(null);
  const [patient, setPatient] = React.useState<Partial<PatientInterface>>({});
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
    const id = event.target.id as keyof typeof patient;
    const { value } = event.target;
    setPatient({ ...patient, [id]: value });
  };
  const [gender, setGender] = React.useState("");

  const handleChange = (
    event: SelectChangeEvent) => {
    const a = event.target.value as keyof typeof patient;
    const name = event.target.name as keyof typeof patient;
    setPatient({ ...patient, [name]: a });
    console.log(name);
  
  };

  function submit() {
    let data = {
      PID: patient.PID ?? "",
      Name: patient.Name ?? "",
      Surname: patient.Surname ?? "",
      Age: patient.Age,
      Gender: patient.Gender ?? "",
      Allergy: patient.Allergy ?? "",
    };
    console.log(patient);

    const apiUrl = "http://localhost:8080/patient";
    const requestOptions = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
        
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
  <h1>Patient</h1>;

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
          <Box >
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom>
              Create Patient
            </Typography>
          </Box>
        </Box>

        <Divider />

        <Grid container spacing={1} sx={{ padding: 1 }}>
          <Grid item xs={12}>
            <FormControl sx={{ width: 200 }} variant="outlined">
              <p>PID</p>
              <TextField
                id="PID"
                variant="outlined"
                type="string"
                size="medium"
                value={patient.PID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <p>Name</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="Name"
                variant="outlined"
                type="string"
                size="medium"
                value={patient.Name || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Surname</p>
              <TextField
                id="Surname"
                variant="outlined"
                type="string"
                size="medium"
                value={patient.Surname || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Age</p>
              <TextField
                id="Age"
                variant="outlined"
                type="number"
                size="medium"
                InputProps={{ inputProps: { min: 1,max:100 } }}
                InputLabelProps={{
                  shrink: true,
                }}
                value={patient.Age || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl  fullWidth variant="outlined">
              <p>Gender</p>
              <Select
                id="gender"
                name="Gender"
                value={patient.Gender || ""}
                onChange={handleChange}
                >
                <MenuItem value={"หญิง"}>Female</MenuItem>
                <MenuItem value={"ชาย"}>Male</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Allergy</p>
              <TextField
                id="Allergy"
                variant="outlined"
                type="string"
                size="medium"
                value={patient.Allergy || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button component={RouterLink} to="/" variant="contained">
              Back
            </Button>
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

export default PatientCreate;

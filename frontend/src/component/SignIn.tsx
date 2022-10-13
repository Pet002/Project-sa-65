import React, { useState } from "react";

import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import TextField from "@material-ui/core/TextField";

import Typography from "@material-ui/core/Typography";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import frame from "../image/Banner.svg";
import logo from "../image/logo.png";
import roommedicine from "../image/roommedicine.jpg";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Box from "@material-ui/core/Box";
import MedicationIcon from "@mui/icons-material/Medication";
import { SigninInterface } from "../models/ISignin";
function Alert(props: AlertProps) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const useStyles = makeStyles((theme) => ({
  paper: {
    // position: "absolute",
    left: "0vw",
    heigth: "100vh",
    border: "none",
    display: "flex",
  },
  paper1: {
    width: "72.65vw",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  form: {
    width: "100%",
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  title: {
    flexGrow: 1,
  },
  formFrame: {
    width: "30vw",
  },
  logo: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
}));

function SignIn() {
  const classes = useStyles();
  const [signin, setSignin] = useState<Partial<SigninInterface>>({});
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const login = () => {
    const apiUrl = "http://localhost:8080/loginP";
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(signin),
    };
    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSuccess(true);
          localStorage.setItem("token", res.data.token);
          localStorage.setItem("uid", res.data.id);
          window.location.reload();
        } else {
          setError(true);
        }
      });
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof signin;
    const { value } = event.target;
    setSignin({ ...signin, [id]: value });
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
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
      {/* <AppBar
        position="fixed"
        style={{ backgroundColor: "#148F77 ", color: "#F4F6F6" }}>
        <Toolbar>
          <MedicationIcon fontSize="large" ></MedicationIcon>
          <Typography variant="h6" className={classes.title} style={{fontWeight:"bold"}}>
             MEDICINE
          </Typography>
          <Button style={{fontWeight:"bold",border:"groove",color: "#F4F6F6"}} >
            Sign in
          </Button>
         
         
          
        </Toolbar>
      </AppBar> */}
      <div className={classes.paper}>
        <img
          style={{ maxHeight: "100vh" }}
          className="img"
          alt="Banner"
          src={frame}
        />
        <div className={classes.paper1}>
          <div className={classes.formFrame}>
            <div className={classes.logo}>
              <img className="img" alt="logo" src={logo} />
            </div>

            <form noValidate>
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="User"
                label="User"
                name="User"
                autoComplete="user"
                autoFocus
                value={signin.User || ""}
                onChange={handleInputChange}
              />
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                label="Password"
                name="Password"
                type="password"
                id="Password"
                autoComplete="current-password"
                value={signin.Password || ""}
                onChange={handleInputChange}
              />
              <Button
                fullWidth
                variant="contained"
                style={{ backgroundColor: "#148F77", color: "#F4F6F6" }}
                className={classes.submit}
                onClick={login}>
                Sign In
              </Button>
              <Button
                fullWidth
                variant="contained"
                style={{ backgroundColor: "rgba(27, 36, 32, 0.5)", color: "#F4F6F6" ,marginTop: 2}}
                className={classes.submit}
                onClick={login}>
                Sign up
              </Button>
            </form>
          </div>
        </div>
      </div>
    </Box>
  );
}

export default SignIn;

import React, { useEffect } from "react";
import clsx from "clsx";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import {
  createStyles,
  makeStyles,
  useTheme,
  Theme,
} from "@material-ui/core/styles";
import Drawer from "@material-ui/core/Drawer";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import List from "@material-ui/core/List";
import CssBaseline from "@material-ui/core/CssBaseline";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft";
import ChevronRightIcon from "@material-ui/icons/ChevronRight";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import Button from "@material-ui/core/Button";
import ClickAwayListener from "@mui/material/ClickAwayListener";
import Grow from "@mui/material/Grow";
import Paper from "@mui/material/Paper";
import Popper from "@mui/material/Popper";
import MenuItem from "@mui/material/MenuItem";
import MenuList from "@mui/material/MenuList";
import Stack from "@mui/material/Stack";
import HomeIcon from "@material-ui/icons/Home";
import MedicationIcon from "@mui/icons-material/Medication";
import LibraryBooksIcon from "@mui/icons-material/LibraryBooks";
import LocalPharmacyIcon from "@mui/icons-material/LocalPharmacy";
import Home from "./Home";
import Prescription from "./Prescription";
import SignIn from "./SignIn";
import Patient from "./Patient";
import { EmployeeInterface } from "../models/IEmployee";
import ArrowDropDownIcon from "@mui/icons-material/ArrowDropDown";
import PrescriptionHistory from "./PrescriptionHistory";
import PatientCreate from "./PatientCreate";
import PersonIcon from "@mui/icons-material/Person";
import AccountBoxIcon from "@mui/icons-material/AccountBox";
const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: "flex",
    },
    title: {
      flexGrow: 1,
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: 36,
    },
    hide: {
      display: "none",
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
      whiteSpace: "nowrap",
    },
    drawerOpen: {
      width: drawerWidth,
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerClose: {
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      overflowX: "hidden",
      width: theme.spacing(7) + 1,
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9) + 1,
      },
    },
    toolbar: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
    },
    a: {
      textDecoration: "none",
      color: "inherit",
    },
    
  })
);

export default function Main() {
  const classes = useStyles();
  const theme = useTheme();
  const [open1, setOpen1] = React.useState(false);
  const [token, setToken] = React.useState<String>("");
  const [open2, setOpen2] = React.useState(false);
  const anchorRef = React.useRef<HTMLButtonElement>(null);
  const [employee, setEmployee] = React.useState<EmployeeInterface>();

  const handleDrawerOpen = () => {
    setOpen1(true);
  };

  const handleDrawerClose = () => {
    setOpen1(false);
  };
  const handleClose = (event: Event | React.SyntheticEvent) => {
    if (
      anchorRef.current &&
      anchorRef.current.contains(event.target as HTMLElement)
    ) {
      return;
    }

    setOpen2(false);
  };

  const menu = [
    { name: "หน้าแรก", icon: <HomeIcon />, path: "/" },

    { name: "ข้อมูลผู้ป่วย", icon: <PersonIcon />, path: "/patient" },
    {
      name: "สร้างข้อมูลผู้ป่วย",
      icon: <AccountBoxIcon />,
      path: "/PatientCreate",
    },
    { name: "สั่งยา", icon: <LocalPharmacyIcon />, path: "/prescription" },
    {
      name: "ประวัติการสั่งยา",
      icon: <LibraryBooksIcon />,
      path: "/PrescriptionHistory",
    },
  ];

  const getEmployee = () => {
    const apiUrl = "http://localhost:8080";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/employees/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log(res.data);
          setEmployee(res.data);
        }
      });
  };

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
      getEmployee();
    }
  }, []);

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };
  

  return (
    <div className={classes.root}>
      <Router>
        <CssBaseline />
        {token && (
          <>
            <AppBar
              position="fixed"
              style={{ backgroundColor: "#148F77 ", color: "#F4F6F6" }}
              className={clsx(classes.appBar, {
                [classes.appBarShift]: open1,
              })}>
              <Toolbar>
                <IconButton
                  color="inherit"
                  aria-label="open drawer"
                  onClick={handleDrawerOpen}
                  edge="start"
                  className={clsx(classes.menuButton, {
                    [classes.hide]: open1,
                  })}>
                  <MenuIcon />
                </IconButton>
                <MedicationIcon fontSize="large"></MedicationIcon>
                <Typography
                  variant="h6"
                  className={classes.title}
                  style={{ fontWeight: "bold" }} >
                  PRESCRIPTION
                </Typography>

                <Stack direction="row" spacing={2}>
                  <div>
                    <Button
                      ref={anchorRef}
                      id="composition-button"
                      aria-controls={open2 ? "composition-menu" : undefined}
                      aria-expanded={open2 ? "true" : undefined}
                      aria-haspopup="true"
                      onClick={() => setOpen2(!open2)} style={{border: '2px solid ',color:"#F4F6F6"}}>
                     {employee?.Name} {employee?.Surname}{" "}
                    </Button>
                    <Popper
                      open={open2}
                      anchorEl={anchorRef.current}
                      role={undefined}
                      placement="bottom-start"
                      transition
                      disablePortal>
                      {({ TransitionProps, placement }) => (
                        <Grow
                          {...TransitionProps}
                          style={{
                            transformOrigin:
                              placement === "bottom-start"
                                ? "left top"
                                : "left bottom",
                          }}>
                          <Paper>
                            <ClickAwayListener onClickAway={handleClose}>
                              <MenuList
                                autoFocusItem={open2}
                                id="composition-menu"
                                aria-labelledby="composition-button">
                                <MenuItem >
                                 Profile
                                </MenuItem>
                                <MenuItem onClick={signout}>Sign Out</MenuItem>
                              </MenuList>
                            </ClickAwayListener>
                          </Paper>
                        </Grow>
                      )}
                    </Popper>
                  </div>
                </Stack>
              </Toolbar>
            </AppBar>
            <Drawer
              variant="permanent"
              className={clsx(classes.drawer, {
                [classes.drawerOpen]: open1,
                [classes.drawerClose]: !open1,
              })}
              classes={{
                paper: clsx({
                  [classes.drawerOpen]: open1,
                  [classes.drawerClose]: !open1,
                }),
              }}>
              <div className={classes.toolbar}>
                <IconButton onClick={handleDrawerClose}>
                  {theme.direction === "rtl" ? (
                    <ChevronRightIcon />
                  ) : (
                    <ChevronLeftIcon />
                  )}
                </IconButton>
              </div>
              <Divider />
              <List>
                {menu.map((item, index) => (
                  <Link to={item.path} key={item.name} className={classes.a}>
                    <ListItem button>
                      <ListItemIcon>{item.icon}</ListItemIcon>
                      <ListItemText primary={item.name} />
                    </ListItem>
                  </Link>
                ))}
              </List>
            </Drawer>
          </>
        )}

        <main className={classes.content}>
          <div className={classes.toolbar} />
          <div>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/Prescription" element={<Prescription />} />
              <Route path="/Patient" element={<Patient />} />
              <Route
                path="/PrescriptionHistory"
                element={<PrescriptionHistory />}
              />
              <Route path="/PatientCreate" element={<PatientCreate />} />
            </Routes>
          </div>
        </main>
      </Router>
    </div>
  );
}

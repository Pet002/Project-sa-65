import React from "react";
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import TextField from "@mui/material/TextField";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
//import dayjs, { Dayjs } from 'dayjs';
import Stack from '@mui/material/Stack';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { TimePicker } from '@mui/x-date-pickers/TimePicker';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { DesktopDatePicker } from '@mui/x-date-pickers/DesktopDatePicker';
import { MobileDatePicker } from '@mui/x-date-pickers/MobileDatePicker';

function App() {
  // export default function MaterialUIPickers() {
  //   const [value, setValue] = React.useState<Dayjs | null>(
  //     dayjs('2014-08-18T21:11:54'),
  //   );

  //   const handleChange = (newValue: Dayjs | null) => {
  //     setValue(newValue);
  //   };


    return (

      <div>
        <Box sx={{ flexGrow: 1 }}>
          <AppBar position="static">
            <Toolbar>
              <IconButton
                size="large"
                edge="start"
                color="inherit"
                aria-label="menu"
                sx={{ mr: 2 }}
              >
                <MenuIcon />
              </IconButton>
              <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                ฉลากยา
              </Typography>
              <Button color="inherit">Login</Button>
            </Toolbar>
          </AppBar>
        </Box>

        <Container maxWidth="md">
          <Paper>
            <Box
              display={"flex"}
              sx={{
                marginTop: 2,
                paddingX: 1,
                paddingY: 1,
              }}
            >
              <h2>ฉลากยา</h2>
            </Box>
            <hr />
            <Grid>

              <Grid container spacing={10}>
                <Grid item xs={2}>
                  <p>วิธีการใช้ยา</p>
                </Grid>
                <Grid item xs={7}>
                  <TextField fullWidth id="วิธีการใช้ยา" variant="outlined" />
                </Grid>
              </Grid>

              <Grid container spacing={10}>
                <Grid item xs={2}>
                  <p>คำเตือน</p>
                </Grid>
                <Grid item xs={7}>
                  <TextField fullWidth id="คำเตือน" variant="outlined" />
                </Grid>
              </Grid>

              <Grid container spacing={10}>
                <Grid item xs={2}>
                  <p>เภสัชกร</p>
                </Grid>
                <Grid item xs={7}>
                  <TextField fullWidth id="เภสัชกร" variant="outlined" />
                </Grid>
              </Grid>

              <Grid container spacing={10}>
                <Grid item xs={2}>
                  <p>วันที่และเวลา</p>
                </Grid>
                <Grid item xs={7}>
                  <TextField fullWidth id="วันที่และเวลา" variant="outlined" />
                </Grid>
                {/* <LocalizationProvider dateAdapter={AdapterDayjs}>
                  <Stack spacing={3}>
                    <DateTimePicker
                      label="วันที่และเวลา"
                      value={value}
                      onChange={handleChange}
                      renderInput={(params) => <TextField {...params} />}
                    />
                  </Stack>
                </LocalizationProvider> */}
              </Grid>

              <Grid container spacing={10}>
                <Grid item xs={12}>
                  <Button variant="contained" color="info">
                    ยกเลิก
                  </Button>
                  <Button variant="contained" color="success" sx={{ float: "right" }}>
                    บันทึกฉลากยา
                  </Button>
                </Grid>
              </Grid>

            </Grid>
          </Paper>
        </Container>
      </div>


    );
  //}
}

export default App;
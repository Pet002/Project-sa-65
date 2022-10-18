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
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
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

function createData(
  Patient: string,
  Symptom: string,
  Medicine: string,
  Pharmacist: string
) {
  return { Patient, Symptom, Medicine, Pharmacist };
}

const rows = [createData("Gingerbread", "bb", "tonphaii", "apple")];

function PrescriptionHistory() {
  <h1>History</h1>;
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table" >
        <TableHead >
          <TableRow>
            <TableCell>Patient</TableCell>
            <TableCell>Symptom</TableCell>
            <TableCell>Medicine</TableCell>
            <TableCell>Pharmacist</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow
              key={row.Patient}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
              <TableCell component="th" scope="row">
                {row.Patient}
              </TableCell>
              <TableCell>{row.Symptom}</TableCell>
              <TableCell>{row.Medicine}</TableCell>
              <TableCell>{row.Pharmacist}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
export default PrescriptionHistory;

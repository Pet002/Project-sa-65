import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import React, { useState } from "react";
import Container from "@material-ui/core/Container";
import Typography from "@material-ui/core/Typography";
import prescription from "../image/prescription.jpg";
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
    gallery: {
      margin: 5,
      border: 1,
      width: 180,
    },paper: {
      marginTop: theme.spacing(8),
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
    
      
    },
  
  })
);

function Home() {
  const classes = useStyles();
  // return focus to the button when we transitioned from !open -> open

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Typography  className={classes.paper}>
          <img
            style={{ width: "500px" }}
            className="img"
            src={prescription}></img>
        </Typography>

        {/* <h1 style={{ textAlign: "center" }}>ระบบสั่งยา</h1> */}
        <h4>Requirements</h4>
        <p>
        ระบบสั่งยา เป็นระบบที่ให้เภสัชกรภายในโรงพยาบาลแห่งหนึ่งสามารถ Login เข้าสู่ระบบมาเพื่อทำหน้าที่สั่งยาให้กับผู้ป่วยแต่ละคน
        โดยการสั่งยาของผู้ป่วยแต่ละคนนั้น เจ้าหน้าที่จะต้องกรอกข้อมูลของผู้ป่วย เลือกชื่อยาที่ต้องจ่าย เก็บไว้รายการยา จากนั้นเมื่อทำการกดบันทึกรายการเรียบร้อยแล้ว
ระบบจะบันทึกข้อมูลที่ทำรายการพร้อมกับช่วงวันเวลาที่ทำรายการไปที่ใบสั่งยา

          <br />
          ระบบสั่งยา เป็นระบบที่ให้เจ้าหน้าที่แต่ละคนสามารถเรียกดูประวัติการทำรายการย้อนหลังได้ว่า
รายการไหนเป็นของผู้ป่วยชื่อว่าอะไร มีประวัติการใช้ยาตัวไหนบ้าง และทำรายการในช่วงวันเวลาไหน
        </p>
      </Container>
    </div>
  );
}
export default Home;

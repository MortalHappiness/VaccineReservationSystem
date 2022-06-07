import { useEffect, useState } from "react";
import TextField from "@mui/material/TextField";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Paper from "@mui/material/Paper";
import CancelIcon from "@mui/icons-material/Cancel";
import { ReservationAPI } from "../api";

export default function ReservationStatus(props) {
  const [completedReservation, setCompletedReservation] = useState([]);
  const [unfinishedReservation, setUnfinishedReservation] = useState([]);
  const cancelReservation = (nationID, reservationID) => {
    ReservationAPI.deleteReservation(nationID, reservationID)
      .then((res) => {
        props.setUserReservations(
          props.userReservations.filter((ele) => ele.id !== res)
        );
      })
      .catch();
  };
  useEffect(() => {
    if (props.userReservations) {
      setCompletedReservation(
        props.userReservations.filter((ele) => ele.completed === true)
      );
      setUnfinishedReservation(
        props.userReservations.filter((ele) => ele.completed === false)
      );
    } else {
      setCompletedReservation([]);
      setUnfinishedReservation([]);
    }
  }, []);
  return (
    <>
      {/* <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignContent: "center",
          alignItems: "center",
          mb: 5,
        }}
      >
        <Typography>身分證字號：</Typography>
        <TextField
          id="reservation_status_id"
          variant="outlined"
          size="small"
          sx={{
            mx: 1,
            width: 140,
          }}
        />
        <Button
          variant="contained"
          sx={{
            mx: 1,
          }}
        >
          查詢
        </Button>
      </Box> */}
      <Typography sx={{ fontSize: 20, my: 2 }}>已接種或過期</Typography>
      {completedReservation.length > 0 ? (
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell sx={{ width: 90, m: 0, px: 0, textAlign: "center" }}>
                  日期時間
                </TableCell>
                <TableCell
                  align="left"
                  sx={{ width: 100, m: 0, px: 0, textAlign: "center" }}
                >
                  醫療院所名稱
                </TableCell>
                <TableCell sx={{ m: 0, px: 0, textAlign: "center" }}>
                  地址
                </TableCell>
                <TableCell
                  align="right"
                  sx={{ width: 80, m: 0, px: 0, textAlign: "center" }}
                >
                  疫苗種類
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {completedReservation.map((row) => {
                return (
                  <TableRow
                    key={row.id}
                    sx={{
                      "&:last-child td, &:last-child th": { border: 0 },
                    }}
                  >
                    <TableCell
                      align="left"
                      sx={{ width: 90, m: 0, px: 0, textAlign: "center" }}
                    >
                      {new Date(row.date).toISOString().substring(0, 10)}
                    </TableCell>
                    <TableCell
                      align="left"
                      sx={{ width: 100, m: 0, px: 0, textAlign: "center" }}
                    >
                      {row.hospital.name}
                    </TableCell>
                    <TableCell
                      align="right"
                      sx={{ m: 0, px: 0, textAlign: "center" }}
                    >
                      {row.hospital.address}
                    </TableCell>
                    <TableCell
                      align="right"
                      sx={{ width: 80, m: 0, px: 0, textAlign: "center" }}
                    >
                      {row.vaccinetype}
                    </TableCell>
                  </TableRow>
                );
              })}
            </TableBody>
          </Table>
        </TableContainer>
      ) : (
        <Typography>尚無已接種或過時紀錄</Typography>
      )}

      <Typography sx={{ fontSize: 20, my: 2 }}>待完成</Typography>
      {unfinishedReservation.length > 0 ? (
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="left">日期時間</TableCell>
                <TableCell align="left">醫療院所名稱</TableCell>
                <TableCell align="right">地址</TableCell>
                <TableCell align="right">疫苗種類</TableCell>
                <TableCell align="right">取消</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {unfinishedReservation.map((row) => {
                return (
                  <TableRow
                    key={row.id}
                    sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                  >
                    <TableCell align="left">{row.date}</TableCell>
                    <TableCell align="left">{row.hospital.name}</TableCell>
                    <TableCell align="right">{row.hospital.address}</TableCell>
                    <TableCell align="right">{row.vaccinetype}</TableCell>
                    <TableCell
                      align="right"
                      onClick={() =>
                        cancelReservation(props.user.nationID, row.id)
                      }
                    >
                      <CancelIcon />
                    </TableCell>
                  </TableRow>
                );
              })}
            </TableBody>
          </Table>
        </TableContainer>
      ) : (
        <Typography>尚無未接種預約紀錄</Typography>
      )}
    </>
  );
}

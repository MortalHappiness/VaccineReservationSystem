import { useEffect, useState } from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Typography from "@mui/material/Typography";
import Paper from "@mui/material/Paper";
import CancelIcon from "@mui/icons-material/Cancel";
// api
import { ReservationAPI } from "../api";
import moment from "moment";

export default function ReservationStatus({ user }) {
  // get reservations
  const [reservations, setReservations] = useState([]);
  useEffect(() => {
    const fetchData = async () => {
      try {
        const reservations = await ReservationAPI.getReservations(
          user.nationID
        );
        setReservations(reservations);
      } catch (err) {
        console.error(err);
      }
    };
    fetchData();
  }, []);

  const completedReservation = reservations.filter(
    (reservation) => reservation.completed || Date.now()/1000 >= reservation.date
  );
  const unfinishedReservation = reservations.filter(
    (reservation) => !reservation.completed && Date.now()/1000 < reservation.date
  );

  // mutation
  const handleCancelReservation = async (reservationID) => {
    try {
      const deletedID = await ReservationAPI.deleteReservation(
        user.nationID,
        reservationID
      );
      setReservations(
        reservations.filter((reservation) => reservation.id !== deletedID)
      );
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <>
      <Typography sx={{ fontSize: 20, my: 2 }}>已接種或過期</Typography>
      {completedReservation.length > 0 ? (
        <TableContainer component={Paper} variant="outlined">
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="left">日期時間</TableCell>
                <TableCell align="left">醫療院所名稱</TableCell>
                <TableCell align="left">地址</TableCell>
                <TableCell align="left">疫苗種類</TableCell>
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
                    <TableCell align="left">
                      {moment.unix(row.date).format("YYYY/MM/DD HH:mm")}
                    </TableCell>
                    <TableCell align="left">{row.hospital.name}</TableCell>
                    <TableCell align="left">{row.hospital.address}</TableCell>
                    <TableCell align="left">{row.vaccineType}</TableCell>
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
        <TableContainer component={Paper} variant="outlined">
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="left">日期時間</TableCell>
                <TableCell align="left">醫療院所名稱</TableCell>
                <TableCell align="left">地址</TableCell>
                <TableCell align="left">疫苗種類</TableCell>
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
                    <TableCell align="left">
                      {moment.unix(row.date).format("YYYY/MM/DD HH:mm")}
                    </TableCell>
                    <TableCell align="left">{row.hospital.name}</TableCell>
                    <TableCell align="left">{row.hospital.address}</TableCell>
                    <TableCell align="left">{row.vaccineType}</TableCell>
                    <TableCell
                      align="right"
                      onClick={() => handleCancelReservation(row.id)}
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

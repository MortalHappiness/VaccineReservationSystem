import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell, { tableCellClasses } from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { createTheme } from "@mui/material";
const theme = createTheme({
  subject: {
    borderRight: "1px solid gray",
    width: 80,
  },
  detail: {
    borderRight: "2px solid black",
  },
});
export default function User(props) {
  return (
    <>
      {props.user === null ? (
        <h1>請先登入</h1>
      ) : (
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 700 }} aria-label="customized table">
            <TableBody>
              <TableRow>
                <TableCell sx={theme.subject}>姓名</TableCell>
                <TableCell sx={theme.detail}>{props.user.name}</TableCell>
                <TableCell sx={theme.subject}>性別</TableCell>
                <TableCell>{props.user.gender}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell sx={theme.subject}>身分證字號</TableCell>
                <TableCell sx={theme.detail}>{props.user.nationID}</TableCell>
                <TableCell sx={theme.subject}>生日</TableCell>
                <TableCell>{props.user.birthDay}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell sx={theme.subject}>地址</TableCell>
                <TableCell sx={theme.detail}>{props.user.address}</TableCell>
                <TableCell sx={theme.subject}>連絡電話</TableCell>
                <TableCell>{props.user.phone}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell sx={theme.subject}>疫苗紀錄</TableCell>
                <TableCell>{props.user.vaccines}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      )}
    </>
  );
}

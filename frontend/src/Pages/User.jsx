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
    width: 120,
  },
  detail: {
    borderRight: "1px solid rgba(224, 224, 224, 1)",
  },
});
export default function User({ user }) {
  return (
    <TableContainer component={Paper} variant="outlined">
      <Table sx={{ minWidth: 700 }} aria-label="customized table">
        <TableBody>
          <TableRow>
            <TableCell sx={theme.subject}>姓名</TableCell>
            <TableCell sx={theme.detail}>{user.name}</TableCell>
            <TableCell sx={theme.subject}>性別</TableCell>
            <TableCell>{user.gender}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell sx={theme.subject}>身分證字號</TableCell>
            <TableCell sx={theme.detail}>{user.nationID}</TableCell>
            <TableCell sx={theme.subject}>生日</TableCell>
            <TableCell>{user.birthDay}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell sx={theme.subject}>地址</TableCell>
            <TableCell sx={theme.detail}>{user.address}</TableCell>
            <TableCell sx={theme.subject}>連絡電話</TableCell>
            <TableCell>{user.phone}</TableCell>
          </TableRow>
          <TableRow>
            <TableCell sx={theme.subject}>疫苗紀錄</TableCell>
            <TableCell>{user.vaccines.join(", ")}</TableCell>
            <TableCell></TableCell>
            <TableCell></TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </TableContainer>
  );
}

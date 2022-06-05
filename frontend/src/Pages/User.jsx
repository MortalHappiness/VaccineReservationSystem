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
export default function User() {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 700 }} aria-label="customized table">
        <TableBody>
          <TableRow>
            <TableCell sx={theme.subject}>姓名</TableCell>
            <TableCell sx={theme.detail}>劉展碩</TableCell>
            <TableCell sx={theme.subject}>性別</TableCell>
            <TableCell>男</TableCell>
          </TableRow>
          <TableRow>
            <TableCell sx={theme.subject}>身分證字號</TableCell>
            <TableCell sx={theme.detail}>A123456789</TableCell>
            <TableCell sx={theme.subject}>生日</TableCell>
            <TableCell>88-12-11</TableCell>
          </TableRow>
          <TableRow>
            <TableCell sx={theme.subject}>地址</TableCell>
            <TableCell sx={theme.detail}>台北市...</TableCell>
            <TableCell sx={theme.subject}>連絡電話</TableCell>
            <TableCell>0988121314</TableCell>
          </TableRow>
          <TableRow>
            <TableCell sx={theme.subject}>疫苗紀錄</TableCell>
            <TableCell>...</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </TableContainer>
  );
}

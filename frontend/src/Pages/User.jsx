import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell, { tableCellClasses } from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
export default function User() {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 700 }} aria-label="customized table">
        <TableBody>
          <TableRow>
            <TableCell>姓名</TableCell>
            <TableCell>劉展碩</TableCell>
            <TableCell>性別</TableCell>
            <TableCell>男</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>身分證字號</TableCell>
            <TableCell>A123456789</TableCell>
            <TableCell>生日</TableCell>
            <TableCell>88-12-11</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>地址</TableCell>
            <TableCell>台北市...</TableCell>
            <TableCell>連絡電話</TableCell>
            <TableCell>0988121314</TableCell>
          </TableRow>
          <TableRow>
            <TableCell>疫苗紀錄</TableCell>
            <TableCell>第一劑</TableCell>
            <TableCell>第二劑</TableCell>
            <TableCell>第三劑</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </TableContainer>
  );
}

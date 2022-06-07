import { useState, useEffect } from "react";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormHelperText from "@mui/material/FormHelperText";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
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
import { Snackbar, Alert } from "@mui/material";
import Paper from "@mui/material/Paper";
import { HospitalAPI, ReservationAPI } from "../api";
import { createTheme } from "@mui/material";
import { area_data, initHospitalData } from "../AreaData";
import useSWR from "swr";

export default function ReservationSearch() {
  const theme = createTheme({
    subject: {
      my: 2,
    },
  });

  const counties = Object.keys(area_data);
  const [county, setCounty] = useState(counties[0]);
  const [alert, setAlert] = useState({});

  const { data, error } = useSWR(
    [county, area_data[county][0]],
    HospitalAPI.getHospital
  );
  const [hospitalData, setHospitalData] = useState(initHospitalData);
  useEffect(() => {
    // console.log("data:", data);

    // console.log("ERR:", error);
    if (error) {
      setAlert({
        open: true,
        severity: "error",
        msg: error,
      });
    }
    // TODO data change to array
    if (data !== undefined) {
      console.log(data);
      let newHospitalData = hospitalData;
      newHospitalData[data.county][data.township] = [data];
      setHospitalData(newHospitalData);
    }
  }, [county, data]);

  const handleChange = (event) => {
    setCounty(event.target.value);
  };

  const handleSearch = () => {
    HospitalAPI.getHospital();
  };

  return (
    <>
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignContent: "center",
          alignItems: "center",
          mb: 1,
        }}
      >
        <Snackbar
          anchorOrigin={{ vertical: "top", horizontal: "center" }}
          open={alert?.open}
          autoHideDuration={3000}
          onClose={() => setAlert({ ...alert, open: false })}
        >
          <Alert variant="filled" severity={alert?.severity}>
            {alert?.msg}
            {
              <Button color="inherit" size="small" onClick={alert?.action}>
                Ok
              </Button>
            }
          </Alert>
        </Snackbar>
        <Typography>縣市：</Typography>
        <FormControl sx={{ mr: 1 }}>
          <Select
            value={county}
            onChange={handleChange}
            displayEmpty
            sx={{ mx: 1, height: 40 }}
          >
            {counties.map((ele) => (
              <MenuItem
                key={ele}
                value={ele}
                sx={{
                  fontSize: 8,
                }}
              >
                {ele}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <Button
          variant="contained"
          sx={{
            mx: 1,
          }}
          onClick={handleSearch}
        >
          查詢
        </Button>
      </Box>
      {area_data[county].map((township) => {
        return (
          <Box key={township}>
            <Typography sx={theme.subject}>{township}</Typography>
            <TableContainer component={Paper}>
              <Table sx={{ minWidth: 650 }} aria-label="simple table">
                <TableBody>
                  {hospitalData[county][township].length > 0 &&
                    hospitalData[county][township].map((hospital) => (
                      <TableRow sx={{ height: 20 }} key={hospital.name}>
                        <TableCell>{hospital.name}</TableCell>
                        <TableCell align="right">
                          {Object.keys(hospital.vaccineCnt)}
                        </TableCell>
                        <TableCell align="right">
                          {hospital.vaccineCnt["BNT"]}
                        </TableCell>
                      </TableRow>
                    ))}
                </TableBody>
              </Table>
            </TableContainer>
          </Box>
        );
      })}
    </>
  );
}

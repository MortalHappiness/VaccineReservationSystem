import { useState } from "react";
// mui
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Select from "@mui/material/Select";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import Snackbar from "@mui/material/Snackbar";
import Alert from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardActions from "@mui/material/CardActions";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
import { AdapterMoment } from "@mui/x-date-pickers/AdapterMoment";
// api
import { HospitalAPI, ReservationAPI } from "../api";
// utils
import { v4 as uuidv4 } from "uuid";
import { areaData } from "../utils/AreaData";

export default function ReservationSearch({ user }) {
  // form input
  const [county, setCounty] = useState("");
  const [township, setTownship] = useState("");
  const handleChangeCounty = (event) => {
    setCounty(event.target.value);
    setTownship("");
  };
  const handleChangeTownship = (event) => {
    setTownship(event.target.value);
  };
  const counties = Object.keys(areaData);
  const townships = areaData[county] || [];

  // search
  const [alert, setAlert] = useState({});
  const [hospitals, setHospitals] = useState([]);
  const handleSearch = async () => {
    if (county == "" || township == "") {
      return;
    }
    try {
      const res = await HospitalAPI.getHospital(county, township);
      setHospitals(res);
    } catch (err) {
      setAlert({
        severity: "warning",
        msg: err.message,
        open: true,
      });
    }
  };

  // reservation dialog
  const [hospital, setHospital] = useState({});
  const [openDialog, setOpenDialog] = useState(false);
  const [vaccineType, setVaccineType] = useState("");
  const [date, setDate] = useState(null);
  const handleOpenDialog = (hosp) => {
    setHospital(hosp);
    setOpenDialog(true);
    setVaccineType("");
    setDate(null);
  };
  const handleCloseDialog = () => setOpenDialog(false);

  const handleChangeVaccineType = (event) => {
    setVaccineType(event.target.value);
  };
  const handleChangeDate = (d) => {
    setDate(d);
  };
  const handleReservation = async () => {
    if (vaccineType == "" || date == null) {
      return;
    }
    try {
      const res = await ReservationAPI.createReservation({
        id: uuidv4(),
        user,
        hospital,
        date: date.unix(),
        completed: false,
        vaccineType,
      });
      setAlert({
        severity: "success",
        msg: "預約成功",
        open: true,
      });
    } catch (err) {
      setAlert({
        severity: "error",
        msg: err.message,
        open: true,
      });
    }
    setOpenDialog(false);
  };

  // utils
  // return vaccines that left > 0
  const getVaccineType = (vaccineCnt) => {
    if (!vaccineCnt) return [];
    return Object.keys(vaccineCnt).filter((key) => vaccineCnt[key] > 0);
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
        <Typography>縣市：</Typography>
        <FormControl sx={{ mr: 1 }}>
          <Select
            value={county}
            onChange={handleChangeCounty}
            displayEmpty
            sx={{ mx: 1, height: 40, width: 100 }}
          >
            {counties.map((ele) => (
              <MenuItem key={ele} value={ele}>
                {ele}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <Typography>鄉鎮市區：</Typography>
        <FormControl sx={{ mr: 1 }}>
          <Select
            value={township}
            onChange={handleChangeTownship}
            displayEmpty
            sx={{ mx: 1, height: 40, width: 100 }}
          >
            {townships.map((ele) => (
              <MenuItem key={ele} value={ele}>
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
      {hospitals.length > 0 &&
        hospitals
          .filter((hosp) => getVaccineType(hosp.vaccineCnt) != "")
          .map((hosp) => (
            <Card key={hosp.id} variant="outlined" sx={{ mb: 2 }}>
              <CardContent>
                <Typography variant="h5" component="h2">
                  {hosp.name}
                </Typography>
                <Typography variant="body2" component="p">
                  {hosp.address}
                </Typography>
                <Typography variant="body2" component="p">
                  疫苗種類：{getVaccineType(hosp.vaccineCnt).join(", ")}
                </Typography>
              </CardContent>
              <CardActions>
                <Button onClick={() => handleOpenDialog(hosp)}>預約</Button>
              </CardActions>
            </Card>
          ))}
      <Snackbar
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
        open={alert?.open}
        autoHideDuration={1000}
        onClose={() => setAlert({ ...alert, open: false })}
      >
        <Alert variant="filled" severity={alert?.severity}>
          {alert?.msg}
        </Alert>
      </Snackbar>
      <Dialog open={openDialog} onClose={handleCloseDialog}>
        <DialogTitle>預約</DialogTitle>
        <DialogContent>
          <DialogContentText>醫院: {hospital.name}</DialogContentText>
          <DialogContentText>地址: {hospital.address}</DialogContentText>

          <FormControl fullWidth sx={{ mt: 2 }} size="small">
            <InputLabel id="vaccine-type-label">疫苗種類</InputLabel>
            <Select
              labelId="vaccine-type-label"
              value={vaccineType}
              onChange={handleChangeVaccineType}
              displayEmpty
              label="疫苗種類"
            >
              {getVaccineType(hospital.vaccineCnt).map((ele) => (
                <MenuItem key={ele} value={ele}>
                  {ele}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <LocalizationProvider dateAdapter={AdapterMoment}>
              <DateTimePicker
                label="預約時間"
                value={date}
                onChange={handleChangeDate}
                disablePast
                renderInput={(params) => <TextField {...params} size="small" />}
              />
            </LocalizationProvider>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDialog} color="primary">
            取消
          </Button>
          <Button onClick={handleReservation} color="primary">
            預約
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}

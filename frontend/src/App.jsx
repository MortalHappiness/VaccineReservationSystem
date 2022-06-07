import { useState, useEffect } from "react";
import useSWR from "swr";
import "./App.css";
import { styled } from "@mui/material/styles";
import Sidebar from "./Sidebar";
import NavigationBar from "./NavigationBar";
import ReservationSearch from "./Pages/ReservationSearch";
import ReservationStatus from "./Pages/ReservationStatus";
import User from "./Pages/User";
import News from "./Pages/News";
import { Navigate, Route, Routes } from "react-router-dom";
import Box from "@mui/material/Box";
import CssBaseline from "@mui/material/CssBaseline";
import { ReservationAPI, SessionAPI } from "./api";
const DrawerHeader = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  justifyContent: "flex-end",
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
}));

function App() {
  const [loginOpen, setLoginOpen] = useState(false);
  const [user, setUser] = useState(null);
  const [userReservations, setUserReservations] = useState([]);
  // const {vaccineCnt, vaccineCntError} = useSWR("", )

  useEffect(() => {
    // 使用瀏覽器 API 更新文件標題
    if (user !== null) {
      ReservationAPI.getReservation(user.nationID)
        .then((res) => setUserReservations(res))
        .catch
        //TODO error
        ();
    } else {
      setUserReservations([]);
    }
  }, [user]);

  useEffect(() => {
    SessionAPI.getSession().then((res) => {
      console.log(res);
      setUser(res);
    });
  }, []);

  return (
    <Box sx={{ display: "flex" }}>
      <NavigationBar
        loginOpen={loginOpen}
        setLoginOpen={setLoginOpen}
        user={user}
        setUser={setUser}
        setUserReservations={setUserReservations}
      />
      <Sidebar />

      <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
        <DrawerHeader />
        <Routes>
          <Route path="/" element={<News />} />
          <Route path="/user" element={<User user={user} />} />
          <Route path="/reservation" element={<ReservationSearch />} />
          <Route
            path="/status"
            element={
              <ReservationStatus
                userReservations={userReservations}
                setUserReservations={setUserReservations}
                user={user}
              />
            }
          />
        </Routes>
      </Box>
    </Box>
  );
}

export default App;

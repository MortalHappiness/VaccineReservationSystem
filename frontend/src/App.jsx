import { useState } from "react";
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
  return (
    <Box sx={{ display: "flex" }}>
      <NavigationBar loginOpen={loginOpen} setLoginOpen={setLoginOpen} />
      <Sidebar />

      <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
        <DrawerHeader />
        <Routes>
          <Route path="/" element={<News />} />
          <Route path="/user" element={<User />} />
          <Route path="/reservation" element={<ReservationSearch />} />
          <Route path="/status" element={<ReservationStatus />} />
        </Routes>
      </Box>
    </Box>
  );
}

export default App;

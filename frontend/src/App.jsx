import "./App.css";
import { Route, Switch, Redirect } from "react-router-dom";
// hooks
import { useState, useEffect } from "react";
import { styled } from "@mui/material/styles";
// mui
import { createTheme } from "@mui/material/styles";
import { ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import CircularProgress from "@mui/material/CircularProgress";
// components
import NavigationBar from "./components/NavigationBar";
import SideBar from "./components/Sidebar";
// pages
import ReservationSearch from "./pages/ReservationSearch";
import ReservationStatus from "./pages/ReservationStatus";
import User from "./pages/User";
import Login from "./pages/Login";
// api
import { SessionAPI } from "./api";

const DrawerHeader = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  justifyContent: "flex-end",
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
}));

const themeOptions = {
  palette: {
    type: "light",
    primary: {
      main: "#0257b8",
    },
    secondary: {
      main: "#f50057",
    },
  },
};

const theme = createTheme(themeOptions);

function App() {
  // drawer functions
  const [openDrawer, setOpenDrawer] = useState(false);
  const handleOpenDrawer = () => setOpenDrawer(true);
  const handleCloseDrawer = () => setOpenDrawer(false);

  // login
  const [fetching, setFetching] = useState(true);
  const [isLogin, setIsLogin] = useState(false);
  const [user, setUser] = useState({});
  useEffect(() => {
    const checkLogin = async () => {
      try {
        const user = await SessionAPI.getSession();
        setIsLogin(true);
        setUser(user);
      } catch (err) {
        console.error(err);
      }
      setFetching(false);
    };
    checkLogin();
  }, []);

  // logout
  const handleLogout = async () => {
    try {
      await SessionAPI.deleteSession();
      setIsLogin(false);
      setUser({});
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Box sx={{ display: "flex" }}>
        <NavigationBar
          isLogin={isLogin}
          handleLogout={handleLogout}
          user={user}
          handleOpenDrawer={handleOpenDrawer}
        />
        <SideBar open={openDrawer} handleCloseDrawer={handleCloseDrawer} />
        <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
          <DrawerHeader />
          <Container>
            {fetching ? (
              <CircularProgress />
            ) : (
              <Switch>
                <Route exact path="/">
                  {isLogin ? (
                    <User user={user} />
                  ) : (
                    <Login setUser={setUser} setIsLogin={setIsLogin} />
                  )}
                </Route>
                <Route path="/reservation">
                  {isLogin ? (
                    <ReservationSearch user={user} />
                  ) : (
                    <Redirect to="/" />
                  )}
                </Route>
                <Route path="/status">
                  {isLogin ? (
                    <ReservationStatus user={user} />
                  ) : (
                    <Redirect to="/" />
                  )}
                </Route>
              </Switch>
            )}
          </Container>
        </Box>
      </Box>
    </ThemeProvider>
  );
}

export default App;

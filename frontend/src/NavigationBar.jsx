import { useState } from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Login from "./Pages/Login";
import MenuItem from "@mui/material/MenuItem";
import Menu from "@mui/material/Menu";
import { Link } from "react-router-dom";
import { SessionAPI } from "./api";

export default function NavigationBar(props) {
  const [anchorEl, setAnchorEl] = useState(null);
  const handleMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };
  const handleLogout = () => {
    SessionAPI.deleteSession()
      .then((res) => {
        setAnchorEl(null);
        props.setUser(null);
        props.setUserReservations([]);
      })
      .catch();
  };
  return (
    <AppBar position="fixed">
      <Toolbar>
        <IconButton
          size="large"
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{ mr: 2 }}
        >
          <MenuIcon />
        </IconButton>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          疫苗系統
        </Typography>
        {props.user === null ? (
          <Button color="inherit" onClick={() => props.setLoginOpen(true)}>
            登入
          </Button>
        ) : (
          <>
            <Button color="inherit" onClick={handleMenu}>
              {props.user.name}
            </Button>
            <Menu
              id="menu-appbar"
              anchorEl={anchorEl}
              anchorOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              open={Boolean(anchorEl)}
              onClose={handleClose}
            >
              <MenuItem onClick={handleClose} component={Link} to="/user">
                個人資料
              </MenuItem>
              <MenuItem onClick={handleLogout}>登出</MenuItem>
            </Menu>
          </>
        )}
        <Login
          open={props.loginOpen}
          setLoginOpen={props.setLoginOpen}
          setUser={props.setUser}
        />
      </Toolbar>
    </AppBar>
  );
}

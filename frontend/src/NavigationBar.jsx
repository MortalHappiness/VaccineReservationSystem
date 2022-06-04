import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Login from "./Pages/Login";

export default function NavigationBar(props) {
  return (
    // <Box sx={{ flexGrow: 1 }}>
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
        <Button color="inherit" onClick={()=>props.setLoginOpen(true)}>登入</Button>
        <Login open={props.loginOpen} setLoginOpen={props.setLoginOpen}/>
      </Toolbar>
    </AppBar>
    // </Box>
  );
}

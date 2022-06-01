import * as React from "react";
import Box from "@mui/material/Box";
import Drawer from "@mui/material/Drawer";
import CssBaseline from "@mui/material/CssBaseline";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import List from "@mui/material/List";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import InboxIcon from "@mui/icons-material/MoveToInbox";
import MailIcon from "@mui/icons-material/Mail";
import { Link } from "react-router-dom";

function Sidebar() {
  return (
    <>
      <Drawer
        sx={{
          width: 240,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: 240,
            boxSizing: "border-box",
          },
        }}
        variant="permanent"
        anchor="left"
      >
        <Toolbar style={{ background: "primary" }}>
          <Typography component={Link} to="/">首頁</Typography>
        </Toolbar>
        {/* <Divider /> */}
        <List>
          <ListItem key="個人資料" disablePadding component={Link} to="/user">
            <ListItemButton>
              <ListItemText primary="個人資料" />
            </ListItemButton>
          </ListItem>
          <ListItem
            key="疫苗預約"
            disablePadding
            component={Link}
            to="/reservation"
          >
            <ListItemButton>
              <ListItemText primary="疫苗預約" />
            </ListItemButton>
          </ListItem>
          <ListItem
            key="預約查詢與取消"
            disablePadding
            component={Link}
            to="/status"
          >
            <ListItemButton>
              <ListItemText primary="預約查詢與取消" />
            </ListItemButton>
          </ListItem>
        </List>
        <Divider />
        {/* <List>
        {["All mail", "Trash", "Spam"].map((text, index) => (
          <ListItem key={text} disablePadding>
            <ListItemButton>
              <ListItemIcon>
                {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
              </ListItemIcon>
              <ListItemText primary={text} />
            </ListItemButton>
          </ListItem>
        ))}
        </List>*/}
      </Drawer>
    </>
  );
}

export default Sidebar;

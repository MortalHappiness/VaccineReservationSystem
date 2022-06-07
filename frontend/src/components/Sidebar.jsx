import Drawer from "@mui/material/Drawer";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import { Link } from "react-router-dom";
import PersonIcon from "@mui/icons-material/Person";
import VaccinesIcon from "@mui/icons-material/Vaccines";
import EventNoteIcon from "@mui/icons-material/EventNote";

// SideBar component
export default function SideBar({ open, handleCloseDrawer }) {
  return (
    <Drawer anchor="left" open={open} onClose={handleCloseDrawer}>
      <List sx={{ minWidth: 300 }}>
        <ListItem key="個人資料" disablePadding>
          <ListItemButton onClick={handleCloseDrawer} component={Link} to="/">
            <ListItemIcon>
              <PersonIcon />
            </ListItemIcon>
            <ListItemText primary="個人資料" />
          </ListItemButton>
        </ListItem>
        <ListItem key="疫苗預約" disablePadding>
          <ListItemButton
            onClick={handleCloseDrawer}
            component={Link}
            to="/reservation"
          >
            <ListItemIcon>
              <VaccinesIcon />
            </ListItemIcon>
            <ListItemText primary="疫苗預約" />
          </ListItemButton>
        </ListItem>
        <ListItem key="預約查詢與取消" disablePadding>
          <ListItemButton
            onClick={handleCloseDrawer}
            component={Link}
            to="/status"
          >
            <ListItemIcon>
              <EventNoteIcon />
            </ListItemIcon>
            <ListItemText primary="預約查詢與取消" />
          </ListItemButton>
        </ListItem>
      </List>
    </Drawer>
  );
}

// mui
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import LogoutIcon from "@mui/icons-material/Logout";

export default function NavigationBar({
  isLogin,
  handleLogout,
  user,
  handleOpenDrawer,
}) {
  return (
    <AppBar position="fixed">
      <Toolbar>
        <IconButton
          size="large"
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{ mr: 2 }}
          onClick={handleOpenDrawer}
        >
          <MenuIcon />
        </IconButton>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          疫苗預約系統
        </Typography>
        {isLogin && (
          <Box sx={{ display: "flex", alignItems: "center" }}>
            <Typography variant="h6" component="div">
              {user.name}
            </Typography>
            <IconButton sx={{ ml: 2, color: "#ffffff" }} onClick={handleLogout}>
              <LogoutIcon />
            </IconButton>
          </Box>
        )}
      </Toolbar>
    </AppBar>
  );
}

import { useState } from "react";
// mui
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
// api
import { SessionAPI } from "../api";

export default function Login({ setIsLogin, setUser }) {
  const [nationID, setNationID] = useState("");
  const [healthCardID, setHealthCardID] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();
    const user = await SessionAPI.createSession(nationID, healthCardID);
    setUser(user);
    setIsLogin(true);
  };

  return (
    <Container component="main" maxWidth="xs" bgcolor="white">
      <Box
        sx={{
          marginTop: 4,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component="h1" variant="h5">
          登入
        </Typography>
        <Box component="form" noValidate sx={{ mt: 1 }} onSubmit={handleSubmit}>
          <TextField
            margin="normal"
            type="password"
            required
            fullWidth
            id="nationID"
            label="身分證字號"
            name="nationID"
            autoComplete="nationID"
            autoFocus
            value={nationID}
            onChange={(event) => setNationID(event.target.value)}
          />
          <TextField
            margin="normal"
            type="password"
            required
            fullWidth
            name="healthCardID"
            label="健保卡號碼"
            id="healthCardID"
            autoComplete="healthCardID"
            value={healthCardID}
            onChange={(event) => setHealthCardID(event.target.value)}
          />
          <Button type="submit" fullWidth variant="contained" sx={{ mt: 3 }}>
            登入
          </Button>
        </Box>
      </Box>
    </Container>
  );
}

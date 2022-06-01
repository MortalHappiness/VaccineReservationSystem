import { useState } from "react";
import "./App.css";
import Sidebar from "./Sidebar";
import NavigationBar from "./NavigationBar";
import { Navigate, Route, Routes } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <NavigationBar />
      <Sidebar />
      
        <Routes>
          <Route path="/" element={<h1>main</h1>} />
          <Route path="/user" element={<h1>user</h1>} />
          <Route path="/reservation" element={<h1>reservation</h1>} />
          <Route path="/status" element={<h1>status</h1>} /> 
        </Routes>

    </div>
  );
}

export default App;

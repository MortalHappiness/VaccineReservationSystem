import { useState } from "react";
import logo from "./logo.svg";
import "./App.css";
import Sidebar from "./Sidebar";
import NavigationBar from "./NavigationBar";

function App() {
  const [count, setCount] = useState(0);

  return (
    <div className="App">
      <NavigationBar />
      <Sidebar />
    </div>
  );
}

export default App;

import React from "react";
import { Routes, Route } from "react-router-dom";

import AppWs from "./AppWs";
import Admin from "./admin";

const App = () => <div className="App">
  <Routes>
        <Route path="/" element={<AppWs />} />
        <Route path="/admin" element={<Admin />} />
  </Routes>
</div>;

export default App;
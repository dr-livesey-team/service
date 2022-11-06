import React from "react";
import MainPage from "./views/MainPage/MainPage";
import Header from "./components/Header/Header";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import RequestPage from "./views/RequestPage/RequestPage";
import StatPage from "./views/StatPage/StatPage";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Header />
        <Routes>
          <Route path="/monitoring" element={<MainPage />} />
          <Route path="/request/" element={<RequestPage id={0} />} />
          <Route path="/statistics/" element={<StatPage />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;

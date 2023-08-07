import React from "react";
import ReactDOM from "react-dom/client";
import "bulma/css/bulma.min.css";
import {  Route,BrowserRouter,Routes} from "react-router-dom";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";

ReactDOM.createRoot(document.getElementById("root")).render(

  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/connect" Component={App} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
  
);

reportWebVitals();
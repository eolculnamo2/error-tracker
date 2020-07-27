import React, { useReducer } from "react";
import { createGlobalStyle } from "styled-components";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import { library } from "@fortawesome/fontawesome-svg-core";
import { fab } from "@fortawesome/free-brands-svg-icons";
import { faUserAlt, faKey } from "@fortawesome/free-solid-svg-icons";
import Login from "./pages/Login";
import NewOrg from "./pages/NewOrg";
import Dashboard from "./pages/Dashboard";
import { StateContext, initialState, reducer } from "./context/StateContext";

const Universal = createGlobalStyle`
  body {
    padding: 0;
    box-sizing: none;
    margin: 0;
    background-color: white;
    font-family: "Segoe UI";

    input {
      font-family: "Segoe UI"; 
    }
  }
`;

function App() {
  library.add(fab, faKey, faUserAlt);
  const [state, dispatch] = useReducer<any>(reducer, initialState);

  fetch("http://localhost:8080/test");

  return (
    <StateContext.Provider value={{ state, dispatch }}>
      <Router>
        <Universal />
        <Switch>
          <Route exact path="/" component={Login} />
          <Route exact path="/neworganization" component={NewOrg} />
          <Route exact path="/dashboard" component={Dashboard} />
        </Switch>
      </Router>
    </StateContext.Provider>
  );
}

export default App;

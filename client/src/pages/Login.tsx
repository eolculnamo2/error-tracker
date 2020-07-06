import React, { useState, useContext } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { NavLink } from "react-router-dom";
import Navbar from "../components/Navbar";
import {
  MainSection,
  SecondarySection,
  LoginBox,
  LoginBtnnsSpacer1,
  LoginBtnnsSpacer2,
  InputStyling,
  IconStyling,
  LoginBtnsDiv,
  InputDiv,
} from "../styling/LoginStyling";
import { StateContext } from "../context/StateContext";

const Login = () => {
  const { state, dispatch } = useContext(StateContext);

  // const { email } = state;
  const [error, setError] = useState<string>("");

  return (
    <div>
      {/* {email}  */}
      <Navbar />
      <MainSection>
        <SecondarySection>
          <LoginBox>
            <InputStyling>
              <InputDiv>
                <IconStyling>
                  <FontAwesomeIcon icon="user-alt" />
                </IconStyling>
                <input type="text" placeholder="Email Address" />
              </InputDiv>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <IconStyling>
                  <FontAwesomeIcon icon="key" />
                </IconStyling>
                <input type="password" placeholder="Password" />
              </InputDiv>
            </InputStyling>
            <LoginBtnsDiv>
              <LoginBtnnsSpacer1>
                <NavLink to="/dashboard">Sign In</NavLink>
              </LoginBtnnsSpacer1>
              <LoginBtnnsSpacer2>
                <NavLink to="/neworganization">New Organization</NavLink>
              </LoginBtnnsSpacer2>
            </LoginBtnsDiv>
          </LoginBox>
        </SecondarySection>
      </MainSection>
    </div>
  );
};

export default Login;

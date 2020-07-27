import React, { useState, useContext } from "react";
import axios from "axios";
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
  Errors,
} from "../styling/LoginStyling";
import { StateContext } from "../context/StateContext";

const Login = () => {
  const { state, dispatch } = useContext(StateContext);

  const { email, password } = state;

  const [passwordError, setPasswordError] = useState<string>("");
  const [emailError, setEmailError] = useState<string>("");

  const signIn = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (!email) {
        setEmailError("Email address is required");
      } else {
        setEmailError("");
      }
      if (!password) {
        setPasswordError("Password is required");
      } else {
        setPasswordError("");
      }
      if (email && password) {
        let user = { email, password };

        const config = {
          headers: {
            "Content-Type": "application/json",
          },
        };

        const body = JSON.stringify(user);

        let res = await axios.post(
          "http://localhost:8080/auth/login",
          body,
          config
        );
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div>
      <Navbar />
      <MainSection>
        <SecondarySection>
          <LoginBox>
            <InputStyling>
              <InputDiv>
                <IconStyling>
                  <FontAwesomeIcon icon="user-alt" />

                </IconStyling>
                <input
                  type="text"
                  placeholder="Email Address"
        
                />
              </InputDiv>
              <Errors>{emailError}</Errors>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <IconStyling>
                  <FontAwesomeIcon icon="key" />
                </IconStyling>
                <input
                  type="password"
                  placeholder="Password"
                />
              </InputDiv>
              <Errors>{passwordError}</Errors>
            </InputStyling>
            <LoginBtnsDiv>
              <LoginBtnnsSpacer1>
                <NavLink to="/dashboard" onClick={signIn}>
                  Sign In
                </NavLink>
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

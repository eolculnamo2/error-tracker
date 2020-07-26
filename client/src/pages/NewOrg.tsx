import React, { useState, useContext } from "react";
import axios from "axios";
import { NavLink } from "react-router-dom";
import Navbar from "../components/Navbar";
import {
  MainSection,
  SecondarySection,
  LoginBox,
  LoginBtnnsSpacer1,
  LoginBtnnsSpacer2,
  InputStyling,
  LoginBtnsDiv,
  InputDiv,
  Errors,
} from "../styling/NewOrgStyling";
import { StateContext, initialState, APP_ACTIONS } from "../context/StateContext";

const NewOrg = () => {
  const { state, dispatch } = useContext(StateContext);

  const { email, firstName, lastName, Name, website, password } = initialState;

  const [emailError, setEmailError] = useState("");
  const [passwordError, setPasswordError] = useState("");
  const [firstNameError, setFirstNameError] = useState("");
  const [lastNameError, setLastNameError] = useState("");
  const [nameError, setNameError] = useState("");
  const [websiteError, setWebsiteError] = useState("");
  const [confirmError, setConfirmError] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const createOrg = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (!state.email)  {
        setEmailError("Valid email address is required");
      } 
      if (state.email.includes("@")) {
        setEmailError("fgfg");
      }
      if (!state.firstName) {
        setFirstNameError("Enter your first name");
      }
      if (state.firstName) {
        setFirstNameError("");
      }
      if (!state.lastName) {
        setLastNameError("Enter your last name");
      }
      if (state.lastName) {
        setLastNameError("");
      }
      if (!state.Name) {
        setNameError("Organization name is required");
      }
      if (state.Name) {
        setNameError("");
      }
      if (!state.website) {
        setWebsiteError("Enter your organization's website url");
      }
      if (state.website) {
        setWebsiteError("");
      }
      if (!state.password) {
        setPasswordError("Password is required");
      }
      if (state.password) {
        setPasswordError("");
      }
      if (confirmPassword === state.password  ) {
        setConfirmError("");
      } if (confirmPassword !== state.password  ) {
        setConfirmError("The passwords do not match");
      }
      else {
        let newOrg = { email, firstName, lastName, Name, website, password };

        const config = {
          headers: {
            "Content-Type": "application/json",
          },
        };

        const body = JSON.stringify(newOrg);

        let res = await axios.post(
          "http://localhost:8080/new-org",
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
                <input
                  type="text"
                  placeholder="Enter your email address"
                  onChange={(e) =>
                    dispatch({
                      type: APP_ACTIONS.UPDATE_EMAIL,
                      payload: e.target.value,
                    })
                  }
                />
              </InputDiv>
              <Errors>{emailError}</Errors>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="First name" onChange={(e) =>
                    dispatch({
                      type: APP_ACTIONS.UPDATE_FIRST_NAME,
                      payload: e.target.value,
                    })
                  } />
              </InputDiv>
              <Errors>{firstNameError}</Errors>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="Last name" onChange={(e) =>
                    dispatch({
                      type: APP_ACTIONS.UPDATE_LAST_NAME,
                      payload: e.target.value,
                    })
                  }  />
              </InputDiv>
              <Errors>{lastNameError}</Errors>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="Organization name" onChange={(e) =>
                    dispatch({
                      type: APP_ACTIONS.UPDATE_NAME,
                      payload: e.target.value,
                    })
                  } />
              </InputDiv>
              <Errors>{nameError}</Errors>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="Website URL" onChange={(e) =>
                    dispatch({
                      type: APP_ACTIONS.UPDATE_WEBSITE,
                      payload: e.target.value,
                    })
                  } />
              </InputDiv>
              <Errors>{websiteError}</Errors>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="password" placeholder="Enter a password" onChange={(e) =>
                    dispatch({
                      type: APP_ACTIONS.UPDATE_PASSWORD,
                      payload: e.target.value,
                    })
                  } />
              </InputDiv>
              <Errors>{passwordError}</Errors>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input
                  type="password"
                  placeholder="Confirm password"
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                />
              </InputDiv>
              <Errors>{confirmError}</Errors>
            </InputStyling>
            <LoginBtnsDiv>
              <LoginBtnnsSpacer1>
                <NavLink to="/dashboard" onClick={createOrg}>
                  Create Organization
                </NavLink>
              </LoginBtnnsSpacer1>
              <LoginBtnnsSpacer2>
                Already have an account?
                <NavLink to="/">Sign in here</NavLink>
              </LoginBtnnsSpacer2>
            </LoginBtnsDiv>
          </LoginBox>
        </SecondarySection>
      </MainSection>
    </div>
  );
};

export default NewOrg;

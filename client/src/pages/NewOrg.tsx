import React from "react";
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
} from "../styling/NewOrgStyling";

const NewOrg = () => {
  return (
    <div>
      <Navbar />
      <MainSection>
        <SecondarySection>
          <LoginBox>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="Enter your email address" />
              </InputDiv>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="First name" />
              </InputDiv>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="Last name" />
              </InputDiv>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="Organization name" />
              </InputDiv>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="text" placeholder="Website URL" />
              </InputDiv>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="password" placeholder="Enter a password" />
              </InputDiv>
            </InputStyling>
            <InputStyling>
              <InputDiv>
                <input type="password" placeholder="Confirm password" />
              </InputDiv>
            </InputStyling>
            <LoginBtnsDiv>
              <LoginBtnnsSpacer1>
                <NavLink to="/dashboard">Create Organization</NavLink>
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

import styled from "styled-components";

export const MainSection = styled.div`
  margin-top: 13rem; //not perfectly centered when working with different resolutions
  background-color: none;
  align-items: center;
`;

export const SecondarySection = styled.div`
  display: flex;
  justify-content: center;
`;

export const LoginBox = styled.div`
  height: 14rem;
  width: 30rem;
  background-color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  border: 0.1rem solid black;

  i {
    font-size: 0.9rem;
  }
`;

export const InputStyling = styled.div`
  padding: 0.6rem;
  align-items: center;

  input {
    outline: none;
    border: none;
    height: 1.3rem;
    font-size: 0.8rem;
    width: 16rem;
  }
`;

export const IconStyling = styled.span`
  font-size: 0.8em;
  padding-right: 0.2rem;
`;

export const InputDiv = styled.div`
  border: none;
  border-bottom: 0.01rem solid #999;
  padding: 0.2rem;
  width: 18rem;
`;

export const LoginBtnsDiv = styled.div`
  display: flex;
  padding-top: 1rem;
  align-items: center;

  button {
    font-size: 0.8rem;
    width: 8rem;
    background-color: transparent;
    outline: none;
    border: none;
    cursor: pointer;
  }

  a {
    text-decoration: none;
    color: black;

    :hover {
      color: green;
    }
  }
`;

export const LoginBtnnsSpacer1 = styled.span`
  padding-right: 1rem;
`;

export const LoginBtnnsSpacer2 = styled.span`
  padding-left: 1rem;
`;

export const Errors = styled.div`
  color: red;
  height: 0.3rem;
  font-size: 0.8rem;
  text-align: center;
`;

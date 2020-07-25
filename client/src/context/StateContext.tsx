import { createContext } from "react";

export interface IState {
  email: string;
  first_name: string;
  last_name: string;
  organization_name: string;
  url: string;
  password: string;
}

export const initialState = {
  email: ".",
  first_name: "",
  last_name: "",
  organization_name: "",
  url: "",
  password: "",
};

export enum APP_ACTIONS {
  UPDATE_EMAIL = "UPDATE_EMAIL",
  UPDATE_FIRST_NAME = "UPDATE_FIRST_NAME",
  UPDATE_LAST_NAME = "UPDATE_LAST_NAME",
}

type AppAction = { [key: string]: (state: IState, action: any) => IState };

export const appActions: AppAction = {
  [APP_ACTIONS.UPDATE_EMAIL]: (state: IState, actions: any) => {
    return { ...state, email: actions.payload };
  },

  [APP_ACTIONS.UPDATE_FIRST_NAME]: (state: IState, actions: any) => {
    return { ...state, first_name: actions.payload };
  },

  [APP_ACTIONS.UPDATE_LAST_NAME]: (state: any, actions: any) => {
    return { ...state, last_name: actions.payload };
  },
};

export const reducer = (state: IState, action: any) => {
  return appActions[action.type](state, action);
};

export const StateContext: any = createContext(initialState);

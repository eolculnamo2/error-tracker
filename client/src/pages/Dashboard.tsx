import React from "react";
import Navbar from "../components/Navbar";
import {
  MainSection,
  SecondarySection,
  SideBar,
} from "../styling/DashboardStyling";

const Dashboard = () => {
  return (
    <div>
      <Navbar />
      <MainSection>
        <SideBar></SideBar>
      </MainSection>
    </div>
  );
};

export default Dashboard;

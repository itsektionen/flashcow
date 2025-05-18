import { Navbar } from "@/components/navbar";
import { Outlet } from "react-router";

function MainLayout() {
  return (
    <>
      <Navbar></Navbar>
      <Outlet />
    </>
  );
}

export default MainLayout;

import MainLayout from "@/layout/MainLayout";
import MobileMainLayout from "@/layout/MainLayout.m";
import { RoutePaths } from "@/router";
import useStore from "@/store";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const navigate = useNavigate();
  useEffect(() => {
    /^\/console\/?$/i.test(window.location.pathname) &&
      navigate(RoutePaths.AutoShare);
  }, []);

  const [isMobile, setIsMobile] = useStore((state) => [
    state.isMobile,
    state.setIsMobile,
  ]);

  useEffect(() => {
    const resize = () => setIsMobile(window.innerWidth < 768);
    resize();
    window.addEventListener("resize", resize);

    return () => window.removeEventListener("resize", resize);
  }, []);

  return isMobile ? <MobileMainLayout /> : <MainLayout />;
};

export default Home;

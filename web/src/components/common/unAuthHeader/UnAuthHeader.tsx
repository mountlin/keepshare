import { theme } from "antd";
import { StyledLink, NavWrapper } from "./style";
import { RoutePaths } from "@/router";

const UnAuthHeader = () => {
  const { token } = theme.useToken();

  return (
    <NavWrapper>
      <StyledLink color={token.colorText} to={"/"}>
        Home
      </StyledLink>
      <StyledLink color={token.colorText} to={"https://github.com/keepshare"}>
        Github
      </StyledLink>
      <StyledLink color={token.colorText} to={RoutePaths.Donation}>
        Donation
      </StyledLink>
    </NavWrapper>
  );
};

export default UnAuthHeader;

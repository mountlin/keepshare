import styled from "styled-components";
import backgroundSvg from "@/assets/images/cloud-bg.svg";

export const Background = styled.div`
  min-width: 100vh;
  min-height: 100vh;
  background: url(${backgroundSvg}) no-repeat center/cover;
`;

export const LogoPng = styled.img`
  margin: 12px 24px;
  height: 64px;
`;

export const ContentWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: calc(100vh - 88px);
`;

export const BannerWrapper = styled.div<{ poster?: string }>`
  position: relative;
  width: 480px;
  height: 335px;
  border-radius: 5px;
  background: ${(props) =>
    props.poster ? `url(${props.poster}) no-repeat center/cover` : "#000"};
`;

export const BannerImage = styled.img`
  position: absolute;
  width: 120px;
  height: 120px;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
`;

export const SharedInfoBox = styled.div`
  position: absolute;
  display: flex;
  flex-direction: column;
  width: 450px;
  left: 15px;
  bottom: 15px;
  color: #fff;
  line-height: 1.4em;
  font-size: 14px;
`;

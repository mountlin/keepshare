import { Space, Typography, message } from "antd";
import {
  Background,
  BannerImage,
  BannerWrapper,
  ContentWrapper,
  LogoPng,
  SharedInfoBox,
} from "./style";
import LogoIcon from "@/assets/images/logo-with-text.png";
import { useEffect, useMemo, useState } from "react";
import { SharedLinkInfo, getSharedLinkInfo } from "@/api/link";
import { useSearchParams } from "react-router-dom";
import PrepareStatusBanner from "@/assets/images/prepare-status-banner.png";
import ExpiredStatusBanner from "@/assets/images/expired-status-banner.png";
import { formatBytes } from "@/util";
import { LinkOutlined } from "@ant-design/icons";

const { Paragraph, Title, Text } = Typography;

const enum SharedStatusType {
  PREPARE = "prepare",
  EXPIRED = "expired",
}

const getStatusDescribeText = (status?: SharedStatusType) => {
  if (status === SharedStatusType.EXPIRED) {
    return {
      title: "The current link has expired",
      subtitle:
        "KeepShare cannot find the link corresponding to the current file",
    };
  }

  return {
    title: "Sharing files in preparation...",
    subtitle:
      "The corresponding sharing files are being prepared and will be generated within hours. You can check it out later, or download it locally via the link above, or faster remote download.",
  };
};

const SharedStatus = () => {
  const { title, subtitle } = useMemo(
    () => getStatusDescribeText(SharedStatusType.PREPARE),
    [],
  );
  const [status, setStatus] = useState<SharedStatusType>(
    SharedStatusType.EXPIRED,
  );
  const [fileInfo, setFileInfo] = useState<Partial<SharedLinkInfo>>({});

  const [params] = useSearchParams();

  useEffect(() => {
    const autoId = params.get("id") || "";
    if (!/^\d+$/i.test(autoId)) {
      return;
    }

    getSharedLinkInfo(autoId).then(({ data: fileInfo, error }) => {
      if (fileInfo) {
        setFileInfo(fileInfo);

        const state = fileInfo.state;
        setStatus(
          state === "PROCESSING"
            ? SharedStatusType.PREPARE
            : SharedStatusType.EXPIRED,
        );

        const hostSharedLink = fileInfo.host_shared_link;
        if (state === "OK" && hostSharedLink) {
          // TODO:
          // location.href = hostSharedLink;
        }
      }
      error && message.error(error.message);
    });
  }, []);

  const { original_link: link, title: filename, size: storage } = fileInfo;
  const size = formatBytes((storage as number) || 0);

  return (
    <Background>
      <LogoPng src={LogoIcon} />
      <ContentWrapper>
        <Paragraph style={{ maxWidth: "720px" }}>
          <Title level={3} style={{ textAlign: "center", color: "#000" }}>
            {title}
          </Title>
          <Text style={{ color: "#000" }}>{subtitle}</Text>
        </Paragraph>

        <BannerWrapper>
          <BannerImage
            src={
              status === SharedStatusType.EXPIRED
                ? ExpiredStatusBanner
                : PrepareStatusBanner
            }
            alt="banner"
          />
          <SharedInfoBox>
            <Text
              ellipsis={{ suffix: filename?.slice(filename?.length - 3) }}
              style={{ maxWidth: "100%", color: "#fff" }}
            >
              {filename}
            </Text>
            <Text
              style={{ color: "rgba(255, 255, 255, 0.60)", fontSize: "12px" }}
            >
              {size}
            </Text>
          </SharedInfoBox>
        </BannerWrapper>
        {link && (
          <Space style={{ marginTop: "16px" }}>
            <LinkOutlined />
            <Text copyable style={{ color: "#000" }}>
              {link}
            </Text>
          </Space>
        )}
      </ContentWrapper>
    </Background>
  );
};

export default SharedStatus;

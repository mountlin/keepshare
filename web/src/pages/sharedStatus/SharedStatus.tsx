import { Space, Typography, message, theme } from "antd";
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
import {
  SharedLinkInfo,
  getLinkInfoFromWhatsLink,
  getSharedLinkInfo,
} from "@/api/link";
import { useSearchParams } from "react-router-dom";
import PrepareStatusBanner from "@/assets/images/prepare-status-banner.png";
import ExpiredStatusBanner from "@/assets/images/expired-status-banner.png";
import { formatBytes } from "@/util";
import { LinkOutlined } from "@ant-design/icons";
import useStore from "@/store";

const { Paragraph, Title, Text, Link } = Typography;

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
      "The corresponding sharing files are being prepared and will be generated within hours. You can check it out later, or download it locally via the link below, or faster remote download.",
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
  const [fileInfo, setFileInfo] = useState<
    Partial<SharedLinkInfo> & { screenshot?: string }
  >({});

  const [params] = useSearchParams();

  useEffect(() => {
    const autoId = params.get("id") || "";
    if (!/^\d+$/i.test(autoId)) {
      return;
    }

    // Get data from keepshare server, but the data may be incomplete (returned from PikPak)
    getSharedLinkInfo(autoId).then(({ data: fileInfo, error }) => {
      if (fileInfo) {
        setFileInfo(fileInfo);

        const state = fileInfo.state;
        setStatus(
          state === "CREATED"
            ? SharedStatusType.PREPARE
            : SharedStatusType.EXPIRED,
        );

        const hostSharedLink = fileInfo.host_shared_link;
        if (state === "OK" && hostSharedLink) {
          // TODO:
          // location.href = hostSharedLink;
        }
      }

      // Get data from whatsLink website
      if (!error) {
        getLinkInfoFromWhatsLink(fileInfo?.original_link || "")
          .then(({ data, error }) => {
            if (error) {
              return;
            }
            setFileInfo(
              Object.assign({}, fileInfo, {
                title: fileInfo?.title || data?.name,
                size: fileInfo?.size || data?.size,
                screenshot: data?.screenshots[0]?.screenshot,
              }),
            );
          })
          .catch(() => {});
      }

      error && message.error(error.message);
    });
  }, []);

  const { original_link: link, title: filename, size: storage } = fileInfo;
  const size = formatBytes((storage as number) || 0);

  const { token } = theme.useToken();
  const isMobile = useStore((state) => state.isMobile);

  return (
    <Background>
      <LogoPng src={LogoIcon} />
      <ContentWrapper>
        <Paragraph
          style={{
            maxWidth: isMobile ? "100%" : "720px",
            padding: token.padding,
          }}
        >
          <Title level={3} style={{ textAlign: "center", color: "#000" }}>
            {title}
          </Title>
          <Text style={{ color: "#000" }}>{subtitle}</Text>
        </Paragraph>

        <BannerWrapper style={{ marginInline: token.margin }}>
          {fileInfo.screenshot ? (
            <BannerImage
              src={fileInfo.screenshot}
              alt="banner"
              style={{
                width: "100%",
                height: "100%",
              }}
            />
          ) : (
            <BannerImage
              src={
                status === SharedStatusType.EXPIRED
                  ? ExpiredStatusBanner
                  : PrepareStatusBanner
              }
              alt="banner"
            />
          )}
          <SharedInfoBox>
            <Text
              ellipsis={{ suffix: filename?.slice(filename?.length - 3) }}
              style={{
                color: "rgba(255, 255, 255, 0.60)",
              }}
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
          <Space style={{ marginTop: "16px", paddingInline: token.padding }}>
            <LinkOutlined />
            <Link
              href={link}
              copyable
              style={{ color: "#000", wordBreak: "break-all" }}
            >
              {link}
            </Link>
          </Space>
        )}
        <Link
          href="https://whatslink.info/"
          style={{
            color: "rgba(0, 0, 0, 0.65)",
            paddingInline: token.padding,
          }}
        >
          The images and information of the link comes from whatslink.info.
          whatslink.info
        </Link>
      </ContentWrapper>
    </Background>
  );
};

export default SharedStatus;

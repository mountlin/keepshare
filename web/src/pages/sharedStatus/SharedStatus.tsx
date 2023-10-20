import { Button, Space, Typography, message, theme } from "antd";
import {
  Background,
  BannerImage,
  BannerWrapper,
  ContentWrapper,
  LogoPng,
  SharedInfoBox,
} from "./style";
import LogoIcon from "@/assets/images/logo-with-text.png";
import { useEffect, useState } from "react";
import {
  SharedLinkInfo,
  SharedLinkStatus,
  getLinkInfoFromWhatsLink,
  getSharedLinkInfo,
} from "@/api/link";
import { useSearchParams } from "react-router-dom";
import ShareIcon from "@/assets/images/prepare-status-banner.png";
import LoadingAPng from "@/assets/images/keepshare-loading.png";
import { copyToClipboard, formatBytes, getSupportLanguage } from "@/util";
import { CopyOutlined, LinkOutlined } from "@ant-design/icons";
import useStore from "@/store";
import { Trans, useTranslation } from "react-i18next";
import { RoutePaths } from "@/router";

const { Title, Text, Link } = Typography;

const getStatusDescribeText = (
  status: SharedLinkStatus,
  holder: JSX.Element,
) => {
  const { t } = useTranslation();

  if (["UNKNOWN", "DELETED", "NOT_FOUND", "BLOCKED"].includes(status)) {
    return {
      title: t("mEjDyHbG9xiu_6NlaegOn"),
      subtitle: (
        <Text>
          <Trans i18nKey="dW_5y60qwkDKvThMDiFl" components={[holder]}></Trans>
        </Text>
      ),
    };
  }

  return {
    title: t("2D0jDl0qeMSqvV0Ly6Iyd"),
    subtitle: (
      <Text>
        <Trans i18nKey="jgAl1qdprLfErnmioSgOu" components={[holder]}></Trans>
      </Text>
    ),
  };
};

const Loading = () => {
  const { t } = useTranslation();

  return (
    <Space
      direction="vertical"
      align="center"
      style={{
        position: "fixed",
        top: "50%",
        left: "50%",
        transform: "translate(-50%, -50%)",
      }}
    >
      <img src={LoadingAPng} alt="loading" width={120} />
      <Text>{t("tHxXtk0qRYf6Kh4qNcuHh")}</Text>
    </Space>
  );
};

const SharedStatus = () => {
  const { t, i18n } = useTranslation();

  const { token } = theme.useToken();

  const remoteDownload = (
    <Link
      href="https://mypikpak.com/drive/url-checker"
      style={{ color: token.colorPrimary }}
    >
      {t("hDvGl13AlFfsLIi2jQ3xP")}
    </Link>
  );
  const [status, setStatus] = useState<SharedLinkStatus>("PENDING");
  const { title, subtitle } = getStatusDescribeText(status, remoteDownload);

  const [fileInfo, setFileInfo] = useState<
    Partial<SharedLinkInfo> & { screenshot?: string }
  >({});

  const [params] = useSearchParams();

  const setThemeMode = useStore((state) => state.setThemeMode);
  // status page keep light mode
  useEffect(() => setThemeMode("light"), []);

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
        setStatus(state);

        const hostSharedLink = fileInfo.host_shared_link;
        if (state === "OK" && hostSharedLink) {
          location.href = hostSharedLink;
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

  const isMobile = useStore((state) => state.isMobile);

  const handleCopyLink = () => {
    try {
      link && copyToClipboard(link);
      message.success(t("xKhHo2JwfdzWgJXiJ0GeI"));
    } catch {
      message.error(t("aiCd4EgbrLDu4cdLlBy"));
    }
  };

  useEffect(() => {
    i18n.changeLanguage(getSupportLanguage() as any);
  }, []);

  return (
    <Background>
      <Link href={RoutePaths.Home}>
        <LogoPng src={LogoIcon} />
      </Link>
      {status === "PENDING" ? (
        <Loading />
      ) : (
        <ContentWrapper>
          {fileInfo.screenshot ? (
            <BannerWrapper style={{ marginInline: token.margin }}>
              {fileInfo.screenshot && (
                <BannerImage
                  src={fileInfo.screenshot}
                  alt="banner"
                  style={{
                    width: "100%",
                    height: "100%",
                  }}
                />
              )}
              <SharedInfoBox>
                {t("whMzAm8sGpQfOTqadiXu")} {size}
              </SharedInfoBox>
            </BannerWrapper>
          ) : (
            <Space
              direction="vertical"
              align="center"
              style={{ marginTop: "auto" }}
            >
              <img src={ShareIcon} style={{ width: "94px" }} alt="shareIcon" />
              <Text style={{ color: token.colorTextSecondary }}>
                {t("whMzAm8sGpQfOTqadiXu")} {size}
              </Text>
            </Space>
          )}
          <Text
            style={{
              maxWidth: "min(600px, 100vw)",
              marginTop: "12px",
              textAlign: "center",
              lineHeight: "1.4em",
            }}
          >
            {filename}
          </Text>
          <Space style={{ marginTop: token.marginLG }}>
            <LinkOutlined style={{ fontSize: "24px" }} />
            <Link
              href={link}
              style={{
                fontSize: token.fontSizeLG,
                wordBreak: "break-all",
                color: token.colorTextBase,
              }}
            >
              {link}
            </Link>
          </Space>
          <Space style={{ marginTop: token.margin }}>
            <Button
              type="primary"
              icon={<CopyOutlined />}
              onClick={handleCopyLink}
            >
              {t("fbWqi7mJuMCxEw3SwCf_0")}
            </Button>
          </Space>
          <Space
            align={isMobile ? "center" : "start"}
            style={{ marginTop: "56px" }}
            direction={isMobile ? "vertical" : "horizontal"}
          >
            <Title level={3} style={{ lineHeight: "1em" }}>
              {title}
            </Title>
            <Text style={{ display: "inline-block", maxWidth: "458px" }}>
              {subtitle}
            </Text>
          </Space>
          <Text
            style={{
              color: token.colorTextSecondary,
              marginTop: "auto",
              marginBottom: token.marginLG,
              textAlign: "center",
            }}
          >
            <Trans
              i18nKey="3BuCl0v1UfHeRbDqnuh0N"
              components={[
                <Link
                  underline
                  style={{ color: token.colorTextSecondary }}
                  href="https://whatslink.info"
                >
                  whatslink.info
                </Link>,
              ]}
            ></Trans>
          </Text>
        </ContentWrapper>
      )}
    </Background>
  );
};

export default SharedStatus;

import { Tabs } from "antd";
import SharedLinksBlacklist from "./SharedLinksBlacklist";
import SharedLinksList from "./SharedLinks";

const tabItems = [
  {
    key: "shared-links",
    label: "Shared Links",
    children: <SharedLinksList />,
  },
  {
    key: "link-blacklist",
    label: "Link Blacklist",
    children: <SharedLinksBlacklist />,
  },
];

const SharedLinks = () => {
  return <Tabs items={tabItems}></Tabs>;
};

export default SharedLinks;

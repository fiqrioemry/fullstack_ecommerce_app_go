import React from "react";
import { Link } from "react-router-dom";

const WebLogo = () => {
  return (
    <Link to="/">
      <img src="/logo.png" alt="logo" className="h-12 md:block hidden" />
      <img src="/mobile-logo.png" alt="logo" className="h-10 block md:hidden" />
    </Link>
  );
};

export { WebLogo };

import React from "react";
import "./Header.css";

function Header({ game }) {
  return (
    <div className="header">
      <b className="logo">{game.name}</b>
      <div className="header-right">
        <b>
          guess [{game.guess1},{game.guess2}]
        </b>
      </div>
    </div>
  );
}

export default Header;

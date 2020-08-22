import React from "react";
import "./Header.css";

function Header({ game }) {
  return (
    <div className="header">
      {game && (
        <>
          <b className="logo">{game.name}</b>
          <div className="header-right">
            {game.isPlayer ? (
              <b>
                guess [{game.guess1},{game.guess2}]
              </b>
            ) : (
              <b>Observer</b>
            )}
          </div>
        </>
      )}
    </div>
  );
}

export default Header;

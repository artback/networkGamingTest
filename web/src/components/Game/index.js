import React, { useEffect } from "react";
import Join from "../Join";
import Board from "../Board";
import Header from "../Header";
import "./Game.css";
const { host } = window.location;
const protocolPrefix = window.location.protocol === "https:" ? "wss:" : "ws:";
function Game() {
  useEffect(() => {
    if (!config) {
      fetch("/api/config")
        .then((response) => response.json())
        .then((data) => setConfig(data))
        .catch((err) => console.error(err));
    }
  });
  const [config, setConfig] = React.useState();
  const [game, setGame] = React.useState({
    name: "",
    guess1: 0,
    guess2: 0,
  });
  const [isJoined, setIsJoined] = React.useState(false);
  const [total, setTotal] = React.useState();
  const [scoreBoard, setScoreBoard] = React.useState();
  const joined = (game) => {
    const { name, guess1, guess2, isPlayer } = game;
    setIsJoined(!isJoined);
    let socketUrl = `${protocolPrefix}//${host}/ws/game?name=${name}`;
    if (isPlayer) {
      socketUrl += `&guess=${guess1}&guess=${guess2}`;
    }

    const socket = new WebSocket(socketUrl);
    socket.onmessage = onmessage;
    socket.onclose = onclose;
  };
  const onmessage = ({ data }) => {
    const { type, payload } = JSON.parse(data);
    switch (type) {
      case "total":
        setTotal(payload);
        break;
      case "result":
        setScoreBoard(payload);
        break;
      case "winner":
        setIsJoined(false);
        setScoreBoard(undefined);
        break;
      case "error":
        console.error(data);
        break;
      default:
      //console.log(data);
    }
  };
  const onclose = () => {
    setIsJoined(false);
    setScoreBoard(undefined);
  };
  return (
    <div className={"container"}>
      {isJoined && <Header game={game} />}
      <div className={"flex"}>
        {!isJoined ? (
          <Join onJoin={joined} config={config} game={game} setGame={setGame} />
        ) : (
          <Board scoreBoard={scoreBoard} title="scoreboard" />
        )}
        <Board scoreBoard={total} title="total" />
      </div>
    </div>
  );
}

export default Game;

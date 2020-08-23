import React from "react";
import Board from "./index";

export default {
  title: "Board",
  component: Board,
  argTypes: {
    backgroundColor: { control: "color" },
  },
};
const scoreboard = {
  ola: { score: 5 },
  per: { score: 5 },
  kurt: { score: 5 },
  mike: { score: 5 },
  k: { score: 5 },
  dsa: { score: 5 },
  sven: { score: 1 },
  hitch: { score: 1 },
  bitch: { score: 1 },
  snitch: { score: 1 },
  witch: { score: 1 },
};

export const Default = () => <Board scoreBoard={scoreboard} title={"total"} />;
export const Empty = () => <Board />;

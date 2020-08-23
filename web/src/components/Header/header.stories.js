import React from "react";
import Header from "./index";


export default {
  title: "header",
  component: Header,
};
const game = {
  name: "ola",
  guess1: 0,
  guess2: 10,
  isPlayer: true,
};
const observer = {
  name: "ola",
  isPlayer: false,
};

export const Default = () => <Header game={game} />;
export const Observer = () => <Header game={observer} />;
export const Empty = () => <Header />;

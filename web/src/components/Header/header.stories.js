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
};

export const Default = () => <Header game={game} />;
export const Empty = () => <Header />;

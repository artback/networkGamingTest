import React from "react";
import Winner from "./index";

export default {
  title: "Winner",
  component: Winner,
};
const winner= {
  name: "ola",
  score: -10
};

export const Default = () => (
  <Winner
      winner={winner}
  />
);

import React from "react";
import Join from "./index";

export default {
  title: "Join",
  component: Join,
};
const config = {
  begin_interval: 0,
  end_interval: 10,
};

export const Default = () => (
  <Join
    config={config}
    onJoin={(join) => console.log(join)}
    game={{ name: "ola", guess1: 1, guess2: 5, isPlayer: true }}
  />
);
export const Empty = () => <Join game={{}} />;

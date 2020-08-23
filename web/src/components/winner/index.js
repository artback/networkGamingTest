import React from "react";
import {Typography} from "@material-ui/core";

function Winner({winner}) {

return (
    <div>
        <Typography variant={"h2"}>Previous games winner is: {winner.name}</Typography>
        <Typography variant={"h3"}>With score {winner.score}</Typography>
    </div>
);
}

export default Winner;

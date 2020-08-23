import React from "react";
import TableContainer from "@material-ui/core/TableContainer";
import Table from "@material-ui/core/Table";
import TableHead from "@material-ui/core/TableHead";
import TableBody from "@material-ui/core/TableBody";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import Paper from "@material-ui/core/Paper";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";

function Board({ scoreBoard, title }) {
  const style = {
    border: "4px solid rgb(0, 94, 156)",
    borderRadius: "10px",
    width: "650px",
    height: "650px",
    margin: "0 auto",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  };
  const center = {
    ...style,
  };

  const sortScoreAsc = (a, b) => b[1] - a[1];
  return scoreBoard ? (
    <TableContainer style={style} component={Paper}>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell align="right">Score</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {scoreBoard &&
            Object.entries(scoreBoard)
              .sort(sortScoreAsc)
              .map(([name, result]) => (
                <TableRow key={name}>
                  <TableCell component="th" scope="row">
                    {name}
                  </TableCell>
                  <TableCell align="right">{result.score}</TableCell>
                </TableRow>
              ))}
        </TableBody>
      </Table>
    </TableContainer>

  ) : (
    <div style={center}>
      <b>No score data</b>
    </div>
  );
}

export default Board;

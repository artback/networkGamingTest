import React from "react";
import FormControl from "@material-ui/core/FormControl";
import TextField from "@material-ui/core/TextField";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Switch from "@material-ui/core/Switch";
import Button from "@material-ui/core/Button";

function Join({ config, onJoin, game, setGame }) {
  const center = {
    border: "4px solid rgb(0, 94, 156)",
    borderRadius: "10px",
    width: "650px",
    height: "650px",
    margin: "0 auto",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  };
  const handleSwitch = () => {
    setGame({ ...game, isPlayer: !game.isPlayer });
  };
  const setGameFields = (event, field) =>
    setGame({ ...game, [field]: event.target.value });
  const validate = ({ name, guess1, guess2 }) => {
    const nOk = name !== undefined && name !== "";

    const guessOk = (guess) =>
      guess !== undefined &&
      guess <= config.end_interval &&
      config.begin_interval <= guess;
    return nOk && (game.isPlayer ? guessOk(guess1) && guessOk(guess2) : true);
  };
  return (
    <div style={center}>
      {config ? (
        <FormControl>
          <TextField
            id="my-input"
            label="Name"
            value={game.name}
            onChange={(event) => setGameFields(event, "name")}
          />
          <FormControlLabel
            control={
              <Switch
                value={game.isPlayer}
                onChange={handleSwitch}
                name="player"
              />
            }
            label={game.isPlayer ? "Player" : "Observer"}
          />
          {game.isPlayer && (
            <>
              <TextField
                id="standard-number"
                label="Guess 1"
                type="number"
                value={game.guess1}
                InputProps={{
                  inputProps: {
                    min: config.begin_interval,
                    max: config.end_interval,
                  },
                }}
                onChange={(event) => setGameFields(event, "guess1")}
                InputLabelProps={{
                  shrink: true,
                }}
              />
              <TextField
                id="standard-number"
                label="Guess 2"
                type="number"
                onChange={(event) => setGameFields(event, "guess2")}
                value={game.guess2}
                InputProps={{
                  inputProps: {
                    min: config.begin_interval,
                    max: config.end_interval,
                  },
                }}
                InputLabelProps={{
                  shrink: true,
                }}
              />
            </>
          )}
          <Button
            variant="contained"
            color="primary"
            onClick={() => onJoin(game)}
            disabled={!validate(game)}
          >
            Join
          </Button>
        </FormControl>
      ) : (
        <b>No config data</b>
      )}
    </div>
  );
}

export default Join;

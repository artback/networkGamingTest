package scoreboard

type Scoreboard map[string]*Result

func Join(s Scoreboard, s2 Scoreboard) Scoreboard {
	res := Scoreboard{}
	for name, result := range s {
		if res[name] == nil {
			res[name] = &Result{Name: name, guess: result.guess}
		}
		res[name].addScore(result.Score)

	}
	for name, result := range s2 {
		if res[name] == nil {
			res[name] = &Result{Name: name, guess: result.guess}
		}
		res[name].addScore(result.Score)
	}
	return res
}

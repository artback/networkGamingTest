package scoreboard

type Scoreboard map[string]int

func Join(s Scoreboard, s2 Scoreboard) Scoreboard {
	res := Scoreboard{}
	for player, score := range s {
		res[player] += score
	}
	for player, score := range s2 {
		res[player] += score
	}
	return res
}

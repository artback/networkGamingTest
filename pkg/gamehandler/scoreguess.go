package gamehandler

func scoreGuess(actual int, guess [2]int) int {
	score := 0
	if actual == guess[0] || actual == guess[1] {
		score += 5
	} else if actual >= guess[0] && actual <= guess[1] {
		score += 5 - (guess[1] - guess[0])
	} else {
		score--
	}
	return score
}
